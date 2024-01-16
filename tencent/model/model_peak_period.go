/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 门店高峰时间段信息结构
type PeakPeriod struct {
	TimeSeries *string                  `json:"time_series,omitempty"`
	Date       LocalStorePeakPeriodDate `json:"date,omitempty"`
}
