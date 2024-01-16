package advertiser

import (
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/core"
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/model/advertiser"
)

// AdvertisersGet   获取罗盘绑定广告主列
func AdvertisersGet(clt *core.SDKClient, accessToken string, advertiserID uint64) (*advertiser.AdvertisersResponse, error) {
	req := &advertiser.AdvertisersRequest{
		AdvertiserID: advertiserID,
	}
	var resp advertiser.AdvertisersResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
