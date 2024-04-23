<template>
  <HeaderComponent/>
  <TableComponent/>
</template>

<script>
import HeaderComponent from './components/Header.vue'
import TableComponent from './components/TableComponent.vue'

export default {
  name: 'App',
  components: {
    HeaderComponent,
    TableComponent
  }
}
</script>

<style>
body {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #fff;
  background: #232526;  /* fallback for old browsers */
  background: -webkit-linear-gradient(to right, #414345, #232526);  /* Chrome 10-25, Safari 5.1-6 */
  background: linear-gradient(to right, #414345, #232526); /* W3C, IE 10+/ Edge, Firefox 16+, Chrome 26+, Opera 12+, Safari 7+ */
  margin: 0;
  padding: 0;
}

</style>
