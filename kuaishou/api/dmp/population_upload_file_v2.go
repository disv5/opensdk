package dmp

import (
	"github.com/disv5/opensdk/kuaishou/core"
	"github.com/disv5/opensdk/kuaishou/model/dmp"
)

// PopulationUpload 人群包上传接口
func PopulationUploadFileV2(clt *core.SDKClient, accessToken string, req *dmp.PopulationUploadFileRequest) (*dmp.FileV2, error) {
	var resp dmp.FileV2
	err := clt.Upload(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
