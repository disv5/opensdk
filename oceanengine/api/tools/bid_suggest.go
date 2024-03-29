package tools

import (
	"github.com/disv5/opensdk/oceanengine/core"
	"github.com/disv5/opensdk/oceanengine/model/tools"
)

// BidSuggest 建议日预算及预期成本
// 通过广告的受众查询广告建议出价，获取建议日预算及预期成本，部分受众暂不支持。
func BidSuggest(clt *core.SDKClient, accessToken string, req *tools.BidSuggestRequest) (*tools.BidSuggest, error) {
	var resp tools.BidSuggestResponse
	if err := clt.Get("2/tools/bid/suggest/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
