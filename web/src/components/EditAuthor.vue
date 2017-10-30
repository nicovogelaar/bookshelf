<template>
  <div>
    <h1>Edit Author</h1>
    <form v-on:submit.prevent="onSubmit">
      <div class="form-group">
        <label for="author-name">Name</label>
        <input type="text" id="author-name" v-model="name" class="form-control" />
      </div>
      <div class="form-group">
        <label for="author-biography">Biography</label>
        <textarea id="author-biography" v-model="biography" class="form-control"></textarea>
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
      this.getAuthor(id)
    }
  },
  created() {
    this.getAuthor(this.$route.params.id)
  },
  computed: {
    name: {
      get() {
        return this.$store.getters.author.name
      },
      set(name) {
        this.$store.commit(types.UPDATE_AUTHOR, { data: { name } })
      },
    },
    biography: {
      get() {
        return this.$store.getters.author.biography
      },
      set(biography) {
        this.$store.commit(types.UPDATE_AUTHOR, { data: { biography } })
      },
    },
  },
  methods: {
    getAuthor(id) {
      this.$store.dispatch('getAuthor', id)
    },
    onSubmit() {
      let id = this.$route.params.id
      let data = this.$store.getters.author
      this.$store.dispatch('updateAuthor', { id, data })
    }
  }
}
</script>