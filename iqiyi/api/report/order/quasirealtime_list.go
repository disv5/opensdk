package report

import (
	"donson.com.cn/draining/internal/pkg/iqiyi/core"
	"donson.com.cn/draining/internal/pkg/iqiyi/model"
	"donson.com.cn/draining/internal/pkg/iqiyi/model/report/order"
)

// QuasiRealTimeOrderList 效果报表
func QuasiRealTimeOrderList(clt *core.SDKClient, accessToken string, req *order.QuasiRealTimeListRequest) (*model.ReportResponse, error) {
	var resp model.ReportResponse
	err := clt.ReportPost(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
