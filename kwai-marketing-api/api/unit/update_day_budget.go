package unit

import (
	"github.com/disv5/opensdk/kwai-marketing-api/core"
	"github.com/disv5/opensdk/kwai-marketing-api/model/unit"
)

// UpdateDayBudget 修改广告组预算
func UpdateDayBudget(clt *core.SDKClient, accessToken string, req *unit.UpdateDayBudgetRequest) error {
	return clt.Post(accessToken, req, nil)
}
