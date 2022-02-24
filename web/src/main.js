import Vue from 'vue'
import App from './App.vue'
import router from './router'
import vuetify from "./plugins/vuetify";
import vueMeta from 'vue-meta'
import VueCookie from 'vue-cookie'

//const VueCookie = require('')
Vue.use(VueCookie,vueMeta)
Vue.config.productionTip = false
new Vue({
    vuetify,
    router,
    render: h => h(App)
}).$mount('#app')