package dmp

import (
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/core"
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/model/dmp"
)

// PopulationDelete 人群包删除接口
func PopulationDelete(clt *core.SDKClient, accessToken string, req *dmp.PopulationDeleteRequest) error {
	return clt.Post(accessToken, req, nil)
}
