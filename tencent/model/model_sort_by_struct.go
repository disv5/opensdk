/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 排序
type SortByStruct struct {
	Field FieldTypeEnum `json:"field,omitempty"`
	Desc  Sort          `json:"desc,omitempty"`
}
