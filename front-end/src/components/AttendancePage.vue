<template>
  <div class="home">
    <v-container>
      <NavigationBar v-bind:userName="this.$store.state.userName"/>
      <div>
        <v-alert
            :value="alert"
            :type="alertType"
            outlined
            prominent
            border="left"
            transition="fade-transition"
        >
        <!-- type="warning" -->
        <!-- dismissible -->
        {{alertMessage}}
        </v-alert>
      </div>
      <v-layout class="d-flex">
          <v-row>
              <v-col>
                  <v-card elevation="4">
                      <v-card-title class="justify-center">
                          <strong id="headingSubmit"> Submit Attendance </strong>
                      </v-card-title>

                        <v-divider></v-divider>

                        <v-date-picker
                            id="datePickingSubmit"
                            v-model="date"
                            scrollable
                            no-title
                            full-width
                        >
                            <!-- scrollable -->
                            <!-- no-title -->
                            <!-- multiple -->
                            <!-- <v-spacer></v-spacer> -->
                            <!-- <v-btn
                                text
                                color="primary"
                                @click="menu = false"
                            >
                                Cancel
                            </v-btn>
                            <v-btn
                                text
                                color="primary"
                                @click="$refs.menu.save(date)"
                            >
                                OK
                            </v-btn> -->

                        </v-date-picker>

                        <v-divider></v-divider>

                        <v-row>
                        <v-col class="pl-4">
                            <v-dialog
                                id="startTimeSubmit"
                                ref="dialog"
                                v-model="startTimeFlag"
                                :return-value.sync="startTime"
                                persistent
                                width="290px"
                            >
                                <template v-slot:activator="{ on, attrs }">
                                    <v-text-field
                                    v-model="startTime"
                                    label="Enter start time"
                                    prepend-icon="mdi-clock-time-four-outline"
                                    readonly
                                    v-bind="attrs"
                                    v-on="on"
                                    ></v-text-field>
                                </template>
                                <v-time-picker
                                    v-if="startTimeFlag"
                                    v-model="startTime"
                                    full-width
                                >
                                    <v-spacer></v-spacer>
                                    <v-btn
                                    text
                                    color="primary"
                                    @click="startTimeFlag = false"
                                    >
                                    Cancel
                                    </v-btn>
                                    <v-btn
                                    text
                                    color="primary"
                                    @click="$refs.dialog.save(startTime)"
                                    >
                                    OK
                                    </v-btn>
                                </v-time-picker>
                            </v-dialog>
                        </v-col>

                        <v-col class="pr-4">
                            <v-dialog
                                id="endTimeSubmit"
                                ref="dialog1"
                                v-model="endTimeFlag"
                                :return-value.sync="endTime"
                                persistent
                                width="290px"
                                >
                                <template v-slot:activator="{ on, attrs }">
                                    <v-text-field
                                    v-model="endTime"
                                    label="Enter end time"
                                    prepend-icon="mdi-clock-time-four-outline"
                                    readonly
                                    v-bind="attrs"
                                    v-on="on"
                                    ></v-text-field>
                                </template>
                                <v-time-picker
                                    v-if="endTimeFlag"
                                    v-model="endTime"
                                    full-width
                                >
                                    <v-spacer></v-spacer>
                                    <v-btn
                                    text
                                    color="primary"
                                    @click="endTimeFlag = false"
                                    >
                                    Cancel
                                    </v-btn>
                                    <v-btn
                                    text
                                    color="primary"
                                    @click="$refs.dialog1.save(endTime)"
                                    >
                                    OK
                                    </v-btn>
                                </v-time-picker>
                            </v-dialog>
                        </v-col>
                        </v-row>

                        <v-row>
                            <v-col class="py-0 mb-4 text-center">
                                <v-btn id="resetButton" @click="resetFields()">Reset </v-btn>
                            </v-col>
                            <v-col class="py-0 mb-4 text-center">
                                <v-btn id="submitButton" @click="submit()">Submit</v-btn>
                            </v-col>
                        </v-row>
                  </v-card>
                  
              </v-col>

              <v-col>
                  <!-- stats -->
                  <v-card elevation="4">
                    <v-card-title class="justify-center">
                        <strong> Current Attendance Stats </strong>
                    </v-card-title>

                    <v-divider></v-divider>

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
                                    v-model="filterStartDate"
                                    label="Filter start date"
                                    prepend-icon="mdi-calendar"
                                    readonly
                                    v-bind="attrs"
                                    v-on="on"
                                ></v-text-field>
                                </template>
                                <v-date-picker
                                v-model="filterStartDate"
                                :active-picker.sync="activePicker"
                                min="1950-01-01"
                                ></v-date-picker>
                                <!-- :max="(new Date(Date.now() - (new Date()).getTimezoneOffset() * 60000)).toISOString().substr(0, 10)" -->
                                <!-- @change="save" -->
                            </v-menu>
                        </v-col>

                        <v-col class="mr-4 pb-0">
                            <v-menu
                                id="filterEndDate"
                                ref="menu2"
                                v-model="menu2"
                                :close-on-content-click="false"
                                transition="scale-transition"
                                offset-y
                                min-width="auto"
                            >
                                <template v-slot:activator="{ on, attrs }">
                                <v-text-field
                                    v-model="filterEndDate"
                                    label="Filter end date"
                                    prepend-icon="mdi-calendar"
                                    readonly
                                    v-bind="attrs"
                                    v-on="on"
                                ></v-text-field>
                                </template>
                                <v-date-picker
                                v-model="filterEndDate"
                                :active-picker.sync="activePicker"                              
                                min="1950-01-01"                               
                                ></v-date-picker>
                                <!-- :max="(new Date(Date.now() - (new Date()).getTimezoneOffset() * 60000)).toISOString().substr(0, 10)" -->
                                <!-- @change="save" -->
                            </v-menu>
                        </v-col>
                    </v-row>

                    <v-row class="py-0 my-0 mr-2 text-right">
                        <v-col>
                            <v-btn id="filterButton" @click="filter()"> Filter </v-btn>
                        </v-col>
                    </v-row>

                    <v-divider></v-divider>

                    <v-row>
                        <v-card-subtitle class="justify-center">
                            <strong> Filtered Stats </strong>
                            <!-- Next code goes here -->
                        </v-card-subtitle>

                        <v-card-text class="justify-left">
                            <ul class="justify-left">
                                <li>Average Hours : {{averageWorkHour}}</li>
                                <li>Number of Holidays : {{holidays}}</li>
                                <li>Days Worked : {{presentDays}}</li>
                                <li>Days Absent : {{absentDays}}</li>
                                <li>Total Regular Time : {{totalRegularHours}}</li>
                                <li>Total Overtime : {{totalOvertimeHours}}</li>
                            </ul>
                        </v-card-text>

                        <v-divider></v-divider>

                        <!-- <v-card-text class="text-h6 text-justify">
                            <ul>
                                <li v-for="cat in categories" :key="cat.Id">{{cat.message}} : {{}}
                            </ul>
                        </v-card-text> -->
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
import PageBottom from './PageBottom.vue'

