package agent

import (
	"donson.com.cn/draining/internal/pkg/oceanengine/core"
	"donson.com.cn/draining/internal/pkg/oceanengine/model/agent"
)

// Info 获取代理商信息
func Info(clt *core.SDKClient, accessToken string, req *agent.InfoRequest) ([]agent.Info, error) {
	var resp agent.InfoResponse
	err := clt.Get("2/agent/info/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
