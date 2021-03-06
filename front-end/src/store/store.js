import Vue from 'vue'
import Vuex from 'vuex'
Vue.use(Vuex)
// const state = {
//   accessToken : "",
//   userName : "",
//   accountType : "",
//   loader : false,
// }
export default new Vuex.Store({
    state : {
        accessToken : "",
        userName : "",
        isHR : false,
        isManager : false,
        isOnboard : false,
        isFinance : false,
        teamMembers : [],
        businessUnits : [],
        notifications : [],
        paidLeaves : 20,
        unpaidLeaves : 20,
        // accountType : "",
        loader : false,
    },

    mutations : {
        LOADER(state, payload){
            state.loader = payload
        }
    },
})