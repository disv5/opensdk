package dpa

import (
	"donson.com.cn/draining/internal/pkg/oceanengine/core"
	"donson.com.cn/draining/internal/pkg/oceanengine/model/dpa"
)

// ProductDetailGet 获取商品详情
func ProductDetailGet(clt *core.SDKClient, accessToken string, req *dpa.ProductDetailGetRequest) (*dpa.DetailGetResponseData, error) {
	var resp dpa.DetailGetResponse
	if err := clt.Get("2/dpa/product/detail/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
