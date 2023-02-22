<template>
  <v-app>
    <v-container>
      <v-row class="mb-5">
        <v-col
            cols="12">
          <v-card>
            <v-card-title>
              用户中心
            </v-card-title>
            <v-card-subtitle>
              User Center
            </v-card-subtitle>
            <v-card-text>
              欢迎回来，{{ email }}
            </v-card-text>
          </v-card>
        </v-col>

        <v-col
            cols="12">
          <v-card>
            <v-card-title>
              修改密码
            </v-card-title>
            <v-card-subtitle>
              Change Password
            </v-card-subtitle>
            <v-card-text>
              <v-form ref="changeForm">
                <v-row>
                  <v-col cols="12">
                    <v-text-field
                        v-model="oldPw"
                        :rules="formRequired"
                        label="原密码"
                        required
                        type="password"
                    ></v-text-field>
                  </v-col>
                  <v-col cols="6">
                    <v-text-field
                        v-model="newPw"
                        :rules="formRequired"
                        label="新密码"
                        required
                        type="password"
                    ></v-text-field>
                  </v-col>
                  <v-col cols="6">
                    <v-text-field
                        v-model="confirmNewPw"
                        :rules="confirmPwRules"
                        label="确认新密码"
                        required
                        type="password"
                    ></v-text-field>
                  </v-col>
                  <v-col cols="3">
                    <v-btn block color="info" v-bind:loading="changeLoading" @click="changePassword">修改</v-btn>
                  </v-col>
                </v-row>
              </v-form>
            </v-card-text>
          </v-card>
        </v-col>

        <v-col
            cols="12">
          <v-card>
            <v-card-title>
              密钥管理
            </v-card-title>
            <v-card-subtitle>
              Secret Manger
            </v-card-subtitle>
            <v-card-text>
              <v-row>
                <v-col cols="12">
                  <v-row>
                    <v-col cols="auto">
                      <v-btn
                          color="info"
                          @click="addSecretDialog=true"
                      >
                        添加
                      </v-btn>
                      <v-dialog
                          v-model="addSecretDialog"
                          max-width="600px"
                      >
                        <v-card>
                          <v-card-title>
                            <span class="text-h5">添加密钥</span>
                          </v-card-title>
                          <v-card-text>
                            <v-container>
                              <v-form ref="addForm">
                                <v-row>
                                  <v-col
                                      cols="12"
                                  >
                                    <v-text-field
                                        v-model="secretName"
                                        :rules="formRequired"
                                        label="名称"
                                        required
                                    ></v-text-field>
                                  </v-col>
                                  <v-col cols="12">
                                    <v-text-field
                                        v-model="secretId"
                                        :rules="formRequired"
                                        label="密钥ID"
                                        required
                                    ></v-text-field>
                                  </v-col>
                                  <v-col cols="12">
                                    <v-text-field
                                        v-model="secret"
                                        :rules="formRequired"
                                        label="密钥"
                                        required></v-text-field>
                                  </v-col>
                                </v-row>
                              </v-form>
                            </v-container>
                          </v-card-text>
                          <v-card-actions>
                            <v-spacer></v-spacer>
                            <v-btn
                                color="blue darken-1"
                                text
                                @click="addSecretDialog = false"
                            >
                              取消
                            </v-btn>
                            <v-btn
                                :loading="addLoading"
                                color="blue darken-1"
                                text
                                @click="addSecret"
                            >
                              添加
                            </v-btn>
                          </v-card-actions>
                        </v-card>
                      </v-dialog>
                    </v-col>
                    <v-col cols="auto">
                      <v-btn
                          color="info"
                          @click="refresh()"
                      >
                        刷新
                      </v-btn>
                    </v-col>
                    <v-spacer></v-spacer>
                  </v-row>
                </v-col>
                <v-col cols="12">
                  <v-data-table
                      disable-sort
                      :footer-props="{
                      itemsPerPageText: '每页密钥数',
                      }"
                      :headers="headers"
                      :items="secrets"
                      :items-per-page=5
                      :loading="secretsLoading"
                      loading-text="正在加载..."
                      class="elevation-1">
                    <template slot="no-data">
                      <div>无任何密钥</div>
                    </template>
                    <template v-slot:item.action="{ item }">
                      <v-tooltip bottom>
                        <template v-slot:activator="{ on, attrs }">
                          <v-btn
                              color="info"
                              icon
                              v-bind="attrs"
                              @click="deleteSecrets(item)"
                              v-on="on"
                          >
                            <v-icon>mdi-delete</v-icon>
                          </v-btn>
                        </template>
                        <span>删除</span>
                      </v-tooltip>
                    </template>
                  </v-data-table>
                </v-col>
              </v-row>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
      <v-overlay v-model="loading">
        <v-progress-circular
            indeterminate
            size="64"
        ></v-progress-circular>
      </v-overlay>
      <v-snackbar
          v-model="message"
      >
        {{ messageText }}
        <template v-slot:action="{ attrs }">
          <v-btn
              color="info"
              text
              v-bind="attrs"
              @click="message = false"
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

