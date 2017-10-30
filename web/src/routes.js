import AddAuthor from './components/AddAuthor.vue'
import EditAuthor from './components/EditAuthor.vue'
import ListAuthors from './components/ListAuthors.vue'
import AddBook from './components/AddBook.vue'
import EditBook from './components/EditBook.vue'
import ListBooks from './components/ListBooks.vue'

export const routes = [
  { name: 'home', path: '/', redirect: { name: 'listAuthors' } },
  { name: 'listAuthors', path: '/authors', component: ListAuthors },
  { name: 'addAuthor', path: '/authors/add', component: AddAuthor },
  { name: 'editAuthor', path: '/authors/edit/:id', component: EditAuthor },
  { name: 'listBooks', path: '/books', component: ListBooks },
  { name: 'addBook', path: '/books/add', component: AddBook },
  { name: 'editBook', path: '/books/edit/:id', component: EditBook },
]
