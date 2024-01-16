package tools

import (
	"github.com/disv5/opensdk/oceanengine/core"
	"github.com/disv5/opensdk/oceanengine/model/tools"
)

// EstimatedPriceGet 获取预估点击成本
func EstimatedPriceGet(clt *core.SDKClient, accessToken string, req *tools.EstimatedPriceGetRequest) (*tools.EstimatedPrice, error) {
	var resp tools.EstimatedPriceGetResponse
	err := clt.Get("2/tools/estimated_price/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
