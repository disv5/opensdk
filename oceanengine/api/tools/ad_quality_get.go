package tools

import (
	"github.com/disv5/opensdk/oceanengine/core"
	"github.com/disv5/opensdk/oceanengine/model/tools"
)

// AdQualityGet 查询广告质量度
// 查询计划的广告质量度，只有产生过投放消耗的计划才会有质量度数据。
func AdQualityGet(clt *core.SDKClient, accessToken string, req *tools.AdQualityGetRequest) ([]tools.AdQuality, error) {
	var resp tools.AdQualityGetResponse
	if err := clt.Get("2/tools/ad_quality/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
