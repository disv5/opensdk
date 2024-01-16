package unit

import (
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/core"
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/model/unit"
)

// Create 创建广告组
func Create(clt *core.SDKClient, accessToken string, req *unit.CreateRequest) (uint64, error) {
	var resp unit.CreateResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return 0, err
	}
	return resp.UnitID, nil
}
