/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type EstimationGetRequest struct {
	CampaignSpec *CampaignTargeting              `json:"campaign_spec,omitempty"`
	SceneSpec    *EstimationSceneTargeting       `json:"scene_spec,omitempty"`
	AccountId    *int64                          `json:"account_id,omitempty"`
	Adcreative   *[]CreativeStruct               `json:"adcreative,omitempty"`
	Targeting    *EstimationReadTargetingSetting `json:"targeting,omitempty"`
	Adgroup      *AdgroupSetting                 `json:"adgroup,omitempty"`
}
