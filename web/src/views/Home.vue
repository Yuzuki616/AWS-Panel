<template>
  <v-container>
    <v-row class="text-center">
      <v-col cols="12">
        <v-img
            :src="require('../assets/aws.png')"
            class="my-3"
            contain
            height="200"
        />
      </v-col>

      <v-col class="mb-4">
        <h1 class="display-2 font-weight-bold mb-3">
          欢迎使用AWS Panel
        </h1>

        <p class="subheading font-weight-regular">
          使用AWS Panel可以方便的管理AWS资源，比如EC2、Lightsail等。
        </p>
      </v-col>

      <v-col
          class="mb-5"
          cols="12"
      >
        <h2 class="headline font-weight-bold mb-3">
          Thanks
        </h2>

        <v-row justify="center">
          <a
              v-for="(link, i) in thanks"
              :key="i"
              :href="link.href"
              class="subheading mx-3"
              target="_blank"
          >
            {{ link.text }}
          </a>
        </v-row>
        <v-row justify="center">
          <p>感谢 <a href="https://www.jetbrains.com/">JetBrains</a> 为本项目提供的 <a
              href="https://www.jetbrains.com/webstorm/">WebStorm</a> 和 <a
              href="https://www.jetbrains.com/go/">Goland</a> IDE的免费使用授权</p>
        </v-row>
        <v-row>
          <v-img
              :src="require('../assets/jb_beam.svg')"
              contain
              height="300"
          />
        </v-row>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: 'homePage',
  metaInfo: {
    title: '主页 - AWS Panel',
  },
  data() {
    return {
      thanks: [
        {
          text: 'Vue.js',
          href: 'https://vuejs.org/',
        },
        {
          text: 'Vuetify',
          href: 'https://vuetifyjs.com/',
        },
        {
          text: 'Axios',
          href: 'https://github.com/axios/axios',
        },
        {
          text: 'Gin',
          href: 'https://github.com/gin-gonic/gin',
        },
        {
          text: 'aws-sdk-go',
          href: 'https://github.com/aws/aws-sdk-go'
        }
      ],
    }
  },
}
</script>