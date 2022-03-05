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
        accountType : "",
        loader : false,
    },

    mutations : {
        LOADER(state, payload){
            state.loader = payload
        }
    },
})