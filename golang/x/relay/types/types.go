package types

import (
	"encoding/hex"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/summa-tx/bitcoin-spv/golang/btcspv"
	"github.com/summa-tx/relays/proto"
)

// ProofHandler is an interface to which the keeper dispatches valid proofs
type ProofHandler interface {
	HandleValidProof(ctx sdk.Context, filled FilledRequests, requests []ProofRequest)
}

// Hash256Digest 32-byte double-sha2 digest
type Hash256Digest = btcspv.Hash256Digest

// Hash160Digest is a 20-byte ripemd160+sha2 hash
type Hash160Digest = btcspv.Hash160Digest

// RawHeader is an 80-byte raw header
type RawHeader = btcspv.RawHeader

// HexBytes is a type alias to make JSON hex ser/deser easier
type HexBytes = btcspv.HexBytes

// BitcoinHeader is a parsed Bitcoin header
type BitcoinHeader = btcspv.BitcoinHeader

// SPVProof is the base struct for an SPV proof
type SPVProof = btcspv.SPVProof

// Origin is an enum of types denoting requests either from the local chain
// or a remote chain
type Origin int

// Origin possible types
const (
	Local  Origin = 0
	Remote Origin = 1
)

// Hash256DigestFromHex converts a hex into a Hash256Digest
func BytesFromHex(hexStr string) ([]byte, *sdkerrors.Error) {
	data := hexStr
	if data[:2] == "0x" {
		data = data[2:]
	}

	bytes, decodeErr := hex.DecodeString(data)
	if decodeErr != nil {
		return []byte{}, ErrBadHex(DefaultCodespace, hexStr)
	}

	return bytes, nil
}

// Hash256DigestFromHex converts a hex into a Hash256Digest
func Hash256DigestFromHex(hexStr string) (btcspv.Hash256Digest, *sdkerrors.Error) {
	data := hexStr
	if data[:2] == "0x" {
		data = data[2:]
	}

	bytes, decodeErr := hex.DecodeString(data)
	if decodeErr != nil {
		return btcspv.Hash256Digest{}, ErrBadHex(DefaultCodespace, hexStr)
	}

	digest, newDigestErr := btcspv.NewHash256Digest(bytes)
	if newDigestErr != nil {
		return btcspv.Hash256Digest{}, FromBTCSPVError(DefaultCodespace, newDigestErr)
	}
	return digest, nil
}

// NullHandler does nothing
type NullHandler struct{}

// HandleValidProof handles a valid proof (by doing nothing)
func (n NullHandler) HandleValidProof(ctx sdk.Context, filled FilledRequests, requests []ProofRequest) {
}

// NewNullHandler instantiates a new null handler
func NewNullHandler() NullHandler {
	return NullHandler{}
}

func stringToError(str string) (*sdkerrors.Error) {
	// TODO: How to handle CodeType?
	err := sdkerrors.New(DefaultCodespace, ExternalError, str)
	return err
}

func bufToH256(buf []byte) (btcspv.Hash256Digest, error) {
	var h btcspv.Hash256Digest;
	if len(buf) != 32 {
		return h, fmt.Errorf("Expected 32 bytes, got %d bytes", len(buf))
	}

	copy(h[:], buf)

	return h, nil
}

func bufToRawHeader(buf []byte) (btcspv.RawHeader, error) {
	var h btcspv.RawHeader;
	if len(buf) != 80 {
		return h, fmt.Errorf("Expected 80 bytes, got %d bytes", len(buf))
	}

	copy(h[:], buf)

	return h, nil
}

func bufToRequestID(buf []byte) (RequestID, error) {
	var h RequestID;
	if len(buf) != 8 {
		return h, fmt.Errorf("Expected 8 bytes, got %d bytes", len(buf))
	}

	copy(h[:], buf)

	return h, nil
}


