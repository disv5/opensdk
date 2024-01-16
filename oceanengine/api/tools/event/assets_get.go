package event

import (
	"donson.com.cn/draining/internal/pkg/oceanengine/core"
	"donson.com.cn/draining/internal/pkg/oceanengine/model/tools/event"
)

// AssetsGet 获取推广内容
// 查询广告主拥有的资产信息
func AssetsGet(clt *core.SDKClient, accessToken string, req *event.AssetsGetRequest) (*event.AssetsGetResponseData, error) {
	var resp event.AssetsGetResponse
	err := clt.Get("2/tools/event/assets/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
