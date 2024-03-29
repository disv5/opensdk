package report

import (
	"github.com/disv5/opensdk/kuaishou/core"
	"github.com/disv5/opensdk/kuaishou/model/report"
)

// UnitReport 广告组数据
func UnitReport(clt *core.SDKClient, accessToken string, req *report.UnitReportRequest) (*report.ReportResponse, error) {
	var resp report.ReportResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
