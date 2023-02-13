<template>
  <v-card>
    <v-card-title>Lightsail实例</v-card-title>
    <v-card-subtitle>Lightsail Instance</v-card-subtitle>
    <v-card-text>
      <v-row>
        <v-col
            cols="auto"
        >
          <v-select
              v-model="zoneSelected"
              :items="zone"
              :rules="formRequired"
              label="区域"
              solo
              dense
              hide-details
              @change="refreshLs"
              required
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
                <span class="text-h5">创建Lightsail</span>
              </v-card-title>
              <v-card-text>
                <v-container>
                  <v-form ref="createLsForm">
                    <v-row>
                      <v-col
                          cols="12"
                      >
                        <v-text-field
                            v-model="lsName"
                            :rules="formRequired"
                            label="名称"
                            required
                        ></v-text-field>
                      </v-col>
                      <v-col cols="12">
                        <v-select
                            v-model="availabilityZoneSelected"
                            :items="availabilityZone"
                            :rules="formRequired"
                            label="可用区"
                        ></v-select>
                      </v-col>
                      <v-col
                          cols="12"
                      >
                        <v-select
                            v-model="bundleSelected"
                            :items="bundle"
                            :rules="formRequired"
                            label="类型"
                            required
                        ></v-select>
                      </v-col>
                      <v-col
                          cols="12"
                      >
                        <v-select
                            v-model="blueprintSelected"
                            :items="blueprint"
                            :rules="formRequired"
                            label="操作系统"
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
                    @click="createLs"
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
        <v-col cols="1">
          <v-btn color="info" @click="refreshLs">刷新</v-btn>
        </v-col>
        <v-spacer></v-spacer>
        <v-col cols="12">
          <v-data-table
              :footer-props="{
              itemsPerPageText: '每页实例数',
            }"
              :headers="lsHeaders"
              :items="lsInstances"
              :items-per-page="5"
              :loading="lsTableLoading"
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
                      @click="openInstancePorts(item)"
                      v-on="on"
                  >
                    <v-icon>mdi-shield-lock-open</v-icon>
                  </v-btn>
                </template>
                <span>开放端口</span>
              </v-tooltip>
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
    <v-overlay v-model="loading">
      <v-progress-circular
          indeterminate
          size="64"
      ></v-progress-circular>
    </v-overlay>
  </v-card>
</template>

<script>
import axios from "axios"

