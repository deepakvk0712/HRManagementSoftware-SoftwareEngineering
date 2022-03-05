<template>
  <v-app>
    <!-- <v-parallax dark src="https://cdn.vuetifyjs.com/images/backgrounds/vbanner.jpg" height ="100%" jumbotron> -->
    <validation-observer ref="observer" v-slot="{ invalid }">
      <form @submit.prevent="submit">
        <v-content>
          <v-container fluid pa-0>
            <v-row align="center" justify="center" style="height: 100vh" dense>
              <v-card width="600" class="justify-center">
                <v-card-title id="regMessage">Dont have an account? Signup here </v-card-title>
                <validation-provider id="regName"
                  v-slot="{ errors }"
                  name="Name"
                  rules="required|max:20"
                >
                  <v-text-field
                    v-model="username"
                    :counter="20"
                    :error-messages="errors"
                    label="Name"
                    required
                    style="margin:25px;"
                  ></v-text-field>
                </validation-provider>

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

                <validation-provider id="regPassword"
                  v-slot="{ errors }"
                  name="Password"
                  rules="required|min:8"
                >
                  <v-text-field
                    v-model="password"
                    :counter="8"
                    :error-messages="errors"
                    label="Password"
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

                <validation-provider id="regPassword1"
                  v-slot="{ errors }"
                  name="Password"
                  rules="required|min:8"
                >
                  <v-text-field
                    v-model="passwordRetype"
                    :counter="8"
                    :error-messages="errors"
                    label="Re-type Password"
                    type="password"
                    :rules="[matchingPasswords]"
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
                <v-btn id="regSubmit" class="mr-4" to="/login" type="submit" style="margin:25px;" :disabled="invalid"
                  >submit</v-btn
                >
                <v-btn id="regClear" @click="clear" style="margin:25px;">clear</v-btn>
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
import { required, email, max, min } from "vee-validate/dist/rules";
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
extend("max", {
  ...max,
  message: "{_field_} cannot be more than {length} characters",
});

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
    username: "",
    email: "",
    password: "",
    passwordRetype: "",
    PasswordCheck: false,
    PasswordCheck2: false,
  }),
  methods: {
    submit() {
      this.$refs.observer.validate();
    },
    clear() {
      this.username = "";
      this.email = "";
      this.password = "";
      this.passwordRetype = "";
      this.$refs.observer.reset();
    },
    matchingPasswords: function () {
      if (this.password === this.passwordRetype) {
        this.PasswordCheck = true;
        this.PasswordCheck2 = true;
        return true;
      } else {
        return "Passwords does not match.";
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