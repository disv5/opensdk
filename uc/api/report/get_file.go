package report

import (
	"donson.com.cn/draining/internal/pkg/uc/core"
	"donson.com.cn/draining/internal/pkg/uc/model/report"
)

func GetFile(clt *core.SDKClient, req *report.GetFileRequest) (*report.GetFileResponse, error) {
	var resp report.GetFileResponse
	err := clt.Post(req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

