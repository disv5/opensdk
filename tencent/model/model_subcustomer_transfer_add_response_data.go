/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type SubcustomerTransferAddResponseData struct {
	FundType       AccountTypeMap `json:"fund_type,omitempty"`
	Amount         *int64         `json:"amount,omitempty"`
	ExternalBillNo *string        `json:"external_bill_no,omitempty"`
	Time           *int64         `json:"time,omitempty"`
}
