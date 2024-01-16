package order

import (
	"encoding/json"
	"time"
)

type ListRequest struct {
	AdvertiserIDs []uint64 `json:"advertiserIds,omitempty"`
	OrderGroupIDs []uint64 `json:"orderGroupIds,omitempty"`
	OrderPlanIDs  []uint64 `json:"orderPlanIds,omitempty"`
	CreativeIds   []uint64 `json:"creativeIds,omitempty"`
	// 必填，格式yyyy-MM-dd，起始时间不得早于2017-01-01
	StartDate time.Time `json:"startDate,omitempty"`
	// 必填，格式yyyy-MM-dd，时间跨度不得大于30天
	EndDate    time.Time `json:"endDate,omitempty"`
	Dimensions []string  `json:"dimensions,omitempty"`

	// Page 页码
	Page int `json:"page,omitempty"`
	// PageSize 每页行数
	PageSize int `json:"pageSize,omitempty"`
}

// Url implement GetRequest interface
func (r ListRequest) Url() string {
	return "openapi/v1/report/order/list"
}

// Encode implement GetRequest interface
func (r ListRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
