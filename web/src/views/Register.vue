<template>
  <v-app>
    <v-container>
        <v-card>
          <v-card-title>注册</v-card-title>
          <v-card-subtitle>Register</v-card-subtitle>
          <v-card-text>
            <v-row>
              <v-col cols="5">
                <v-form ref="registerForm">
                  <v-text-field
                      v-model="username"
                      :rules="usernameRules"
                      label="用户名"
                      name="username"
                      prepend-inner-icon="mdi-account-outline"
                      required
                      type="text"
                  ></v-text-field>
                  <v-text-field
                      id="password"
                      v-model="password"
                      :rules="passwordRules"
                      label="密码"
                      name="Password"
                      prepend-inner-icon="mdi-lock"
                      required
                      type="password"
                  ></v-text-field>
                  <v-text-field
                      id="confirmPassword"
                      v-model="confirmPassword"
                      :rules="confirmPasswordRules"
                      label="确认密码"
                      name="Confirm Password"
                      prepend-inner-icon="mdi-lock"
                      required
                      type="password"></v-text-field>
                </v-form>
              </v-col>
            </v-row>
          </v-card-text>
          <v-card-actions>
            <v-row>
              <v-col cols="3">
                <v-btn block color="info" v-bind:loading="loading" @click="submit">注册</v-btn>
              </v-col>
              <v-spacer></v-spacer>
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
import axios from '../api'

export default {
  name: "register",
  metaInfo: {
    title: '注册 - AWS Panel',
  },
  data() {
    return {
      text: '注册失败！',
      snackbar: false,
      username: "",
      usernameRules: [
        v => !!v || "用户名为必填项",
        v => /^[a-zA-Z0-9_-]{4,8}$/.test(v) || "用户名无效，请输入4-8位字母、数字、下划线或减号"
      ],
      password: '',
      passwordRules: [v => !!v || "密码为必填项"],
      confirmPassword: "",
      confirmPasswordRules: [
        v => !!v || "确认密码为必填项",
        v => v === this.password || "密码不匹配"
      ],
      loading: false,
    }
  },
  methods: {
    submit() {
      if (this.$refs.registerForm.validate()) {
        this.loading = true
        let data = new FormData();
        data.append("username", this.username)
        data.append("password", this.password)
        axios.post("/api/v1/User/Register",
            data).then((response) => {
          if (response.data.code === 200) {
            this.text = '注册成功！即将跳转至登陆页面...'
            setTimeout(() => {
              this.$router.push('/Login')
            }, 2000)
          } else {
            this.text = response.data.msg
          }
        }).catch((error) => {
          console.log(error)
          this.loading = false
        }).finally(() => {
          this.snackbar = true
          this.loading = false
        })
      }
    }
  }
};
</script>