export default {
  name: 'userPage',
  metaInfo: {
    title: '用户中心 - AWS Panel',
  },
  data() {
    return {
      email: '未知用户',
      oldPw: '',
      newPw: '',
      confirmNewPw: '',
      formRequired: [
        v => !!v || "必填项！"
      ],
      confirmPwRules: [
        v => !!v || "必填项!",
        v => v === this.newPw || "密码不匹配"
      ],
      changeLoading: false,
      addSecretDialog: false,
      secretName: null,
      secretId: null,
      secret: null,
      addLoading: false,
      headers: [
        {text: "操作", value: "action"},
        {text: "密钥名称", value: "name"},
      ],
      secrets: [],
      secretsLoading: false,
      messageText: '',
      message: false,
      loading: false
    }
  },
  mounted() {
    axios.get('/api/v1/User/Info', {withCredentials: true}).then(response => {
      this.email = response.data.data
    })
    axios.get('/api/v1/Secret/List', {withCredentials: true}).then(response => {
      let tmp = []
      for (const v of response.data.data) {
        tmp.push({
          name: v.name,
          id: v.id,
          secret: v.secret
        })
        this.secrets = tmp
      }
    })
  },
  methods: {
    refresh() {
      this.secretsLoading = true
      axios.get('/api/v1/Secret/List', {withCredentials: true}).then(response => {
        let tmp = []
        if (response.data.data!=null){
          response.data.data.forEach(v=>{
            tmp.push({
              name: v.name,
              id: v.id,
              secret: v.secret
            })
          })
        }
        this.secrets = tmp
      }).finally(() => {
        this.secretsLoading = false
      })
    },
    changePassword() {
      if (this.$refs.changeForm.validate()) {
        this.changeLoading = true
        let data = new FormData()
        data.append('oldPassword', this.oldPw)
        data.append('newPassword', this.newPw)
        axios.post('/api/v1/User/ChangePassword', data, {withCredentials: true}).then(response => {
          if (response.data.code === 200) {
            this.$cookie.delete('loginSession')
            this.messageText = '修改成功！'
            this.message = true
            setTimeout(() => {
              this.$router.push("/Login")
            }, 2000);
          } else {
            this.messageText = response.data.msg
          }
        }).finally(()=>{
          this.$refs.changeForm.reset()
          this.changeLoading = false
          this.message = true
        })
      }
    },
    deleteSecrets(item) {
      this.loading = true
      let data = new FormData()
      data.append('name', item.name)
      axios.post('api/v1/Secret/Delete', data, {withCredentials: true}).then(response => {
        if (response.data.code === 200) {
          this.messageText = '删除成功！'
        } else {
          this.messageText = response.data.msg
        }
      }).finally(() => {
        this.loading = false
        this.message = true
        this.refresh()
      })
    },
    addSecret() {
      this.addLoading = true
      let data = new FormData()
      data.append('name', this.secretName)
      data.append('id', this.secretId)
      data.append('secret', this.secret)
      axios.post('api/v1/Secret/Add', data, {withCredentials: true}).then(response => {
        if (response.data.code === 200) {
          this.messageText = '添加成功'
        } else {
          this.messageText = response.data.msg
        }
      }).finally(() => {
        this.$refs.addForm.reset()
        this.addLoading = false
        this.addSecretDialog = false
        this.message = true
        this.refresh()
      })
    }
  }
}
</script>