package creative

import (
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/core"
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/model/creative"
)

// List 获取广告创意信息
func List(clt *core.SDKClient, accessToken string, req *creative.ListRequest) (*creative.ListResponse, error) {
	var resp creative.ListResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
