package advertiser

import (
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/core"
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/model/advertiser"
)

// UpdateBudget 修改账户预算
func UpdateBudget(clt *core.SDKClient, accessToken string, req *advertiser.UpdateBudgetRequest) error {
	return clt.Post(accessToken, req, nil)
}
