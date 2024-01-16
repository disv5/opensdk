package shop

import (
	"donson.com.cn/draining/internal/pkg/oceanengine/core"
	"donson.com.cn/draining/internal/pkg/oceanengine/model/qianchuan/shop"
)

// Get 获取店铺账户信息
func Get(clt *core.SDKClient, accessToken string, req *shop.GetRequest) ([]shop.Shop, error) {
	var resp shop.GetResponse
	err := clt.Get("v1.0/qianchuan/shop/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
