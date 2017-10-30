import Vue from 'vue'
import authors from '../../api/authors'
import types from '../mutation-types'

const author = {
  name: '',
  biography: ''
}

const state = {
  all: [],
  author: author
}

const getters = {
  allAuthors: state => state.all,
  author: state => state.author
}

const actions = {
  getAllAuthors ({ commit }, { limit, offset }) {
    authors.getAllAuthors(limit, offset).then(data => {
      Vue.$log.debug(types.LIST_AUTHORS + ' with limit: ' + (limit || null) + ' and offset: ' + (offset || null))
      commit(types.LIST_AUTHORS, { data })
    })
  },

  getAuthor ({ commit }, id) {
    authors.getAuthor(id).then(data => {
      Vue.$log.debug(types.GET_AUTHOR + ' with id: ' + id)
      commit(types.GET_AUTHOR, { data })
    })
  },

  createAuthor (context, data) {
    return authors.createAuthor(data).then(data => {
      Vue.$log.debug(types.CREATE_AUTHOR)
      return data
    })
  },

  updateAuthor ({ commit }, { id, data }) {
    return authors.updateAuthor(id, data).then(data => {
      Vue.$log.debug(types.UPDATE_AUTHOR + ' with id: ' + id)
      commit(types.UPDATE_AUTHOR, { data })
      return data
    })
  },

  deleteAuthor (context, id) {
    return authors.deleteAuthor(id).then(() => {
      Vue.$log.debug(types.DELETE_AUTHOR + ' with id: ' + id)
    })
  }
}

const mutations = {
  [types.LIST_AUTHORS] (state, { data }) {
    state.all = data
  },

  [types.GET_AUTHOR] (state, { data }) {
    state.author = data
  },

  [types.UPDATE_AUTHOR] (state, { data }) {
    state.author = Object.assign(state.author, data)
  }
}

export default {
  state,
  getters,
  actions,
  mutations
}
