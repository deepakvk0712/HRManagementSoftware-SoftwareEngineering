<template>
  <form>
    <div>
    <h1 class = ma-9> Welcome to the Training Session </h1>
    </div>

    <div>
    <h4 class = ma-9> Please watch this video and answer the below questions</h4>
    </div>


    <section class= ma-5>
        <iframe width="570" height="380" src="https://www.youtube.com/embed/sB2iQSvrcG0" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen>
        </iframe>
    </section>

    <!-- <video-embed src="https://www.youtube.com/watch?v=sB2iQSvrcG0"></video-embed> -->

    <v-btn id="startID" color = "#6495ED" @click="isHidden = true" class="mb-4">Start Quiz</v-btn>

    <div>
    <h2 class = ma-5> Get 7/10 correct answers to pass</h2>
    </div>

    <div>
    <h2 class = ma-5> All the Best!</h2>
    </div>

    <!-- <div v-if="isHidden == true && questionIndex < questions.length">
    <div v-for="q in questions" :key="q.rightAnswer"> -->

    <div v-if="isHidden == true && questionIndex < questions.length">
      <label>{{ question.question }}</label>
      <div v-for="(c, index) of question.choices" :key="c">
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
    <!-- </div> -->

    <!-- <div v-else-if="isHidden == true">
      <button type="button" @click="restart">restart</button>
    </div> -->

    <div v-if="isHidden == true">
    <button type="button" :disabled="picked == -1" @click="submit(question)">Submit</button>
    </div>

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
      answers : []
    };
  },
  methods: {
    submit() {
        const { answer, question, questions, questionIndex } = this;

        const obj = {
            [question.questionNumber] : this.picked
        }
        this.answers.push(obj)

        if (answer === question.rightAnswer) {
            this.score++;
        }

        if (questionIndex < questions.length) {
            this.questionIndex++;
            this.question = { ...questions[this.questionIndex] };
        }

        else{
            //call the api
        }

        this.picked = -1
    },
    restart() {
      this.question = questions[0];
      this.answer = "";
      this.questionIndex = 0;
      this.score = 0;
      this.answers = []
    },
  },
};
</script>