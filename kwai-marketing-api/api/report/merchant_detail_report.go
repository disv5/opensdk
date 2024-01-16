package report

import (
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/core"
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/model/report"
)

// MerchantDetailReport 小店通转化数据报表
func MerchantDetailReport(clt *core.SDKClient, accessToken string, req *report.MerchantDetailReportRequest) (*report.MerchantDetailReportResponse, error) {
	var resp report.MerchantDetailReportResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
