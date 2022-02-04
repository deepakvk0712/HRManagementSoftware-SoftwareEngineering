<template>
  <v-row justify="center">
    <v-dialog v-model="closePopup" persistent max-width="800px">
      <template v-slot:activator="{ on }">
        <v-btn v-on="on" text>
          <span style="margin-right: 10px">Start Onboarding</span>
          <v-icon justify="center"> folder_shared </v-icon>
        </v-btn>
      </template>
      <v-card>
        <v-card-title>
          <span class="headline">Onboarding Form</span>
        </v-card-title>
        <v-form class="px-3" ref="form">
          <v-card-text>
            <v-text-field
              label="Prefered Name"
              v-model="preferedName"
              prepend-icon="person"
            ></v-text-field>

            <!-- Address Fields -->
            <v-text-field
              label="Permanent Address"
              v-model="permanentAddress"
              prepend-icon="home"
            ></v-text-field>
            <v-col class="d-flex" cols="12" sm="6">
                <v-select
                :items="states"
                label="Permanent State"
                dense
                outlined
                v-model="permanentState"
                ></v-select>
            </v-col>
            <v-col class="d-flex" cols="12" sm="6">
                <v-select
                :items="countries"
                label="Permanent Country"
                dense
                outlined
                v-model="permanentCountry"
                ></v-select>
            </v-col>
            <v-text-field
              label="Mailing Address"
              v-model="permanentAddress"
              prepend-icon="home"
            ></v-text-field>
            <v-col class="d-flex" cols="12" sm="6">
                <v-select
                :items="states"
                label="Mailing State"
                dense
                outlined
                v-model="mailingState"
                ></v-select>
            </v-col>
            <v-col class="d-flex" cols="12" sm="6">
                <v-select
                :items="countries"
                label="Mailing Country"
                dense
                outlined
                v-model="mailingCountry"
                ></v-select>
            </v-col>
            <v-col class="d-flex" cols="12" sm="6">
                <VueTelInputVuetify v-model = "mobileNumber" label="Mobile Number"/>
            </v-col>
            <v-col class="d-flex" cols="12" sm="6">
                <VueTelInputVuetify v-model = "alternateNumber" label="Alternate Number"/>
            </v-col>

          </v-card-text>
          
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="grey darken-1" text @click="closePopup = false"
              >Cancel</v-btn
            >
            <v-btn
              color="green darken-1"
              text
              outlined
              @click="closePopup = false"
              >Apply</v-btn
            >
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
    states: [
        'California', 'Florida', 'Washington DC', 'Wisconsin', 'Portland', 'Arizona', 'Texas', 'Georgia', 'North Carolina', 'South Carolina', 'Alabama', 'Michigan',
    ],
    countries : [
        'United States', 'India', 'France', 'Canada', 'United Kingdom', 'Bangladesh', 'Australia', 'Germany', 'Russia',
    ],
    mobileNumber : "",
    AlternateNumber: "",
    permanentState: 1,
    permanentCountry:1,
    mailingState: 1,
    mailingCountry:1,
    dropdownItems:[
        'Paid Leave', 'Un-Paid Leave'
    ],
    permanentAddress : "",
    leaveTypeSelection:1,
    closePopup: false,
    preferedName: "",
    inputRules: [(v) => v.length >= 3 || "Minimum lenght is 3 charachters"],
  }),
  methods: {
    
  },
  computed: {
    
  },
  watch: {
    
  },
};
</script>