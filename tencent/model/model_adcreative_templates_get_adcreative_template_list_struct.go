/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 创意形式和投放权限数据结构
type AdcreativeTemplatesGetAdcreativeTemplateListStruct struct {
	AdcreativeTemplateId          *int64                         `json:"adcreative_template_id,omitempty"`
	AdcreativeTemplateName        *string                        `json:"adcreative_template_name,omitempty"`
	AdcreativeTemplateDescription *string                        `json:"adcreative_template_description,omitempty"`
	AdcreativeTemplateSize        *string                        `json:"adcreative_template_size,omitempty"`
	AdcreativeTemplateStyle       *string                        `json:"adcreative_template_style,omitempty"`
	AdcreativeTemplateAppellation *string                        `json:"adcreative_template_appellation,omitempty"`
	SiteSet                       SiteSetDefinition              `json:"site_set,omitempty"`
	PromotedObjectType            PromotedObjectTypeWithoutJd    `json:"promoted_object_type,omitempty"`
	AdcreativeSampleImageList     *[]AdcreativeSampleImage       `json:"adcreative_sample_image_list,omitempty"`
	AdAttributes                  *[]AdcreativeElement           `json:"ad_attributes,omitempty"`
	AdcreativeAttributes          *[]AdcreativeElement           `json:"adcreative_attributes,omitempty"`
	AdcreativeElements            *[]AdcreativeElement           `json:"adcreative_elements,omitempty"`
	SupportPageType               *[]string                      `json:"support_page_type,omitempty"`
	LandingPageConfig             *LandingPageConfig             `json:"landing_page_config,omitempty"`
	SupportBillingSpecList        *[]SupportBillingSpec          `json:"support_billing_spec_list,omitempty"`
	SupportDynamicAbilitySpecList *SupportDynamicAbilitySpecList `json:"support_dynamic_ability_spec_list,omitempty"`
	SupportBidModeList            *[]string                      `json:"support_bid_mode_list,omitempty"`
}
