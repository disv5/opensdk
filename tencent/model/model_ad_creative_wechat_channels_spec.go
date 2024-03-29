/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 微信视频号信息
type AdCreativeWechatChannelsSpec struct {
	ExportId         *string          `json:"export_id,omitempty"`
	Username         *string          `json:"username,omitempty"`
	LivePromotedType LivePromotedType `json:"live_promoted_type,omitempty"`
}
