package report

import (
	"encoding/json"
	"github.com/disv5/opensdk/uc/model"
)

type DownloadFileRequest struct {
	Header model.HeaderTarget `json:"header"`
	Body   DownloadFileBody   `json:"body"`
}

type DownloadFileBody struct {
	TaskId int64 `json:"taskId"`
}

// Url implement GetRequest interface
func (r DownloadFileRequest) Url() string {
	return "api/report/downloadFile"
}

// Encode implement GetRequest interface
func (r DownloadFileRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
