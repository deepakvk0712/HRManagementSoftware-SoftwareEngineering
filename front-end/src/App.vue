<template>
  <div id = "app">
    <NavigationBar 
      :tab-info = "tabInfo"
      :image-path = "require('./assets/logo.png')"
    />
    <router-view/>
  </div>
</template>

<script>
import NavigationBar from "./components/NavigationBar.vue"

export default {
  name: 'App',
  components: {
    NavigationBar
  },
  data: () => ({
    tabInfo: [
      {
        text: 'Home',
        path: '/',
        icon: 'ion-ios-home'
      },
      {
        text: 'My Information',
        path: '/myinfo',
        icon: 'ion-ios-business'
      },
      {
        text: 'Apply Leave',
        path: '/leave',
        icon: 'ion-ios-medkit-outline'
      },
      {
        text: 'Pay Slips',
        path: '/payslip',
        icon: 'ion-ios-cash'
      }
    ]
  })
}
</script>

<style lang="scss">
@import 'https://unpkg.com/ionicons@4.2.2/dist/css/ionicons.min.css';

figure {
	margin-block-start: 0;
	margin-block-end: 0;
	margin-inline-start: 10px;
	margin-inline-end: 0;
}

body {
  margin: 0;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  /* margin-top: 60px; */
}
</style>
