// Code generated by goctl. DO NOT EDIT.
package types

type CreateRuleReq struct {
	LintingRule
}

type CreateRuleResp struct {
	Id string `json:"id"`
}

type LintingRule struct {
	Name    string            `json:"name"`
	Project string            `json:"project"`
	Content map[string]string `json:"content"`
}

type UpdateRuleReq struct {
	Id string `json:"id"`
	LintingRule
}

type GetRuleReq struct {
	Id      string `form:"id,optional"`
	Name    string `form:"name,optional"`
	Project string `form:"project,optional"`
}

type GetRuleResp struct {
	Id string `json:"id"`
	LintingRule
}