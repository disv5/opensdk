/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type MaterialLabelsAddRequest struct {
	AccountId   *int64               `json:"account_id,omitempty"`
	ImageIdList *[]string            `json:"image_id_list,omitempty"`
	MediaIdList *[]string            `json:"media_id_list,omitempty"`
	Labels      *[]CreateLabelStruct `json:"labels,omitempty"`
}
