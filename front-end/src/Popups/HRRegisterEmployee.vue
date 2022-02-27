<template>
  <v-row justify="center"> 
    <v-dialog v-model="closePopup" persistent max-width="800px">
      <template v-slot:activator="{ on }">
        <v-btn v-on="on" text>
          <span style="margin-right: 10px"> Register Employee </span>
          <v-icon color="#D22B2B" justify="center"> folder_shared </v-icon>
        </v-btn>
      </template>
      <v-card>
        <v-card-title>
          <span class="headline">HR Register Employee</span>
        </v-card-title>
        <v-form class="px-3" ref="form">
          <v-card-text>
            <v-text-field
              label="Enter employee first name"
              v-model="firstName"
              prepend-icon="edit"
              :rules="inputRules"
            ></v-text-field>
            <v-text-field
              label="Enter employee last name"
              v-model="lastName"
              prepend-icon="edit"
              :rules="inputRules"
            ></v-text-field>
            <v-text-field
              label="Enter buissness unit"
              v-model="businessUnit"
              prepend-icon="edit"
              :rules="inputRules"
            ></v-text-field>
            <v-text-field
              label="Enter employee title"
              v-model="title"
              prepend-icon="edit"
              :rules="inputRules"
            ></v-text-field>
            <v-text-field
              label="Enter employee type"
              v-model="type"
              prepend-icon="edit"
              :rules="inputRules"
            ></v-text-field>
            <v-text-field
                v-model="managerId"
                hide-details
                single-line
                type="number"
            />
            <v-text-field
              label="Enter employee grade"
              v-model="grade"
              prepend-icon="edit"
              :rules="inputRules"
            ></v-text-field>
            <v-text-field
              label="Enter employee location"
              v-model="location"
              prepend-icon="edit"
              :rules="inputRules"
            ></v-text-field>
            <v-col class="d-flex" cols="12" sm="6">
                <v-select
                :items="countries"
                label="Enter country"
                dense
                outlined
                v-model="country"
                ></v-select>
            </v-col>
            <v-text-field
              label="Enter employee personal email"
              v-model="personalEmail"
              prepend-icon="edit"
              :rules="inputRules"
            ></v-text-field>
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
              @click="submit()"
              >Submit</v-btn
            >
          </v-card-actions>
        </v-form>
      </v-card>
    </v-dialog>
  </v-row>
</template>

<script>
export default {
  data: () => ({
    closePopup: false,
    inputRules: [(v) => v.length >= 3 || "Minimum lenght is 3 charachters"],
    firstName : "",
    lastName : "",
    businessUnit : "",
    managerId : 0,
    title : "",
    type : "",
    grade : "",
    location : "",
    country : "",
    personalEmail : "",
  }),

  methods : {
        submit() {
            const requestObj = {
                "firstName" : this.firstName,
                "lastName" : this.lastName,
                "businessUnit" : this.businessUnit,
                "managerId" : this.managerId, //not added
                "title" : this.title,
                "type" : this.type,
                "grade" : this.grade,
                "location" : this.location,
                "country" : this.country,
                "personalEmail" : this.personalEmail,
            }
            // closePopup = false
            this.$axios.post("http://localhost:8080/register/HR", requestObj)
                .then(response => {
                    console.log(response)
                    this.closePopup = false
                    // this.$router.push('/landing')
                    // let token = response.data.token;
                    // this.$axios.defaults.headers.common["Authorization"] = "Bearer " + token;
                    
                })
        },

    }
};
</script>