package report

import (
	"github.com/disv5/opensdk/uc/core"
	"github.com/disv5/opensdk/uc/model/report"
)

func Campaign(clt *core.SDKClient, req *report.CampaignRequest) (*report.CampaignResponse, error) {
	var resp report.CampaignResponse
	err := clt.Post(req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
