<template>
  <v-app>
    <v-container>
      <v-card>
        <v-card-title>登陆</v-card-title>
        <v-card-subtitle>Login</v-card-subtitle>
        <v-card-text>
          <v-row>
            <v-col cols="5">
              <v-form ref="loginForm">
                <v-text-field
                    v-model="email"
                    :rules="emailRules"
                    label="邮箱"
                    name="email"
                    prepend-inner-icon="mdi-account-outline"
                    required
                    type="text"
                ></v-text-field>
                <v-text-field
                    id="password"
                    v-model="password"
                    :rules="passwordRules"
                    label="密码"
                    name="password"
                    prepend-inner-icon="mdi-lock"
                    required
                    type="password"
                ></v-text-field>
              </v-form>
            </v-col>
          </v-row>
        </v-card-text>
        <v-card-actions>
          <v-row>
            <v-col cols="3">
              <v-btn block color="info" v-bind:loading="loading" @click="submit">登陆</v-btn>
            </v-col>
          </v-row>
        </v-card-actions>
      </v-card>
      <v-snackbar
          v-model="snackbar"
      >
        {{ text }}

        <template v-slot:action="{ attrs }">
          <v-btn
              color="info"
              text
              v-bind="attrs"
              @click="snackbar = false"
          >
            Close
          </v-btn>
        </template>
      </v-snackbar>
    </v-container>
  </v-app>
</template>

<script>
import axios from "../utils/api";

export default {
  name: "loginPage",
  metaInfo: {
    title: '登陆 - AWS Panel',
  },
  data() {
    return {
      text: "登陆失败！",
      snackbar: false,
      email: "",
      emailRules: [
        v => !!v || "邮箱为必填项",
        v => /^[a-zA-Z0-9_-]{4,8}$/.test(v) || "邮箱无效"
      ],
      password: "",
      passwordRules: [v => !!v || "密码为必填项"],
      loading: false,
    }
  },
  methods: {
    submit() {
      if (this.$refs.loginForm.validate()) {
        this.loading = true
        let data = new FormData()
        data.append("email", this.email)
        data.append("password", this.password)
        axios.post("/api/v1/User/Login",
            data, {withCredentials: true}).then((response) => {
          if (response.data.code === 200) {
            if (response.data.isAdmin) {
              this.$cookie.set('isAdmin', 'true')
            }
            this.text = "登陆成功，即将跳转到主页..."
            setTimeout(() => {
              this.$router.push("/")
            }, 2000);
          } else {
            this.text = response.data.msg
          }
        }).catch((error) => {
          console.error(error)
        }).finally(() => {
          this.snackbar = true
          this.loading = false
        })
      }
    }
  }
};
</script>