export default {
  name: 'attendancePage',
  components: {
    NavigationBar, 
    PageBottom,
  },
  data: () => ({
    startTime : null,
    endTime : null,
    startTimeFlag : false,
    endTimeFlag : false,
    date : null,
    filterStartDate : null,
    filterEndDate : null,
    alert : false,
    alertType : "",
    alertMessage : "",
    totalDays: 0,
    holidays : 0,
    presentDays : 0,
    averageWorkHour : 0,
    absentDays : 0,
    totalWorkHours : 0,
    totalOvertimeHours : 0,
    totalRegularHours : 0,
  }),

  methods : {
        resetFields() {
            this.date = null
            this.startTime = null
            this.endTime = null
        },

        submit() {
            const reqObj = {
                "Date" : this.date,
                "StartTime" : this.startTime,
                "EndTime" : this.endTime,
            }

            this.$axios.post("http://localhost:8080/working/insertWorkingRecord", reqObj)
                .then(response => {
                    console.log(response)
                    this.alertMessage = "Successfully submitted the hours worked on " + this.date + " from " + this.startTime + " to " + this.endTime + ".";
                    this.alert = true
                    this.alertType = "success"
                    console.log(this.alertMessage)
                    // this.$forceUpdate()
                    setTimeout(() => {
                        this.alert = false;
                        this.alertType = "success"
                        // this.alertMessage="Default message is this!"
                        // this.$forceUpdate();
                    }, 7000)
                    
                })
        },

        filter() {
            this.$axios.get("http://localhost:8080/working/getWorkingDetails"+"?StartDate="+this.filterStartDate+"&EndDate="+this.filterEndDate)
                .then(response => {
                    // console.log(response)
                    let respObj = response.data.data
                    this.totalDays = respObj.totaldays
                    this.holidays = respObj.holidays
                    this.presentDays = respObj.presentDays
                    this.absentDays = respObj.absentDays
                    this.averageWorkHour = respObj.averageWorkHour
                    this.totalWorkHours = respObj.totalWorkHour
                    this.totalOvertimeHours = respObj.totalOvertimeHour
                    this.totalRegularHours = respObj.totalRegularHour
                    console.log(this.totalRegularHours)
                })
        }
  }
}
</script>




