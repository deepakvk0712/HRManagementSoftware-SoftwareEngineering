<template>
  <div class="team">
    <v-container>
      <NavigationBar v-bind:userName="this.$store.state.userName"/>

      <v-layout class="mt-4" row wrap>
        <v-divider></v-divider>
      </v-layout>

      <div class="mt-4">
          
        <v-container
        class="grey lighten-5 mb-6"
        >
          <v-row
          :align="align"
          no-gutters
          style="height: auto;"
          >
            <v-col class="col-6" align-center>
              <p id="quicklink2" class="text-left font-weight-black">
                Notifications
              </p>
            </v-col>

            <v-col class="col-6 text-right">
              <v-card-actions>
                <SendNotificationPopup/>
                <!-- style="margin:15px;" -->
              </v-card-actions>
            </v-col>
          </v-row>
          <v-layout class="mt-4" row wrap>
            <v-divider></v-divider>
          </v-layout>
          <v-row v-if="$store.state.notifications.length == 0" class="mt-4 text-center">
            No new notifications to show!
          </v-row>
          <!-- The message displaying code goes here. -->
          <v-row
          v-for="item in $store.state.notifications" :key="item.messageID"
          class="mt-4"
          :align="align"
          no-gutters
          style="height: auto;"
          >
            <v-col class="col-10" align-center>
              <p id="quicklink2" class="text-left font-weight-black">
                {{item.sender}} :   {{item.message}}
              </p>
            </v-col>

            <v-col class="col-2 text-right">
              <v-btn block @click="markRead(item.messageID)"> Mark Read </v-btn>
            </v-col>
          </v-row>
        </v-container>
      </div>

      <v-layout class="mt-4 mb-4" row wrap>
        <v-divider></v-divider>
      </v-layout>

      <div v-if="!this.$store.state.isOnboard || !this.$store.state.isFinance">
        <p id="quicklink1" class="h1 text-left font-weight-black">
          Required Forms
        </p>
      </div>
      <v-layout v-if="!this.$store.state.isOnboard || !this.$store.state.isFinance" row wrap>
        <v-divider></v-divider>
      </v-layout>

      <v-layout row wrap>
        <v-flex v-if="!this.$store.state.isOnboard" sm6 xs12 md6 lg3>
          <v-card class="ma-3">
            <v-list-item>
              <v-list-item-avatar tile class="mt-n7">
                <v-sheet color="white" elevation="10">
                  <v-icon color="#D22B2B" style="font-size: 42px; margin-top: 10px" large>feed</v-icon>
                </v-sheet>
              </v-list-item-avatar>
              <v-list-item-content>
                <div id="onboardForm" class="text-right mb-1" style="font-size: 20px">Onboarding Form</div>
                <v-list-item-title id="subOnboardForm" class="text-right"
                  >Employee Input Needed</v-list-item-title
                >
                <div><v-divider></v-divider></div>
              </v-list-item-content>
            </v-list-item>
            <v-card-actions>
              <OnboardingFormPopup style="margin:15px;"/>
            </v-card-actions>
          </v-card>
        </v-flex>

        <v-flex v-if="!this.$store.state.isFinance" sm6 xs12 md6 lg3>
          <v-card class="ma-3">
            <v-list-item>
              <v-list-item-avatar tile class="mt-n7">
                <v-sheet color="white" elevation="10">
                  <v-icon color="#D22B2B" style="font-size: 44px; margin-top: 10px" large>account_balance</v-icon>
                </v-sheet>
              </v-list-item-avatar>
              <v-list-item-content>
                <div id="finForm" class="text-right mb-1" style="font-size: 20px">Financial Form</div>
                <v-list-item-title id="subFinForm" class="text-right"
                  >Banking Information</v-list-item-title
                >
                <div><v-divider></v-divider></div>
              </v-list-item-content>
            </v-list-item>
            <v-card-actions>
              <FinancialFormPopup style="margin:15px;"/>
            </v-card-actions>
          </v-card>
        </v-flex>
        
        <!-- <v-flex xs12 sm6 md4 lg3 v-for="person in team" :key="person.name">
          <v-card class="text-center ma-3">
            <v-responsive class="pt-4">
              <v-avatar size="100" class="red lighten-2">
                <img :src="person.avatar" alt="" />
              </v-avatar>
            </v-responsive>
            <v-card-text>
              <div class="subheading">{{ person.name }}</div>
              <div class="grey--text">{{ person.role }}</div>
            </v-card-text>
            <v-card-actions>
              <v-btn outlined color="orange">
                <v-icon small left>message</v-icon>
                <span>Message</span>
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-flex> -->
      </v-layout>

      <v-layout v-if="!this.$store.state.isOnboard || !this.$store.state.isFinance" class="mt-4" row wrap>
        <v-divider></v-divider>
      </v-layout>

      <div class="mt-4">
        <p id="quicklink2" class="text-left font-weight-black">
          Quick Links
        </p>
      </div>
      <v-layout row wrap>
        <v-flex v-if="this.$store.state.isHR" sm6 xs12 md6 lg3>
          <!-- Add the check to show this form element only if the account type is admin -->
          <v-card class="ma-3">
            <v-list-item>
              <v-list-item-avatar tile class="mt-n7">
                <v-sheet color="white" elevation="10">
                  <v-icon color="#D22B2B" style="font-size: 44px; margin-top: 10px" large>person</v-icon>
                </v-sheet>
              </v-list-item-avatar>
              <v-list-item-content>
                <div id="regEmpHead" class="text-right mb-1" style="font-size: 20px">Register Employee</div>
                <v-list-item-title id="subRegEmpHead" class="text-right"
                  >Register new employee</v-list-item-title
                >
                <div><v-divider></v-divider></div>
              </v-list-item-content>
            </v-list-item>
            <v-card-actions>
              <RegisterEmployee style="margin:15px;"/>
            </v-card-actions>
          </v-card>
        </v-flex>

        <v-flex sm6 xs12 md6 lg3>
          <v-card class="ma-3" >
            <v-list-item>
              <v-list-item-avatar tile class="mt-n7">
                <v-sheet color="white" elevation="10">
                  <v-icon color="#D22B2B" style="font-size: 48px; margin-top: 15px" large
                    >person</v-icon
                  >
                </v-sheet>
              </v-list-item-avatar>
              <v-list-item-content>
                <div class="text-right mb-1" style="font-size: 20px">
                  Leave Information
                </div>
                <v-list-item-title class="text-right"
                  >Paid Leave: {{ paidLeavesRemaining }}/{{
                    maxPaidLeave
                  }}</v-list-item-title
                >
                <v-list-item-title class="text-right"
                  >Un-Paid Leave: {{ unpaidLeavesRemaining }}/{{
                    maxUnpaidLeaves
                  }}</v-list-item-title
                >
                <div><v-divider></v-divider></div>
              </v-list-item-content>
            </v-list-item>
            <v-card-actions>
              <ApplyLeavePopup style="margin:5px;"/>
            </v-card-actions>
          </v-card>
        </v-flex>
      </v-layout>

      <v-layout class="mt-4" row wrap>
        <v-divider></v-divider>
      </v-layout>
      
      <PageBottom/>
    </v-container>
  </div>
