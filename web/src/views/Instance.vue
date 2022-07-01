<template>
  <v-app>
    <v-container>
      <v-row class="mb-5">
        <v-col cols="12">
          <v-select
              v-model="secretSelected"
              :items="secrets"
              label="密钥"
              solo
              dense
              hide-details
              @change="secretSelect"
          ></v-select>
        </v-col>
        <v-col cols="12">
          <v-card>
            <v-card-title>
              Ec2实例
            </v-card-title>
            <v-card-subtitle>
              Ec2 Instances
            </v-card-subtitle>
            <v-card-text>
              <v-row>
                <v-col cols="auto">
                  <v-select
                      v-model="regionSelected"
                      :items="regions"
                      label="区域"
                      solo
                      dense
                      hide-details
                      @change="refresh"
                  ></v-select>
                </v-col>
                <v-col cols="auto">
                  <v-btn
                      color="info"
                      @click="createCheck"
                  >
                    创建
                  </v-btn>
                  <v-dialog
                      v-model="createDialog"
                      max-width="600px"
                  >
                    <v-card>
                      <v-card-title>
                        <span class="text-h5">创建Ec2</span>
                      </v-card-title>
                      <v-card-text>
                        <v-container>
                          <v-form ref="createForm">
                            <v-row>
                              <v-col
                                  cols="12"
                              >
                                <v-text-field
                                    v-model="ec2Name"
                                    :rules="formRequired"
                                    label="名称"
                                    required
                                ></v-text-field>
                              </v-col>
                              <v-col
                                  cols="12"
                              >
                                <v-select
                                    v-model="typeSelected"
                                    :items="type"
                                    :rules="formRequired"
                                    label="类型"
                                    required
                                ></v-select>
                              </v-col>
                              <v-col
                                  cols="12"
                              >
                                <v-select
                                    v-model="amiSelected"
                                    :items="ami"
                                    :rules="formRequired"
                                    label="操作系统"
                                    required
                                ></v-select>
                              </v-col>
                              <v-col
                                  cols="12"
                              >
                                <v-select
                                    v-model="diskSelected"
                                    :items="disk"
                                    :rules="formRequired"
                                    label="硬盘大小"
                                    required
                                ></v-select>
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
                            @click="createDialog = false"
                        >
                          取消
                        </v-btn>
                        <v-btn
                            color="blue darken-1"
                            text
                            @click="create"
                        >
                          创建
                        </v-btn>
                      </v-card-actions>
                    </v-card>
                  </v-dialog>
                  <v-dialog
                      v-model="sshKeyDialog"
                      max-width="600px"
                  >
                    <v-card>
                      <v-card-title>
                        <span class="text-h5">SSH密钥</span>
                      </v-card-title>
                      <v-card-text>
                        <v-textarea readonly v-model="sshKey"></v-textarea>
                      </v-card-text>
                      <v-card-actions>
                        <v-spacer></v-spacer>
                        <v-btn
                            color="blue darken-1"
                            text
                            @click="copySshKey"
                        >
                          复制
                        </v-btn>
                      </v-card-actions>
                    </v-card>
                  </v-dialog>
                </v-col>
                <v-col cols="auto">
                  <v-btn color="info" @click="refresh">刷新</v-btn>
                </v-col>
                <v-spacer></v-spacer>
                <v-col cols="12">
                  <v-data-table
                      :footer-props="{
              itemsPerPageText: '每页实例数',
            }"
                      :headers="headers"
                      :items="Instances"
                      :items-per-page="5"
                      :loading="tableLoading"
                      loading-text="正在加载..."
                      class="elevation-1"
                      disable-sort
                  >
                    <template slot="no-data">
                      <div>无任何实例</div>
                    </template>
                    <template v-slot:item.Action="{ item }">
                      <v-tooltip bottom>
                        <template v-slot:activator="{ on, attrs }">
                          <v-btn
                              color="info"
                              icon
                              v-bind="attrs"
                              @click="startInstance(item)"
                              v-on="on"
                          >
                            <v-icon>mdi-play</v-icon>
                          </v-btn>
                        </template>
                        <span>启动</span>
                      </v-tooltip>

                      <v-tooltip bottom>
                        <template v-slot:activator="{ on, attrs }">
                          <v-btn
                              color="info"
                              icon
                              v-bind="attrs"
                              @click="stopInstance(item)"
                              v-on="on"
                          >
                            <v-icon>mdi-stop</v-icon>
                          </v-btn>
                        </template>
                        <span>停止</span>
                      </v-tooltip>

                      <v-tooltip bottom>
                        <template v-slot:activator="{ on, attrs }">
                          <v-btn
                              color="info"
                              icon
                              v-bind="attrs"
                              @click="restartInstance(item)"
                              v-on="on"
                          >
                            <v-icon>mdi-restart</v-icon>
                          </v-btn>
                        </template>
                        <span>重启</span>
                      </v-tooltip>

                      <v-tooltip bottom>
                        <template v-slot:activator="{ on, attrs }">
                          <v-btn
                              color="info"
                              icon
                              v-bind="attrs"
                              @click="deleteInstance(item)"
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
        <v-col cols="12">
          <lightsail
              ref="lightsail"
              :loading="loading"
              :message="message"
              :messageText="messageText"
          ></lightsail>
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
import axios from "../api"
import Lightsail from "../components/lightsail";

