package account

import (
	"donson.com.cn/draining/internal/pkg/uc/model"
	"encoding/json"
)

type GetAccountRequest struct {
	Header model.HeaderTarget `json:"header"`
}


// Url implement GetRequest interface
func (r GetAccountRequest) Url() string {
	return "api/account/getAccount"
}

// Encode implement GetRequest interface
func (r GetAccountRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
