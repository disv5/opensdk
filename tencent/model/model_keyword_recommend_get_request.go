/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type KeywordRecommendGetRequest struct {
	SiteSets          *[]string            `json:"site_sets,omitempty"`
	RecommendCategory RecommendCategory    `json:"recommend_category,omitempty"`
	AccountId         *int64               `json:"account_id,omitempty"`
	SystemIndustryId  *int64               `json:"system_industry_id,omitempty"`
	QueryWord         *[]string            `json:"query_word,omitempty"`
	BusinessPointId   *string              `json:"business_point_id,omitempty"`
	AdgroupId         *int64               `json:"adgroup_id,omitempty"`
	CampaignId        *int64               `json:"campaign_id,omitempty"`
	IncludeWord       *[]string            `json:"include_word,omitempty"`
	ExcludeWord       *[]string            `json:"exclude_word,omitempty"`
	FilterAdWord      *bool                `json:"filter_ad_word,omitempty"`
	FilterAccountWord *bool                `json:"filter_account_word,omitempty"`
	RecommendReasons  *[]string            `json:"recommend_reasons,omitempty"`
	Province          *[]int64             `json:"province,omitempty"`
	City              *[]int64             `json:"city,omitempty"`
	OrderBy           *[]OrderByStructInfo `json:"order_by,omitempty"`
}
