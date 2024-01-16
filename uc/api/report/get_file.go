package report

import (
	"github.com/disv5/opensdk/uc/core"
	"github.com/disv5/opensdk/uc/model/report"
)

func GetFile(clt *core.SDKClient, req *report.GetFileRequest) (*report.GetFileResponse, error) {
	var resp report.GetFileResponse
	err := clt.Post(req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
