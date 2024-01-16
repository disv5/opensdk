/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 支持的落地页类型
type LandingPageConfigPermit struct {
	Required            *bool                          `json:"required,omitempty"`
	SupportPageTypeList *[]SupportPageTypeStructPermit `json:"support_page_type_list,omitempty"`
}
