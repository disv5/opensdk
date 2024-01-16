package coupon

import (
	"donson.com.cn/draining/internal/pkg/oceanengine/core"
	"donson.com.cn/draining/internal/pkg/oceanengine/model/clue/coupon"
)

// CodeConsume 核销券码
func CodeConsume(clt *core.SDKClient, accessToken string, req *coupon.CodeConsumeRequest) error {
	return clt.Post("2/clue/coupon/code/consume/", req, nil, accessToken)
}
