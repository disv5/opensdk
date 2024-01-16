package report

import (
	"donson.com.cn/draining/internal/pkg/uc/core"
	"donson.com.cn/draining/internal/pkg/uc/model/report"
)

func Campaign(clt *core.SDKClient, req *report.CampaignRequest) (*report.CampaignResponse, error) {
	var resp report.CampaignResponse
	err := clt.Post(req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

