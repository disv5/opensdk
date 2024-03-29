/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 头像点击跳转信息
type HeadClickSpec struct {
	BrandAppId             *string `json:"brand_app_id,omitempty"`
	SearchBrandAreaKeyword *string `json:"search_brand_area_keyword,omitempty"`
	FinderUsername         *string `json:"finder_username,omitempty"`
}
