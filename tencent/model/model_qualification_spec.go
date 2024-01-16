/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 资质信息
type QualificationSpec struct {
	IndustrySpec           *IndustryQualificationsSpec           `json:"industry_spec,omitempty"`
	AdSpec                 *AdQualificationsSpec                 `json:"ad_spec,omitempty"`
	AdditionalIndustrySpec *AdditionalIndustryQualificationsSpec `json:"additional_industry_spec,omitempty"`
}
