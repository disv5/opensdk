package creative

import (
	"donson.com.cn/draining/internal/pkg/oceanengine/core"
	"donson.com.cn/draining/internal/pkg/oceanengine/model/qianchuan/creative"
)

// RejectReason 获取创意审核建议
// 通过此接口用于获取广告创意纬度的审核意见
func RejectReason(clt *core.SDKClient, accessToken string, req *creative.RejectReasonRequest) ([]creative.RejectReasonList, error) {
	var resp creative.RejectReasonResponse
	err := clt.Get("v1.0/qianchuan/creative/reject_reason/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
