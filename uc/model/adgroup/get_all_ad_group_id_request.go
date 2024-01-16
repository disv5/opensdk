package adgroup

import (
	"donson.com.cn/draining/internal/pkg/uc/model"
	"encoding/json"
)

type GetAllAdGroupIdRequest struct {
	Header model.HeaderTarget `json:"header"`
}


// Url implement GetRequest interface
func (r GetAllAdGroupIdRequest) Url() string {
	return "api/adgroup/getAllAdGroupID"
}

// Encode implement GetRequest interface
func (r GetAllAdGroupIdRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
