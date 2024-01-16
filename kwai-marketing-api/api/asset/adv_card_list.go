package asset

import (
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/core"
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/model/asset"
)

// AdvCardList 获取高级创意列表
func AdvCardList(clt *core.SDKClient, accessToken string, req *asset.AdvCardListRequest) (*asset.AdvCardListResponse, error) {
	var resp asset.AdvCardListResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
