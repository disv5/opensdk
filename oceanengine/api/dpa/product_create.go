package dpa

import (
	"donson.com.cn/draining/internal/pkg/oceanengine/core"
	"donson.com.cn/draining/internal/pkg/oceanengine/model/dpa"
)

// ProductCreate 创建DPA商品（无商品id）
func ProductCreate(clt *core.SDKClient, accessToken string, req *dpa.ProductCreateRequest) (uint64, error) {
	var resp dpa.ProductCreateResponse
	if err := clt.Post("2/dpa/product/create/", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.ProductID, nil
}
