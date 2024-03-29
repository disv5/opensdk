/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// FieldTypeEnum : 排序条件字段，最新上传：derive_on_first_time，消耗：template_stat.cost，点击率：template_stat.click_rate，使用量：template_stat.use_cnt
type FieldTypeEnum string

// List of FieldTypeEnum
const (
	FieldTypeEnum_DERIVE_ON_FIRST_TIME    FieldTypeEnum = "derive_on_first_time"
	FieldTypeEnum_TEMPLATE_STATCOST       FieldTypeEnum = "template_stat.cost"
	FieldTypeEnum_TEMPLATE_STATCLICK_RATE FieldTypeEnum = "template_stat.click_rate"
	FieldTypeEnum_TEMPLATE_STATUSE_CNT    FieldTypeEnum = "template_stat.use_cnt"
)