export default {
  name: 'Instances',
  components: {Lightsail},
  metaInfo: {
    title: '实例 - AWS Panel',
  },
  data: () => ({
    formRequired: [
      v => !!v || "必填项！"
    ],
    createDialog: false,
    sshKeyDialog: false,
    sshKey: '',
    //secretsLoading: false,
    tableLoading: false,
    regions: [
      {text: '弗吉尼亚北部', value: 'us-east-1'},
      {text: '俄亥俄', value: 'us-east-2'},
      {text: '加利福尼亚北部', value: 'us-west-1'},
      {text: '俄勒冈', value: 'us-west-2'},
      {text: "香港", value: "ap-east-1"},
      {text: '日本', value: 'ap-northeast-1'},
      {text: "大阪", value: "ap-northeast-3"},
      {text: '韩国', value: 'ap-northeast-2'},
      {text: '加拿大', value: 'ca-central-1'},
      {text: '墨西哥', value: 'us-east-3'},
      {text: '澳大利亚', value: 'ap-southeast-6'},
      {text: '新加坡', value: 'ap-southeast-1'},
      {text: '澳洲', value: 'ap-southeast-2'},
      {text: '法国', value: 'eu-central-1'},
      {text: '冰岛', value: 'eu-west-1'},
      {text: '德国', value: 'eu-west-2'},
      {text: '俄罗斯', value: 'eu-west-3'},
      {text: '印度', value: 'ap-south-1'},
      {text: '新西兰', value: 'ap-southeast-3'},
      {text: '马来西亚', value: 'ap-southeast-5'},
      {text: '泰国', value: 'ap-northeast-3'},
      {text: '菲律宾', value: 'ap-southeast-4'},
      {text: '西班牙', value: 'eu-west-4'},
      {text: '瑞士', value: 'eu-west-6'},
    ],
    regionSelected: '',
    secrets: [],
    secretSelected: '',
    headers: [
      {text: '操作', value: 'Action'},
      {text: '名称', value: 'Name'},
      {text: '实例ID', value: 'InstanceId'},
      {text: '状态', value: 'Status'},
      {text: '类型', value: 'Type'},
      {text: 'IP', value: 'Ip'}
    ],
    ec2Name: '',
    type: [
      't2.micro',
      't2.small',
      't2.medium',
      't2.large',
      't2.xlarge',
      't2.2xlarge',
    ],
    typeSelected: '',
    ami: [
      {text: 'Debian10', value: 'debian-10-amd64-20210329-591'},
      {text: 'Ubuntu20.04', value: 'ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-20210430'},
      {text: 'Redhat8', value: 'RHEL_HA-8.4.0_HVM-20210504-x86_64-2-Hourly2-GP2'}
    ],
    amiSelected: '',
    disk: [
      {text: '8GB', value: 8},
      {text: '16GB', value: 16},
      {text: '32GB', value: 32},
      {text: '64GB', value: 64},
      {text: '128GB', value: 128},
      {text: '256GB', value: 256},
      {text: '512GB', value: 512},
      {text: '1024GB', value: 1024},
    ],
    diskSelected: '',
    Instances: [],
    message: false,
    messageText: '',
    loading: false,
  }),
  mounted() {
    axios.get("/api/v1/Secret/List", {withCredentials: true}).then(response => {
      let tmp = []
      if (response.data.code === 200) {
        for (const v of response.data.data) {
          tmp.push({
            text: v.name,
            id: v.id,
            secret: v.secret
          })
        }
        this.secrets = tmp
      } else {
        console.error("load secret error: ", response.data.msg)
      }
    })
  },
  methods:
      {
        secretSelect() {
          this.refresh()
          this.$refs.lightsail.$emit("secretSelect", this.secretSelected)
        },
        createCheck() {
          if ((this.regionSelected === '') && (this.secretSelected === '')) {
            console.log(this.regionSelected, this.secretSelected)
            this.messageText = '请先选择密钥和区域'
            this.message = true
          } else {
            this.createDialog = true
          }
        },
        refresh() {
          if ((this.regionSelected !== '') && (this.secretSelected !== '')) {
            this.tableLoading = true
            let data = new FormData()
            data.append("region", this.regionSelected)
            data.append("secretName", this.secretSelected)
            axios.post('/api/v1/Ec2/List', data, {withCredentials: true}).then(response => {
              if (response.data.code === 200) {
                if (response.data.data == null) {
                  this.Instances = []
                } else {
                  this.Instances = response.data.data
                }
              } else {
                this.Instances = []
                console.error(response.data.msg)
              }
            }).finally(() => {
              this.tableLoading = false
            });
          }
        },
        create() {
          if (this.$refs.createForm.validate()) {
            this.createDialog = false
            this.loading = true
            let data = new FormData()
            data.append("region", this.regionSelected)
            data.append("secretName", this.secretSelected)
            data.append("ec2Name", this.ec2Name)
            data.append("ec2Type", this.typeSelected)
            data.append("ami", this.amiSelected)
            data.append("disk", this.diskSelected)
            axios.post("/api/v1/Ec2/Create", data, {withCredentials: true}).then(response => {
              if (response.data.code === 200) {
                this.messageText = '已添加至创建队列！'
                this.sshKey = response.data.data
              }
            }).catch(rsp =>{
              if (rsp.response.data.msg !== undefined) {
            console.error(rsp.response.data.msg)
          }
              this.messageText="操作失败"
            }).finally(() => {
              this.$refs.createForm.reset()
              this.loading = false
              this.sshKeyDialog = true
              this.message = true
            })
          }
        },
        copySshKey() {
          this.sshKeyDialog = false
          navigator.clipboard.writeText(this.sshKey).then(() => {
            this.messageText = '已复制到剪贴板'
            this.message = true
            this.sshKey = ''
            this.refresh()
          })
        },
        startInstance(item) {
          if (item.Status !== '') {
            this.loading = true
            let data = new FormData()
            data.append("region", this.regionSelected)
            data.append("secretName", this.secretSelected)
            data.append("ec2Id", item.InstanceId)
            axios.post("/api/v1/Ec2/Start", data, {withCredentials: true}).then(response => {
              if (response.data.code === 200) {
                this.messageText = '已添加至启动队列！'
                this.refresh()
              }
            }).catch(rsp =>{
              if (rsp.response.data.msg !== undefined) {
            console.error(rsp.response.data.msg)
          }
              this.messageText="操作失败"
            }).finally(() => {
              this.loading = false
              this.message = true
            })
          }
        },
        stopInstance(item) {
          if (item.Status !== '') {
            this.loading = true
            let data = new FormData()
            data.append("region", this.regionSelected)
            data.append("secretName", this.secretSelected)
            data.append("ec2Id", item.InstanceId)
            axios.post("/api/v1/Ec2/Stop", data, {withCredentials: true}).then(response => {
              if (response.data.code === 200) {
                this.messageText = '已添加至停止队列！'
                this.refresh()
              }
            }).catch(rsp =>{
              if (rsp.response.data.msg !== undefined) {
            console.error(rsp.response.data.msg)
          }
              this.messageText="操作失败"
            }).finally(() => {
              this.loading = false
              this.message=true
            })
          }
        },
        restartInstance(item) {
          if (item.Status !== '') {
            this.loading = true
            let data = new FormData()
            data.append("region", this.regionSelected)
            data.append("secretName", this.secretSelected)
            data.append("ec2Id", item.InstanceId)
            axios.post("/api/v1/Ec2/Reboot", data, {withCredentials: true}).then(response => {
              if (response.data.code === 200) {
                this.messageText = '已添加至重启队列！'
                this.refresh()
              }
            }).catch(rsp =>{
              if (rsp.response.data.msg !== undefined) {
            console.error(rsp.response.data.msg)
          }
              this.messageText="操作失败"
            }).finally(() => {
              this.message=true
              this.loading = false
            })
          }
        },
        deleteInstance(item) {
          if (item.Status !== '') {
            this.loading = true
            let data = new FormData()
            data.append("region", this.regionSelected)
            data.append("secretName", this.secretSelected)
            data.append("ec2Id", item.InstanceId)
            axios.post("/api/v1/Ec2/Delete", data, {withCredentials: true}).then(response => {
              if (response.data.code === 200) {
                this.messageText = '已添加至删除队列！'
                this.refresh()
              }
            }).catch(rsp =>{
              if (rsp.response.data.msg !== undefined) {
            console.error(rsp.response.data.msg)
          }
              this.messageText="操作失败"
            }).finally(() => {
              this.message=true
              this.loading = false
            })
          }
        },
      },
}


</script>