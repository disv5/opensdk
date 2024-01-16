/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type AdParamGetRequest struct {
	AccountId            *int64             `json:"account_id,omitempty"`
	CampaignType         CampaignType       `json:"campaign_type,omitempty"`
	PromotedObjectType   PromotedObjectType `json:"promoted_object_type,omitempty"`
	AdcreativeTemplateId *int64             `json:"adcreative_template_id,omitempty"`
	SiteSet              SiteSetDefinition  `json:"site_set,omitempty"`
	ProductCatalogId     *int64             `json:"product_catalog_id,omitempty"`
}