package creative

import (
	"donson.com.cn/draining/internal/pkg/oceanengine/core"
	"donson.com.cn/draining/internal/pkg/oceanengine/model/creative"
)

// MaterialRead 创意素材信息
func MaterialRead(clt *core.SDKClient, accessToken string, req *creative.MaterialReadRequest) ([]creative.Material, error) {
	var resp creative.MaterialReadResponse
	err := clt.Get("2/creative/material/read/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
