package tool

import (
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/core"
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/model/tool"
)

// OperationRecordList 账户操作记录信息查询
func OperationRecordList(clt *core.SDKClient, accessToken string, req *tool.OperationRecordListRequest) (*tool.OperationRecordListResponse, error) {
	var resp tool.OperationRecordListResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
