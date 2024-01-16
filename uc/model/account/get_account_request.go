package account

import (
	"encoding/json"
	"github.com/disv5/opensdk/uc/model"
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