</template>

<script>
import ApplyLeavePopup from "../Popups/ApplyLeavePopup.vue"
import OnboardingFormPopup from '../Popups/OnboardingFormPopup.vue'
import FinancialFormPopup from '../Popups/FinancialFormPopup.vue'
import NavigationBar from './NavigationBar.vue'
import PageBottom from './PageBottom.vue'
import RegisterEmployee from '../Popups/HRRegisterEmployee.vue'
import SendNotificationPopup from '../Popups/sendNotification.vue'

export default {
  name: "LandingPage",
  async mounted() {
    await this.$axios.get('http://localhost:8080/dashboard/')
      .then(response => {
        // console.log(response);
        let respObj = JSON.parse(response.data.data)
        this.$store.state.userName = respObj.username
        // this.$store.state.accountType = respObj.accountType
        this.$store.state.isManager = respObj.isManager
        this.$store.state.isHR = respObj.isHR
        // this.$store.state.teamMembers = respObj.teamMembers
        this.$store.state.businessUnits = respObj.businessUnits
        this.$store.state.isOnboard = respObj.isOnboard
        this.$store.state.isFinance = respObj.isFinance

        let tMembers = []
        for (let i=0; i < respObj.teamMembers.length; i++) {
          let cat = {}
          cat.name = respObj.teamMembers[i].name
          cat.emailID = respObj.teamMembers[i].emailID
          tMembers.push(cat)
        }
        
        this.$store.state.teamMembers = tMembers
        if(respObj.messages == null){
          let nMembers = []
          this.$store.state.notifications = nMembers
        }
        else{ 
          let nMembers = []
          for (let i=0; i < respObj.messages.length; i++) {
            let cat = {}
            cat.sender = respObj.messages[i].sender
            cat.message = respObj.messages[i].message
            cat.messageID = respObj.messages[i].messageID
            nMembers.push(cat)
          }
          this.$store.state.notifications = nMembers
        }

        // console.log(this.$store.state.teamMembers)
        // this.$store.state.accountType = false
        // this.$store.state.isManager = false
    
        // console.log(this.$store.state.userName);
      })
  },
  components: {
    ApplyLeavePopup, OnboardingFormPopup, FinancialFormPopup, NavigationBar, PageBottom, RegisterEmployee, SendNotificationPopup,
  },
  data: () => ({
    maxPaidLeave: 20,
    paidLeavesRemaining: 12,
    maxUnpaidLeaves: 20,
    unpaidLeavesRemaining: 14,
    isLoading:false,
    userName : this.$store.state.userName,

  }),

  methods: {
    markRead(messageID){
      console.log(messageID)
      this.$axios.put("http://192.168.0.35:8080/notify/markRead?id=" + messageID)
        .then(response => {
            console.log(response)
            // let respObj = JSON.parse(response.data.data)
        })
    }
  }
};
</script>