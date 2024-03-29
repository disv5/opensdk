/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 能力项详情
type PropertyDetailCopy struct {
	StringDetail     *TextRestriction                `json:"string_detail,omitempty"`
	IntegerDetail    *NumberRestriction              `json:"integer_detail,omitempty"`
	EnumDetail       *AdcreativeElementEnumProperty  `json:"enum_detail,omitempty"`
	ArrayRestriction *AdcreativeElementArrayProperty `json:"array_restriction,omitempty"`
}
