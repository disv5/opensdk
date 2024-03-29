package file

import (
	"github.com/disv5/opensdk/kuaishou/core"
	"github.com/disv5/opensdk/kuaishou/model/file"
)

// AdAppCreate 创建应用
func AdAppCreate(clt *core.SDKClient, accessToken string, req *file.AdAppCreateRequest) (*file.App, error) {
	var resp file.App
	err := clt.Upload(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
