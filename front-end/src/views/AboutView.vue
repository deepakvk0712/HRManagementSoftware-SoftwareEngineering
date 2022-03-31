<template>
  <div style="top:0px; height: 100%;" class="about">
    <v-container fluid class="p-0" style="height: 100%">
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
                  v-model="tabs"
                  centered
                >
                  <v-tab
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
                    <h1 class="text-h2 font-weight-thin mb-2">
                      Meet Axle United!
                    </h1>
                  </v-card-title>
                  <!-- <v-card-subtitle style="font-size: 1.8em;" class="pb-2 justify-center font-weight-black">
                    Meet Axle United!
                  </v-card-subtitle> -->

                  <v-card-text>
                    <p style="font-size: 1.4em;" class="text-justify">
                      Morbi mattis ullamcorper velit. Donec orci lectus, aliquam ut, faucibus non, euismod id, nulla. Fusce convallis metus id felis luctus adipiscing. Aenean massa. Vestibulum purus quam, scelerisque ut, mollis sed, nonummy id, metus.

                      Nulla consequat massa quis enim. Praesent venenatis metus at tortor pulvinar varius. Donec venenatis vulputate lorem. Phasellus accumsan cursus velit. Pellentesque ut neque.
                      Inference is a powerful tool. You do it without thinking — your brain fills in the missing vowels in our company’s name effortlessly. It knows what’s missing. So do we. We have designed our technology to make these kinds of connections, and at lightning speed.

                      <br><br>Infrrd represents a new way of thinking about the data hidden in your documents. We use AI to extract data that can be used to fuel corporate transformation.

                      <br><br>From old documents to stamps and handwritten notes to receipts, our system can infer what’s missing and fill in the blanks to get you more valuable extracted data. That provides the springboard you need to automate your process, which in turn can transform your business quickly and efficiently.

                      <br><br>Our AI knows how to find what others don’t. We make it quick and easy for you to get the most from the data you have, and to move your enterprise closer to incredibly valuable end-to-end automation.
                    </p>
                  </v-card-text>

                  <v-card-title class="justify-center">
                    <h1 class="text-h3 font-weight-thin mb-2">
                      Meet The Team
                    </h1>
                  </v-card-title>

                  <v-card-text>
                    <v-row>
                      <!-- align-content="center" -->
                      <!-- align="center" -->
                      <v-col style="text-align:center; display:inline-block;" align-content="center" cols="4" class="fill-height">
                        <!-- <span> -->
                          <v-icon color="#D22B2B">account_group_outline</v-icon>
                          <v-card-text class="text-h4 font-weight-bold">
                            2000+ Employees
                          </v-card-text>
                        <!-- </span> -->
                        <!-- large -->
                      </v-col>

                      <v-col style="text-align:center; display:inline-block;" align-content="center" cols="4" class="fill-height">
                        <!-- <span> -->
                          <v-icon color="#D22B2B">feed</v-icon>
                          <v-card-text class="text-h4 font-weight-bold">
                            50000+ Records Processed
                          </v-card-text>
                        <!-- </span> -->
                        <!-- style="font-size: 44px;" -->
                      </v-col>

                      <v-col style="text-align:center; display:inline-block;" align-content="center" cols="4" class="fill-height">
                        <!-- <span> -->
                          <v-icon color="#D22B2B">account_balance</v-icon>
                          <v-card-text class="text-h4 font-weight-bold">
                            50% Employess from Harvard/Stanford University
                          </v-card-text>
                        <!-- </span> -->
                         <!-- margin-top: 10px -->
                         <!-- style="font-size: 44px;" -->
                      </v-col>

                    </v-row>
                  </v-card-text>

                  <v-card-title class="justify-center">
                    <h1 class="text-h3 font-weight-thin mb-2">
                      Our Accomplishments!
                    </h1>
                  </v-card-title>

                  <v-card-text>
                    <v-row>
                      <!-- align-content="center" -->
                      <!-- align="center" -->
                      <v-col style="text-align:center; display:inline-block;" align-content="center" cols="4" class="fill-height">
                        <!-- <span> -->
                          <v-icon color="#D22B2B">feed</v-icon>
                          <v-card-text class="text-h4 font-weight-bold">
                            Over 300+ Clients
                          </v-card-text>
                        <!-- </span> -->
                        <!-- large -->
                      </v-col>

                      <v-col style="text-align:center; display:inline-block;" align-content="center" cols="4" class="fill-height">
                        <!-- <span> -->
                          <v-icon color="#D22B2B">account_group_outline</v-icon>
                          <v-card-text class="text-h4 font-weight-bold">
                            Voted Most Promising Statup 2022 by the Times group!
                          </v-card-text>
                        <!-- </span> -->
                        <!-- style="font-size: 44px;" -->
                      </v-col>

                      <v-col style="text-align:center; display:inline-block;" align-content="center" cols="4" class="fill-height">
                        <!-- <span> -->
                          <v-icon color="#D22B2B">account_balance</v-icon>
                          <v-card-text class="text-h4 font-weight-bold">
                            100% Customer Satisfaction
                          </v-card-text>
                        <!-- </span> -->
                         <!-- margin-top: 10px -->
                         <!-- style="font-size: 44px;" -->
                      </v-col>

                    </v-row>
                  </v-card-text>

                  <v-card-title class="justify-left">
                    <h1 class="text-h3 font-weight-thin mb-2">
                      Leadership
                    </h1>
                  </v-card-title>

                  <v-card-text class="text-h6 text-justify">
                    Our leaders reflect our values of teamwork, transparency, and innovation.
                    And they’re committed to making the impossible possible. <br> <br>

                    <a href="">Meet the leadership team ---></a>
                  </v-card-text>
                
                </v-parallax>
              </v-tab-item>

              <v-tab-item>
                <v-parallax
                  height="100%"
                  src="https://cdn.vuetifyjs.com/images/parallax/material2.jpg"
                >
                  <v-card-title class="justify-center">
                    <h1 class="text-h2 font-weight-thin mb-2">
                      Company Policy
                    </h1>
                  </v-card-title>

                  <v-list-item v-for="compPol in companyPolicy" :key="compPol.Id" class="text-justify">
                    <v-list-item-content>
                        <v-card-text>
                          <v-list-item-sub-title class="text-h6 font-weight-thin mb-2">{{compPol.content}}</v-list-item-sub-title>
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
                    <h1 class="text-h2 font-weight-thin mb-2">
                      Employee Policy
                    </h1>
                  </v-card-title>

                  <v-list-item v-for="empPol in employeePolicy" :key="empPol.Id" class="text-justify">
                    <v-list-item-content>
                        <v-card-text>
                          <v-list-item-sub-title class="text-h6 font-weight-thin mb-2">{{empPol.content}}</v-list-item-sub-title>
                        </v-card-text>
                      </v-list-item-content>
                  </v-list-item>
                  
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
        content : "Morbi mattis ullamcorper velit. Donec orci lectus, aliquam ut, faucibus non, euismod id, nulla. Fusce convallis metus id felis luctus adipiscing. Aenean massa. Vestibulum purus quam, scelerisque ut, mollis sed, nonummy id, metus. Nulla consequat massa quis enim. Praesent venenatis metus at tortor pulvinar varius. Donec venenatis vulputate lorem. Phasellus accumsan cursus velit. Pellentesque ut neque. Inference is a powerful tool. You do it without thinking — your brain fills in the missing vowels in our company’s name effortlessly. It knows what’s missing. So do we. We have designed our technology to make these kinds of connections, and at lightning speed."
      },
      {
        Id : 2,
        content : "Morbi mattis ullamcorper velit. Donec orci lectus, aliquam ut, faucibus non, euismod id, nulla. Fusce convallis metus id felis luctus adipiscing. Aenean massa. Vestibulum purus quam, scelerisque ut, mollis sed, nonummy id, metus. Nulla consequat massa quis enim. Praesent venenatis metus at tortor pulvinar varius. Donec venenatis vulputate lorem. Phasellus accumsan cursus velit. Pellentesque ut neque. Inference is a powerful tool. You do it without thinking — your brain fills in the missing vowels in our company’s name effortlessly. It knows what’s missing. So do we. We have designed our technology to make these kinds of connections, and at lightning speed."
      },
      {
        Id : 3,
        content : "Morbi mattis ullamcorper velit. Donec orci lectus, aliquam ut, faucibus non, euismod id, nulla. Fusce convallis metus id felis luctus adipiscing. Aenean massa. Vestibulum purus quam, scelerisque ut, mollis sed, nonummy id, metus. Nulla consequat massa quis enim. Praesent venenatis metus at tortor pulvinar varius. Donec venenatis vulputate lorem. Phasellus accumsan cursus velit. Pellentesque ut neque. Inference is a powerful tool. You do it without thinking — your brain fills in the missing vowels in our company’s name effortlessly. It knows what’s missing. So do we. We have designed our technology to make these kinds of connections, and at lightning speed."
      },
      {
        Id : 4,
        content : "Morbi mattis ullamcorper velit. Donec orci lectus, aliquam ut, faucibus non, euismod id, nulla. Fusce convallis metus id felis luctus adipiscing. Aenean massa. Vestibulum purus quam, scelerisque ut, mollis sed, nonummy id, metus. Nulla consequat massa quis enim. Praesent venenatis metus at tortor pulvinar varius. Donec venenatis vulputate lorem. Phasellus accumsan cursus velit. Pellentesque ut neque. Inference is a powerful tool. You do it without thinking — your brain fills in the missing vowels in our company’s name effortlessly. It knows what’s missing. So do we. We have designed our technology to make these kinds of connections, and at lightning speed."
      },
      {
        Id : 5,
        content : "Morbi mattis ullamcorper velit. Donec orci lectus, aliquam ut, faucibus non, euismod id, nulla. Fusce convallis metus id felis luctus adipiscing. Aenean massa. Vestibulum purus quam, scelerisque ut, mollis sed, nonummy id, metus. Nulla consequat massa quis enim. Praesent venenatis metus at tortor pulvinar varius. Donec venenatis vulputate lorem. Phasellus accumsan cursus velit. Pellentesque ut neque. Inference is a powerful tool. You do it without thinking — your brain fills in the missing vowels in our company’s name effortlessly. It knows what’s missing. So do we. We have designed our technology to make these kinds of connections, and at lightning speed."
      },
      {
        Id : 6,
        content : "Morbi mattis ullamcorper velit. Donec orci lectus, aliquam ut, faucibus non, euismod id, nulla. Fusce convallis metus id felis luctus adipiscing. Aenean massa. Vestibulum purus quam, scelerisque ut, mollis sed, nonummy id, metus. Nulla consequat massa quis enim. Praesent venenatis metus at tortor pulvinar varius. Donec venenatis vulputate lorem. Phasellus accumsan cursus velit. Pellentesque ut neque. Inference is a powerful tool. You do it without thinking — your brain fills in the missing vowels in our company’s name effortlessly. It knows what’s missing. So do we. We have designed our technology to make these kinds of connections, and at lightning speed."
      },
    ],
    employeePolicy : [
      {
        Id : 1,
        content : "Morbi mattis ullamcorper velit. Donec orci lectus, aliquam ut, faucibus non, euismod id, nulla. Fusce convallis metus id felis luctus adipiscing. Aenean massa. Vestibulum purus quam, scelerisque ut, mollis sed, nonummy id, metus. Nulla consequat massa quis enim. Praesent venenatis metus at tortor pulvinar varius. Donec venenatis vulputate lorem. Phasellus accumsan cursus velit. Pellentesque ut neque. Inference is a powerful tool. You do it without thinking — your brain fills in the missing vowels in our company’s name effortlessly. It knows what’s missing. So do we. We have designed our technology to make these kinds of connections, and at lightning speed."
      },
      {
        Id : 2,
        content : "Morbi mattis ullamcorper velit. Donec orci lectus, aliquam ut, faucibus non, euismod id, nulla. Fusce convallis metus id felis luctus adipiscing. Aenean massa. Vestibulum purus quam, scelerisque ut, mollis sed, nonummy id, metus. Nulla consequat massa quis enim. Praesent venenatis metus at tortor pulvinar varius. Donec venenatis vulputate lorem. Phasellus accumsan cursus velit. Pellentesque ut neque. Inference is a powerful tool. You do it without thinking — your brain fills in the missing vowels in our company’s name effortlessly. It knows what’s missing. So do we. We have designed our technology to make these kinds of connections, and at lightning speed."
      },
      {
        Id : 3,
        content : "Morbi mattis ullamcorper velit. Donec orci lectus, aliquam ut, faucibus non, euismod id, nulla. Fusce convallis metus id felis luctus adipiscing. Aenean massa. Vestibulum purus quam, scelerisque ut, mollis sed, nonummy id, metus. Nulla consequat massa quis enim. Praesent venenatis metus at tortor pulvinar varius. Donec venenatis vulputate lorem. Phasellus accumsan cursus velit. Pellentesque ut neque. Inference is a powerful tool. You do it without thinking — your brain fills in the missing vowels in our company’s name effortlessly. It knows what’s missing. So do we. We have designed our technology to make these kinds of connections, and at lightning speed."
      },
      {
        Id : 4,
        content : "Morbi mattis ullamcorper velit. Donec orci lectus, aliquam ut, faucibus non, euismod id, nulla. Fusce convallis metus id felis luctus adipiscing. Aenean massa. Vestibulum purus quam, scelerisque ut, mollis sed, nonummy id, metus. Nulla consequat massa quis enim. Praesent venenatis metus at tortor pulvinar varius. Donec venenatis vulputate lorem. Phasellus accumsan cursus velit. Pellentesque ut neque. Inference is a powerful tool. You do it without thinking — your brain fills in the missing vowels in our company’s name effortlessly. It knows what’s missing. So do we. We have designed our technology to make these kinds of connections, and at lightning speed."
      },
      {
        Id : 5,
        content : "Morbi mattis ullamcorper velit. Donec orci lectus, aliquam ut, faucibus non, euismod id, nulla. Fusce convallis metus id felis luctus adipiscing. Aenean massa. Vestibulum purus quam, scelerisque ut, mollis sed, nonummy id, metus. Nulla consequat massa quis enim. Praesent venenatis metus at tortor pulvinar varius. Donec venenatis vulputate lorem. Phasellus accumsan cursus velit. Pellentesque ut neque. Inference is a powerful tool. You do it without thinking — your brain fills in the missing vowels in our company’s name effortlessly. It knows what’s missing. So do we. We have designed our technology to make these kinds of connections, and at lightning speed."
      },
      {
        Id : 6,
        content : "Morbi mattis ullamcorper velit. Donec orci lectus, aliquam ut, faucibus non, euismod id, nulla. Fusce convallis metus id felis luctus adipiscing. Aenean massa. Vestibulum purus quam, scelerisque ut, mollis sed, nonummy id, metus. Nulla consequat massa quis enim. Praesent venenatis metus at tortor pulvinar varius. Donec venenatis vulputate lorem. Phasellus accumsan cursus velit. Pellentesque ut neque. Inference is a powerful tool. You do it without thinking — your brain fills in the missing vowels in our company’s name effortlessly. It knows what’s missing. So do we. We have designed our technology to make these kinds of connections, and at lightning speed."
      },
    ],
  }),
}

</script>


<style lang="scss" scoped>
</style>