/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 返回结构
type PlayablePagesGetListStruct struct {
	LandingPageId          *string                 `json:"landing_page_id,omitempty"`
	PlayablePageId         *string                 `json:"playable_page_id,omitempty"`
	PlayablePageMaterialId *string                 `json:"playable_page_material_id,omitempty"`
	PlayablePageName       *string                 `json:"playable_page_name,omitempty"`
	PlayablePageCdnBaseUrl *string                 `json:"playable_page_cdn_base_url,omitempty"`
	PlayablePageDirection  PlayablePageDirection   `json:"playable_page_direction,omitempty"`
	AuditStatus            PlayablePageAuditStatus `json:"audit_status,omitempty"`
	AuditMsg               *string                 `json:"audit_msg,omitempty"`
	CreatedTime            *int64                  `json:"created_time,omitempty"`
	LastModifiedTime       *int64                  `json:"last_modified_time,omitempty"`
}
