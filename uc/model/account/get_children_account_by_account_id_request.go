package account

import (
	"donson.com.cn/draining/internal/pkg/uc/model"
	"encoding/json"
)

type GetChildrenAccountByAccountIdRequest struct {
	Header model.Header                      `json:"header"`
	Body   GetChildrenAccountByAccountIdBody `json:"body"`
}

type GetChildrenAccountByAccountIdBody struct {
	AgentType uint `json:"agentType"` // 代理类型(非必填，默认为商务代理 0-商务代理 ;1-服务代理)
}

// Url implement GetRequest interface
func (r GetChildrenAccountByAccountIdRequest) Url() string {
	return "api/account/getChildrenAccountByAccountId"
}

// Encode implement GetRequest interface
func (r GetChildrenAccountByAccountIdRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
