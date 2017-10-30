<template>
  <div>
    <h1>Books</h1>
    <router-link to="/books/add" class="btn btn-success">Add Book</router-link>
    <br /><br />
    <table class="table table-bordered">
      <thead>
        <th>ID</th>
        <th class="w-50">Title</th>
        <th>ISBN</th>
        <th>Author</th>
        <th>Actions</th>
      </thead>
      <tr v-for="(book, index) in books.books" :key="index">
        <td>{{ book.id }}</td>
        <td>{{ book.title }}</td>
        <td>{{ book.isbn }}</td>
        <td><router-link :to="{ name: 'editAuthor', params: { id: book.author.id }}">{{ book.author.name }}</router-link></td>
        <td>
          <router-link :to="{ name: 'editBook', params: { id: book.id }}" class="btn btn-warning">Edit</router-link>
          <a href="#" @click.prevent="deleteBook(book.id)" class="btn btn-danger">Delete</a>
        </td>
      </tr>
    </table>
  </div>
</template>

<script>
export default {
  created() {
    this.getAllBooks()
  },
  computed: {
    books() {
      return this.$store.getters.allBooks
    }
  },
  methods: {
    getAllBooks() {
      this.$store.dispatch('getAllBooks', {})
    },
    deleteBook(id) {
      this.$store.dispatch('deleteBook', id).then(() => {
        this.getAllBooks()
      })
    }
  }
}
</script>