import axios from 'axios'
import * as types from '@/store/mutation-types'
import { reverseEndianness, convertUnixTimestamp } from '@/utils/utils'
const relayURL = '/relay'

const state = {
  connected: true
}

const mutations = {
  [types.SET_CONNECTED] (state, connected) {
    state.connected = connected
  }
}

const actions = {
  getBKD ({ commit, dispatch }) {
    axios.get(`${relayURL}/getbestdigest`)
      .then((res) => {
        commit(types.SET_CONNECTED, true)

        const hashBE = reverseEndianness(res.data.result.result)
        console.log('get BKD: ', hashBE)

        dispatch('info/setBKD', { hash: hashBE }, { root: true })
        dispatch('verifyHash', { hash: hashBE, type: 'BKD' })
      })
      .catch((e) => {
        console.error('relay/getBKD:\n', e)
        if (e.message === 'Request failed with status code 500') {
          commit(types.SET_CONNECTED, false)
        }
      })
  },

  getLCA ({ commit, dispatch }) {
    axios.get(`${relayURL}/getlastreorglca`)
      .then((res) => {
        commit(types.SET_CONNECTED, true)

        const hashBE = reverseEndianness(res.data.result.result)
        console.log('get LCA: ', hashBE)

        dispatch('info/setLCA', { hash: hashBE }, { root: true })
        dispatch('verifyHash', { hash: hashBE, type: 'LCA'})
      })
      .catch((e) => {
        console.error('relay/getLCA:\n', e)
        if (e.message === 'Request failed with status code 500') {
          commit(types.SET_CONNECTED, false)
        }
      })
  },

  verifyHash ({ rootState, dispatch }, data) {
    // data.hash, data.type = 'BKD', 'LCA'
    console.log({ data })
    axios.get(`${rootState.blockchainURL}/block/${data.hash}`)
      .then((block) => {
        console.log('block', block)
        dispatch(
          `info/set${data.type}`,
          {
            height: block.data.height,
            time: convertUnixTimestamp(block.data.timestamp),
            updatedAt: new Date()
          },
          { root: true }
        )
        dispatch(
          'info/setLastComms',
          { source: 'relay', date: new Date() },
          { root: true }
        )
      }).catch((e) => {
        console.error('relay/verifyHash:\n', e)
      })
  },

  // NB: Verify height does not actually verify height. This is for updating only. See `info/getExternalInfo`
  verifyHeight ({ dispatch }, data) {
    console.log('verify height', data)
    if (data) {
      dispatch(
        'info/setCurrentBlock',
        { updatedAt: new Date() },
        { root: true }
      )
      dispatch(
        'info/setLastComms',
        { source: 'relay', date: new Date() },
        { root: true }
      )
    }
  }
}

// state,
// mutations,
export default {
  namespaced: true,
  state,
  mutations,
  actions
}
