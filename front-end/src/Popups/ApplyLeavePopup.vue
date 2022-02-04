<template>
  <v-row justify="center">
    <v-dialog v-model="closePopup" persistent max-width="800px">
      <template v-slot:activator="{ on }">
        <v-btn v-on="on" text>
          <span style="margin-right: 10px"> Apply Leave </span>
          <v-icon justify="center"> exit_to_app </v-icon>
        </v-btn>
      </template>
      <v-card>
        <v-card-title>
          <span class="headline">Leave Application</span>
        </v-card-title>
        <v-form class="px-3" ref="form">
          <v-card-text>
            <v-text-field
              label="Subject"
              v-model="leaveApplicationTitle"
              prepend-icon="exit_to_app"
              :rules="inputRules"
            ></v-text-field>
            <v-textarea
              label="Please state reason for requesting leave."
              v-model="leaveRequestDescription"
              prepend-icon="edit"
              :rules="inputRules"
            ></v-textarea>
            <v-col cols="12" lg="6">
              <v-menu
                ref="menu1"
                v-model="menu1"
                :close-on-content-click="false"
                transition="scale-transition"
                offset-y
                max-width="290px"
                min-width="290px"
              >
                <template v-slot:activator="{ on }">
                  <v-text-field
                    v-model="fromDateFormatted"
                    label="From Date"
                    hint="MM/DD/YYYY format"
                    persistent-hint
                    prepend-icon="event"
                    @blur="date = parseDate(fromDateFormatted)"
                    v-on="on"
                  ></v-text-field>

                  <v-text-field
                    v-model="toDateFormatted"
                    label="To Date"
                    hint="MM/DD/YYYY format"
                    persistent-hint
                    prepend-icon="event"
                    @blur="date = parseDate(toDateFormatted)"
                    v-on="on"
                  ></v-text-field>
                </template>
                <v-date-picker
                  v-model="date"
                  no-title
                  @input="menu1 = false"
                ></v-date-picker>
              </v-menu>
            </v-col>
          </v-card-text>
          <v-col class="d-flex" cols="12" sm="6">
            <v-select
              :items="dropdownItems"
              label="Leave Type"
              dense
              outlined
              v-model="leaveTypeSelection"
            ></v-select>
          </v-col>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="grey darken-1" text @click="closePopup = false"
              >Cancel</v-btn
            >
            <v-btn
              color="green darken-1"
              text
              outlined
              @click="closePopup = false"
              >Apply</v-btn
            >
          </v-card-actions>
        </v-form>
      </v-card>
    </v-dialog>
  </v-row>
</template>

<script>
export default {
  data: (vm) => ({
    dropdownItems:[
        'Paid Leave', 'Un-Paid Leave'
    ],
    leaveTypeSelection:1,
    closePopup: false,
    leaveApplicationTitle: "",
    leaveRequestDescription: "",
    date: new Date().toISOString().substr(0, 10),
    fromDateFormatted: vm.formatDate(new Date().toISOString().substr(0, 10)),
    toDateFormatted: vm.formatDate(new Date().toISOString().substr(0, 10)),
    menu1: false,
    inputRules: [(v) => v.length >= 3 || "Minimum lenght is 3 charachters"],
  }),
  methods: {
    formatDate(date) {
      if (!date) return null;
      const [year, month, day] = date.split("-");
      return `${year}/${month}/${day}`;
    },
    parseDate(date) {
      if (!date) return null;
      const [year, month, day] = date.split("/");
      return `${year}-${month.padStart(2, "0")}-${day.padStart(2, "0")}`;
    },
  },
  computed: {
    computedDateFormatted() {
      return this.formatDate(this.date);
    },
  },
  watch: {
    date() {
      this.dateFormatted = this.formatDate(this.date);
    },
  },
};
</script>