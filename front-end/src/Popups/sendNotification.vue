<template>
  <v-row justify="right">
      <!-- justify="center" -->
    <v-dialog v-model="closePopup" persistent max-width="800px">
      <template v-slot:activator="{ on }">
        <v-btn v-on="on" block text> New Notification
          <!-- <span style="margin-right: 10px"> New Notification </span> -->
          <!-- <v-icon color="#D22B2B"> folder_shared </v-icon> -->
        </v-btn>
      </template>
      <v-card>
        <v-card-title>
          <span id="finCardHead" class="headline">Send Notification</span>
        </v-card-title>
        <v-form class="px-3" ref="form">
          <v-card-text>
            <v-col class="d-flex" cols="12" sm="12">
                <v-select
                @change="getTeamMembers()"
                v-if="this.$store.state.isHR"
                id="hrBuissnessUnit"
                :items="businessUnits"
                label="Select Buissness Unit"
                dense
                outlined
                v-model="selectedBuissnessUnit"
                required
                ></v-select>
            </v-col>
            <v-col class="d-flex" cols="12" sm="12">
                <v-select
                id="teamMembersNames"
                :items="teamMembers"
                label="Select the team member to notify"
                dense
                outlined
                @change="selectedMember = data.item.emailID"
                v-model="selectedMember"
                required
                >
                    <template slot="selection" slot-scope="data">
                        {{ data.item.name }} - {{ data.item.emailID }}
                    </template>
                    <template slot="item" slot-scope="data">
                        {{ data.item.name }} - {{ data.item.emailID }}
                    </template>
                </v-select>
            </v-col>
            <v-text-field
              id="message"
              label="Enter the notification to send"
              v-model="message"
              prepend-icon="edit"
              :rules="inputRules"
            ></v-text-field>
          </v-card-text>
          
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn id="finCancelButton" color="grey darken-1" text @click="closePopup = false"
              >Cancel</v-btn
            >
            <v-btn
              id = "finSubmitButton"
              color="green darken-1"
              text
              outlined
              @click="sendNotification()"
              >Send</v-btn
            >
          </v-card-actions>
        </v-form>
      </v-card>
    </v-dialog>
  </v-row>
</template>

<script>
import { mapState } from 'vuex'
export default {
    computed: {
        ...mapState({
            teamMembers: 'teamMembers',
            businessUnits : 'businessUnits'
        })
    },
    data: () => ({
        closePopup: false,
        // businessUnits : [],
        // teamMembers : [],
        selectedMember : "",
        message : "",
        selectedBuissnessUnit : "",
        empID : 0,
    }),

    methods: {
        getTeamMembers() {
            this.$axios.get("http://localhost:8080/notify/teamMembers?businessUnit=" + this.selectedBuissnessUnit)
            .then(response => {
                console.log(response)
                let respObj = JSON.parse(response.data.data)
                let tMembers = []
                for (var i=0; i < respObj.teamMembers.length; i++) {
                    let cat = {}
                    cat.name = respObj.teamMembers[i].name
                    cat.emailID = respObj.teamMembers[i].emailID
                    tMembers.push(cat)
                }
                
                this.$store.state.teamMembers = tMembers
                console.log("other team member business unit: ")
                console.log(this.$store.state.teamMembers)         
            })
        },
        sendNotification(){
            console.log(this.selectedMember.emailID)
            const reqObj = {
                'receiver' : this.selectedMember.emailID,
                'message' : this.message
            }

            this.$axios.post("http://localhost:8080/notify/", reqObj)
            .then(response => {
                console.log(response)
                // let respObj = JSON.parse(response.data.data)
            })
            this.$emit('notif', 'A new notification has been sent to ' + this.selectedMember.emailID, "success")
            this.closePopup = false
        }
    }
};
</script>