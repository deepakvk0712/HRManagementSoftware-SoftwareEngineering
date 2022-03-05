<template>
  <v-app>
    <!-- <v-parallax dark src="https://cdn.vuetifyjs.com/images/backgrounds/vbanner.jpg" height ="100%" jumbotron> -->
    <validation-observer ref="observer" v-slot="{ invalid }">
        <!-- v-slot="{ invalid }" -->
      <form @submit.prevent="submit">
        <v-content>
          <v-container fluid pa-0>

            <!-- <v-flex v-if="displayError.length > 0" fluid>
                xs12 -->
                <!-- <v-alert width="90%" class="text-center" type="warning" :value="displayError" dismissible -->
                <!-- >{{displayError}}</v-alert -->
                <!-- > -->
            <!-- </v-flex> --> 

            <v-row align="center" justify="center" style="height: 100vh" dense>
              <v-card width="600" class="justify-center">
                <v-card-title id="regChangePassword">Change your Password here!</v-card-title>
                
                <!-- <validation-provider
                  v-slot="{ errors }"
                  name="Name"
                  rules="required|max:20">

                  <v-text-field
                    v-model="username"
                    :counter="20"
                    :error-messages="errors"
                    label="Name"
                    required
                    style="margin:25px;"
                  ></v-text-field>
                </validation-provider> -->

                <validation-provider id="regEmailID"
                  v-slot="{ errors }"
                  name="EmailID"
                  rules="required|email"
                >
                  <v-text-field
                    v-model="email"
                    :error-messages="errors"
                    label="E-mail"
                    required
                    style="margin:25px;"
                  ></v-text-field>
                </validation-provider>

                <validation-provider id="regOldPassword"
                  v-slot="{ errors }"
                  name="Old Password"
                  rules="required|min:8"
                >
                  <v-text-field
                    v-model="password"
                    :counter="8"
                    :error-messages="errors"
                    label="Old Password"
                    type="password"
                    
                    required
                    style="margin:25px;"
                  >
                    <template v-slot:append>
                      <v-icon v-if="PasswordCheck" color="green">
                      </v-icon>
                      <v-icon v-if="!PasswordCheck" color="red">
                      </v-icon>
                    </template>
                  </v-text-field>
                </validation-provider>

                <validation-provider id="regNewPassword"
                  v-slot="{ errors }"
                  name="New Password"
                  rules="required|min:8"
                >
                  <v-text-field
                    v-model="newPassword"
                    :counter="8"
                    :error-messages="errors"
                    label="New Password"
                    type="password"
                    :rules="[notMatchingPasswords]"
                    required
                    style="margin:25px;"
                  >
                    <template v-slot:append>
                      <v-icon v-if="PasswordCheck2" color="green">
                      </v-icon>
                      <v-icon v-if="!PasswordCheck2" color="red">
                      </v-icon>
                    </template>
                  </v-text-field>
                </validation-provider>

                <v-divider></v-divider>
                <v-btn id="regSubmitButton" @click="changePassword" class="mr-4" type="submit" style="margin:25px;" :disabled="invalid">submit</v-btn>
                <v-btn id="regClearButton" @click="clear" style="margin:25px;">clear</v-btn>
              </v-card>
            </v-row>
          </v-container>
        </v-content>
      </form>
    </validation-observer>
    <!-- </v-parallax> -->
  </v-app>
</template>

<script>
import { required, email, min } from "vee-validate/dist/rules";
import {
  extend,
  ValidationObserver,
  ValidationProvider,
  setInteractionMode,
} from "vee-validate";
setInteractionMode("eager");

extend("required", {
  ...required,
  message: "{_field_} is required",
});
// extend("max", {
//   ...max,
//   message: "{_field_} cannot be more than {length} characters",
// });

extend("min", {
  ...min,
  message: "{_field_} cannot be less than {length} characters",
});

extend("email", {
  ...email,
  message: "Email-ID should be valid",
});
export default {
  components: {
    ValidationProvider,
    ValidationObserver,
  },
  data: () => ({
    // username: "",
    email: "",
    password: "",
    newPassword: "",
    PasswordCheck: false,
    PasswordCheck2: false,
    // valid: false,
    // displayError : "",
  }),
  methods: {
//    async submit() {
//       const isValid = await this.$refs.observer.validate();
//       if(isValid){
//         this.$router.push("/login");
//       } 
//     },

    // enableButton() {
    //     if(this.valid==true && this.invalid==false){
    //         return false
    //     }
    //     else{
    //         return true
    //     }
    // },

    submit() {
      this.$refs.observer.validate();
      this.$router.push("/login");
    },

    clear() {
    //   this.username = "";
      this.email = "";
      this.password = "";
      this.newPassword = "";
      this.$refs.observer.reset();
    },

    changePassword() {
            const requestObj = {
                "email" : this.email,
                "oldPassword" : this.password,
                "newPassword" : this.newPassword,
            }

            this.$axios.post("http://localhost:8080/login", requestObj)
                .then(response => {

                    let jsonData = JSON.parse(response.data.data)
                    console.log(jsonData)
                    
                    if(response.error == null){
                        this.$router.push('/login')
                    }
                    let token = jsonData.AccessToken
                    this.$axios.defaults.headers.common["Authorization"] = "Bearer " + token
                })
        },


    
    // submitBut: function() {
    //   alert("Password has been reset! Login with new Password");
    // console.log()
        // if(!this.$refs.observer.validate()){
        //     this.displayError = "Password has been reset! Login with new Password";
        // }
        // else{
        //     this.$router.push("/login");
        // }
        // this.displayError = "Password has been reset! Login with new Password";
        
    // },
    notMatchingPasswords: function () {
      if (this.password != this.newPassword) {
            this.PasswordCheck = true;
            this.PasswordCheck2 = true;
            return true;
      } else {
            return "Old Password and New Password cannot be the same.";
      }
    },
  },
};
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

    width: 100%;
    height: 100%;
    position: absolute;
    top: 0;
    left: 0;
    background-size: cover;
    background-color: rgb(146, 215, 228);
    transform: scale(1.1);
}
</style>