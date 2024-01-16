package budget

import (
	"net/url"
)

// GetAccountBudgetReq 获取账户日预算
type GetAccountBudgetReq struct {
	AdvertiserIds []uint64 `json:"advertiser_ids" binding:"required"` // 广告账户IDs
}

// Encode implement GetRequest interface
func (r GetAccountBudgetReq) Encode() string {
	values := &url.Values{}
	return values.Encode()
}
