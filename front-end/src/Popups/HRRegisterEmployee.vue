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
          <span id="regEmployeeHead" class="headline">HR Register Employee</span>
        </v-card-title>
        <v-form class="px-3" ref="form">
          <v-card-text>
            <v-text-field
              id="regEmpNameInput"
              label="Enter employee first name"
              v-model="firstName"
              prepend-icon="edit"
              :rules="inputRules"
              required
            ></v-text-field>
            <v-text-field
              id="regEmpLastNameInput"
              label="Enter employee last name"
              v-model="lastName"
              prepend-icon="edit"
              :rules="inputRules"
              required
            ></v-text-field>
            <v-text-field
              id="regEmpBusUnitInput"
              label="Enter buissness unit"
              v-model="businessUnit"
              prepend-icon="edit"
              :rules="inputRules"
              required
            ></v-text-field>
            <v-text-field
              id="regEmpTitleInput"
              label="Enter employee title"
              v-model="title"
              prepend-icon="edit"
              :rules="inputRules"
              required
            ></v-text-field>
            <v-text-field
              id="regEmpTypeInput"
              label="Enter employee type"
              v-model="type"
              prepend-icon="edit"
              :rules="inputRules"
              required
            ></v-text-field>
            <v-text-field
              id = "regEmpManagerIdInput"
              v-model="managerId"
              hide-details
              label="Enter Manager ID"
              prepend-icon="edit"
              single-line
              type="number"
              required
            />
            <v-text-field
              id="regEmpGradeInput"
              label="Enter employee grade"
              v-model="grade"
              prepend-icon="edit"
              :rules="inputRules"
              required
            ></v-text-field>
            <v-text-field
              id="regEmpLocationInput"
              label="Enter employee location"
              v-model="location"
              prepend-icon="edit"
              :rules="inputRules"
              required
            ></v-text-field>
            <v-col class="d-flex" cols="12" sm="6">
                <v-select
                id="regEmpCountryInput"
                :items="countries"
                label="Enter country"
                dense
                outlined
                v-model="country"
                required
                ></v-select>
            </v-col>
            <v-text-field
              id= "regEmpPersEmailInput"
              label="Enter employee personal email"
              v-model="personalEmail"
              prepend-icon="edit"
              :rules="inputRules"
              required
            ></v-text-field>

            <v-text-field
              id= "Salary"
              label="Enter employee negotiated salary"
              v-model="salary"
              prepend-icon="edit"
              :rules="inputRules"
              single-line
              type="number"
              required
            ></v-text-field>

            <v-col class="d-flex" cols="12" sm="6">
                <v-select
                id="regEmpIsHrInput"
                :items="HrPossibility"
                label="Is this a new HR?"
                dense
                outlined
                v-model="IsHR"
                required
                ></v-select>
            </v-col>

            <v-col class="d-flex" cols="12" sm="6">
                <v-select
                id="regEmpIsManagerInput"
                :items="ManagerPossibility"
                label="Is this a new Manger?"
                dense
                outlined
                v-model="IsManager"
                required
                ></v-select>
            </v-col>
          </v-card-text>
          
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn id="regEmpCancelButton" color="grey darken-1" text @click="closePopup = false"
              >Cancel</v-btn
            >
            <v-btn
              id="regEmpSubmitButton"
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
    countries : ['United States', 'India', 'France', 'Canada', 'United Kingdom', 'Bangladesh', 'Australia', 'Germany', 'Russia',],
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
    salary: 0,
    IsHR:false,
    HrPossibility : [true, false],
    IsManager: false,
    ManagerPossibility: [true, false],
  }),

  methods : {
        submit() {
          const requestObj = {
              "firstName" : this.firstName,
              "lastName" : this.lastName,
              "businessUnit" : this.businessUnit,
              "managerId" : parseInt(this.managerId),
              "title" : this.title,
              "type" : this.type,
              "grade" : this.grade,
              "location" : this.location,
              "country" : this.country,
              "personalEmail" : this.personalEmail,
              "isHR" : this.IsHR,
              "isManager" : this.IsManager,
              "salary" : parseInt(this.salary)
          }
          this.$axios.post("http://localhost:8080/users/registerHR", requestObj)
            .then(response => {
                console.log(response)
                this.$emit('notif', 'Successfully registered a new employee.', "success")
                this.closePopup = false
                                  
            })
        },
    }
};
</script>