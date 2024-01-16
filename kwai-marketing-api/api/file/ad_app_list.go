package file

import (
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/core"
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/model/file"
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
