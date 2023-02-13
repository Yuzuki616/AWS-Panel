<template>
  <v-app>
    <v-navigation-drawer v-model="drawer" app clipped>
      <v-list
          nav
          dense
      >

        <v-list-item-group active-class="nattier-blue--text text--accent-4">
          <v-list-item onclick="window.location.href='/'">
            <v-list-item-icon>
              <v-icon>mdi-home</v-icon>
            </v-list-item-icon>
            <v-list-item-title>主页｜Home</v-list-item-title>
          </v-list-item>

          <v-list-item onclick="window.location.href='/Instance'">
            <v-list-item-icon>
              <v-icon>mdi-server</v-icon>
            </v-list-item-icon>
            <v-list-item-title>实例｜Instances</v-list-item-title>
          </v-list-item>

          <!--<v-list-item onclick="window.location.href='/Quota'">
            <v-list-item-icon>
              <v-icon>mdi-car-speed-limiter</v-icon>
            </v-list-item-icon>
            <v-list-item-title>配额 ｜ Quota</v-list-item-title>
          </v-list-item>-->

          <v-list-item onclick="window.location.href='/User'">
            <v-list-item-icon>
              <v-icon>mdi-account-details</v-icon>
            </v-list-item-icon>
            <v-list-item-title>用户｜User</v-list-item-title>
          </v-list-item>
          <v-list-item v-if="this.$cookie.get('isAdmin')==='true'" onclick="window.location.href='/Manger'">
            <v-list-item-icon>
              <v-icon>mdi-information-outline</v-icon>
            </v-list-item-icon>
            <v-list-item-title>管理 | Manger</v-list-item-title>
          </v-list-item>
        </v-list-item-group>
      </v-list>
    </v-navigation-drawer>

    <v-app-bar app clipped-left color="white">
      <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>

      <v-app-bar-title>
        <v-icon large>mdi-aws</v-icon>
        AWS Panel
      </v-app-bar-title>

      <v-spacer></v-spacer>

      <v-btn v-if="this.$cookie.get('loginSession')==null" href="/Register" text>
        <span class="mr-2">注册</span>
        <v-icon small>mdi-account-plus</v-icon>
      </v-btn>
      <v-btn v-else href="/User" text>
        <span class="mr-2">用户</span>
        <v-icon small>mdi-account-details</v-icon>
      </v-btn>

      <v-btn v-if="this.$cookie.get('loginSession')==null" href="/Login" text>
        <span class="mr-2">登陆</span>
        <v-icon small>mdi-account-key</v-icon>
      </v-btn>
      <v-btn v-else text @click="logout()">
        <span class="mr-2">注销</span>
        <v-icon small>mdi-logout</v-icon>
      </v-btn>

    </v-app-bar>

    <!-- 根据应用组件来调整你的内容 -->
    <v-main>

      <!-- 给应用提供合适的间距 -->
      <v-container fluid>

        <!-- 如果使用 vue-router -->
        <router-view></router-view>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import axios from './utils/api'
export default {
  name: 'App',
  data: () => ({
    drawer: null,
  }),
  methods: {
    logout() {
      this.$cookie.delete('loginSession');
      this.$cookie.delete('isAdmin');
      axios.get("/api/v1/User/Logout")
      this.$router.push('/Login')
    },
  },
  mounted() {
    if (this.$cookie.get('loginSession')!=null){
      axios.get("/api/v1/User/IsAdmin").then(rsp=>{
        if (this.$cookie.get('loginSession')!=null){
          if (rsp.status===401){
            this.$cookie.delete('loginSession');
            location.reload()
          }
        }
        if (rsp.data.msg===true){
          if (this.$cookie.get('isAdmin')==null){
            this.$cookie.set('isAdmin', 'true')
            location.reload()
          }
        }else{
          if (this.$cookie.get('isAdmin')!=null){
            this.$cookie.delete('isAdmin')
            location.reload()
          }
        }
      })
    }
  }
}
</script>

<style>
</style>