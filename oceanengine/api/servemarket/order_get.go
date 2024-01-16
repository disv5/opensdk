package servemarket

import (
	"github.com/disv5/opensdk/oceanengine/core"
	"github.com/disv5/opensdk/oceanengine/model/servemarket"
)

// OrderGet 获取应用订单数据
func OrderGet(clt *core.SDKClient, accessToken string, req *servemarket.OrderGetRequest) (*servemarket.OrderGetResponseData, error) {
	var resp servemarket.OrderGetResponse
	if err := clt.OpenGet("v1.0/serve_market/order/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
