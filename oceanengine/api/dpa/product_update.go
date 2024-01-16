package dpa

import (
	"github.com/disv5/opensdk/oceanengine/core"
	"github.com/disv5/opensdk/oceanengine/model/dpa"
)

// ProductUpdate 创建DPA商品（已有商品id）/修改DPA商品
func ProductUpdate(clt *core.SDKClient, accessToken string, req *dpa.ProductUpdateRequest) (uint64, error) {
	var resp dpa.ProductUpdateResponse
	if err := clt.Post("2/dpa/product/update/", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.ProductID, nil
}
