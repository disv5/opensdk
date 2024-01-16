package tool

import (
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/core"
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/model/tool"
)

// AppSearch 获取可选的应用定向
func AppSearch(clt *core.SDKClient, accessToken string, req *tool.AppSearchRequest) (*tool.TargetingApp, error) {
	var resp tool.TargetingApp
	err := clt.Get(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
