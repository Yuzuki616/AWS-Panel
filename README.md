# Aws-Panel

一个可以管理AWS资源的Web面板

[![](https://img.shields.io/github/license/Yuzuki999/AWS-Panel?style=for-the-badge)](https://www.gnu.org/licenses/gpl-3.0.html)

暂时未实现用户权限及用户管理，为避免风险，建议仅在本地使用

# 已实现的功能

- [x] 分用户多密钥
- [x] EC2管理
- [x] Lightsail管理
- [ ] AGA管理
- [ ] Wavelength管理
- [ ] 配额相关操作
- [x] 用户管理

# 使用

从Releases里下载可执行文件并运行，然后访问http://127.0.0.1:8011
初始管理员帐号密码为admin admin123456

# 构建

``` bash
git clone https://github.com/Yuzuki616/AWS-Panel.git
cd AWS-Panel
bash build.sh
```

# Thanks

- [Gin](https://github.com/gin-gonic/gin)
- [Aws-Sdk-Go](https://github.com/aws/aws-sdk-go)
- [Vue.js](https://vuejs.org/)
- [Vuetify](https://vuetifyjs.com/)
- [Axios](https://github.com/axios/axios)
