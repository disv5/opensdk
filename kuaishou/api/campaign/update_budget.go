package campaign

import (
	"github.com/disv5/opensdk/kuaishou/core"
	"github.com/disv5/opensdk/kuaishou/model/campaign"
)

// UpdateBudget 修改广告计划预算
func UpdateBudget(clt *core.SDKClient, accessToken string, req *campaign.UpdateBudgetRequest) error {
	return clt.Post(accessToken, req, nil)
}
