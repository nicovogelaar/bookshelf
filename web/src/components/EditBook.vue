<template>
  <div>
    <h1>Edit Book</h1>
    <form v-on:submit.prevent="onSubmit">
      <div class="form-group">
        <label for="book-title">Title</label>
        <input type="text" id="book-title" v-model="title" class="form-control" />
      </div>
      <div class="form-group">
        <label for="book-isbn">ISBN</label>
        <input type="text" id="book-isbn" v-model="isbn" class="form-control" />
      </div>
      <div class="form-group">
        <label for="book-description">Description</label>
        <textarea id="book-description" v-model="description" class="form-control"></textarea>
      </div>
      <div class="form-group">
        <label for="book-author">Author</label>
        <select class="form-control" v-model="authorId">
          <option v-for="(author, index) in authors.authors" :key="index" :value="author.id">
            {{ author.name }}
          </option>
        </select>
      </div>
      <button type="submit" class="btn btn-primary">Update</button>
    </form>
  </div>
</template>

<script>
import types from '../store/mutation-types'

export default {
  watch: {
    '$route.params.id': function (id) {
      this.getBook(id)
    }
  },
  created() {
    this.getBook(this.$route.params.id)
    this.getAllAuthors()
  },
  computed: {
    title: {
      get() {
        return this.$store.getters.book.title
      },
      set(title) {
        this.$store.commit(types.UPDATE_BOOK, { data: { title } })
      },
    },
    isbn: {
      get() {
        return this.$store.getters.book.isbn
      },
      set(isbn) {
        this.$store.commit(types.UPDATE_BOOK, { data: { isbn } })
      },
    },
    description: {
      get() {
        return this.$store.getters.book.description
      },
      set(description) {
        this.$store.commit(types.UPDATE_BOOK, { data: { description } })
      },
    },
    authorId: {
      get() {
        return this.$store.getters.book.authorId
      },
      set(authorId) {
        this.$store.commit(types.UPDATE_BOOK, { data: { authorId } })
      },
    },
    authors() {
      return this.$store.getters.allAuthors
    },
  },
  methods: {
    getAllAuthors() {
      this.$store.dispatch('getAllAuthors', {})
    },
    getBook(id) {
      this.$store.dispatch('getBook', id)
    },
    onSubmit() {
      let id = this.$route.params.id
      let data = this.$store.getters.book
      this.$store.dispatch('updateBook', { id, data })
    }
  }
}
</script>