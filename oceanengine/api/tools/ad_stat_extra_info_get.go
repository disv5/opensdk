package tools

import (
	"donson.com.cn/draining/internal/pkg/oceanengine/core"
	"donson.com.cn/draining/internal/pkg/oceanengine/model/tools"
)

// AdStatExtraInfoGet 查询广告计划学习期状态
func AdStatExtraInfoGet(clt *core.SDKClient, accessToken string, req *tools.AdStatExtraInfoGetRequest) ([]tools.AdStatExtraInfo, error) {
	var resp tools.AdStatExtraInfoGetResponse
	err := clt.Get("2/tools/ad_stat_extra_info/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}