<template>
  <div style="top:0px; height: 100%;" class="about">
    <v-container fluid class="p-0" >
      <!-- style="height: 100%" -->
      <!-- 100vh -->
      <NavigationBar class="pb-0" v-bind:userName="userName"/>
      <v-layout fill-height class="d-flex">
        <v-row>
          <v-card class="mx-auto p-0" width="100%" elevation="4">
            <v-toolbar
              color="#F94D00"
              dark
              flat
            >
            <!-- prominent -->
              <!-- <v-text-field
                append-icon="mdi-microphone"
                class="mx-4"
                flat
                hide-details
                label="Search"
                prepend-inner-icon="mdi-magnify"
                solo-inverted
              ></v-text-field> -->
        
              <template v-slot:extension>
                <v-tabs
                  id="tabsId"
                  v-model="tabs"
                  centered
                >
                  <v-tab
                    class="black--text"
                    v-for="t in tabName"
                    :key="t"
                  >
                    {{t}}
                  </v-tab>
                </v-tabs>
              </template>
            </v-toolbar>

            <v-tabs-items v-model="tabs">
              <v-tab-item>
                <v-parallax
                  height="100%"
                  src="https://cdn.vuetifyjs.com/images/parallax/material2.jpg"
                >
                  <v-card-title class="justify-center">
                    <h1 id="AboutUsHeading" class="black--text text-h2 font-weight-thin mb-2">
                      Meet Axle United!
                    </h1>
                  </v-card-title>
                  <!-- <v-card-subtitle style="font-size: 1.8em;" class="pb-2 justify-center font-weight-black">
                    Meet Axle United!
                  </v-card-subtitle> -->

                  <v-card-text>
                    <p id="textAboutUs" style="font-size: 1.4em;" class="black--text text-justify">
                      Axle United represents a new way of thinking about the data hidden in your documents. We use AI to extract data that can be used to fuel corporate transformation.

                      <br><br>From old documents to stamps and handwritten notes to receipts, our system can infer what’s missing and fill in the blanks to get you more valuable extracted data. That provides the springboard you need to automate your process, which in turn can transform your business quickly and efficiently.

                      <br><br>Our AI knows how to find what others don’t. We make it quick and easy for you to get the most from the data you have, and to move your enterprise closer to incredibly valuable end-to-end automation.
                      
                      <br><br>Morbi mattis ullamcorper velit. Donec orci lectus, aliquam ut, faucibus non, euismod id, nulla. Fusce convallis metus id felis luctus adipiscing. Aenean massa. Vestibulum purus quam, scelerisque ut, mollis sed, nonummy id, metus.

                      Nulla consequat massa quis enim. Praesent venenatis metus at tortor pulvinar varius. Donec venenatis vulputate lorem. Phasellus accumsan cursus velit. Pellentesque ut neque.
                      Inference is a powerful tool. You do it without thinking — your brain fills in the missing vowels in our company’s name effortlessly. It knows what’s missing. So do we. We have designed our technology to make these kinds of connections, and at lightning speed.
                    </p>
                  </v-card-text>

                  <v-card-title class="justify-center">
                    <h1 id="teamAboutUs" class="black--text text-h3 font-weight-thin mb-2">
                      Meet The Team
                    </h1>
                  </v-card-title>

                  <v-card-text>
                    <v-row>
                      <!-- align-content="center" -->
                      <!-- align="center" -->
                      <v-col style="text-align:center; display:inline-block;" align-content="center" cols="4" class="text-center fill-height">
                        <!-- <span> -->
                          <v-icon  color="#000080">account_group_outline</v-icon>
                          <v-card-text class="black--text text-h5 font-weight-bold">
                            2000+ Employees
                          </v-card-text>
                        <!-- </span> -->
                        <!-- large -->
                      </v-col>

                      <v-col style="text-align:center; display:inline-block;" align-content="center" cols="4" class="fill-height">
                        <!-- <span> -->
                          <v-icon color="#000080">feed</v-icon>
                          <v-card-text class="black--text text-h4 font-weight-bold">
                            50000+ Records Processed
                          </v-card-text>
                        <!-- </span> -->
                        <!-- style="font-size: 44px;" -->
                      </v-col>

                      <v-col style="text-align:center; display:inline-block;" align-content="center" cols="4" class="fill-height">
                        <!-- <span> -->
                          <v-icon color="#000080">account_balance</v-icon>
                          <v-card-text class="black--text text-h4 font-weight-bold">
                            50% Employess from Harvard/Stanford University
                          </v-card-text>
                        <!-- </span> -->
                         <!-- margin-top: 10px -->
                         <!-- style="font-size: 44px;" -->
                      </v-col>

                    </v-row>
                  </v-card-text>

                  <v-card-title class="justify-center">
                    <h1 id="accomplishmentsTeams" class="black--text text-h3 font-weight-thin mb-2">
                      Our Accomplishments!
                    </h1>
                  </v-card-title>

                  <v-card-text>
                    <v-row>
                      <!-- align-content="center" -->
                      <!-- align="center" -->
                      <v-col style="text-align:center; display:inline-block;" align-content="center" cols="4" class="fill-height">
                        <!-- <span> -->
                          <v-icon color="#000080">feed</v-icon>
                          <v-card-text class="black--text text-h4 font-weight-bold">
                            Over 300+ Clients
                          </v-card-text>
                        <!-- </span> -->
                        <!-- large -->
                      </v-col>

                      <v-col style="text-align:center; display:inline-block;" align-content="center" cols="4" class="fill-height">
                        <!-- <span> -->
                          <v-icon color="#000080">account_group_outline</v-icon>
                          <v-card-text class="black--text text-h4 font-weight-bold">
                            Voted Most Promising Statup 2022 by the Times group!
                          </v-card-text>
                        <!-- </span> -->
                        <!-- style="font-size: 44px;" -->
                      </v-col>

                      <v-col style="text-align:center; display:inline-block;" align-content="center" cols="4" class="fill-height">
                        <!-- <span> -->
                          <!-- #D22B2B -->
                          <!-- #000080 -->
                          <v-icon color="#000080">account_balance</v-icon>
                          <v-card-text class="black--text text-h4 font-weight-bold">
                            100% Customer Satisfaction
                          </v-card-text>
                        <!-- </span> -->
                         <!-- margin-top: 10px -->
                         <!-- style="font-size: 44px;" -->
                      </v-col>

                    </v-row>
                  </v-card-text>

                  <v-card-title class="justify-left">
                    <h1 class="black--text text-h3 font-weight-thin mb-2">
                      Leadership
                    </h1>
                  </v-card-title>

                  <v-card-text class="black--text text-h6 text-justify">
                    Our leaders reflect our values of teamwork, transparency, and innovation.
                    And they’re committed to making the impossible possible. <br> <br>

                    <a href="https://github.com/deepakvk0712/HRManagementSoftware-SoftwareEngineering" class="black--text">Meet the leadership team ---></a>
                  </v-card-text>
                
                </v-parallax>
              </v-tab-item>

              <v-tab-item>
                <v-parallax
                  height="100%"
                  src="https://cdn.vuetifyjs.com/images/parallax/material2.jpg"
                >
                  <v-card-title class="justify-center">
                    <h1 id="companyPolicy" class="black--text text-h2 font-weight-thin mb-2">
                      Company Policy
                    </h1>
                  </v-card-title>

                  <v-list-item id="companyPolicyList" v-for="compPol in companyPolicy" :key="compPol.Id" class="text-justify">
                    <v-list-item-content>
                        <v-card-title class="justify-left">
                          <h1 class="text-h5 font-weight-medium mb-2">
                            {{compPol.heading}}
                          </h1>
                        </v-card-title>
                        <v-card-text>
                          <v-list-item-sub-title class="text-h6 font-weight-medium mb-2">{{compPol.content}}</v-list-item-sub-title>
                        </v-card-text>
                      </v-list-item-content>
                  </v-list-item>
                  
                </v-parallax>
              </v-tab-item>

              <v-tab-item>
                <v-parallax
                  height="100%"
                  src="https://cdn.vuetifyjs.com/images/parallax/material2.jpg"
                >
                <!-- :src="require('@/assets/mountain.png')" -->
                <!-- https://cdn.vuetifyjs.com/images/parallax/material2.jpg -->
                  <v-card-title class="justify-center">
                    <h1 id="employeePolicy" class="black--text text-h2 font-weight-thin mb-2">
                      Employee Policy
                    </h1>
                  </v-card-title>

                  <v-list-item id="employeePolicyList" v-for="empPol in employeePolicy" :key="empPol.Id" class="text-justify">
                    <v-list-item-content>
                        <v-card-title class="justify-left">
                          <h1 class="text-h5 font-weight-medium mb-2">
                            {{empPol.heading}}
                          </h1>
                        </v-card-title>
                        <v-card-text>
                          <v-list-item-sub-title class="text-h6 font-weight-medium mb-2">{{empPol.content}}</v-list-item-sub-title>
                        </v-card-text>
                      </v-list-item-content>
                  </v-list-item>
                  
                </v-parallax>
              </v-tab-item>

              <v-tab-item>
                <v-parallax
                  height="100%"
                  src="https://cdn.vuetifyjs.com/images/parallax/material2.jpg"
                >
                <!-- :src="require('@/assets/mountain.png')" -->
                <!-- https://cdn.vuetifyjs.com/images/parallax/material2.jpg -->
                  <v-card-title class="justify-center">
                    <h1 id="employeePolicy" class="black--text text-h2 font-weight-thin mb-2">
                      Core Values
                    </h1>
                  </v-card-title>

                  <v-list-item id="coreValuesId" v-for="val in coreValues" :key="val.Id" class="text-justify">
                    <v-list-item-content>
                        <v-card-title class="justify-left">
                          <h1 class="text-h5 font-weight-medium mb-2">
                            {{val.heading}}
                          </h1>
                        </v-card-title>
                        <v-card-text>
                          <v-list-item-sub-title class="text-h6 font-weight-medium mb-2">{{val.content}}</v-list-item-sub-title>
                        </v-card-text>
                      </v-list-item-content>
                  </v-list-item>
                  
                </v-parallax>
              </v-tab-item>

              <v-tab-item>
                <v-parallax
                  height="100%"
                  src="https://cdn.vuetifyjs.com/images/parallax/material2.jpg"
                >
                <!-- :src="require('@/assets/mountain.png')" -->
                <!-- https://cdn.vuetifyjs.com/images/parallax/material2.jpg -->
                  <v-card-title class="black--text justify-center">
                    <h1 id="employeePolicy" class="text-h2 font-weight-thin mb-2">
                      Mission Statement
                    </h1>
                  </v-card-title>

                  <v-card-title class="justify-left">
                    <h1 class="black--text text-h3 font-weight-medium mb-2">
                      {{missionStatement}}
                    </h1>
                  </v-card-title>                                           
                </v-parallax>
              </v-tab-item>
            </v-tabs-items>
          </v-card>
        </v-row>
      </v-layout>
    </v-container>
  </div>
