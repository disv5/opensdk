package file

import (
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/core"
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/model/file"
)

// AdAppUpdate 修改应用
func AdAppUpdate(clt *core.SDKClient, accessToken string, req *file.AdAppUpdateRequest) (*file.App, error) {
	var resp file.App
	err := clt.Upload(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
