package unit

import (
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/core"
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/model/unit"
)

// Update 修改广告组
func Update(clt *core.SDKClient, accessToken string, req *unit.UpdateRequest) (uint64, error) {
	var resp unit.UpdateResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return 0, err
	}
	return resp.UnitID, nil
}
