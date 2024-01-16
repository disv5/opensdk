package dmp

import (
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/core"
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/model/dmp"
)

// PopulationAccountsPush 人群包跨账户推送
func PopulationAccountsPush(clt *core.SDKClient, accessToken string, req *dmp.PopulationAccountsPushRequest) (*dmp.PopulationAccountsPushResponse, error) {
	var resp dmp.PopulationAccountsPushResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
