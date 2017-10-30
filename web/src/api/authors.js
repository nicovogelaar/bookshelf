import Vue from 'vue'

export default {
  getAllAuthors (limit, offset) {
    return Vue.http.get('authors', { params: { limit, offset } }).then(({ body }) => {
      return body
    })
  },

  getAuthor (id) {
    return Vue.http.get('authors/' + id).then(({ body }) => {
      return body
    })
  },

  createAuthor (data) {
    return Vue.http.post('authors', data).then(({ body }) => {
      return body
    })
  },

  updateAuthor (id, data) {
    return Vue.http.post('authors/' + id, data).then(({ body }) => {
      return body
    })
  },

  deleteAuthor (id) {
    return Vue.http.delete('authors/' + id)
  }
}
