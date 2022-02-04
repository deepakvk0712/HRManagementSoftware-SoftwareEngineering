<template>
  <v-app>
    <NavigationBar/>
    <v-content class="ma-4">
      <router-view></router-view>
    </v-content>
    <PageBottom/>
  </v-app>
</template>

<script>
import NavigationBar from './components/NavigationBar.vue'
import PageBottom from './components/PageBottom.vue'

export default {
  components: {
    NavigationBar, PageBottom
  }
}
</script>


<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

nav {
  padding: 30px;

  a {
    font-weight: bold;
    color: #2c3e50;

    &.router-link-exact-active {
      color: #42b983;
    }
  }
}
</style>
