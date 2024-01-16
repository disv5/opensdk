/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 返回结构
type FundsGetListStruct struct {
	Balance      *int64         `json:"balance,omitempty"`
	FundStatus   FundStatus     `json:"fund_status,omitempty"`
	RealtimeCost *int64         `json:"realtime_cost,omitempty"`
	FundType     AccountTypeMap `json:"fund_type,omitempty"`
}