</template>

<script>
import NavigationBar from '../components/NavigationBar.vue'

export default{
  name: 'aboutMe',
  components: {
    NavigationBar,
  },
  data: () => ({
    isClicked: false,
    tabs:0,
    tabName : ["About Us", "Company Policy", "Employee Policy", "Core Values", "Mission Statement"],
    companyPolicy : [
      {
        Id : 1,
        heading : "Equal opportunity policy",
        content : "An equal opportunity policy (EOP) prevents companies from discriminating against job applicants or employees if they are a member of a protected class (e.g, race, gender, age, religion, familial status, color). The EOP is essential for any anti-harassment, workplace violence, non-discrimination or diversity policies your company may consider developing."
      },
      {
        Id : 2,
        heading : "Workplace health and safety",
        content : "Your business should be proactive and write a health and safety policy that is designed for each workplace. For example, you might specify what employees should do in case of office emergencies or how to handle unsafe materials. The Occupational Safety and Health Administration (OSHA) has guidelines on how to create a safe workplace and protect workers from occupational hazards that you can base your policy on."
      },
      {
        Id : 3,
        heading : "Employee code of conduct policy ",
        content : "Axle United is committed to eradicating discrimination and unlawful harassment in our workplaces. Any actions, jokes or comments based on an employee or client’s race, religion ethnicity, sex, age or any other legally protected class are not tolerated and will be met with significant disciplinary action."
      },
      {
        Id : 4,
        heading : "Attendance, vacation and time-off policies",
        content : "Employees are expected to be on time and regular in attendance. This means being at your workspace and ready to work at your scheduled time each day. You will be given a 10-minute grace period after the start of your shift before you will be considered tardy. Employees who are tardy on more than five occasions will be subject to disciplinary action. Absenteeism and tardiness are burdensome to your co-workers and leaders, and will not be tolerated without just cause."
      },
      {
        Id : 5,
        heading : "Employee disciplinary action policy",
        content : "A simple step-by-step list of what happens regarding disciplinary action can make it easy for employees to know what to expect if they violate a company policy. Describe a specific process you will follow to ensure every employee is treated fairly when it comes to discipline. Have a lawyer review this information before you include it in your employee handbook to make sure all disciplinary action is legal."
      },
      {
        Id : 6,
        heading : "Employee complaint policies",
        content : "Grievances are formal complaints your employees can file to document their concerns with an aspect of their workplace. These grievances might be filed as a result of an incident or conflict with a fellow employee. A grievance can be filed for nearly any reason, including physical workplace complaints, financial issues like payroll and social circumstances like harassment or bullying. It’s important to outline a formal process for resolving complaints within your company so that employees know how to handle their concerns in a professional way."
      },
      // {
      //   Id : 1,
      //   heading : "Equal opportunity policy",
      //   content : ""
      // },
    ],
    employeePolicy : [
      {
        Id : 1,
        heading : "At-will employment:",
        content : "This statement reiterates that either you or the employee can terminate the employment relationship at any time and for any reason, as long as the reason is a lawful one. It is a best practice to prominently display this statement in the beginning of your employee handbook (except in Montana, where at-will employment is not recognized). Reinforce at-will status in your handbook acknowledgment form as well."
      },
      {
        Id : 2,
        heading : "Anti-harassment and non-discrimination:",
        content : "These policies prohibit harassment and discrimination in the workplace. Non-discrimination laws are governed by federal, state and local provisions, so review your applicable law and account for all appropriate protections."
      },
      {
        Id : 3,
        heading : "Employment classifications:",
        content : "It is a best practice to clearly define employment classifications, such as full-time, part-time, exempt or non-exempt since an employee's classification can dictate eligibility for benefits and overtime pay."
      },
      {
        Id : 4,
        heading : "Leave and time off benefits:",
        content : "These policies address a company's rules and procedures regarding holidays, vacation, sick, and other types of time off benefits, or leave required by law (such as voting leave, family leave, and domestic violence leave) or company policy. Check your state and local law to ensure all leave requirements are included in your employee handbook."
      },
      {
        Id : 5,
        heading : "Meal and break periods:",
        content : "A policy on meal and break periods informs employees of the frequency and duration of such breaks as well as any rules or restrictions related to break periods. Rest periods, lactation breaks, and meal periods must be provided in accordance with federal, state and local laws."
      },
      {
        Id : 6,
        heading : "Timekeeping and pay:",
        content : "A timekeeping policy informs employees of the method for recording time worked and the importance of accurately recording their time. A policy on paydays lets employees know the frequency of paydays, the methods available for receiving pay, and any special procedures for when a payday falls on a holiday or when an employee is absent from work."
      },
      {
        Id : 7,
        heading : "Safety and health:",
        content : "Safety policies describe safety and emergency procedures and require employees to report work-related injuries immediately. Additionally, some regulations under the Occupational Safety and Health Act require employers to have specific policies and programs in place if certain workplace hazards exist (such as a hazard communication program if certain chemicals are present in the workplace)."
      },
      {
        Id : 8,
        heading : "Employee conduct, attendance and punctuality. :",
        content : "Attendance policies make it clear that employees must be ready to work at their scheduled start time each day and provide procedures for informing the company of an unscheduled absence or late arrival. It is also a best practice to have policies on standards of conduct, drug and alcohol abuse, disciplinary action, confidentiality, conflicts of interest, and workplace violence."
      },
      // {
      //   Id : 2,
      //   heading : "Anti-harassment and non-discrimination:",
      //   content : ""
      // },
    ],
    coreValues : [
      {
        Id : 1,
        heading : "Integrity",
        content : "Acting with strong ethics is a priority for everyone representing the organization as well as the company’s behavior as a whole."
      },
      {
        Id : 2,
        heading : "Honesty",
        content : "It’s not just the best policy. It’s a core business practice to act in a transparent, trustworthy manner that earns the respect of colleagues, customers, and the public."
      },
      {
        Id : 3,
        heading : "Fairness",
        content : "Treating everyone with the common decency we all deserve and expect. "
      },
      {
        Id : 4,
        heading : "Accountability",
        content : "Accepting responsibility for your actions (and inactions) is the ultimate way to build trust internally and externally. "
      },
      {
        Id : 5,
        heading : "Promise to Customers",
        content : "Creating a great customer experience begins with staying true to the words we speak and the bonds we make."
      },
      {
        Id : 6,
        heading : "Diversity and Inclusion",
        content : "Organizations succeed by bringing different lived experiences and a range of backgrounds into a shared environment where everyone has equal opportunity."
      },
      {
        Id : 7,
        heading : "Learning",
        content : "No one has all the answers. A culture of humility and continuous learning is a bedrock principle of successful companies. "
      },
      {
        Id : 8,
        heading : "Teamwork",
        content : "When people work together, they can create something greater than themselves as individuals. "
      },
      {
        Id : 9,
        heading : "Passion",
        content : "Having a joy not just for the work itself but also the people around us, so that everyone can be bold, innovative, and creative. "
      },
      {
        Id : 10,
        heading : "Quality",
        content : "Companies are judged by the craftsmanship of their products and services, so the highest standards must be maintained. "
      },
    ],
    missionStatement : '“Shape the future of the Internet by creating unprecedented value and opportunity for our customers, employees, investors, and ecosystem partners.”'
  }),
}

</script>


<style lang="scss" scoped>
</style>