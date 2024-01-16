package ad

import (
	"donson.com.cn/draining/internal/pkg/oceanengine/core"
	"donson.com.cn/draining/internal/pkg/oceanengine/model/ad"
)

// RejectReason 获取计划审核建议
func RejectReason(clt *core.SDKClient, accessToken string, req *ad.RejectReasonRequest) ([]ad.RejectReason, error) {
	var resp ad.RejectReasonResponse
	if err := clt.Get("2/ad/reject_reason/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
