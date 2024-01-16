/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 自定义地理位置规则
type CustomLocationRule struct {
	PoiType       LbsPoiType               `json:"poi_type,omitempty"`
	DateRange     *CustomLocationDateRange `json:"date_range,omitempty"`
	FrequencySpec *FrequencySpec           `json:"frequency_spec,omitempty"`
	AreaList      *[]AreaSpec              `json:"area_list,omitempty"`
}