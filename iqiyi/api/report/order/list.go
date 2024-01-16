package report

import (
	"donson.com.cn/draining/internal/pkg/iqiyi/core"
	"donson.com.cn/draining/internal/pkg/iqiyi/model"
	"donson.com.cn/draining/internal/pkg/iqiyi/model/report/order"
)

// OrderList 效果报表
func OrderList(clt *core.SDKClient, accessToken string, req *order.ListRequest) (*model.ReportResponse, error) {
	var resp model.ReportResponse
	err := clt.ReportPost(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
