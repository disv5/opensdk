/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 智能定向功能,功能灰度开放，如需使用可联系您的运营接口同学。<br>智能定向功能与自动扩量/系统优选相关字段不可同时设置。2022年6月30日起，智能定向无法与行为兴趣意向、“二方人群”人群包同时使用
type SmartTargeting struct {
	SmartTargetingVersion *int64                       `json:"smart_targeting_version,omitempty"`
	SmartTargetingSwitch  *bool                        `json:"smart_targeting_switch,omitempty"`
	StartAudience         *[]int64                     `json:"start_audience,omitempty"`
	UnbreakableTargeting  *UnbreakableTargetingSetting `json:"unbreakable_targeting,omitempty"`
}