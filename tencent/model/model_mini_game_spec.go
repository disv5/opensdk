/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 小游戏落地页信息
type MiniGameSpec struct {
	MiniGameTrackingParameter *string                             `json:"mini_game_tracking_parameter,omitempty"`
	MiniGameOpenlink          *string                             `json:"mini_game_openlink,omitempty"`
	MiniGameOpenlinkPageSpec  *AdCreativeMiniGameOpenlinkPageSpec `json:"mini_game_openlink_page_spec,omitempty"`
	MiniGameOpenlinkSwitch    *bool                               `json:"mini_game_openlink_switch,omitempty"`
}