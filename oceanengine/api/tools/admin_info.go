package tools

import (
	"donson.com.cn/draining/internal/pkg/oceanengine/core"
	"donson.com.cn/draining/internal/pkg/oceanengine/model/tools"
)

// AdminInfo 获取行政信息
func AdminInfo(clt *core.SDKClient, accessToken string, req *tools.AdminInfoRequest) (*tools.AdminInfoResponseData, error) {
	var resp tools.AdminInfoResponse
	err := clt.Get("2/tools/admin/info/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
