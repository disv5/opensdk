package dpa

import (
	"donson.com.cn/draining/internal/pkg/oceanengine/core"
	"donson.com.cn/draining/internal/pkg/oceanengine/model/dpa"
)

// ProductDelete 删除DPA商品
func ProductDelete(clt *core.SDKClient, accessToken string, req *dpa.ProductDeleteRequest) error {
	return clt.Post("2/dpa/product/delete/", req, nil, accessToken)
}
