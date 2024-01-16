package asset

import (
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/core"
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/model/asset"
)

// AdvCardRemove 删除高级创意接口
func AdvCardRemove(clt *core.SDKClient, accessToken string, req *asset.AdvCardRemoveRequest) ([]uint64, error) {
	var resp asset.AdvCardRemoveResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp.AdvCardID, nil
}