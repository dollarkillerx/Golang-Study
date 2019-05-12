Beego 开发
===
- cache
  - 文件
  - 内存
  - memcache
  - redis 推荐
- config
  - ini
  - json 推荐
  - xml
  - yaml 推荐
- context
  - request
  - response
- httplibs
  - 支持get post put delete head
  - 支持https
  - 支持超时设置
  - 支持文件上传
- logs
  - 多种输出引擎
  - 异步输出
- orm
- session
- toolbox
  - 定时任务
  - 监控
#### 工具引用
- bee new 新建项目结构
- bee run 自动编译部署
- bee generate 自动生成代码
```
bee new dollarkiller 创建项目
bee run 启动

bee generate scaffold脚手架 user表名称 -fields字段="id:int64,name:string,gender:int,age:int" -driver引擎=mysql -conn链接="root:passwprd@tcp(127.0.0.1:3306)/beego"
```
- go get github.com/go-sql-driver/mysql 添加mysql依赖
- 目录结构
  ```
    ├── conf  配置目录
    │   └── app.conf
    ├── controllers 控制器
    │   └── default.go
    ├── dollarkiller 二进制文件
    ├── main.go  文件入库
    ├── models  模型
    ├── routers 路由
    │   └── router.go
    ├── static 静态文件
    │   ├── css
    │   ├── img
    │   └── js
    │       └── reload.min.js
    ├── tests 测试目录
    │   └── default_test.go
    └── views 视图目录
        └── index.tpl

  ```
- 执行流程 go->router.init->controller
