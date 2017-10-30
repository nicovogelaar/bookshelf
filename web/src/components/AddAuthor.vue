<template>
  <div>
    <h1>Add Author</h1>
    <form v-on:submit.prevent="onSubmit">
      <div class="form-group">
        <label for="author-name">Name</label>
        <input type="text" id="author-name" v-model="name" class="form-control" />
      </div>
      <div class="form-group">
        <label for="author-biography">Biography</label>
        <textarea id="author-biography" v-model="biography" class="form-control"></textarea>
      </div>
      <button type="submit" class="btn btn-primary">Create</button>
    </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      name: '',
      biography: '',
    }
  },
  methods: {
    onSubmit() {
      this.$store.dispatch('createAuthor', this.$data).then(author => {
        this.$router.push('/authors/edit/' + author.id)
      })
    }
  }
}
</script>