<template>
  <div>
    <h1>Add Book</h1>
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
      <button type="submit" class="btn btn-primary">Create</button>
    </form>
  </div>
</template>

<script>
export default {
  created() {
    this.getAllAuthors()
  },
  data() {
    return {
      title: '',
      isbn: '',
      description: '',
      authorId: ''
    }
  },
  computed: {
    authors() {
      return this.$store.getters.allAuthors
    }
  },
  methods: {
    getAllAuthors() {
      this.$store.dispatch('getAllAuthors', {})
    },
    onSubmit() {
      this.$store.dispatch('createBook', this.$data).then(book => {
        this.$router.push('/books/edit/' + book.id)
      })
    }
  }
}
</script>