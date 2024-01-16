package report

import (
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/core"
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/model/report"
)

// CampaignReport 广告计划数据（实时）
func CampaignReport(clt *core.SDKClient, accessToken string, req *report.CampaignReportRequest) (*report.ReportResponse, error) {
	var resp report.ReportResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