func headerFromProto(m *proto.BitcoinHeader) (btcspv.BitcoinHeader, error) {
	var header btcspv.BitcoinHeader

	raw, err := bufToRawHeader(m.Raw)
	if err != nil {
		return header, err
	}

	hash, err := bufToH256(m.Hash)
	if err != nil {
		return header, err
	}

	root, err := bufToH256(m.MerkleRoot)
	if err != nil {
		return header, err
	}


	header.Raw = raw
	header.Hash = hash
	header.Height = m.Height
	header.MerkleRoot = root

	return header, nil
}

func headerToProto(m *btcspv.BitcoinHeader) (proto.BitcoinHeader) {
	var header proto.BitcoinHeader

	header.Raw = []byte{m.Raw}
	header.Hash = []byte{m.Hash}
	header.Height = m.Height
	header.MerkleRoot = []byte{m.Root}

	return header
}

func headerSliceFromProto(m []*proto.BitcoinHeader) ([]btcspv.BitcoinHeader, error) {
	headers := make([]btcspv.BitcoinHeader, len(m))
	for	i, h := range m {
		header, err := headerFromProto(h)
		headers[i] = header
		if err != nil {
			return nil, err
		}
	}

	return headers, nil
}

func headersSliceToProto(m []*btcspv.BitcoinHeader) ([]proto.BitcoinHeader) {
	headers := make([]proto.BitcoinHeader, len(m))
	for i, h := range m {
		header := headerToProto(h)
		headers[i] = header
	}

	return headers
}

func spvProofFromProto(m *proto.SPVProof) (btcspv.SPVProof, error) {
	var spvProof btcspv.SPVProof

	txID, err := bufToH256(m.TxID)
	if err != nil {
		return spvProof, err
	}

	header, err := headerFromProto(m.ConfirmingHeader)
	if err != nil {
		return spvProof, err
	}

	spvProof.Version = HexBytes(m.Version)
	spvProof.Vin = HexBytes(m.Vin)
	spvProof.Vout = HexBytes(m.Vout)
	spvProof.Locktime = HexBytes(m.Locktime)
	spvProof.TxID = txID
	spvProof.Index = m.Index
	spvProof.ConfirmingHeader = header
	spvProof.IntermediateNodes = HexBytes(m.IntermediateNodes)

	return spvProof, nil
}

func spvProofToProto(m *btcspv.SPVProof) (proto.SPVProof) {
	spvProof.Version = []byte{m.Version}
	spvProof.Vin = []byte{m.Vin}
	spvProof.Vout = []byte{m.Vout}
	spvProof.Locktime = []byte{m.Locktime}
	spvProof.TxID = []byte{m.TxID}
	spvProof.Index = m.Index
	spvProof.ConfirmingHeader = headerToProto(m.header)
	spvProof.IntermediateNodes = []byte{m.IntermediateNodes}

	return spvProof
}

func proofRequestFromProto(m *proto.ProofRequest) (ProofRequest, error) {
	var proofRequest ProofRequest

	spends, err := bufToH256(m.Spends)
	if err != nil {
		return proofRequest, err
	}

	pays, err := bufToH256(m.Pays)
	if err != nil {
		return proofRequest, err
	}

	proofRequest.Spends = spends
	proofRequest.Pays = pays
	proofRequest.PaysValue = m.PaysValue
	proofRequest.ActiveState = m.ActiveState
	proofRequest.NumConfs = uint8(m.NumConfs)
	proofRequest.Origin = Origin(m.Origin)
	proofRequest.Action = btcspv.HexBytes(m.Action)

	return proofRequest, nil
}

func proofRequestToProto(m *ProofRequest) (proto.ProofRequest) {
	var proofRequest proto.ProofRequest

	proofRequest.Spends = []byte{m.Spends}
	proofRequest.Pays = []byte{m.Pays}
	proofRequest.PaysValue = m.PaysValue
	proofRequest.ActiveState = m.ActiveState
	proofRequest.NumConfs = uint32(m.NumConfs)
	proofRequest.Origin = []byte{m.Origin}
	proofRequest.Action = []byte{m.Action}

	return proofRequest
}