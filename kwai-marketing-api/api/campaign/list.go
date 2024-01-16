package campaign

import (
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/core"
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/model/campaign"
)

// List 获取广告计划信息
func List(clt *core.SDKClient, accessToken string, req *campaign.ListRequest) (*campaign.ListResponse, error) {
	var resp campaign.ListResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
