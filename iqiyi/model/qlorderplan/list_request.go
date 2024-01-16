package qlorderplan

import (
	"encoding/json"
)

type ListRequest struct {
	AdvertiserID uint64   `json:"advertiserId,omitempty"`
	OrderGroupId uint64   `json:"orderGroupId,omitempty"`
	Ids          []uint64 `json:"ids,omitempty"`
	Name         string   `json:"name,omitempty"`
	Status       []string `json:"status,omitempty"`
	// Page 页码
	Page int `json:"page,omitempty"`
	// PageSize 每页行数
	PageSize int `json:"pageSize,omitempty"`
}

// Url implement GetRequest interface
func (r ListRequest) Url() string {
	return "openapi/v2/qlOrderPlan/list"
}

// Encode implement GetRequest interface
func (r ListRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
