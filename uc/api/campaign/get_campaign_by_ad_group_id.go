package campaign

import (
	"donson.com.cn/draining/internal/pkg/uc/core"
	"donson.com.cn/draining/internal/pkg/uc/model/campaign"
)

func GetCampaignByAdGroupId(clt *core.SDKClient, req *campaign.GetCampaignByAdGroupIdRequest) (*campaign.GetCampaignByAdGroupIdResponse, error) {
	var resp campaign.GetCampaignByAdGroupIdResponse
	err := clt.Post(req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
