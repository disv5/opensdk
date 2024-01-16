package campaign

import (
	"github.com/disv5/opensdk/kwai-marketing-api/core"
	"github.com/disv5/opensdk/kwai-marketing-api/model/campaign"
)

// UpdateBudget 修改广告计划预算
func UpdateBudget(clt *core.SDKClient, accessToken string, req *campaign.UpdateBudgetRequest) error {
	return clt.Post(accessToken, req, nil)
}
