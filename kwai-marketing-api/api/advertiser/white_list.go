package advertiser

import (
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/core"
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/model/advertiser"
)

// WhiteList 获取可选白名单接口
func WhiteList(clt *core.SDKClient, accessToken string, advertiserID uint64) (*advertiser.WhiteListResponse, error) {
	req := &advertiser.WhiteListRequest{
		AdvertiserID: advertiserID,
	}
	var resp advertiser.WhiteListResponse
	err := clt.Get(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
