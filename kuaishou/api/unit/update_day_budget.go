package unit

import (
	"github.com/disv5/opensdk/kuaishou/core"
	"github.com/disv5/opensdk/kuaishou/model/unit"
)

// UpdateDayBudget 修改广告组预算
func UpdateDayBudget(clt *core.SDKClient, accessToken string, req *unit.UpdateDayBudgetRequest) error {
	return clt.Post(accessToken, req, nil)
}
