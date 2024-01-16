package dpa

import (
	"donson.com.cn/draining/internal/pkg/oceanengine/core"
	"donson.com.cn/draining/internal/pkg/oceanengine/model/dpa"
)

// DetailGet 获取商品列表
func DetailGet(clt *core.SDKClient, accessToken string, req *dpa.DetailGetRequest) (*dpa.DetailGetResponseData, error) {
	var resp dpa.DetailGetResponse
	if err := clt.Get("2/dpa/detail/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}