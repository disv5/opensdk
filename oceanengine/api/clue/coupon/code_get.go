package coupon

import (
	"donson.com.cn/draining/internal/pkg/oceanengine/core"
	"donson.com.cn/draining/internal/pkg/oceanengine/model/clue/coupon"
)

// CodeGet 查询券码记录
func CodeGet(clt *core.SDKClient, accessToken string, req *coupon.CodeGetRequest) (*coupon.CodeGetResponseData, error) {
	var resp coupon.CodeGetResponse
	if err := clt.Get("2/clue/coupon/code/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
