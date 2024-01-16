package campaign

import (
	"github.com/disv5/opensdk/kuaishou/core"
	"github.com/disv5/opensdk/kuaishou/model/campaign"
)

// Create 创建广告计划
func Create(clt *core.SDKClient, accessToken string, req *campaign.CreateRequest) (uint64, error) {
	var resp campaign.CreateResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return 0, err
	}
	return resp.CampaignID, nil
}