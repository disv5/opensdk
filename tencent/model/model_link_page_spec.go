/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 文字链跳转信息
type LinkPageSpec struct {
	PageId          *int64                     `json:"page_id,omitempty"`
	PageUrl         *string                    `json:"page_url,omitempty"`
	MiniProgramSpec *AdcreativeMiniProgramSpec `json:"mini_program_spec,omitempty"`
	MiniGameSpec    *LinkMiniGameSpec          `json:"mini_game_spec,omitempty"`
}