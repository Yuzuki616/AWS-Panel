<template>
  <v-app>
    <v-container>
      <v-card>
        <v-card-title>注册</v-card-title>
        <v-card-subtitle>Register</v-card-subtitle>
        <v-card-text>
          <v-row>
            <v-col cols="5">
              <v-form ref="form">
                <v-form ref="codeForm">
                  <v-text-field
                      v-model="email"
                      :rules="emailRules"
                      label="邮箱"
                      name="email"
                      prepend-inner-icon="mdi-account-outline"
                      required
                      type="text"
                  ></v-text-field>
                </v-form>
                <v-row align="center" v-if="enableEmailVerity">
                  <v-col cols=8>
                    <v-text-field
                        v-model="emailCode"
                        :rules="emailRules"
                        label="验证码"
                        name="username"
                        prepend-inner-icon="mdi-account-outline"
                        required
                        type="text"
                    ></v-text-field>
                  </v-col>
                  <v-col cols="4">
                    <v-btn
                        block color="info"
                        v-bind:loading="loading"
                        v-bind:disabled="codeBtnDisable"
                        @click="sendEmailCode">
                      {{ codeBtnText }}
                    </v-btn>
                  </v-col>
                </v-row>
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
import axios from '../utils/api'
import sleep from '../utils/sleep'

export default {
  name: "registerPage",
  metaInfo: {
    title: '注册 - AWS Panel',
  },
  data() {
    return {
      text: '注册失败！',
      snackbar: false,
      email: "",
      emailRules: [
        v => !!v || "邮箱为必填项",
        v => /^[\w-.]+@([\w-]+\.)+[\w-]{2,4}$/.test(v) || "邮箱无效",
      ],
      enableEmailVerity: false,
      emailCode: '',
      emailCodeRules: [
        v => !!v || "验证码为必填项",
        v => /^[a-zA-Z0-9]{6}$/.test(v) || "验证码无效"
      ],
      codeBtnText: '发送验证码',
      codeBtnDisable: false,
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
  mounted() {
    axios.get('/api/v1/Config/IsEnableEmailVerify',).then(response => {
      this.enableEmailVerity = response.data.data
    })
  },
  methods: {
    sendEmailCode() {
      if (!this.$refs.codeForm.validate()) {
        return
      }
      this.loading = true
      axios.post('/api/v1/User/SendMailVerify', {
        email: this.email
      }).then(res => {
        if (res.data.code === 200) {
          this.text = res.data.msg
        } else {
          this.text = res.data.msg
        }
      }).catch(err => {
        this.text = err
      }).finally(() => {
        this.loading = false
        this.snackbar = true
        this.codeBtnDisable = true
        for (let i = 60; i >= 0; i--) {
          let s = i
          sleep.sleep(1000 * (60 - s)).then(() => {
            if (s === 0) {
              this.codeBtnText = "发送验证码"
              this.codeBtnDisable = false
            } else {
              this.codeBtnText = s
            }
          })
        }
      })
    },
    submit() {
      if (!this.$refs.codeForm.validate() || !this.$refs.form.validate()) {
        return
      }
      this.loading = true
      let data = new FormData();
      data.append("email", this.email)
      data.append("password", this.password)
      data.append("code", this.emailCode)
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
};
</script>