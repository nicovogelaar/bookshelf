<template>
	<nav class="navbar navbar-expand-md navbar-dark bg-dark fixed-top">
    <router-link :to="{ name: 'home' }" class="navbar-brand">Bookshelf</router-link>
    <button class="navbar-toggler"
            type="button"
            data-toggle="collapse"
            data-target="#navbar-default"
            aria-controls="navbar-default"
            aria-expanded="false"
            aria-label="Toggle navigation">
      <span class="navbar-toggler-icon"></span>
    </button>

    <div class="collapse navbar-collapse" id="navbar-default">
      <ul class="navbar-nav mr-auto">
        <li class="nav-item"><router-link to="/authors" class="nav-link">Authors</router-link></li>
        <li class="nav-item"><router-link to="/books" class="nav-link">Books</router-link></li>
      </ul>
    </div>
  </nav>
</template>