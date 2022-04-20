<template>
  <v-row justify="center">
    <v-dialog v-model="closePopup" persistent max-width="800px">
      <template v-slot:activator="{ on }">
        <v-btn v-on="on" text>
          <span style="margin-right: 10px">Start Onboarding</span>
          <v-icon color="#D22B2B" justify="center"> folder_shared </v-icon>
        </v-btn>
      </template>
      <v-card>
        <v-card-title>
          <span class="headline">Onboarding Form</span>
        </v-card-title>
        <v-form class="px-3" ref="form">
          <v-card-text>
            <!-- <v-text-field
              label="Enter Employee ID"
              v-model="empID"
              hide-details
              single-line
              type="number"
            /> -->
            <v-text-field
              label="First Name"
              v-model="firstName"
              prepend-icon="person"
            ></v-text-field>
            <v-text-field
              label="Last Name"
              v-model="lastName"
              prepend-icon="person"
            ></v-text-field>

            <!-- Address Fields -->
            <!-- <v-text-field
              label="Permanent Address"
              v-model="permanentAddress"
              prepend-icon="home"
            ></v-text-field> -->
            <!-- <v-col class="d-flex" cols="12" sm="6">
                <v-select
                :items="states"
                label="Permanent State"
                dense
                outlined
                v-model="permanentState"
                ></v-select>
            </v-col> -->
            <!-- <v-col class="d-flex" cols="12" sm="6">
                <v-select
                :items="countries"
                label="Permanent Country"
                dense
                outlined
                v-model="permanentCountry"
                ></v-select>
            </v-col> -->
            <v-text-field
              label="Mailing Address"
              v-model="permanentAddress"
              prepend-icon="home"
            ></v-text-field>
            <!-- <v-col class="d-flex" cols="12" sm="6">
                <v-select
                :items="states"
                label="Mailing State"
                dense
                outlined
                v-model="mailingState"
                ></v-select>
            </v-col> -->
            <!-- <v-col class="d-flex" cols="12" sm="6">
                <v-select
                :items="countries"
                label="Mailing Country"
                dense
                outlined
                v-model="mailingCountry"
                ></v-select>
            </v-col> -->
            <v-text-field
              label="Alternate Email"
              v-model="alternateEmail"
              prepend-icon="person"
            ></v-text-field>

            <v-text-field
              label="Enter Drivers License Number"
              v-model="driversLicense"
              prepend-icon="person"
            ></v-text-field>
            <v-text-field
              label="Enter SSN"
              v-model="ssn"
              prepend-icon="person"
            ></v-text-field>
            <v-text-field
              label="Enter State ID"
              v-model="stateId"
              prepend-icon="person"
            ></v-text-field>

            <v-col class="d-flex" cols="12" sm="6">
                <v-select
                :items="genderList"
                label="Enter Gender"
                dense
                outlined
                v-model="empGender"
                ></v-select>
            </v-col>

            <v-col class="d-flex" cols="12" sm="6">
                <v-select
                :items="raceList"
                label="Enter Race"
                dense
                outlined
                v-model="empRace"
                ></v-select>
            </v-col>

            <v-col class="d-flex" cols="12" sm="6">
                <v-select
                :items="ethnicityList"
                label="Enter Ethnicity"
                dense
                outlined
                v-model="empEthnicity"
                ></v-select>
            </v-col>

            <v-col class="d-flex" cols="12" sm="6">
                <v-select
                :items="citizenshipList"
                label="Enter Country of Citizenship"
                dense
                outlined
                v-model="empCitizenship"
                ></v-select>
            </v-col>

            <v-col class="d-flex" cols="12" sm="6">
                <v-select
                :items="nationalityList"
                label="Enter Nationality"
                dense
                outlined
                v-model="empNationality"
                ></v-select>
            </v-col>

            <v-col class="d-flex" cols="12" sm="6">
                <v-select
                :items="pronounsList"
                label="Enter Your Pronouns"
                dense
                outlined
                v-model="empPronoun"
                ></v-select>
            </v-col>


            <v-col class="d-flex" cols="12" sm="6">
                <VueTelInputVuetify v-model = "mobileNumber" label="Mobile Number"/>
            </v-col>

            <v-divider></v-divider>

            <v-date-picker
                v-model="dob"
                scrollable
                no-title
                full-width
            />

            
            <!-- <v-col class="d-flex" cols="12" sm="6">
                <VueTelInputVuetify v-model = "alternateNumber" label="Alternate Number"/>
            </v-col> -->

          </v-card-text>
          
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="grey darken-1" text @click="cancel()"
              >Cancel</v-btn
            >
            <v-btn
              color="green darken-1"
              text
              outlined
              @click="submit()"
              >Apply</v-btn
            >
            <!-- closePopup = false -->
          </v-card-actions>
        </v-form>
      </v-card>
    </v-dialog>
  </v-row>
