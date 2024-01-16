package report

import (
	"encoding/json"
	"github.com/disv5/opensdk/uc/model"
)

type GetFileRequest struct {
	Header model.HeaderTarget `json:"header"`
	Body   GetFileBody        `json:"body"`
}

type GetFileBody struct {
	TaskId int64 `json:"taskId"`
}

// Url implement GetRequest interface
func (r GetFileRequest) Url() string {
	return "api/report/getFile"
}

// Encode implement GetRequest interface
func (r GetFileRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
