<template>
  <div class="team">
    <v-container>
      <NavigationBar/>
      <v-layout row wrap>
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
        <v-flex sm6 xs12 md6 lg3>
          <v-card class="ma-3">
            <v-list-item>
              <v-list-item-avatar tile class="mt-n7">
                <v-sheet color="white" elevation="10">
                  <v-icon color="#D22B2B" style="font-size: 42px; margin-top: 10px" large>feed</v-icon>
                </v-sheet>
              </v-list-item-avatar>
              <v-list-item-content>
                <div class="text-right mb-1" style="font-size: 20px">Onboarding Form</div>
                <v-list-item-title class="text-right"
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

        <v-flex sm6 xs12 md6 lg3>
          <v-card class="ma-3">
            <v-list-item>
              <v-list-item-avatar tile class="mt-n7">
                <v-sheet color="white" elevation="10">
                  <v-icon color="#D22B2B" style="font-size: 44px; margin-top: 10px" large>account_balance</v-icon>
                </v-sheet>
              </v-list-item-avatar>
              <v-list-item-content>
                <div class="text-right mb-1" style="font-size: 20px">Financial Form</div>
                <v-list-item-title class="text-right"
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

export default {
  name: "LandingPage",
  components: {
    ApplyLeavePopup, OnboardingFormPopup, FinancialFormPopup, NavigationBar, PageBottom,
  },
  data: () => ({
    team: [
      { name: "Iyad", role: "web developer", avatar: "/img1.png" },
      { name: "Reda", role: "Graphic designer", avatar: "/img2.png" },
      { name: "Zineb", role: "web developer", avatar: "/img3.png" },
      { name: "Hu TechGroup", role: "Desktop developer", avatar: "/img4.png" },
    ],
    maxPaidLeave: 20,
    paidLeavesRemaining: 12,
    maxUnpaidLeaves: 20,
    unpaidLeavesRemaining: 14,
  }),
};
</script>