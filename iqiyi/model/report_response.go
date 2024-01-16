package model

import (
	"github.com/disv5/opensdk/iqiyi/model/report/order"
)

type ReportResponse struct {
	Code           int                `json:"code,omitempty"`    // 返回码
	Message        string             `json:"message,omitempty"` // 返回信息
	Data           []order.ReportData `json:"data,omitempty"`    // JSON返回值
	TraceId        string             `json:"traceId,omitempty"` // 请求id
	LastUpdateTime string             `json:"lastUpdateTime,omitempty"`
	TotalCount     int                `json:"totalCount,omitempty"`
}
