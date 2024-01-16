package region

import (
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/core"
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/model/region"
)

// DistrictList 获取商圈地域定向
func DistrictList(clt *core.SDKClient, accessToken string, advertiserID uint64) (map[string]region.District, error) {
	req := &region.DistrictListRequest{
		AdvertiserID: advertiserID,
	}
	resp := make(map[string]region.District)
	err := clt.Get(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
