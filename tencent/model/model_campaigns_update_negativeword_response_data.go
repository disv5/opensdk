/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type CampaignsUpdateNegativewordResponseData struct {
	CampaignId        *int64                          `json:"campaign_id,omitempty"`
	Status            NegativeWordOperStatus          `json:"status,omitempty"`
	DuplicateWords    *DuplicateNegativeWordStruct    `json:"duplicate_words,omitempty"`
	ExceedLengthWords *ExceedLengthNegativeWordStruct `json:"exceed_length_words,omitempty"`
	ExceedLimitWords  *ExceedLimitNegativeWordStruct  `json:"exceed_limit_words,omitempty"`
	HasSpecialWords   *HasSpecialNegativeWordStruct   `json:"has_special_words,omitempty"`
	SuccessWords      *SuccessNegativeWordStruct      `json:"success_words,omitempty"`
}
