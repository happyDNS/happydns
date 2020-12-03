// Copyright or © or Copr. happyDNS (2020)
//
// contact@happydns.org
//
// This software is a computer program whose purpose is to provide a modern
// interface to interact with DNS systems.
//
// This software is governed by the CeCILL license under French law and abiding
// by the rules of distribution of free software.  You can use, modify and/or
// redistribute the software under the terms of the CeCILL license as
// circulated by CEA, CNRS and INRIA at the following URL
// "http://www.cecill.info".
//
// As a counterpart to the access to the source code and rights to copy, modify
// and redistribute granted by the license, users are provided only with a
// limited warranty and the software's author, the holder of the economic
// rights, and the successive licensors have only limited liability.
//
// In this respect, the user's attention is drawn to the risks associated with
// loading, using, modifying and/or developing or reproducing the software by
// the user in light of its specific status of free software, that may mean
// that it is complicated to manipulate, and that also therefore means that it
// is reserved for developers and experienced professionals having in-depth
// computer knowledge. Users are therefore encouraged to load and test the
// software's suitability as regards their requirements in conditions enabling
// the security of their systems and/or data to be ensured and, more generally,
// to use and operate it in the same conditions as regards security.
//
// The fact that you are presently reading this means that you have had
// knowledge of the CeCILL license and that you accept its terms.

import Vue from 'vue'
import UserApi from '@/api/user'

export default {
  namespaced: true,

  state: {
    session: null
  },

  getters: {
    user_getSession: state => state.session,
    user_getSettings: state => state.session ? state.session.settings : null,
    user_isLogged: state => state.session != null
  },

  actions: {
    login ({ commit }, { email, password }) {
      return new Promise((resolve, reject) => {
        UserApi.login(email, password)
          .then(
            (response) => {
              commit('auth_success', response.data)
              resolve()
            },
            (error) => {
              reject(error)
            }
          )
      })
    },

    logout ({ commit }) {
      return new Promise((resolve, reject) => {
        UserApi.logout()
          .then(
            (response) => {
              commit('logout')
              resolve()
            },
            (error) => {
              reject(error)
            }
          )
      })
    },

    retrieveSession ({ commit }) {
      commit('auth_backup')
      return this.dispatch('user/updateSession')
    },

    updateSession ({ commit }) {
      return new Promise((resolve, reject) => {
        UserApi.getSession()
          .then(
            (response) => {
              commit('auth_success', response.data)
              resolve()
            },
            (error) => {
              if (error.response) {
                commit('logout')
                reject(error)
              }
              resolve()
            }
          )
      })
    },

    updateSettings ({ commit }) {
      return this.dispatch('user/updateSession')
    }
  },

  mutations: {
    auth_backup (state) {
      if (sessionStorage.loggedUser) {
        try {
          Vue.set(state, 'session', JSON.parse(sessionStorage.loggedUser))
        } catch {
          Vue.set(state, 'session', null)
        }
      }
    },

    auth_success (state, session) {
      Vue.set(state, 'session', session)
      sessionStorage.loggedUser = JSON.stringify(session)
    },

    logout (state) {
      Vue.set(state, 'session', null)
    }
  }
}
