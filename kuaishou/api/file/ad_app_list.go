package file

import (
	"github.com/disv5/opensdk/kuaishou/core"
	"github.com/disv5/opensdk/kuaishou/model/file"
)

// AdAppList 获取应用列表
func AdAppList(clt *core.SDKClient, accessToken string, req *file.AdAppListRequest) (*file.AdAppListResponse, error) {
	var resp file.AdAppListResponse
	err := clt.GetOnBody(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
