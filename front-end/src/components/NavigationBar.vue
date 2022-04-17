<template>
  <nav>
    <v-app-bar color="#444" dark app>
      <v-app-bar-nav-icon @click.stop="isClicked = !isClicked">
      </v-app-bar-nav-icon>
      <v-spacer></v-spacer>
      <v-btn @click="signOut()"  text>
        <!-- to="/login" -->
        <span> Sign Out </span>
        <v-icon right> exit_to_app </v-icon>
      </v-btn>
    </v-app-bar>
    <v-navigation-drawer v-model="isClicked" dark app class="#444 darken-4">
      <v-layout column align-center>
        <v-flex class="mt-5">
          <v-avatar size="100">
            <img src="../assets/user.png" alt="" />
          </v-avatar>
          <p class="white--text subheading mt-1 text-center">{{userName}}</p>
          <v-divider color="white"></v-divider>
        </v-flex>
      </v-layout>
      <v-list flat>
        <v-list-item
          v-for="itm in sideBarItems"
          :key="itm.title"
          router
          :to="itm.path"
          active-class="border"
        >
          <v-list-item-action>
            <v-icon> {{ itm.icon }} </v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title>{{ itm.title }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
  </nav>
</template>

<script>
export default {
  props: {
    userName : {
      type: String,
    }
  },
  data: () => ({
    isClicked: false,
    sideBarItems: [
      { icon: "dashboard", title: "Landing Page", path: "/landing" },
      { icon: "people", title: "My Profile", path: "/user-profile" },
      { icon: "account_balance", title: "About", path: "/about" },
      { icon: "money", title: "Pay Slip", path: "/payslip" },
      { icon: "people", title: "Attendance", path: "/attendance" },
      { icon: "people", title: "Training", path: "/training" },
    ],
  }),

  methods : {
    signOut(){
      this.$axios.defaults.headers.common['Authorization'] = ""
      this.$router.push("/login")
    }
  }
};
</script>

<style scoped lang="scss">
.border {
  border-left: 3px solid red;
}
</style>
