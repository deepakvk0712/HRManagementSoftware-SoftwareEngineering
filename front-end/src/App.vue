<template>
  <v-app>
    <v-content class="ma-4" v-show="this.$store.state.loader == false">
      <router-view></router-view>
    </v-content>
    <!-- <v-content class="ma-4">
      <PageLoader v-show="this.$store.state.loader"/>
    </v-content> -->
  </v-app>
  
</template>

<script>
// import PageLoader from '../src/components/PageLoader.vue'
export default {
  components: {
    // PageLoader,
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
