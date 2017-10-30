import Vue from 'vue'
import Vuex from 'vuex'
import authors from './modules/authors'
import books from './modules/books'

Vue.use(Vuex)

const debug = process.env.NODE_ENV !== 'production' // eslint-disable-line no-undef

export default new Vuex.Store({
  modules: {
    authors,
    books
  },
  strict: debug,
  plugins: []
})
