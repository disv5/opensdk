package advertiser

import (
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/core"
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/model/advertiser"
)

// FundGet 获取广告账户余额信息
func FundGet(clt *core.SDKClient, accessToken string, advertiserID uint64) (float64, error) {
	req := &advertiser.FundGetRequest{
		AdvertiserID: advertiserID,
	}
	var resp advertiser.FundGetResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return 0, err
	}
	return resp.Balance, nil
}
