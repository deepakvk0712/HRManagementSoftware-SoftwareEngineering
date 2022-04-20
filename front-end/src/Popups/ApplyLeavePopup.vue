<template>
  <v-row justify="center">
    <v-dialog v-model="closePopup" persistent max-width="800px">
      <template v-slot:activator="{ on }">
        <v-btn v-on="on" text>
          <span style="margin-right: 10px"> Apply Leave </span>
          <v-icon color="#D22B2B" justify="center"> exit_to_app </v-icon>
        </v-btn>
      </template>
      <v-card>
        <v-card-title>
          <span id="leaveApplyHead" class="headline">Leave Application</span>
        </v-card-title>
        <v-form class="px-3" ref="form">
          <v-card-text>
            <v-text-field
              id="leaveApplySubjectInput"
              label="Subject"
              v-model="leaveApplicationTitle"
              prepend-icon="exit_to_app"
              :rules="inputRules"
            ></v-text-field>
            <v-textarea
              id="leaveApplyReasonInput"
              label="Please state reason for requesting leave."
              v-model="leaveRequestDescription"
              prepend-icon="edit"
              :rules="inputRules"
            ></v-textarea>
            <!-- <v-col cols="12" lg="6">
              <v-menu
                id="leaveApplyInput"
                ref="menu1"
                v-model="menu1"
                :close-on-content-click="false"
                transition="scale-transition"
                offset-y
                max-width="290px"
                min-width="290px"
              >
                <template v-slot:activator="{ on }">
                  <v-text-field
                    v-model="fromDateFormatted"
                    label="From Date"
                    hint="MM/DD/YYYY format"
                    persistent-hint
                    prepend-icon="event"
                    @blur="date = parseDate(fromDateFormatted)"
                    v-on="on"
                  ></v-text-field>

                  <v-text-field
                    v-model="toDateFormatted"
                    label="To Date"
                    hint="MM/DD/YYYY format"
                    persistent-hint
                    prepend-icon="event"
                    @blur="date = parseDate(toDateFormatted)"
                    v-on="on"
                  ></v-text-field>
                </template>
                <v-date-picker
                  id="leaveApplyEndDateInput"
                  v-model="date"
                  no-title
                  @input="menu1 = false"
                ></v-date-picker>
              </v-menu>
            </v-col> -->


            <v-row>
              <v-col class="ml-4 pb-0">
                  <v-menu
                      id="filterStartDate"
                      ref="menu1"
                      v-model="menu1"
                      :close-on-content-click="false"
                      transition="scale-transition"
                      offset-y
                      min-width="auto"
                  >
                      <template v-slot:activator="{ on, attrs }">
                      <v-text-field
                          v-model="leaveStartDate"
                          label="Leave start date"
                          prepend-icon="mdi-calendar"
                          readonly
                          v-bind="attrs"
                          v-on="on"
                      ></v-text-field>
                      </template>
                      <v-date-picker
                      v-model="leaveStartDate"
                      :active-picker.sync="activePicker"
                      min="1950-01-01"
                      
                      ></v-date-picker>
                      <!-- :max="(new Date(Date.now() - (new Date()).getTimezoneOffset() * 60000)).toISOString().substr(0, 10)" -->
                      <!-- @change="save" -->
                  </v-menu>
              </v-col>

              <v-col class="mr-4 pb-0">
                  <v-menu
                      id="leaveEndDate"
                      ref="menu2"
                      v-model="menu2"
                      :close-on-content-click="false"
                      transition="scale-transition"
                      offset-y
                      min-width="auto"
                  >
                      <template v-slot:activator="{ on, attrs }">
                      <v-text-field
                          v-model="leaveEndDate"
                          label="Leave end date"
                          prepend-icon="mdi-calendar"
                          readonly
                          v-bind="attrs"
                          v-on="on"
                      ></v-text-field>
                      </template>
                      <v-date-picker
                      id="leaveApplyEndDateInput"
                      v-model="leaveEndDate"
                      :active-picker.sync="activePicker"
                      min="1950-01-01"
                      ></v-date-picker>
                      <!-- :max="(new Date(Date.now() - (new Date()).getTimezoneOffset() * 60000)).toISOString().substr(0, 10)" -->
                      <!-- @change="save" -->
                  </v-menu>
              </v-col>
          </v-row>
          </v-card-text>
          <v-col class="d-flex" cols="12" sm="6">
            <v-select
              id="leaveApplyLeaveTypeInput"
              :items="dropdownItems"
              label="Leave Type"
              dense
              outlined
              v-model="leaveTypeSelection"
            ></v-select>
          </v-col>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn id="leaveApplyCancelInput" color="grey darken-1" text @click="cancel()"
              >Cancel</v-btn
            >
            <v-btn
              id="leaveApplySubmitInput"
              color="green darken-1"
              text
              outlined
              @click="apply()"
              >Apply</v-btn
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
    dropdownItems:[
        'Paid Leave', 'Unpaid Leave'
    ],
    leaveTypeSelection:"",
    closePopup: false,
    leaveApplicationTitle: "",
    leaveRequestDescription: "",
    leaveStartDate : null,
    leaveEndDate : null,
    menu1: false,
    inputRules: [(v) => v.length >= 3 || "Minimum lenght is 3 charachters"],
  }),
  methods: {
    cancel(){
      this.leaveApplicationTitle = ""
      this.leaveRequestDescription = ""
      this.leaveTypeSelection = ""
      this.leaveStartDate = null
      this.leaveEndDate = null
      this.closePopup = false
    },
    apply(){
      console.log(this.leaveTypeSelection)
      const reqObj = {
        "subject" : this.leaveApplicationTitle,
        "reason" : this.leaveRequestDescription,
        "leaveType" : this.leaveTypeSelection,
        "startDate" : this.leaveStartDate,
        "endDate" : this.leaveEndDate
      }

      this.$axios.post("http://localhost:8080/leave/apply", reqObj)
        .then(response => {
            console.log(response)

            //Clearing the input after successful request
            this.leaveApplicationTitle = ""
            this.leaveRequestDescription = ""
            this.leaveTypeSelection = ""
            this.leaveStartDate = null
            this.leaveEndDate = null

            this.$emit('notif', 'Requested ' + this.leaveTypeSelection +' from ' + this.leaveStartDate + ' to ' + this.leaveEndDate + ' submitted', "success")
            this.closePopup = false            
        }).catch(error => {
            console.log(error)

            //Clearing the input after failed request
            this.leaveApplicationTitle = ""
            this.leaveRequestDescription = ""
            this.leaveTypeSelection = ""
            this.leaveStartDate = null
            this.leaveEndDate = null
            this.$emit('notif', 'Failed to send a leave request to the HR department', "error")
            this.closePopup = false
        })
    }
  },
};
</script>