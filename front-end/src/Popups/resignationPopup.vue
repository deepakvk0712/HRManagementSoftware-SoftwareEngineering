<template>
  <v-row justify="center">
    <v-dialog v-model="closePopup" persistent max-width="800px">
      <template v-slot:activator="{ on }">
        <v-btn v-on="on" text>
          <span style="margin-right: 10px"> Resign </span>
          <v-icon color="#D22B2B" justify="center"> people </v-icon>
        </v-btn>
      </template>
      <v-card>
        <v-card-title>
          <span id="finCardHead" class="headline">Resignation Form</span>
        </v-card-title>
        <v-form class="px-3" ref="form">
          <v-card-text>
            <v-card-text>
                On a scale of 1-5, with 5 being very good, how much do you rate the training you were given?
                <div class="text-center mt-12">
                <v-rating
                    id="train-rate"
                    length="5"
                    v-model="trainingRating"
                    color="yellow darken-3"
                    background-color="grey darken-1"
                    empty-icon="$ratingFull"
                    half-increments
                    hover
                    large
                ></v-rating>
                </div>
            </v-card-text>
            <v-card-text>
                On a scale of 1-5, with 5 being completely, do you believe this company helped you grow in your career?
                <div class="text-center mt-12">
                <v-rating
                    id="growth-rating"
                    length="5"
                    v-model="growthRating"
                    color="yellow darken-3"
                    background-color="grey darken-1"
                    empty-icon="$ratingFull"
                    half-increments
                    hover
                    large
                ></v-rating>
                </div>
            </v-card-text>
            <v-card-text>
                On a scale of 1-5, with 5 being certainly would recommend, would you recommend a friend to join this company?
                <div class="text-center mt-12">
                <v-rating
                    id="recommend-other"
                    length="5"
                    v-model="recommendOthers"
                    color="yellow darken-3"
                    background-color="grey darken-1"
                    empty-icon="$ratingFull"
                    half-increments
                    hover
                    large
                ></v-rating>
                </div>
            </v-card-text>
            <v-card-text>
                On a scale of 1-5, with 5 being excellent, how much would you rate the product we make?
                <div class="text-center mt-12">
                <v-rating
                    id="product-rating"
                    length="5"
                    v-model="productRating"
                    color="yellow darken-3"
                    background-color="grey darken-1"
                    empty-icon="$ratingFull"
                    half-increments
                    hover
                    large
                ></v-rating>
                </div>
            </v-card-text>
            <v-card-text>
                On a scale of 1-5, with 5 being excellent, how much would you rate your stay here at this company?
                <div class="text-center mt-12">
                <v-rating
                    id="stay-rating"
                    length="5"
                    v-model="stayRating"
                    color="yellow darken-3"
                    background-color="grey darken-1"
                    empty-icon="$ratingFull"
                    half-increments
                    hover
                    large
                ></v-rating>
                </div>
            </v-card-text>
            <v-text-field
              id="resignFeedback"
              label="Leave any feedback you have about your time here with us!"
              v-model="feedback"
              prepend-icon="edit"
              :rules="inputRules"
            ></v-text-field>
            <v-text-field
              id="changeMind"
              label="Could we have doen anything different to change your mind?"
              v-model="changeMind"
              prepend-icon="edit"
              :rules="inputRules"
            ></v-text-field>
            <v-checkbox
                id="agree-checkbox"
                v-model="agreeCheckbox"
                :label="`Do you wish to go ahead and submit you resignation?`"
            ></v-checkbox>

            <v-checkbox
                id="document-checkbox"
                v-model="documentCheckbox"
                :label="`Do you agree to submit all documents requested to relieve you permanently?`"
            ></v-checkbox>
            
            <v-checkbox
                id="terms-checkbox"
                v-model="termsCheckbox"
                :label="`By checking this box, I agree to all the terms and conditions and policy of the company.`"
            ></v-checkbox>

            <v-row>
                <v-col
                    cols="12"
                    sm="6"
                >
                    <v-text-field
                        id="full-name"
                        v-model="fullName"
                        label="Enter Full Name"
                        outlined
                    ></v-text-field>
                </v-col>
                <v-col
                    cols="12"
                    sm="6"
                >
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
                            v-model="currentDate"
                            label="Enter Today's Date"
                            prepend-icon="mdi-calendar"
                            readonly
                            v-bind="attrs"
                            v-on="on"
                        ></v-text-field>
                        </template>
                        <v-date-picker
                        v-model="currentDate"
                        :active-picker.sync="activePicker"                              
                        min="1950-01-01"                               
                        ></v-date-picker>
                    </v-menu>
                </v-col>
            </v-row>
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
              @click="resign()"
              >Resign</v-btn
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
    feedback : "",
    changeMind : "",
    trainingRating : 5,
    growthRating : 5,
    recommendOthers : 5,
    productRating : 5,
    stayRating : 5,
    agreeCheckbox : false,
    documentCheckbox :  false,
    termsCheckbox : false,
    currentDate : null,
    fullName : "",
  }),

  methods: {
      cancel(){
        this.trainingRating = 5
        this.growthRating = 5
        this.recommendOthers = 5
        this.productRating = 5
        this.stayRating = 5
        this.feedback = ""
        this.currentDate = null
        this.termsCheckbox = false
        this.documentCheckbox = false
        this.agreeCheckbox = false
        this.changeMind = ""
        this.fullName = ""
        this.closePopup = false
      },
      resign(){
        // console.log("resignation call")
        const requestObj = {
          "q1" : this.trainingRating,
          "q2" : this.growthRating,
          "q3" : this.recommendOthers,
          "q4" : this.productRating,
          "q5" : this.stayRating,
          "feedback" : this.feedback,
          "resigndate" : this.currentDate
        }
        // console.log(requestObj)
        this.$axios.post("http://localhost:8080/resign/insertFeedback", requestObj)
          .then(response => {
              console.log(response)

              this.trainingRating = 5
              this.growthRating = 5
              this.recommendOthers = 5
              this.productRating = 5
              this.stayRating = 5
              this.feedback = ""
              this.currentDate = null
              this.termsCheckbox = false
              this.documentCheckbox = false
              this.agreeCheckbox = false
              this.changeMind = ""
              this.fullName = ""
              
              this.$emit('notif', 'Successfully submitted your resignation', "success")
              this.closePopup = false
                                
          }).catch(error => {
              console.log(error)
              this.trainingRating = 5
              this.growthRating = 5
              this.recommendOthers = 5
              this.productRating = 5
              this.stayRating = 5
              this.feedback = ""
              this.currentDate = null
              this.termsCheckbox = false
              this.documentCheckbox = false
              this.agreeCheckbox = false
              this.changeMind = ""
              this.fullName = ""
              this.$emit('notif', 'Failed to submit resignation to HR department!', "error")
              this.closePopup = false
          })
      }
  }
};
</script>