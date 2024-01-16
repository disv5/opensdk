package coupon

import (
	"donson.com.cn/draining/internal/pkg/oceanengine/core"
	"donson.com.cn/draining/internal/pkg/oceanengine/model/clue/coupon"
)

// Detail 获取卡券详情
func Detail(clt *core.SDKClient, accessToken string, req *coupon.DetailRequest) (*coupon.DetailResponseData, error) {
	var resp coupon.DetailResponse
	if err := clt.Get("2/clue/coupon/detail/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
