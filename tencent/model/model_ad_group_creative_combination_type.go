/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// AdGroupCreativeCombinationType : 广告类型，支持普通广告、集装箱广告和动态创意广告
type AdGroupCreativeCombinationType string

// List of AdGroupCreativeCombinationType
const (
	AdGroupCreativeCombinationType_NORMAL   AdGroupCreativeCombinationType = "COMBINATION_TYPE_NORMAL"
	AdGroupCreativeCombinationType_CAROUSEL AdGroupCreativeCombinationType = "COMBINATION_TYPE_CAROUSEL"
	AdGroupCreativeCombinationType_DYNAMIC  AdGroupCreativeCombinationType = "COMBINATION_TYPE_DYNAMIC"
)