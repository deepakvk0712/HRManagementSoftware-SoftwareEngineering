<template>
    <v-app>
        <v-parallax dark src="https://cdn.vuetifyjs.com/images/backgrounds/vbanner.jpg" height ="100%" jumbotron>
            <v-content>
                <v-container fluid pa-0>
                <v-row align="center" justify="center" style="height:100vh" dense>
                    <v-card width = "600" class ="justify-center">
                        <v-card-title> Login </v-card-title>
                        <v-card-text class="text-center">
                            <v-text-field label = "Username" v-model="userName" prepend-icon="mdi-account-circle"/>
                            <v-text-field label = "Password" 
                            :type="showPassword ? 'text' : 'password'"
                            v-model="password"
                            prepend-icon="mdi-lock"
                            :append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
                            @click-append="showPassword = !showPassword"/>
                        </v-card-text>

                        <v-divider></v-divider>

                        <v-card-actions>
                            <v-btn to="/signup" color = "#89CFF0" > SignUp </v-btn>
                            <v-btn  color = "#6495ED" @click="login"> Login </v-btn> 
                            <!-- to="/landing" -->
                        </v-card-actions>
                    </v-card>
                </v-row>
                </v-container>
            </v-content>
        </v-parallax> 
    </v-app>
</template>

<script>
export default {
    name: "LoginPage",
    data: () => ({
      showPassword:false,
      userName : "",
      password : "",
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
                    let token = jsonData.accessToken
                    this.$store.state.accessToken = token
                    console.log(token)
                    this.$axios.defaults.headers.common['Authorization'] = "Bearer " + token
                    this.$store.commit('LOADER', false)
                    this.$router.push('/landing')           
                    
                })
        },

    }

}
</script>
