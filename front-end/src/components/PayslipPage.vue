<template>
    <v-app>
        <div>
          <h1 id="titleId">Payslip Page</h1>
        </div>
        <v-data-table
            id="empTableId"
            dense
            :headers="headers"
            :items="payslipObj"
            item-key="name"
            class="elevation-1"
        >
        <!-- <template v-slot:[`item.payBeginDate`]="{ item }">
            <span>{{ new Date(item.payBeginDate).toLocaleString() }}</span>
        </template> -->

        </v-data-table>

      <div>
      <h3> Start Date </h3>
      </div>
        <v-date-picker
            v-model="startDate"
            width="390"
            class="mt-4"
            no-title
        >
        </v-date-picker>


        <div>
        <!-- <b-form-datepicker id="example-datepicker" v-model="value" class="mb-2"></b-form-datepicker> -->
        <p>Value: '{{ startDate }}'</p>
        </div>


      <div>
      <h3> End Date </h3>
      </div>
        <v-date-picker
            v-model="endDate"
            width="390"
            no-title
        >
        </v-date-picker>


        <div>
        <!-- <b-form-datepicker id="example-datepicker" v-model="value" class="mb-2"></b-form-datepicker> -->
        <p>Value: '{{ endDate }}'</p>
        </div>

      <v-btn id="filterId" color = "#6495ED" @click="filter()" class="mb-4">Filter Dates</v-btn>

      <v-btn id="teamButtonId" color = "#6495ED" @click="teamSalaries()" class="mb-4">Show Team Payslip</v-btn>
       
        <div v-if="isManager == true">
          <input class="mr-4" v-model="empID" placeholder="Employee ID">
          <input class="mr-4" v-model="salaryUpdate" placeholder="New Salary">
          <v-btn id="salaryUpdate" color = "#6495ED" @click="updateSalary()" class="mb-4">Update Salary</v-btn>
        </div>

        <v-data-table v-show="isManager == true"
            dense
            :headers="headers1"
            :items="teamPayslipObj"
            item-key="name"
            class="elevation-1"
        ></v-data-table>


        <!-- <div v-if="isManager === 'true'"> -->
        <!-- <b-table striped hover :items="items"></b-table> -->

      <!-- <div v-for="slip in payslipObj" :key="slip.employeeID">
        {{ slip.name }}: {{ value }}
      </div> -->

        <!-- <div id="app">
          <v-simple-table>
            <template v-slot:default>
              <thead>
                <tr>
                  <th class="text-left">
                    Name
                  </th>
                  <th class="text-left">
                    Calories
                  </th>
                </tr>
              </thead>
              <tbody>
                <tr
                  v-for="item in desserts"
                  :key="item.name"
                >
                  <td>{{ item.name }}</td>
                  <td>{{ item.calories }}</td>
                </tr>
              </tbody>
            </template>
          </v-simple-table>
        </div> -->

    </v-app>
</template>