export default {
  name: 'lightsailPage',
  props: ['loading', 'message', 'messageText'],
  data() {
    return {
      formRequired: [
        v => !!v || "必填项！"
      ],
      secretName: '',
      createDialog: false,
      sshKeyDialog: false,
      sshKey: "",
      lsName: "",
      zone: [
        {text: 'Tokyo', value: 'ap-northeast-1a'},
        {text: 'Seoul', value: 'ap-northeast-2a'},
        {text: 'Singapore', value: 'ap-southeast-1a'}
      ],
      zoneSelected: "",
      availabilityZone: [],
      availabilityZoneSelected: "",
      availabilityZoneStore: new Map(),
      bundle: [
        {text: 'Nano', value: 'nano_2_0'},
        {text: 'Micro', value: 'micro_2_0'},
        {text: 'Small', value: 'small_2_0'}
      ],
      bundleSelected: "",
      blueprint: [
        {text: 'Debian10', value: 'debian_10'},
        {text: 'Ubuntu20.04', value: 'ubuntu_20_04'}
      ],
      blueprintSelected: "",
      lsHeaders: [
        {text: '操作', value: 'Action'},
        {text: '名称', value: 'Name'},
        {text: '状态', value: 'Status'},
        {text: '类型', value: 'Type'},
        {text: 'IP', value: 'Ip'}
      ],
      lsInstances: [],
      lsTableLoading: false
    }
  },
  methods: {
    createCheck() {
      if (this.secretName === "" || this.zoneSelected === "") {
        this.messageText = '请先选择密钥和区域'
        this.message = true
      } else {
        this.createDialog = true
      }
    },
    createLs() {
      if (this.$refs.createLsForm.validate()) {
        this.createDialog = false
        this.loading = true
        let data = new FormData()
        data.append("secretName", this.secretName)
        data.append("name", this.lsName)
        data.append("zone", this.zoneSelected)
        data.append("availabilityZone",this.availabilityZoneSelected)
        data.append("bundleId", this.bundleSelected)
        data.append("blueprintId", this.blueprintSelected)
        axios.post("/api/v1/LightSail/Create", data, {withCredentials: true}).then(rsp => {
          if (rsp.data.code === 200) {
            this.messageText = "已添加至创建队列"
            this.sshKey=rsp.data.data
            this.sshKeyDialog = true
          }else{
            this.messageText = "操作失败"
          }
        }).catch(rsp => {
          this.messageText = "操作失败"
          if (rsp.response.data.msg !== undefined) {
            console.error(rsp.response.data.msg)
          }
        }).finally(() => {
          this.$refs.createLsForm.reset()
          this.loading = false
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
        this.refreshLs()
      })
    },
    getRegions() {
      this.availabilityZoneStore.clear()
      let data = new FormData()
      data.append("secretName", this.secretName)
      axios.post("/api/v1/LightSail/GetRegions", data, {withCredentials: true}).then(rsp => {
        if (rsp.data.code === 200) {
          let tmp = [];
          rsp.data.data.forEach((item) => {
            tmp.push({text: item.DisplayName, value: item.Name})
            if (item.AvailabilityZones.length === 0) {
              item.AvailabilityZones.push({"ZoneName": item.Name + "a"})
            }
            this.availabilityZoneStore.set(item.Name, item.AvailabilityZones)
          })
          this.zone = tmp
        }
      })
    },
    refreshLs() {
      if ((this.zoneSelected === '') || (this.secretName === '')) {
        return
      }
      this.availabilityZone = []
      this.availabilityZoneStore.get(this.zoneSelected).forEach(item => {
        this.availabilityZone.push(item.ZoneName)
      })
      this.lsTableLoading = true
      let data = new FormData()
      data.append("secretName", this.secretName)
      data.append("zone", this.zoneSelected)
      axios.post("/api/v1/LightSail/List", data, {withCredentials: true}).then(rsp => {
        if (rsp.data.code === 200) {
          if (rsp.data.data == null) {
            this.lsInstances = []
          } else {
            this.lsInstances = rsp.data.data
          }
        } else{
            this.messageText = "操作失败"
          }
      }).catch(rsp => {
        this.lsInstances = []
        if (rsp.response.data.msg !== undefined) {
          console.error(rsp.response.data.msg)
        }
        this.messageText = "操作失败"
      }).finally(() => {
        this.lsTableLoading = false
      })
    },
    openInstancePorts(item){
      if (item.Status !== '') {
        this.loading = true
        let data = new FormData()
        data.append('secretName', this.secretName)
        data.append('name', item.Name)
        data.append("zone", this.zoneSelected)
        axios.post("/api/v1/LightSail/OpenPorts", data,
            {withCredentials: true}).then(rsp => {
          if (rsp.data.code === 200) {
            this.messageText = "开放端口成功"
          }else{
            this.messageText = "操作失败"
          }
        }).catch(rsp => {
          if (rsp.response.data.msg !== undefined) {
            console.error(rsp.response.data.msg)
          }
          this.messageText = "操作失败"
        }).finally(() => {
          this.loading = false
          this.message = true
        })
      }
    },
    startInstance(item) {
        if (item.Status !== '') {
        this.loading = true
        let data = new FormData()
        data.append('secretName', this.secretName)
        data.append('name', item.Name)
        data.append("zone", this.zoneSelected)
        axios.post("/api/v1/LightSail/Start", data,
            {withCredentials: true}).then(rsp => {
          if (rsp.data.code === 200) {
            this.messageText = "已添加至启动队列"
          }else{
            this.messageText = "操作失败"
          }
        }).catch(rsp => {
          if (rsp.response.data.msg !== undefined) {
            console.error(rsp.response.data.msg)
          }
          this.messageText = "操作失败"
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
        data.append('secretName', this.secretName)
        data.append('name', item.Name)
        data.append("zone", this.zoneSelected)
        axios.post("/api/v1/LightSail/Stop", data,
            {withCredentials: true}).then(rsp => {
          if (rsp.data.code === 200) {
            this.messageText = "已添加至停止队列"
          }else{
            this.messageText = "操作失败"
          }
        }).catch(rsp => {
          if (rsp.response.data.msg !== undefined) {
            console.error(rsp.response.data.msg)
          }
          this.messageText = "操作失败"
        }).finally(() => {
          this.message = true
          this.loading = false
        })
      }
    },
    restartInstance(item) {
      if (item.Status !== '') {
        this.loading = true
        let data = new FormData()
        data.append('secretName', this.secretName)
        data.append('name', item.Name)
        data.append("zone", this.zoneSelected)
        axios.post("/api/v1/LightSail/Reboot", data,
            {withCredentials: true}).then(rsp => {
          if (rsp.data.code === 200) {
            this.messageText = "已添加至重启队列"
          }
        }).catch(rsp => {
          if (rsp.response.data.msg !== undefined) {
            console.error(rsp.response.data.msg)
          }
          this.messageText = "操作失败"
        }).finally(() => {
          this.message = true
          this.loading = false
        })
      }
    },
    deleteInstance(item) {
      if (item.Status !== '') {
        this.loading = true
        let data = new FormData()
        data.append('secretName', this.secretName)
        data.append('name', item.Name)
        data.append("zone", this.zoneSelected)
        data.append("resourceName", item.ResourceName)
        axios.post("/api/v1/LightSail/Delete", data,
            {withCredentials: true}).then(rsp => {
          if (rsp.data.code === 200) {
            this.messageText = "已添加至删除队列"
          }else{
            this.messageText = "操作失败"
          }
        }).catch(rsp => {
          if (rsp.response.data.msg !== undefined) {
            console.error(rsp.response.data.msg)
          }
          this.messageText = "操作失败"
        }).finally(() => {
          this.message = true
          this.loading = false
        })
      }
    }
  },
  mounted() {
    this.$on("secretSelect", secret => {
      this.secretName = secret
      this.getRegions()
      this.refreshLs()
    })
  }
}
</script>