<template id = "formid">
  <form class = "container-fluid">
    <div>
    <h1 class = ma-9 style="color: black"> Welcome to the Training Session </h1>
    </div>

    <div>
    <h4 class = ma-9> Please watch this video and answer the below questions</h4>
    </div>


    <section class= ma-5>
        <iframe width="570" height="380" src="https://www.youtube.com/embed/sB2iQSvrcG0" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen>
        </iframe>
    </section>

    <!-- <video-embed src="https://www.youtube.com/watch?v=sB2iQSvrcG0"></video-embed> -->

    <v-btn id="startID" color = "#6495ED" @click="isHidden = true;click()" class="mb-4">Start Quiz</v-btn>

    <div>
    <h2 class = ma-5> Get 7/10 correct answers to complete training</h2>
    </div>

    <div>
    <h2 class = ma-5> All the Best!</h2>
    </div>

    <!-- <div v-if="isHidden == true && questionIndex < questions.length">
    <div v-for="q in questions" :key="q.rightAnswer"> -->
    
    <v-card class="mx-auto"
      max-width="700" text-align="left">
      <v-img
        src="https://cdn.vuetifyjs.com/images/cards/sunshine.jpg"
        height="200px"
      ></v-img>

      <v-card-actions class = "justify-left">
        <div v-if="isHidden == true && questionIndex < questions.length && trainComp == false ">
          <label class="text-h6">{{ question.question }}</label>
          <div class = "styled" v-for="(c, index) of question.choices" :key="c">
            <input type="radio" name="choice" :value="index+1" v-model="picked" />
          <!-- :value="c" -->
          {{ c }}
          <!-- {{picked}} -->

        <!-- <v-radio-group v-model="radioGroup"> -->

            <!-- <v-radio
                v-model="q1"
                :key="n"
                :label="`${c}`"
                :value="n"
            ></v-radio> -->

            <!-- v-for="n in c" -->
            <!-- {{c}} -->
        <!-- </v-radio-group> -->
        
        <!-- <select v-model="selected" multiple> -->
            <!-- <option>A</option>
            <option>B</option>
            <option>C</option> -->
        <!-- </select> -->

      </div>
    </div>
    </v-card-actions>
    </v-card>
    <!-- </div> -->

    <!-- <div v-else-if="isHidden == true">
      <button type="button" @click="restart">restart</button>
    </div> -->


    <div v-if="isHidden == true">
    <button v-if="questionIndex < 9 && trainComp == false" color = "#6495ED" type="button" :disabled="picked == -1" @click="submit(question)" >Submit</button>
    </div>

   <v-btn v-if="questionIndex == 9" id="finishID" color = "#6495ED" @click="scored()" class="mb-4">Final Submit</v-btn>

    <div v-if="trainComp == true">
      <h2 class = ma-5> Thanks for completing the test! Your score is {{sco}}</h2>
    </div>


    <div v-if="finalSubmit == true && sco <7">
    <button type="button" :disabled="picked == -1" color = "#6495ED" @click="submit(question)">Retake Quiz</button>
    </div>

    <!-- <div v-if="trainComp == false && clicked == true">
      <h2 class = ma-5> Your score is {{sco}}</h2>
    </div> -->
    
    <!-- <h2> {{this.questionIndex}} </h2> 

    <h3>{{this.questions.length}} </h3> -->

    <!-- :disabled="picked == -1" -->
    
    <!-- <div>
        {{answers}}
    </div> -->
    
  </form>
  <!-- <p>score: {{ score }}</p> -->
</template>

