import Vue from 'vue'
import books from '../../api/books'
import types from '../mutation-types'

const book = {
  title: '',
  isbn: '',
  description: '',
  authorId: ''
}

const state = {
  all: [],
  book: book
}

const getters = {
  allBooks: state => state.all,
  book: state => state.book
}

const actions = {
  getAllBooks ({ commit }, { limit, offset }) {
    books.getAllBooks(limit, offset).then(data => {
      Vue.$log.debug(types.LIST_BOOKS + ' with limit: ' + (limit || null) + ' and offset: ' + (offset || null))
      commit(types.LIST_BOOKS, { data })
    })
  },

  getBook ({ commit }, id) {
    books.getBook(id).then(data => {
      Vue.$log.debug(types.GET_BOOK + ' with id: ' + id)
      commit(types.GET_BOOK, { data })
    })
  },

  createBook (context, data) {
    return books.createBook(data).then(data => {
      Vue.$log.debug(types.CREATE_BOOK)
      return data
    })
  },

  updateBook ({ commit }, { id, data }) {
    return books.updateBook(id, data).then(data => {
      Vue.$log.debug(types.UPDATE_BOOK + ' with id: ' + id)
      commit(types.UPDATE_BOOK, { data })
      return data
    })
  },

  deleteBook (context, id) {
    return books.deleteBook(id).then(() => {
      Vue.$log.debug(types.DELETE_BOOK + ' with id: ' + id)
    })
  }
}

const mutations = {
  [types.LIST_BOOKS] (state, { data }) {
    state.all = data
  },

  [types.GET_BOOK] (state, { data }) {
    state.book = data
  },

  [types.UPDATE_BOOK] (state, { data }) {
    state.book = Object.assign(state.book, data)
  }
}

export default {
  state,
  getters,
  actions,
  mutations
}
