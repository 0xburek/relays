pragma solidity ^0.5.10;

/** @title Relay */
/** @author Summa (https://summa.one) */

import {SafeMath} from "@summa-tx/bitcoin-spv-sol/contracts/SafeMath.sol";
import {TypedMemView} from "@summa-tx/bitcoin-spv-sol/contracts/TypedMemView.sol";
import {ViewBTC} from "@summa-tx/bitcoin-spv-sol/contracts/ViewBTC.sol";
import {ViewSPV} from "@summa-tx/bitcoin-spv-sol/contracts/ViewSPV.sol";
import {IRelay} from "./Interfaces.sol";

contract Relay is IRelay {
    using SafeMath for uint256;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;
    using ViewBTC for bytes29;
    using ViewSPV for bytes29;

    /* using BytesLib for bytes;
    using BTCUtils for bytes;
    using ValidateSPV for bytes; */

    int256 constant IDEAL_BLOCK_TIME = 600;
    int256 constant HALFLIFE = 172800;
    int256 constant RADIX = 2 ** 16;
    uint256 constant MAX_BITS = 0x1d00ffff;

    // How often do we store the height?
    // A higher number incurs less storage cost, but more lookup cost
    uint32 public constant HEIGHT_INTERVAL = 4;

    bytes32 internal relayGenesis;
    bytes32 internal bestKnownDigest;
    bytes32 internal lastReorgCommonAncestor;
    mapping (bytes32 => bytes32) internal previousBlock;
    mapping (bytes32 => uint256) internal blockHeight;

    
    uint256 anchorHeight;
    uint256 anchorParentTime;
    uint256 anchorNBits;

    /// @notice                   Gives a starting point for the relay
    /// @dev                      We don't check this AT ALL really. Don't use relays with bad genesis
    /// @param  _genesisHeader    The starting header
    /// @param  _height           The starting height
    constructor(
                bytes memory _genesisHeader,
                uint256 _height,
                uint256 _anchorHeight,
                uint256 _anchorParentTime,
                uint256 _anchorNBits
                ) public {
        bytes29 _genesisView = _genesisHeader.ref(0).tryAsHeader();
        require(_genesisView.notNull(), "Stop being dumb");
        bytes32 _genesisDigest = _genesisView.hash256();

        relayGenesis = _genesisDigest;
        bestKnownDigest = _genesisDigest;
        lastReorgCommonAncestor = _genesisDigest;
        blockHeight[_genesisDigest] = _height;

        anchorHeight = _anchorHeight;
        anchorParentTime = _anchorParentTime;
        anchorNBits = _anchorNBits;
    }

    /// @notice     Getter for relayGenesis
    /// @dev        This is an initialization parameter
    /// @return     The hash of the first block of the relay
    function getRelayGenesis() public view returns (bytes32) {
        return relayGenesis;
    }

    /// @notice     Getter for bestKnownDigest
    /// @dev        This updated only by calling markNewHeaviest
    /// @return     The hash of the best marked chain tip
    function getBestKnownDigest() public view returns (bytes32) {
        return bestKnownDigest;
    }

    /// @notice     Getter for relayGenesis
    /// @dev        This is updated only by calling markNewHeaviest
    /// @return     The hash of the shared ancestor of the most recent fork
    function getLastReorgCommonAncestor() public view returns (bytes32) {
        return lastReorgCommonAncestor;
    }

    /// @notice         Finds the height of a header by its digest
    /// @dev            Will fail if the header is unknown
    /// @param _digest  The header digest to search for
    /// @return         The height of the header, or error if unknown
    function findHeight(bytes32 _digest) external view returns (uint256) {
        return _findHeight(_digest);
    }

    /// @notice         Finds an ancestor for a block by its digest
    /// @dev            Will fail if the header is unknown
    /// @param _digest  The header digest to search for
    /// @return         The height of the header, or error if unknown
    function findAncestor(bytes32 _digest, uint256 _offset) external view returns (bytes32) {
        return _findAncestor(_digest, _offset);
    }

    /// @notice             Checks if a digest is an ancestor of the current one
    /// @dev                Limit the amount of lookups (and thus gas usage) with _limit
    /// @param _ancestor    The prospective ancestor
    /// @param _descendant  The descendant to check
    /// @param _limit       The maximum number of blocks to check
    /// @return             true if ancestor is at most limit blocks lower than descendant, otherwise false
    function isAncestor(bytes32 _ancestor, bytes32 _descendant, uint256 _limit) external view returns (bool) {
        return _isAncestor(_ancestor, _descendant, _limit);
    }

    /// @notice             Adds headers to storage after validating
    /// @dev                We check integrity and consistency of the header chain
    /// @param  _anchor     The header immediately preceeding the new chain
    /// @param  _headers    A tightly-packed list of 80-byte Bitcoin headers
    /// @return             True if successfully written, error otherwise
    function addHeaders(bytes calldata _anchor, bytes calldata _headers) external returns (bool) {
        bytes29 _headersView = _headers.ref(0).tryAsHeaderArray();
        bytes29 _anchorView = _anchor.ref(0).tryAsHeader();

        require(_headersView.notNull(), "Header array length must be divisible by 80");
        require(_anchorView.notNull(), "Anchor must be 80 bytes");

        return _addHeaders(_anchorView, _headersView);
    }

    /// @notice                   Gives a starting point for the relay
    /// @dev                      We don't check this AT ALL really. Don't use relays with bad genesis
    /// @param  _ancestor         The digest of the most recent common ancestor
    /// @param  _currentBest      The 80-byte header referenced by bestKnownDigest
    /// @param  _newBest          The 80-byte header to mark as the new best
    /// @param  _limit            Limit the amount of traversal of the chain
    /// @return                   True if successfully updates bestKnownDigest, error otherwise
    function markNewHeaviest(
        bytes32 _ancestor,
        bytes calldata _currentBest,
        bytes calldata _newBest,
        uint256 _limit
    ) external returns (bool) {
        bytes29 _new = _newBest.ref(0).tryAsHeader();
        bytes29 _current = _currentBest.ref(0).tryAsHeader();
        require(
            _new.notNull() && _current.notNull(),
            "Bad args. Check header and array byte lengths."
        );
        return _markNewHeaviest(_ancestor, _current, _new, _limit);
    }

    /// @notice             Adds headers to storage after validating
    /// @dev                We check integrity and consistency of the header chain
    /// @param  _anchor     The header immediately preceeding the new chain
    /// @param  _headers    A tightly-packed list of new 80-byte Bitcoin headers to record
    /// @return             True if successfully written, error otherwise
    function _addHeaders(bytes29 _anchor, bytes29 _headers) internal returns (bool) {
        /// Extract basic info
        bytes32 _previousDigest = _anchor.hash256();
        uint256 _anchorHeight = _findHeight(_previousDigest);  /* NB: errors if unknown */
        uint256 _previousTime = _anchor.time();

        /*
        NB:
        1. check that the header has sufficient work
        2. check that headers are in a coherent chain (no retargets, hash links good)
        3. Store the block connection
        4. Store the height
        */
        uint256 _height;
        bytes32 _currentDigest;
        uint256 _target;
        uint256 _expectedTarget;
        for (uint256 i = 0; i < _headers.len() / 80; i += 1) {
            bytes29 _header = _headers.indexHeaderArray(i);
            _target = _header.target();
            _height = _anchorHeight.add(i + 1);
            _currentDigest = _header.hash256();
            
            /*
            NB:
            if the block is already authenticated, we don't need to a work check
            Or write anything to state. This saves gas
            */
            if (previousBlock[_currentDigest] == bytes32(0)) {
                require(
                    TypedMemView.reverseUint256(uint256(_currentDigest)) <= _target,
                    "Header work is insufficient"
                );
                previousBlock[_currentDigest] = _previousDigest;
                if (_height % HEIGHT_INTERVAL == 0) {
                    /*
                    NB: We store the height only every 4th header to save gas
                    */
                    blockHeight[_currentDigest] = _height;
                }
            }

            _expectedTarget = _nextTarget(int(anchorHeight),
                                          int(anchorParentTime),
                                          anchorNBits,
                                          int(_height.sub(1)),
                                          int(_previousTime));
            require((_target & _expectedTarget) == _target,
                    "Invalid retarget provided");

            /* NB: we do still need to make chain level checks tho */
            require(_header.checkParent(_previousDigest), "Headers do not form a consistent chain");

            _previousDigest = _currentDigest;
            _previousTime = _header.time();
        }

        emit Extension(
            _anchor.hash256(),
            _currentDigest);
        return true;
    }

    /// @notice         Finds the height of a header by its digest
    /// @dev            Will fail if the header is unknown
    /// @param _digest  The header digest to search for
    /// @return         The height of the header
    function _findHeight(bytes32 _digest) internal view returns (uint256) {
        uint256 _height = 0;
        bytes32 _current = _digest;
        for (uint256 i = 0; i < HEIGHT_INTERVAL + 1; i = i.add(1)) {
            _height = blockHeight[_current];
            if (_height == 0) {
                _current = previousBlock[_current];
            } else {
                return _height.add(i);
            }
        }
        revert("Unknown block");
    }

    /// @notice         Finds an ancestor for a block by its digest
    /// @dev            Will fail if the header is unknown
    /// @param _digest  The header digest to search for
    /// @return         The height of the header, or error if unknown
    function _findAncestor(bytes32 _digest, uint256 _offset) internal view returns (bytes32) {
        bytes32 _current = _digest;
        for (uint256 i = 0; i < _offset; i = i.add(1)) {
            _current = previousBlock[_current];
        }
        require(_current != bytes32(0), "Unknown ancestor");
        return _current;
    }

    /// @notice             Checks if a digest is an ancestor of the current one
    /// @dev                Limit the amount of lookups (and thus gas usage) with _limit
    /// @param _ancestor    The prospective ancestor
    /// @param _descendant  The descendant to check
    /// @param _limit       The maximum number of blocks to check
    /// @return             true if ancestor is at most limit blocks lower than descendant, otherwise false
    function _isAncestor(bytes32 _ancestor, bytes32 _descendant, uint256 _limit) internal view returns (bool) {
        bytes32 _current = _descendant;
        /* NB: 200 gas/read, so gas is capped at ~200 * limit */
        for (uint256 i = 0; i < _limit; i = i.add(1)) {
            if (_current == _ancestor) {
                return true;
            }
            _current = previousBlock[_current];
        }
        return false;
    }

    /// @notice                   Marks the new best-known chain tip
    /// @param  _ancestor         The digest of the most recent common ancestor
    /// @param  _current          The 80-byte header referenced by bestKnownDigest
    /// @param  _new              The 80-byte header to mark as the new best
    /// @param  _limit            Limit the amount of traversal of the chain
    /// @return                   True if successfully updates bestKnownDigest, error otherwise
    function _markNewHeaviest(
        bytes32 _ancestor,
        bytes29 _current,  // Header
        bytes29 _new,      // Header
        uint256 _limit
    ) internal returns (bool) {
        require(_limit <= 2016, "Requested limit is greater than 1 difficulty period");

        bytes32 _newBestDigest = _new.hash256();
        bytes32 _currentBestDigest = _current.hash256();
        require(_currentBestDigest == bestKnownDigest, "Passed in best is not best known");
        require(
            previousBlock[_newBestDigest] != bytes32(0),
            "New best is unknown");
        require(
            _isMostRecentAncestor(_ancestor, bestKnownDigest, _newBestDigest, _limit),
            "Ancestor must be heaviest common ancestor");
        require(
            _heaviestFromAncestor(_ancestor, _current, _new) == _newBestDigest,
            "New best hash does not have more work than previous");

        bestKnownDigest = _newBestDigest;
        lastReorgCommonAncestor = _ancestor;

        uint256 _newDiff = _new.diff();
        /* if (_newDiff != currentEpochDiff) { */
        /*     currentEpochDiff = _newDiff; */
        /* } */

        emit NewTip(
            _currentBestDigest,
            _newBestDigest,
            _ancestor);
        return true;
    }

    /// @notice             Checks if a digest is an ancestor of the current one
    /// @dev                Limit the amount of lookups (and thus gas usage) with _limit
    /// @param _ancestor    The prospective shared ancestor
    /// @param _left        A chain tip
    /// @param _right       A chain tip
    /// @param _limit       The maximum number of blocks to check
    /// @return             true if it is the most recent common ancestor within _limit, false otherwise
    function _isMostRecentAncestor(
        bytes32 _ancestor,
        bytes32 _left,
        bytes32 _right,
        uint256 _limit
    ) internal view returns (bool) {
        /* NB: sure why not */
        if (_ancestor == _left && _ancestor == _right) {
            return true;
        }

        bytes32 _leftCurrent = _left;
        bytes32 _rightCurrent = _right;
        bytes32 _leftPrev = _left;
        bytes32 _rightPrev = _right;

        for(uint256 i = 0; i < _limit; i = i.add(1)) {
            if (_leftPrev != _ancestor) {
                _leftCurrent = _leftPrev;  // cheap
                _leftPrev = previousBlock[_leftPrev];  // expensive
            }
            if (_rightPrev != _ancestor) {
                _rightCurrent = _rightPrev;  // cheap
                _rightPrev = previousBlock[_rightPrev];  // expensive
            }
        }
        if (_leftCurrent == _rightCurrent) {return false;} /* NB: If the same, they're a nearer ancestor */
        if (_leftPrev != _rightPrev) {return false;} /* NB: Both must be ancestor */
        return true;
    }

    /// @notice             Decides which header is heaviest from the ancestor
    /// @dev                Does not support reorgs above 2017 blocks (:
    /// @param _ancestor    The prospective shared ancestor
    /// @param _left        A chain tip
    /// @param _right       A chain tip
    /// @return             true if it is the most recent common ancestor within _limit, false otherwise
    function _heaviestFromAncestor(
        bytes32 _ancestor,
        bytes29 _left,
        bytes29 _right
    ) internal view returns (bytes32) {
        uint256 _ancestorHeight = _findHeight(_ancestor);
        uint256 _leftHeight = _findHeight(_left.hash256());
        uint256 _rightHeight = _findHeight(_right.hash256());

        require(
            _leftHeight >= _ancestorHeight && _rightHeight >= _ancestorHeight,
            "A descendant height is below the ancestor height");

        /* NB: we can shortcut if one block is in a new difficulty window and the other isn't */
        uint256 _nextPeriodStartHeight = _ancestorHeight.add(2016).sub(_ancestorHeight % 2016);
        bool _leftInPeriod = _leftHeight < _nextPeriodStartHeight;
        bool _rightInPeriod = _rightHeight < _nextPeriodStartHeight;

        /*
        NB:
        1. Left is in a new window, right is in the old window. Left is heavier
        2. Right is in a new window, left is in the old window. Right is heavier
        3. Both are in the same window, choose the higher one
        4. They're in different new windows. Choose the heavier one
        */
        if (!_leftInPeriod && _rightInPeriod) {return _left.hash256();}
        if (_leftInPeriod && !_rightInPeriod) {return _right.hash256();}
        if (_leftInPeriod && _rightInPeriod) {
            return _leftHeight >= _rightHeight ? _left.hash256() : _right.hash256();
        } else {  // if (!_leftInPeriod && !_rightInPeriod) {
            if (((_leftHeight % 2016).mul(_left.diff())) <
                (_rightHeight % 2016).mul(_right.diff())) {
                return _right.hash256();
            } else {
                return _left.hash256();
            }
        }
    }

    function _bitsToTarget(uint256 nBits) internal pure returns (uint256) {
      uint256 target;
      
      uint256 exponent = nBits >> 24;
      uint256 mantissa = nBits & 0x007fffff;
      
      if (exponent <= 3) {
        target = mantissa >> (8 * (3 - exponent));
      } else {
        target = mantissa << (8 * (exponent - 3));
      }
      return target;
    }
    
    function _nextTarget(int256 anchorHeigth,
                        int256 anchorParentTime,
                        uint256 anchorNBits,
                        int256 currentHeigth,
                        int256 currentTime
                        ) internal pure returns (uint256) {
      
      uint256 maxTarget = _bitsToTarget(MAX_BITS);
      uint256 anchorTarget = _bitsToTarget(anchorNBits);
      
      int256 timeDelta = currentTime - anchorParentTime;
      int256 heigthDelta = currentHeigth - anchorHeigth;
      
      int256 exponent =
        ((timeDelta - IDEAL_BLOCK_TIME * (heigthDelta + 1)) * RADIX) / HALFLIFE;
      int256 numShifts = exponent >> 16;
      exponent = exponent - numShifts * RADIX;
      
      // I think factor can be negative...
      int256 factor =
        ((195766423245049 * exponent + 971821376 * exponent * exponent +
          5127 * exponent * exponent * exponent + 2 ** 47) >> 48) + RADIX;
      
      // if factor is negative, is this cast a problem?
      uint256 nextTarget = anchorTarget * uint256(factor);
      
      if (numShifts < 0) {
        nextTarget = nextTarget >> (-numShifts);
      } else {
        nextTarget = nextTarget << numShifts;
      }
      
      nextTarget = nextTarget >> 16;
      
      if (nextTarget == 0) {
        return 1;
      }
      
      if (nextTarget > maxTarget) {
        return maxTarget;
      }
      
      return nextTarget;
    }
}

// For unittests
contract TestRelay is Relay {

    /// @notice                   Gives a starting point for the relay
    /// @dev                      We don't check this AT ALL really. Don't use relays with bad genesis
    /// @param  _genesisHeader    The starting header
    /// @param  _height           The starting height
    constructor(bytes memory _genesisHeader,
                uint256 _height,
                uint256 _anchorHeight,
                uint256 _anchorParentTime,
                uint256 _anchorNBits)
      Relay(_genesisHeader, _height, _anchorHeight, _anchorParentTime, _anchorNBits)
    public {}

    function heaviestFromAncestor(
        bytes32 _ancestor,
        bytes calldata _left,
        bytes calldata _right
    ) external view returns (bytes32) {
        return _heaviestFromAncestor(
            _ancestor,
            _left.ref(0).tryAsHeader(),
            _right.ref(0).tryAsHeader()
        );
    }

    function isMostRecentAncestor(
        bytes32 _ancestor,
        bytes32 _left,
        bytes32 _right,
        uint256 _limit
    ) external view returns (bool) {
        return _isMostRecentAncestor(_ancestor, _left, _right, _limit);
    }
}
