package adraise

import (
	"github.com/disv5/opensdk/oceanengine/core"
	"github.com/disv5/opensdk/oceanengine/enum"
	"github.com/disv5/opensdk/oceanengine/model/tools/adraise"
)

// Set 设置一键起量: 用来启动或关停一键起量服务
func Set(clt *core.SDKClient, accessToken string, req *adraise.SetRequest) (enum.AdRaiseStatus, error) {
	var resp adraise.SetResponse
	err := clt.Post("2/tools/ad_raise/set/", req, &resp, accessToken)
	if err != nil {
		return enum.AdRaiseStatus_UNKNOWN, err
	}
	return resp.Data.Status, nil
}
