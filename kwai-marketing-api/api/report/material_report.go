package report

import (
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/core"
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/model/report"
)

// MaterialReport 素材报表
func MaterialReport(clt *core.SDKClient, accessToken string, req *report.MaterialReportRequest) (*report.ReportResponse, error) {
	var resp report.MaterialReportResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.ReportResponse, nil
}
