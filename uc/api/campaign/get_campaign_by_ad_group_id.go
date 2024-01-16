package campaign

import (
	"github.com/disv5/opensdk/uc/core"
	"github.com/disv5/opensdk/uc/model/campaign"
)

func GetCampaignByAdGroupId(clt *core.SDKClient, req *campaign.GetCampaignByAdGroupIdRequest) (*campaign.GetCampaignByAdGroupIdResponse, error) {
	var resp campaign.GetCampaignByAdGroupIdResponse
	err := clt.Post(req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
