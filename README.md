### 测试时环境
- go 1.18
- mongo 4.4.0
- redis 不需要（model层未使用缓存）

### 项目框架
- go-zero 1.5.1

### 准备工作
- 默认组织为kong，第一个admin用户需要手动插入。命令如下
```
db.user.insert({uid:1,name:"admin",project:"insomnia",organization:"kong",isAdmin:true,actionType:100})
```
- 上述命令字段含义参考 cmd/internal/model/user.go
- 接口文档 cmd/task.md
- 接口调用示例（Insomnia导出） cmd/Insomnia_2023-04-20.yaml 
- 执行命令:
  ```
    git clone https://github.com/qiushye/kong-task.git
    cd kong-task/cmd
    go run kong-task.go
  ```

 ### 访问控制原理
 通过header中的信息解析出uid，目前uid是明文传输，根据uid在db中查到的权限来决定接口的返回。