<script>
  export default {
    name : "PaySlipPage",
    async created() {
      await this.$axios.get('http://localhost:8080/paycheck/')
        .then(response => {
          let jsonData = JSON.parse(response.data.data)
          console.log(jsonData)
          this.payslipObj = jsonData.paychecks
        })
    },
    data: () => ({

        //   desserts: [
        // {
        //   name: 'Frozen Yogurt',
        //   calories: 159,
        // },
        // ],
        

          // items: [
          //   { age: 40, first_name: 'Dickerson', last_name: 'Macdonald' },
          //   { age: 21, first_name: 'Larsen', last_name: 'Shaw' },
          //   { age: 89, first_name: 'Geneva', last_name: 'Wilson' },
          //   { age: 38, first_name: 'Jami', last_name: 'Carney' }
          // ],


      // PayCheckDate: [
      //   {
      //     name: '03 March 2022 ',
      //     ViewPaycheck: 'View',
      //     Company: 'HRManagement',
      //     BeginDate: '03/01/2022',
      //     EndDate: '04/01/2022',
      //     TotalPay: '$2000',
      //     PaycheckNumber: 1234567,
      //   },
      //   {
      //     name: '03 April 2022 ',
      //     ViewPaycheck: 'View',
      //     Company: 'HRManagement',
      //     BeginDate: '04/01/2022',
      //     EndDate: '05/01/2022',
      //     TotalPay: '$2000',
      //     PaycheckNumber: 123456,
      //   },
      //   {
      //     name: '03 May 2022 ',
      //     ViewPaycheck: 'View',
      //     Company: 'HRManagement',
      //     BeginDate: '05/01/2022',
      //     EndDate: '06/01/2022',
      //     TotalPay: '$2000',
      //     PaycheckNumber: 1234568,
      //   },
      // ],

      isManager : false,
      payslipObj : [],
      teamPayslipObj : [],
      startDate: '',
      endDate: '',
      ViewPaycheck:'view',
      empID:'',
      salaryUpdate:'',


      headers: [
        {
          text: 'Paycheck Date',
          align: 'start',
          sortable: false,
          value: 'checkDate',
        },
        { text: 'Employee ID', value: 'employeeID' },
        // { text: 'View Paycheck', value: 'ViewPaycheck' },
        // { text: 'Company', value: 'Company' },
        { text: 'Begin Date', value: 'payBeginDate'},
        { text: 'End Date', value: 'payEndDate'},
        { text: 'Total Pay', value: 'amountPaid' },
        { text: 'Paycheck Number', value: 'checkNumber' },
      ],


      headers1: [
        {
          text: 'Paycheck Date',
          align: 'start',
          sortable: false,
          value: 'checkDate',
        },
        { text: 'Employee Name', value: 'employeeName' },
        { text: 'Employee ID', value: 'employeeID' },
        // { text: 'View Paycheck', value: 'ViewPaycheck' },
        // { text: 'Company', value: 'Company' },
        { text: 'Begin Date', value: 'payBeginDate' },
        { text: 'End Date', value: 'payEndDate' },
        { text: 'Total Pay', value: 'amountPaid' },
        { text: 'Updated Salary', value: 'salary' },
        // { text: 'Paycheck Number', value: 'checkNumber' },
      ],

      // items: [
      //   { age: 40, first_name: 'Dickerson', last_name: 'Macdonald' },
      //   { age: 21, first_name: 'Larsen', last_name: 'Shaw' },
      //   { age: 89, first_name: 'Geneva', last_name: 'Wilson' },
      //   { age: 38, first_name: 'Jami', last_name: 'Carney' }
      // ],

    }),


    // uncomment this method for putting filter on date
    methods : {
      teamSalaries() {
        console.log("coming here")
        // this.$store.commit('LOADER', true)
        // const requestObj = {
        //     "email" : this.userName,
        //     "password" : this.password
        // }
        this.$axios.get("http://localhost:8080/paycheck/")
          .then(response => {
              let jsonData = JSON.parse(response.data.data)
              console.log(jsonData)
              this.teamPayslipObj = jsonData.teamSalaries
              this.isManager = jsonData.isManager

              // this.$store.state.accessToken = jsonData.accessToken
              // let token = jsonData.accessToken
              // this.$axios.defaults.headers.common['Authorization'] = "Bearer " + token

        //       // this.$store.commit('LOADER', false)
        //       // this.$router.push('/landing')
        //       // }
          })
        },

      filter(){
        console.log("coming inside filter date")
        // this.$store.commit('LOADER', true)
        // const requestObj = {
        //     "email" : this.userName,
        //     "password" : this.password
        // }
        this.$axios.get("http://localhost:8080/paycheck/", { params: { startDate: this.startDate , endDate:this.endDate } })
          .then(response => {
              let jsonData = JSON.parse(response.data.data)
              console.log(jsonData)
              this.teamPayslipObj = jsonData.teamSalaries
              this.payslipObj = jsonData.paychecks

              // this.$store.state.accessToken = jsonData.accessToken
              // let token = jsonData.accessToken
              // this.$axios.defaults.headers.common['Authorization'] = "Bearer " + token

        //       // this.$store.commit('LOADER', false)
        //       // this.$router.push('/landing')
        //       // }
          })
        },

        updateSalary(){
            const requestObj = {
                "employeeID" : parseInt(this.empID),
                "newSalary" : parseFloat(this.salaryUpdate)
            }
            this.$axios.put("http://localhost:8080/paycheck/updateSalary", requestObj)
                .then(response => {
                    let jsonData = JSON.parse(response.data.data)
                    console.log(jsonData)
                    this.$store.state.accessToken = jsonData.accessToken
          })
        },

          // teamSalaries(){
          //     this.$axios.post("http://localhost:8080/paycheck/")
          //     .then(response => {
          //         let jsonData = JSON.parse(response.data.data)
          //         console.log(jsonData)
          //         this.teamPayslipObj = jsonData.teamSalaries

          //     },
    }

  }
</script>