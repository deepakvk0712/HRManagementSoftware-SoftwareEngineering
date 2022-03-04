<template>
  <div class="home">
    <v-container>

      <NavigationBar v-bind:userName="userName"/>
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

            <v-divider class="pb-2"></v-divider>

            <v-row fluid>
              <v-col>
                <v-card-subtitle class="pl-2 py-0 text-left">
                Productivity
                </v-card-subtitle>
              </v-col>
              <v-vol>
                <v-card-subtitle class="py-2 text-right">
                  <strong>{{ perProductivity }}%</strong>
                </v-card-subtitle>
              </v-vol>
            </v-row>

            <v-progress-linear
              v-model="perProductivity"
              height="20"
            >
              
            </v-progress-linear>

            <v-divider class="pt-2"></v-divider>

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
            <v-card-title id="usrProfHead" class="align-center">
              Account Details
            </v-card-title>

            <v-divider class="py-2"></v-divider>

            <v-row>
              <v-col>
                <v-card-subtitle id="usrProFName" class="pb-0 text-left font-weight-black">
                  First Name
                </v-card-subtitle>
                <v-text-field
                  id="usrProFNameInput"
                  label="First Name"
                  v-model="firstName"
                  class="pt-0 ps-4 mr-4"
                ></v-text-field>
              </v-col>

              <v-col>
                <v-card-subtitle id="usrProLName" class="pb-0 text-left font-weight-black">
                  Last Name
                </v-card-subtitle>
                <v-text-field
                  id="usrProLNameInput"
                  label="Last Name"
                  v-model="lastName"
                  class="pt-0 ps-4 mr-4"
                ></v-text-field>
              </v-col>

            </v-row>

            <v-row>
              <v-col class="py-0">
                <v-card-subtitle id="usrPhoneNumber" class="py-0 text-left font-weight-black">
                  Phone Number
                </v-card-subtitle>
                <!-- <v-text-field
                  label="Phone Number"
                  v-model="firstName"
                  class="pt-0 ps-4 mr-4"
                ></v-text-field> -->
                <VueTelInputVuetify id="usrProNumberInput" style="padding-top:0px;" class="ml-4" v-model="mobileNumber" label="Mobile Number"/>
              </v-col>

              <v-col class="py-0">
                <v-card-subtitle id="usrEmail" class="py-0 text-left font-weight-black">
                  Email Address
                </v-card-subtitle>
                <v-text-field
                  id="usrProEmailInput"
                  label="Email Address"
                  v-model="emailId"
                  class="pt-0 ps-4 mr-4"
                ></v-text-field>
              </v-col>

            </v-row>

            <v-row>
              <v-col class="py-0">
                <v-card-subtitle id="usrAddr" class="py-0 text-left font-weight-black">
                  Address
                </v-card-subtitle>
                <v-text-field
                  id="usrProAddressInput"
                  label="Employee Address"
                  v-model="address"
                  class="pt-0 ps-4 mr-4"
                ></v-text-field>
              </v-col>
            </v-row>

            <v-row>
              <v-col class="py-0">
                <v-card-subtitle id="usrCity" class="py-0 text-left font-weight-black">
                  City
                </v-card-subtitle>
                <v-text-field
                  id="usrProCityInput"
                  label="City "
                  v-model="city"
                  class="pt-0 ps-4 mr-4"
                ></v-text-field>
              </v-col>

              <v-col class="py-0">
                <v-card-subtitle id="usrState" class="py-0 text-left font-weight-black">
                  State
                </v-card-subtitle>
                <v-select
                id="usrProStateInput"
                :items="states"
                label="State"
                dense
                outlined
                v-model="empState"
                ></v-select>
              </v-col>

              <v-col class="py-0">
                <v-card-subtitle id="usrZip" class="py-0 text-left font-weight-black">
                  Zip
                </v-card-subtitle>
                <v-text-field
                  id="usrProZipInput"
                  label="Zip "
                  v-model="zipCode"
                  class="pt-0 ps-4 mr-4"
                ></v-text-field>
              </v-col>

            </v-row>

            <v-row>
              <v-col class="py-0">
                <v-card-subtitle id="usrDesc" class="py-0 text-left font-weight-black">
                  Description
                </v-card-subtitle>
                <v-text-field
                  id="usrProDescInput"
                  label="Description"
                  v-model="description"
                  class="pt-0 ps-4 mr-4"
                ></v-text-field>
              </v-col>
            </v-row>

            <v-row class="ps-6 pb-4">
              <v-btn
                id="usrUpdateButton"
                color="green darken-1"
                text
                outlined
                @click="submit()"
                ><span class="pl-4 pr-4">Update</span></v-btn
              >

            </v-row>
          </v-card>
        </v-col>
        </v-row>
      </v-layout>

      <PageBottom/>
    </v-container>
  </div>
</template>

<script>
import NavigationBar from '../components/NavigationBar.vue'
import VueTelInputVuetify from "vue-tel-input-vuetify/lib/vue-tel-input-vuetify.vue"
import PageBottom from '../components/PageBottom.vue'
import store from '../store/store'
export default {
  name: 'userProfile',
  components: {
    NavigationBar, VueTelInputVuetify, PageBottom,
  },
  data: () => ({
    perProductivity : 35,
    states: [
        'California', 'Florida', 'Washington DC', 'Wisconsin', 'Portland', 'Arizona', 'Texas', 'Georgia', 'North Carolina', 'South Carolina', 'Alabama', 'Michigan',
    ],
    firstName : "",
    lastName : "",
    mobileNumber : null,
    emailId : "",
    address : "",
    city : "",
    empState : "",
    zipCode: "",
    description : "",
    userName : store.state.userName,
  }),

  methods : {
    submit() {
      // const reqObj = {
      //   "firstName" : this.firstName,
      //   "lastName" : this.lastName,
      //   "mobileNumber" : this.mobileNumber,
      //   "emailId" : this.emailId,
      //   "address" : this.address,
      //   "city" : this.city,
      //   "empState" : this.empState,
      //   "zipCode" : this.zipCode,
      //   "descrption" : this.descrption,
      // }

      //Making a Update call to the backend to update user profile.
      // this.$axios.put("http://localhost:8080/users/profile", reqObj)
      //   .then(response => {
      //       console.log(response)
                    
      //   })
    }
  },
}
</script>
