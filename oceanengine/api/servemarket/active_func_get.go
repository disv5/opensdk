package servemarket

import (
	"github.com/disv5/opensdk/oceanengine/core"
	"github.com/disv5/opensdk/oceanengine/model/servemarket"
)

// ActiveFuncGet 获取用户已购功能点列表
func ActiveFuncGet(clt *core.SDKClient, accessToken string, req *servemarket.ActiveFuncGetRequest) ([]servemarket.OrderFunction, error) {
	var resp servemarket.ActiveFuncGetResponse
	if err := clt.OpenGet("v1.0/serve_market/active_func/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.FuncList, nil
}
