package report

import (
	"github.com/disv5/opensdk/uc/core"
	"github.com/disv5/opensdk/uc/model/report"
)

// DownloadFile 解析csv文件，返回多维数组
func DownloadFile(clt *core.SDKClient, req *report.DownloadFileRequest) (*[][]string, error) {
	var resp [][]string
	err := clt.Post2(req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
