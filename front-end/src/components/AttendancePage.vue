<template>
  <div class="home">
    <v-container>
      <NavigationBar v-bind:userName="this.$store.state.userName"/>
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
                                :max="(new Date(Date.now() - (new Date()).getTimezoneOffset() * 60000)).toISOString().substr(0, 10)"
                                min="1950-01-01"
                                @change="save"
                                ></v-date-picker>
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
                                :max="(new Date(Date.now() - (new Date()).getTimezoneOffset() * 60000)).toISOString().substr(0, 10)"
                                min="1950-01-01"
                                @change="save"
                                ></v-date-picker>
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


  }),

  methods : {
        resetFields() {
            this.date = null
            this.startTime = null
            this.endTime = null
        },

        submit() {

            // console.log(typeof this.startTime)
            // console.log(typeof this.endTime)
            // console.log(typeof this.date)
            // console.log(this.startTime + ":" + this.endTime)
            // this.date = new Date(this.date).toISOString().substring(0,10);
            // console.log("Date : " + typeof this.date + " Value : " + this.date);
            // this.startTime = new Date(this.startTime).toISOString().substring(0,10);
            // console.log("Date : " + typeof this.startTime + " Value : " + this.startTime);
            const reqObj = {
                "Date" : this.date,
                "StartTime" : this.startTime,
                "EndTime" : this.endTime,
            }

            // console.log(reqObj)

            this.$axios.post("http://localhost:8080/working/insertWorkingRecord", reqObj)
                .then(response => {
                    console.log(response)
                    // this.closePopup = false
                    console.log("I came till here")
                    // this.$router.push('/landing')
                    
                })
        },

        filter() {
            // console.log(this.filterStartDate + "     " + this.filterEndDate)
            const reqObj = {
                "StartDate" : this.filterStartDate,
                "EndDate" : this.filterEndDate
            }

            this.$axios.get("http://localhost:8080/working/getWorkingDetails", reqObj)
                .then(response => {
                    console.log(response)
                    console.log("got details")
                    // this.closePopup = false
                    // this.$router.push('/landing')
                    
                })
        }
  }
}
</script>