<script>
const questions = [
  {
    questionNumber : "questionOne",
    question: "Which is the following is the most important feature of the spiral model?",
    choices: ["Efficiency management", "Time management", "Risk management", "Quality management"],
    // rightAnswer: "Risk management",
  },
  {
    questionNumber : "questionTwo",
    question: "Identify the incorrect phase of the prototype model.",
    choices: ["Prototype refinement", "Engineer product", "Coding", "Quick design"],
    //rightAnswer: "Coding",
  },
  {
    questionNumber : "questionThree",
    question: "The software life cycle can be said to consist of a series of phases. The classical model is referred to as the waterfall model. Which phase defined as “The concept is explored and refined, and the client’s requirements are elicited?",
    choices: ["Requirements", "Specification", "Design", "Implementation"],
    //rightAnswer: Requirements,
  },

  {
    questionNumber : "questionFour",
    question: "Which of the following models will not result in the desired output, when the user participation isn’t involved?",
    choices: ["RAD and protoyping", "Waterfall and prototyping", "Prototyping and spiral", "Spiral and RAD"],
    //rightAnswer: RAD and protoyping,
  },

    {
    questionNumber : "questionFive",
    question: "In the classical chief programmer team approach, the team member responsible for maintaining the detailed design and coding is",
    choices: ["The chief programmer", "The programming secretary", "A specialized function that exists outside ‘the team’", "The individual coder (i.e. programmer)"],
    //rightAnswer: The individual coder (i.e. programmer),
  },

      {
    questionNumber : "questionSix",
    question: "What are the major activities of the spiral model of software engineering?",
    choices: ["Planning, Risk Analysis, Engineering, Customer Evaluation", "Defining, Prototyping, Testing, Delivery", "Requirements", "Quick Design, Build Prototype, Evaluate Prototype, Refine Prototype"],
    //rightAnswer: Planning, Risk Analysis, Engineering, Customer Evaluation,
  },

      {
    questionNumber : "questionSeven",
    question: "A process view in software engineering would consider which of the following",
    choices: ["Product performance", "Staffing", "Functionality", "Reliability"],
    //rightAnswer: Staffing,
  },

      {
    questionNumber : "questionEight",
    question: "Which of the following is not a ‘concern’ during the management of a software project?",
    choices: ["Money", "Time", "Product quality", "Product quantity"],
    //rightAnswer: Product quantity,
  },

      {
    questionNumber : "questionNine",
    question: "Which of the following could be a deliverable for a software system?",
    choices: ["Source Code", "Reference Manual", "Requirements Document", "All of the above"],
    //rightAnswer: All of the above,
  },

      {
    questionNumber : "questionTen",
    question: "Which of the following is not a logical layer of the application in client server system?",
    choices: ["Presentation layer", "Application layer", "Data Management layer", "Programming layer"],
    //rightAnswer: Programming layer,
  },

//       {
//     questionNumber : "questionEleven",
//     question: "In the classical chief programmer team approach, the team member responsible for maintaining the detailed design and coding is",
//     choices: ["The chief programmer", "The programming secretary", "A specialized function that exists outside ‘the team’", "The individual coder (i.e. programmer)"],
//     //rightAnswer: The individual coder (i.e. programmer),
//   },

//       {
//     questionNumber : "questionFive",
//     question: "In the classical chief programmer team approach, the team member responsible for maintaining the detailed design and coding is",
//     choices: ["The chief programmer", "The programming secretary", "A specialized function that exists outside ‘the team’", "The individual coder (i.e. programmer)"],
//     //rightAnswer: The individual coder (i.e. programmer),
//   },

//       {
//     questionNumber : "questionFive",
//     question: "In the classical chief programmer team approach, the team member responsible for maintaining the detailed design and coding is",
//     choices: ["The chief programmer", "The programming secretary", "A specialized function that exists outside ‘the team’", "The individual coder (i.e. programmer)"],
//     //rightAnswer: The individual coder (i.e. programmer),
//   },

//       {
//     questionNumber : "questionFive",
//     question: "In the classical chief programmer team approach, the team member responsible for maintaining the detailed design and coding is",
//     choices: ["The chief programmer", "The programming secretary", "A specialized function that exists outside ‘the team’", "The individual coder (i.e. programmer)"],
//     //rightAnswer: The individual coder (i.e. programmer),
//   },

//       {
//     questionNumber : "questionFive",
//     question: "In the classical chief programmer team approach, the team member responsible for maintaining the detailed design and coding is",
//     choices: ["The chief programmer", "The programming secretary", "A specialized function that exists outside ‘the team’", "The individual coder (i.e. programmer)"],
//     //rightAnswer: The individual coder (i.e. programmer),
//   },

//       {
//     questionNumber : "questionFive",
//     question: "In the classical chief programmer team approach, the team member responsible for maintaining the detailed design and coding is",
//     choices: ["The chief programmer", "The programming secretary", "A specialized function that exists outside ‘the team’", "The individual coder (i.e. programmer)"],
//     //rightAnswer: The individual coder (i.e. programmer),
//   },

//       {
//     questionNumber : "questionFive",
//     question: "In the classical chief programmer team approach, the team member responsible for maintaining the detailed design and coding is",
//     choices: ["The chief programmer", "The programming secretary", "A specialized function that exists outside ‘the team’", "The individual coder (i.e. programmer)"],
//     //rightAnswer: The individual coder (i.e. programmer),
//   },

//       {
//     questionNumber : "questionFive",
//     question: "In the classical chief programmer team approach, the team member responsible for maintaining the detailed design and coding is",
//     choices: ["The chief programmer", "The programming secretary", "A specialized function that exists outside ‘the team’", "The individual coder (i.e. programmer)"],
//     //rightAnswer: The individual coder (i.e. programmer),
//   },

//       {
//     questionNumber : "questionFive",
//     question: "In the classical chief programmer team approach, the team member responsible for maintaining the detailed design and coding is",
//     choices: ["The chief programmer", "The programming secretary", "A specialized function that exists outside ‘the team’", "The individual coder (i.e. programmer)"],
//     //rightAnswer: The individual coder (i.e. programmer),
//   },

//       {
//     questionNumber : "questionFive",
//     question: "In the classical chief programmer team approach, the team member responsible for maintaining the detailed design and coding is",
//     choices: ["The chief programmer", "The programming secretary", "A specialized function that exists outside ‘the team’", "The individual coder (i.e. programmer)"],
//     //rightAnswer: The individual coder (i.e. programmer),
//   },

//       {
//     questionNumber : "questionFive",
//     question: "In the classical chief programmer team approach, the team member responsible for maintaining the detailed design and coding is",
//     choices: ["The chief programmer", "The programming secretary", "A specialized function that exists outside ‘the team’", "The individual coder (i.e. programmer)"],
//     //rightAnswer: The individual coder (i.e. programmer),
//   },

];
export default {
  name: "TrainingPage",
  data() {
    return {
      questions,
      score: 0,
      questionIndex: 0,
      question: questions[0],
      answer: "",
      isHidden:false,
      picked : -1,
      answers : {},
      trainComp : false,
      sc: 0,
      sco:0,
      clicked: false,
      finalSubmit: false,

    };
  },
  methods: {
    submit() {

        const { answer, question, questions, questionIndex } = this;

        let key = [question.questionNumber]

        this.answers[key] =  this.picked

        if (answer === question.rightAnswer) {
            this.score++;
        }

        // if(this.finalSubmit == true){
        //     this.questionIndex = 0
        //     this.question = questions[0]
        // }
        
        console.log("final")
        console.log(this.finalSubmit)

        if (questionIndex < questions.length) {
            this.questionIndex++;
            this.question = { ...questions[this.questionIndex] };
            
        }
        console.log(this.answers)
        console.log(questionIndex)
        console.log("submitting")

        // else if(this.questionIndex == 10)
        //   {
        //        console.log("coming to else")
        //         this.$axios.post("http://192.168.0.16:8080/training/send", this.answers)
        //           .then(response => {
        //               let jsonData = JSON.parse(response.data.data)
        //               console.log(jsonData)
        //               this.$store.state.accessToken = jsonData.accessToken
        //               this.sco = score
        //               this.clicked = true
        //         })
        //   }

      this.picked = -1

    },

    scored(){

      const { question, } = this;
      let key = [question.questionNumber]
      this.answers[key] =  this.picked

      console.log(this.answers)
        console.log("coming to score")
        this.$axios.post("http://localhost:8080/training/send", this.answers)
          .then(response => {
              let jsonData = JSON.parse(response.data.data)
              console.log(jsonData)
              this.$store.state.accessToken = jsonData.accessToken
              this.sco = jsonData.score
              this.trainComp = true
              // this.clicked = true
              this.finalSubmit = true
              this.questionIndex = 0
              this.question = questions[0]


        })
    },

    restart() {
      this.question = questions[0];
      this.answer = "";
      this.questionIndex = 0;
      this.score = 0;
      this.answers = []
    },

    click(){
        console.log("coming here to click")
        // this.$store.commit('LOADER', true)
        // const requestObj = {
        //     "email" : this.userName,
        //     "password" : this.password
        // }
        this.$axios.get("http://localhost:8080/training/")
          .then(response => {
            // console.log(response)
            console.log("here")
              let jsonData = JSON.parse(response.data.data)
              console.log(jsonData)
              this.trainComp = jsonData.trainingCompleted
              console.log(this.trainComp)
              this.sco = jsonData.score
              // this.$forceUpdate()

              //this.teamPayslipObj = jsonData.teamSalaries
              //this.isManager = jsonData.isManager

              // this.$store.state.accessToken = jsonData.accessToken
              // let token = jsonData.accessToken
              // this.$axios.defaults.headers.common['Authorization'] = "Bearer " + token

        //       // this.$store.commit('LOADER', false)
        //       // this.$router.push('/landing')
        //       // }
          })


    }

  },
};
</script>

// <style scoped lang="scss">
// .formid {
//   background-color: solid red;
// }
// </style>

<style scoped>
    .container-fluid {
        background-color: rgb(238, 243, 245);
        height: 100%;
        width: 100%;
        background-attachment: fixed;
        min-height: 100%;
        min-width: 100%;
        margin: 0;
        padding: 10px 75px
    }
    .styled{
      text-align: left ;
      font-style: oblique;
      font-size: large;
    }
</style>