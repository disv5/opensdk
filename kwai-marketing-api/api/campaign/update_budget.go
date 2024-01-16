package campaign

import (
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/core"
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/model/campaign"
)

// UpdateBudget 修改广告计划预算
func UpdateBudget(clt *core.SDKClient, accessToken string, req *campaign.UpdateBudgetRequest) error {
	return clt.Post(accessToken, req, nil)
}
