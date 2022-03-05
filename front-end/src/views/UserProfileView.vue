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
              <!-- Ashley Harrington -->
              {{empName}}
            </v-card-subtitle>

            <v-card-subtitle class="py-0 justify-center font-weight-black">
              <!-- Project Manager -->
              {{empDesignation}}
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
                  <strong>{{ prodScore }}%</strong>
                </v-card-subtitle>
              </v-vol>
            </v-row>

            <v-progress-linear
              v-model="prodScore"
              height="20"
            >
              
            </v-progress-linear>

            <v-divider class="pt-2"></v-divider>

            <v-divider></v-divider>

            <v-card-subtitle class="text-left font-weight-black">
              About Me
            </v-card-subtitle>
            <v-card-text class="text-left">
              <!-- This is Ashely Harrington. I am an avid snowboarder and I love to go on long hikes with my dogs. I enjoy the mountains more than the beach. -->
              {{descriptionEmp}}
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
                  v-model="firstName"
                  class="pt-0 ps-4 mr-4"
                ></v-text-field>
                <!-- label="First Name" -->
              </v-col>

              <v-col>
                <v-card-subtitle id="usrProLName" class="pb-0 text-left font-weight-black">
                  Last Name
                </v-card-subtitle>
                <v-text-field
                  id="usrProLNameInput"
                  v-model="lastName"
                  class="pt-0 ps-4 mr-4"
                ></v-text-field>
                <!-- label="Last Name" -->
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
                <VueTelInputVuetify id="usrProNumberInput" style="padding-top:0px;" class="ml-4" v-model="mobileNumber" />
                <!-- label="Mobile Number" -->
              </v-col>

              <v-col class="py-0">
                <v-card-subtitle id="usrEmail" class="py-0 text-left font-weight-black">
                  Email Address
                </v-card-subtitle>
                <v-text-field
                  id="usrProEmailInput"
                  v-model="emailId"
                  class="pt-0 ps-4 mr-4"
                ></v-text-field>
                <!-- label="Email Address" -->
              </v-col>

            </v-row>

            <v-row>
              <v-col class="py-0">
                <v-card-subtitle id="usrAddr" class="py-0 text-left font-weight-black">
                  Address
                </v-card-subtitle>
                <v-text-field
                  id="usrProAddressInput"               
                  v-model="address"
                  class="pt-0 ps-4 mr-4"
                ></v-text-field>
                <!-- label="Employee Address" -->
              </v-col>
            </v-row>

            <v-row>
              <v-col class="py-0">
                <v-card-subtitle id="usrCity" class="py-0 text-left font-weight-black">
                  <!-- City -->
                  City
                </v-card-subtitle>
                <v-text-field
                  id="usrProCityInput"                 
                  v-model="city"
                  class="pt-0 ps-4 mr-4"
                ></v-text-field>
                <!-- label="Designation " -->
              </v-col>

              <v-col class="py-0">
                <v-card-subtitle id="usrState" class="py-0 text-left font-weight-black">
                  State
                </v-card-subtitle>
                <v-select
                id="usrProStateInput"
                :items="states"      
                dense
                outlined
                v-model="empState"
                ></v-select>
                <!-- label="State" -->
              </v-col>

              <v-col class="py-0">
                <v-card-subtitle id="usrZip" class="py-0 text-left font-weight-black">
                  Zip
                </v-card-subtitle>
                <v-text-field
                  id="usrProZipInput"
                  v-model="zipCode"
                  class="pt-0 ps-4 mr-4"
                ></v-text-field>
                <!-- label="Zip " -->
              </v-col>

            </v-row>

            <v-row>
              <v-col class="py-0">
                <v-card-subtitle id="usrDesc" class="py-0 text-left font-weight-black">
                  Description
                </v-card-subtitle>
                <v-text-field
                  id="usrProDescInput"
                  v-model="description"
                  class="pt-0 ps-4 mr-4"
                ></v-text-field>
                <!-- label="Description" -->
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
  async mounted() {
    await this.$axios.get('http://10.20.205.4:8080/users/profile')
      .then(response => {
        // console.log(response);
        let respObj = JSON.parse(response.data.data)
        this.empName = respObj.firstName + " " + respObj.lastName
        this.empDesignation = respObj.title
        this.descriptionEmp = respObj.aboutMe
        this.prodScore = respObj.productivityScore
        // console.log(this.$store.state.userName);
      })
  },
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
    // designation : "",
    city : "",
    empState : "",
    zipCode: "",
    description : "",
    userName : store.state.userName,

    empName : "",
    empDesignation : "",
    prodScore : 20,
    descriptionEmp : "",
  }),

  methods : {
    submit() {
      const reqObj = {
        "firstName" : this.firstName,
        "lastName" : this.lastName,
        "phone" : this.mobileNumber,
        "personalEmail" : this.emailId,
        "address" : this.address,
        "title" : this.empDesignation,
        // "empState" : this.empState,
        // "zipCode" : this.zipCode,
        "aboutMe" : this.description,
      }

      // Making a Update call to the backend to update user profile.
      this.$axios.put("http://10.20.205.4:8080/users/profile", reqObj)
        .then(response => {
            console.log(response)
            this.empName = this.firstName + " " + this.lastName
            // this.empDesignation = this.designation
            this.descriptionEmp = this.description   
        })
    }
  },
}
</script>
