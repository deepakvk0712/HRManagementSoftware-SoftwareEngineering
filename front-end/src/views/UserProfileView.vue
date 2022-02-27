<template>
  <div class="home">
    <v-container>

      <NavigationBar/>
      <v-layout class="d-flex">
        <v-row>
        <v-col>
          <v-card elevation="4">
            <v-card-title class="justify-center">
              <v-avatar
                size="100px"
                
              >
                <img alt = "Avatar" src="../assets/logo.png"/>

              </v-avatar>
            </v-card-title>
            <v-card-subtitle class="pb-0 justify-center font-weight-black">
              Ashley Harrington
            </v-card-subtitle>

            <v-card-subtitle class="py-0 justify-center font-weight-black">
              Project Manager
            </v-card-subtitle>

            <v-divider></v-divider>

            <v-card-subtitle class="text-left font-weight-black">
              About Me
            </v-card-subtitle>
            <v-card-text class="text-left">
              This is Ashely Harrington. I am an avid snowboarder and I love to go on long hikes with my dogs. I enjoy the mountains more than the beach.
            </v-card-text>

          </v-card>
        </v-col>

        <v-col>
          <v-card>
            <v-card-title class="align-center">
              <v-avatar
                size="36px"
              >
                <img alt = "Avatar" src="../assets/logo.png"/>

              </v-avatar>
            </v-card-title>
          </v-card>
        </v-col>
        </v-row>
      </v-layout>

    </v-container>
  </div>
</template>

<script>
import NavigationBar from '../components/NavigationBar.vue'
export default {
  name: 'HomeView',
  components: {
    NavigationBar,
  }
}
</script>
