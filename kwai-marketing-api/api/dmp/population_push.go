package dmp

import (
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/core"
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/model/dmp"
)

// PopulationPush 人群包上线接口
func PopulationPush(clt *core.SDKClient, accessToken string, req *dmp.PopulationPushRequest) error {
	return clt.Post(accessToken, req, nil)
}
