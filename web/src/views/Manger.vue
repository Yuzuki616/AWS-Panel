<template>
  <v-app>
    <v-container>
      <v-card>
        <v-card-title>
          用户列表
        </v-card-title>
        <v-card-subtitle>
          User List
        </v-card-subtitle>
        <v-card-text>
          <v-col cols="12">
            <v-row>
              <!--<v-col cols="auto">
                <v-btn
                    color="info"
                    @click="addUserDialog=true"
                >
                  添加
                </v-btn>
                <v-dialog
                    v-model="addUserDialog"
                    max-width="600px"
                >
                  <v-card>
                    <v-card-title>
                      <span class="text-h5">添加用户</span>
                    </v-card-title>
                    <v-card-text>
                      <v-container>
                        <v-form ref="addForm">
                          <v-row>
                            <v-col
                                cols="12"
                            >
                              <v-text-field
                                  v-model="addUserName"
                                  :rules="formRequired"
                                  label="用户名"
                                  required
                              ></v-text-field>
                            </v-col>
                            <v-col cols="12">
                              <v-text-field
                                  v-model="addPassword"
                                  :rules="formRequired"
                                  label="密码"
                                  required
                              ></v-text-field>
                            </v-col>
                            <v-col cols="12">
                              <v-text-field
                                  v-model="addConfirmPassword"
                                  :rules="formRequired"
                                  label="确认密码"
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
                          @click="addUserDialog = false"
                      >
                        取消
                      </v-btn>
                      <v-btn
                          :loading="addLoading"
                          color="blue darken-1"
                          text
                          @click="addUser()"
                      >
                        添加
                      </v-btn>
                    </v-card-actions>
                  </v-card>
                </v-dialog>
              </v-col>-->
              <v-col cols="auto">
                <v-btn color="info" @click="refresh">刷新</v-btn>
              </v-col>
              <v-spacer></v-spacer>
              <v-col cols="12">
                <v-data-table
                    disable-sort
                    :footer-props="{
                      itemsPerPageText: '每页用户数',
                      }"
                    :headers="headers"
                    :items="users"
                    :items-per-page=5
                    :loading="userLoading"
                    loading-text="正在加载..."
                    class="elevation-1">
                  <template slot="no-data">
                    <div>无任何用户</div>
                  </template>
                  <template v-slot:item.Action="{ item }">
                    <v-tooltip bottom>
                      <template v-slot:activator="{ on, attrs }">
                        <v-btn
                            color="info"
                            icon
                            v-bind="attrs"
                            @click="deleteUser(item)"
                            v-on="on"
                        >
                          <v-icon>mdi-delete</v-icon>
                        </v-btn>
                      </template>
                      <span>删除</span>
                    </v-tooltip>
                    <v-tooltip bottom>
                      <template v-slot:activator="{ on, attrs }">
                        <v-btn
                            color="info"
                            icon
                            v-bind="attrs"
                            @click="banUser(item)"
                            v-on="on"
                        >
                          <v-icon>mdi-lock</v-icon>
                        </v-btn>
                      </template>
                      <span>封禁</span>
                    </v-tooltip>
                    <v-tooltip bottom>
                      <template v-slot:activator="{ on, attrs }">
                        <v-btn
                            color="info"
                            icon
                            v-bind="attrs"
                            @click="unBanUser(item)"
                            v-on="on"
                        >
                          <v-icon>mdi-lock-open</v-icon>
                        </v-btn>
                      </template>
                      <span>解封</span>
                    </v-tooltip>
                  </template>
                </v-data-table>
              </v-col>
            </v-row>
          </v-col>
        </v-card-text>
      </v-card>
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
import api from '../utils/api'

export default {
  name: 'mangerPage',
  metaInfo: {
    title: '管理中心 - AWS Panel',
  },
  data() {
    return {
      loading: false,
      message: false,
      messageText: '',
      formRequired: [
        v => !!v || "必填项！"
      ],
      addUserDialog: false,
      addUserName: '',
      addPassword: '',
      addConfirmPassword: '',
      addLoading: false,
      userLoading: false,
      headers: [
        {text: '操作', value: 'Action'},
        {text: '用户名', value: 'UserName'},
        {text: '状态', value: 'Status'},
        {text: '管理权限', value: 'IsAdmin'},
      ],
      users: [],
    }
  },
  methods: {
    /*addUser(){
    },*/
    refresh() {
      this.userLoading = true;
      api.get("/api/v1/User/List").then(rsp => {
        if (rsp.data.code === 200) {
          let tmp = []
          for (const v of rsp.data.data) {
            tmp.push({
              Email: v.Email,
              Status: v.Status ? '封禁' : '正常',
              IsAdmin: v.IsAdmin ? '是' : '否',
            })
          }
          this.users = tmp
        } else {
          console.log("list user list error: ", rsp.data.msg)
        }
      }).finally(() => {
        this.userLoading = false
      })
    },
    deleteUser(item) {
      let data = new FormData()
      data.append("email", item.Email)
      api.post("/api/v1/User/Delete",data).then(rsp => {
        this.messageText = rsp.data.msg
        this.message = true
      }).finally(()=>{
        this.refresh()
      })
    },
    banUser(item) {
      let data = new FormData()
      data.append("email", item.Email)
      api.post("/api/v1/User/Ban",data).then(rsp => {
        this.messageText = rsp.data.msg
        this.message = true
      }).finally(()=>{
        this.refresh()
      })
    },
    unBanUser(item){
      let data = new FormData()
      data.append("email", item.Email)
      api.post("/api/v1/User/UnBan",data).then(rsp => {
        this.messageText = rsp.data.msg
        this.message = true
      }).finally(()=>{
        this.refresh()
      })
    }
  },
  mounted() {
    this.refresh()
  }
}

</script>