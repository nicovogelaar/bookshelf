<template>
	<div>
    <h1>Authors</h1>
    <router-link to="/authors/add" class="btn btn-success">Add Author</router-link>
    <br /><br />
    <table class="table table-bordered">
      <thead>
        <th>ID</th>
        <th>Name</th>
        <th>Actions</th>
      </thead>
      <tr v-for="(author, index) in authors.authors" :key="index">
        <td>{{ author.id }}</td>
        <td>{{ author.name }}</td>
        <td>
          <router-link :to="{ name: 'editAuthor', params: { id: author.id }}" class="btn btn-warning">Edit</router-link>
          <a href="#" @click.prevent="deleteAuthor(author.id)" class="btn btn-danger">Delete</a>
        </td>
      </tr>
    </table>
	</div>
</template>

<script>
export default {
  created() {
    this.getAllAuthors()
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
    deleteAuthor(id) {
      this.$store.dispatch('deleteAuthor', id).then(() => {
        this.getAllAuthors()
      })
    }
  }
}
</script>