# linting rule
author: Xiayu Qiu

version: v1

## kong-task-api

group: rule

middleware: Auth

### 1. 创建规则

##### POST /v1/rule

Request:

```golang
type CreateRuleReq struct {
	LintingRule
}
```


Response:

```golang
type CreateRuleResp struct {
	Id string `json:"id"`
}
```


### 2. 修改规则

##### PUT /v1/rule/:id

Request:

```golang
type UpdateRuleReq struct {
	Id string `path:"id"`
	LintingRule
}
```


Response:


### 3. 获取规则

##### GET /v1/rule

Request:

```golang
type GetRuleReq struct {
	Name    string `form:"name,optional"`
	Project string `form:"project,optional"`
}
```


Response:

```golang
type GetRuleResp struct {
	Id string `json:"id"`
	LintingRule
}
```


### 4. 添加用户到项目中

##### POST /v1/user

Request:

```golang
type UseInfo struct {
	Uid          uint32 `json:"uid"`                       // 添加到project的用户id
	Name         string `json:"name"`                      // 添加到project的用户名
	Project      string `json:"project"`                   // 项目名
	Organization string `json:"organization,default=kong"` // 组织名，默认为kong
	IsAdmin      bool   `json:"isAdmin"`                   // 是否为管理员
	ActionType   uint   `json:"actionType"`                // 权限
}
```


Response:


### 5. 更新用户在项目中的权限

##### PUT /v1/user/:uid

Request:

```golang
type UpdateUserReq struct {
	Uid          uint32 `path:"uid"`                       // 添加到project的用户id
	Project      string `json:"project"`                   // 项目名
	Organization string `json:"organization,default=kong"` // 组织名，默认为kong
	IsAdmin      bool   `json:"isAdmin"`                   // 是否为管理员
	ActionType   uint   `json:"actionType"`                // 权限
}
```


Response:


### 6. 从项目中移除用户

##### DELETE /v1/user/:uid

Request:

```golang
type RemoveUserReq struct {
	Uid          uint32 `path:"uid"`                       // 添加到project的用户id
	Project      string `json:"project"`                   // 项目名
	Organization string `json:"organization,default=kong"` // 组织名，默认为kong
}
```


Response:


### 7. 获取项目的用户列表

##### GET /v1/user

Request:

```golang
type GetUsersReq struct {
	Project      string `form:"project"`                   // 项目名
	Organization string `form:"organization,default=kong"` // 组织名，默认为kong
}
```


Response:

```golang
type GetUsersResp struct {
	Users []UseInfo `json:"users"`
}
type UseInfo struct {
	Uid          uint32 `json:"uid"`                       // 添加到project的用户id
	Name         string `json:"name"`                      // 添加到project的用户名
	Project      string `json:"project"`                   // 项目名
	Organization string `json:"organization,default=kong"` // 组织名，默认为kong
	IsAdmin      bool   `json:"isAdmin"`                   // 是否为管理员
	ActionType   uint   `json:"actionType"`                // 权限
}
```

