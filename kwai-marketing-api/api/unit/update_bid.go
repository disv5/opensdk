package unit

import (
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/core"
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/model/unit"
)

// UpdateBid 修改广告组出价
func UpdateBid(clt *core.SDKClient, accessToken string, req *unit.UpdateBidRequest) error {
	return clt.Post(accessToken, req, nil)
}
