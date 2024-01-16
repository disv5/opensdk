package adgroup

import (
	"encoding/json"
	"github.com/disv5/opensdk/uc/model"
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
