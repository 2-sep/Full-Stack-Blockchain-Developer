# Node.js

https://www.bilibili.com/video/BV1PV411d7XV/?spm_id_from=333.337.search-card.all.click&vd_source=5c46b11c8d760605427b8431fb93d551

## npm 包管理工具

npm init

npm install xxx

npm unintall xxx

mpm list

package.json 是 npm init 时生成的

作用：描述项目以及项目所依赖的模块信息

package-lock.json 是运行 npm install 时生成的一份文件

记录当前状态下实际安装的各个 npm 包的具体来源和版本号

本地安装

全局安装

开启 ES6 模块

package.json 添加 "type":"module"

## Yarn 包管理工具

## nodejs 与数据库交互

npm install mysql

**交互的基本模式**

导入库

定义 connection

connection.connect()

connection.query()

connection.end()

**封装执行 sql 语句函数**

https://www.bilibili.com/video/BV1KX4y1K7uz?p=15&vd_source=5c46b11c8d760605427b8431fb93d551