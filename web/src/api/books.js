import Vue from 'vue'

export default {
  getAllBooks (limit, offset) {
    return Vue.http.get('books', { params: { limit, offset } }).then(({ body }) => {
      return body
    })
  },

  getBook (id) {
    return Vue.http.get('books/' + id).then(({ body }) => {
      return body
    })
  },

  createBook (data) {
    return Vue.http.post('books', data).then(({ body }) => {
      return body
    })
  },

  updateBook (id, data) {
    return Vue.http.post('books/' + id, data).then(({ body }) => {
      return body
    })
  },

  deleteBook (id) {
    return Vue.http.delete('books/' + id)
  }
}
