<template>
    <v-app>
        <v-parallax dark src="https://cdn.vuetifyjs.com/images/backgrounds/vbanner.jpg" height ="100%" jumbotron>
            <!-- <v-row align="center" justify="center">
                <v-col class="text-center" cols="12"></v-col>
            </v-row> -->
            <v-content>
                <v-container fluid pa-0>
                <v-row align="center" justify="center" style="height:100vh" dense>
                    <v-card width = "600" class ="justify-center">
                        <v-card-title> Login </v-card-title>
                        <v-card-text class="text-center">
                            <v-text-field label = "Username" prepend-icon="mdi-account-circle"/>
                            <v-text-field label = "Password" 
                            :type="showPassword ? 'text' : 'password'"
                            prepend-icon="mdi-lock"
                            :append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
                            @click-append="showPassword = !showPassword"/>
                        </v-card-text>

                        <v-divider></v-divider>

                        <v-card-actions>
                            <v-btn to="/signup" color = "#89CFF0"> SignUp </v-btn>
                            <v-btn to="/landing" color = "#6495ED"> Login </v-btn>
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
      showPassword:false
    }),
}
</script>