</template>

<script>
import VueTelInputVuetify from "vue-tel-input-vuetify/lib/vue-tel-input-vuetify.vue"
export default {
  components: {
      VueTelInputVuetify,
  },
  data: () => ({
    driversLicense : "",
    ssn : "",
    stateId : "",
    empID: 0,
    states: [
        'California', 'Florida', 'Washington DC', 'Wisconsin', 'Portland', 'Arizona', 'Texas', 'Georgia', 'North Carolina', 'South Carolina', 'Alabama', 'Michigan',
    ],
    countries : [
        'United States', 'India', 'France', 'Canada', 'United Kingdom', 'Bangladesh', 'Australia', 'Germany', 'Russia',
    ],
    genderList : [
      'Male' , 'Female' , 'Other', 'Prefer Not to say'
    ],
    empGender : "",
    raceList : [
      'African American', 'Native American', 'White', 'Hispanic'
    ],
    empRace : "",
    ethnicityList : [
      'African American', 'Native American', 'White', 'Hispanic'
    ],
    empEthnicity : "",
    citizenshipList : [
      'India', 'United States', 'Mexico', 'Canada', 'Italy', 'France', 'China', 'Syria', 'Great Britain', 'Switzerland'
    ],
    empCitizenship : "",

    nationalityList : [
      'India', 'United States', 'Mexico', 'Canada', 'Italy', 'France', 'China', 'Syria', 'Great Britain', 'Switzerland'
    ],
    empNationality : "",
    pronounsList : [
      'He or Him', 'She', 'Ze', 'Others'
    ],
    empPronoun : "",
    dob : null,
    mobileNumber : "",
    AlternateNumber: "",
    permanentState: 1,
    permanentCountry:1,
    mailingState: 1,
    mailingCountry:1,
    dropdownItems:[
        'Paid Leave', 'Un-Paid Leave'
    ],
    firstName : "",
    lastName : "",
    alternateEmail : "",
    permanentAddress : "",
    leaveTypeSelection:1,
    closePopup: false,
    preferedName: "",
    inputRules: [(v) => v.length >= 3 || "Minimum lenght is 3 charachters"],
  }),

  methods : {
    cancel(){
      this.driversLicense = ""
      this.ssn = ""
      this.stateId = ""
      this.permanentAddress = ""
      this.alternateEmail = ""
      this.mobileNumber = ""
      this.firstName = ""
      this.lastName = ""
      this.empGender = ""
      this.empRace = ""
      this.empEthnicity = ""
      this.empCitizenship = ""
      this.empNationality = ""
      this.empPronoun = ""
      this.dob = null
      this.closePopup = false
    },
    submit(){
      const reqObj = {
        "EmployeeID" : parseInt(this.empID),
        "DriversLicense" : this.driversLicense,
        "SSN" : this.ssn,
        "StateID" : this.stateId,
        "Address" : this.permanentAddress,
        "AlternateEmails" : this.alternateEmail,
        "Phone" : this.mobileNumber,
        "FirstName" : this.firstName,
        "LastName" : this.lastName,
        "Gender" : this.empGender,
        "Race" : this.empRace,
        "Ethnicity" : this.empEthnicity,
        "Citizenship" : this.empCitizenship,
        "Nationality" : this.empNationality,
        "Pronouns" : this.empPronoun,
        "DateOfBirth" : this.dob,
      }

      this.$axios.post("http://localhost:8080/users/UpdateEmployeeInfo1", reqObj)
        .then(response => {
          console.log(response)

          //Clearing input after successful submission
          this.driversLicense = ""
          this.ssn = ""
          this.stateId = ""
          this.permanentAddress = ""
          this.alternateEmail = ""
          this.mobileNumber = ""
          this.firstName = ""
          this.lastName = ""
          this.empGender = ""
          this.empRace = ""
          this.empEthnicity = ""
          this.empCitizenship = ""
          this.empNationality = ""
          this.empPronoun = ""
          this.dob = null

          this.$emit('notif', 'Onboarding form submitted successfully.', "success")
          this.closePopup = false
        }).catch(error => {
          console.log(error)

          //Clearing input after failed submission
          this.driversLicense = ""
          this.ssn = ""
          this.stateId = ""
          this.permanentAddress = ""
          this.alternateEmail = ""
          this.mobileNumber = ""
          this.firstName = ""
          this.lastName = ""
          this.empGender = ""
          this.empRace = ""
          this.empEthnicity = ""
          this.empCitizenship = ""
          this.empNationality = ""
          this.empPronoun = ""
          this.dob = null
          this.$emit('notif', 'Failed to submit onboarding form', "error")
          this.closePopup = false
        })
    }
  },
};
</script>