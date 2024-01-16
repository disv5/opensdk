package report

import (
	"github.com/disv5/opensdk/iqiyi/core"
	"github.com/disv5/opensdk/iqiyi/model"
	"github.com/disv5/opensdk/iqiyi/model/report/order"
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
