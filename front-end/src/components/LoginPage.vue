<template>
    <v-app>
        <!-- <v-parallax dark src="https://cdn.vuetifyjs.com/images/backgrounds/vbanner.jpg" height ="100%" jumbotron> -->
            <!-- <v-row align="center" justify="center">
                <v-col class="text-center" cols="12"></v-col>
            </v-row> -->
            <v-content>
                <v-container fluid pa-0>

                <v-row align="center" justify="center" style="height:170vh" dense>
                    <v-card width = "600" class ="justify-center">
                        <v-card-title id="regLoginDisplay"> Login </v-card-title>
                        <v-card-text class="text-center">
                            <v-text-field id="userNameInput" label = "Username" v-model="userName" prepend-icon="mdi-account-circle"/>
                            <v-text-field id="userPasswordInput" label = "Password" prepend-icon="mdi-lock"
                            :append-icon="value ? 'visibility' : 'visibility_off'"
                            @click:append="() => (value = !value)"
                            :type="value ? 'text' : 'password'" v-model="password"/>
                            <!-- :type="showPassword ? 'text' : 'password'"
                            v-model="password"
                            prepend-icon="mdi-lock"
                            :append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
                            @click-append="showPassword = !showPassword"/> -->
                        </v-card-text>

                        <v-divider></v-divider>

                        <v-card-actions id="regSubmitButton">
                            <!-- <v-btn to="/signup" color = "#89CFF0" > SignUp </v-btn> -->
                            <v-btn  color = "#6495ED" @click="login"> Login </v-btn>
                            <!-- to="/landing" -->
                            
                        </v-card-actions>
                        <!-- <p class="forgot-password text-center">
                                <router-link to = "/ChangePassword">Change Password</router-link>
                        </p> -->
                    </v-card>
                </v-row>
                </v-container>
                <PageBottom/>
            </v-content>
        <!-- </v-parallax>  -->
    </v-app>
</template>

<script>
import PageBottom from './PageBottom.vue'

export default {
    components: {
    PageBottom,
  },
    name: "LoginPage",
    data: () => ({
      showPassword:false,
      userName : "",
      password : "",
      value : "",
    }),

    methods : {
        login() {
            this.$store.commit('LOADER', true)
            const requestObj = {
                "email" : this.userName,
                "password" : this.password
            }
            this.$axios.post("http://localhost:8080/login", requestObj)
                .then(response => {
                    let jsonData = JSON.parse(response.data.data)
                    console.log(jsonData)
                    let firstLogin = jsonData.firstLogin;
                    this.$store.state.accessToken = jsonData.accessToken
                    // if(firstLogin == true){
                    //     this.$store.commit('LOADER', false)
                    //     // this.$router.push('/ChangePassword')
                    // }
                    // else{
                    let token = jsonData.accessToken
                    this.$axios.defaults.headers.common['Authorization'] = "Bearer " + token
                    this.$store.commit('LOADER', false)
                    this.$router.push('/landing')
                    // }
                })
        },
    }
}
</script>

<style scoped>
.v-application {
  /* background-color: #0081a8; */
  /* border: 2px solid black;
  padding: 150px;
  margin: auto 10px;
  width: auto;
  height: 2000px; */
  /* background: #0081a8;
  background-size: cover;
  height: 100%;
  overflow: hidden;
  background-repeat: no-repeat; */

    /* width: 100%;
    height: 100%;
    position: absolute;
    top: 0;
    left: 0; */
    background: url( '../assets/H.png') no-repeat center center;
    background-size: cover;
    height: 100%;
    overflow: hidden;
    transform: scale(1.04);
    /* background-size: cover;
    background-color: rgb(202, 198, 198);
    transform: scale(1.1); */
}
</style>