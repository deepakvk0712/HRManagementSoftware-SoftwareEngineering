<template>
  <v-row justify="center">
    <v-dialog v-model="closePopup" persistent max-width="800px">
      <template v-slot:activator="{ on }">
        <v-btn v-on="on" text>
          <span style="margin-right: 10px"> Fill Form </span>
          <v-icon color="#D22B2B" justify="center"> folder_shared </v-icon>
        </v-btn>
      </template>
      <v-card>
        <v-card-title>
          <span id="finCardHead" class="headline">Banking Information</span>
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
              id="bankNameInput"
              label="Enter Bank Name"
              v-model="bankName"
              prepend-icon="edit"
              :rules="inputRules"
            ></v-text-field>
            <v-text-field
              id="routingNumberInput"
              label="Enter Routing Number"
              v-model="routingNumber"
              prepend-icon="edit"
              :rules="inputRules"
            ></v-text-field>
            <v-text-field
              id="accountNumberInput"
              label="Enter Account Number"
              v-model="accountNumber"
              prepend-icon="edit"
              :rules="inputRules"
            ></v-text-field>
            <!-- <v-col class="d-flex" cols="12" sm="6">
                <v-select
                :items="typeOfAccount"
                label="Account Type"
                dense
                outlined
                v-model="accountType"
                ></v-select>
            </v-col> -->
          </v-card-text>
          
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn id="finCancelButton" color="grey darken-1" text @click="cancel()"
              >Cancel</v-btn
            >
            <v-btn
              id = "finSubmitButton"
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
    // typeOfAccount : [
    //     'Checking Account', 'Savings Account',
    // ],
    inputRules: [(v) => v.length >= 3 || "Minimum lenght is 3 charachters"],
    bankName : "",
    routingNumber : "",
    accountNumber : "",
    // accountType : "",
    empID : 0,
  }),

  methods: {
    cancel(){
      this.bankName = ""
      this.routingNumber = ""
      this.accountNumber = ""
      this.closePopup = false
    },
    submit() {
      const requestObj = {
        "EmployeeID" : parseInt(this.empID),
        "Bank" : this.bankName,
        "RoutingNumber" : this.routingNumber,
        "AccountNumber" : this.accountNumber,
      }
      // closePopup = false
      this.$axios.post("http://localhost:8080/users/UpdateEmployeeInfo2", requestObj)
        .then(response => {
            console.log(response)

            //Clearing the financial input after successful submission
            this.bankName = ""
            this.routingNumber = ""
            this.accountNumber = ""

            this.$emit('notif', 'Financial form submitted successfully.', "success")
            this.closePopup = false
            // this.$router.push('/landing')
            
        }).catch(error => {
          console.log(error)

          //Clearing the financial input after failed submission
          this.bankName = ""
          this.routingNumber = ""
          this.accountNumber = ""
          this.$emit('notif', 'Failed to submit the financial form', "error")
          this.closePopup = false
        })
    }
  }
};
</script>