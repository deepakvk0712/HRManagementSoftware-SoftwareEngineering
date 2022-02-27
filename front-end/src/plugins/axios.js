import ax from 'axios'

export const axios = ax
/* eslint-disable */
export default {
    install (Vue) {
        Vue.prototype.$axios = ax
    }
}
/* eslint-enable */
