package asset

import (
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/core"
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/model/asset"
)

// AdvCardCreate 创建高级创意接口
func AdvCardCreate(clt *core.SDKClient, accessToken string, req *asset.AdvCardCreateRequest) ([]uint64, error) {
	var resp asset.AdvCardCreateResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp.AdvList, nil
}
