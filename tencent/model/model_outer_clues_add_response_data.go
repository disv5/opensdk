/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type OuterCluesAddResponseData struct {
	FailOuterLeadIdList *[]OuterCluesAddListStruct      `json:"fail_outer_lead_id_list,omitempty"`
	SuccessLeadIdList   *[]SuccessLeadsIdListRespStruct `json:"success_lead_id_list,omitempty"`
}
