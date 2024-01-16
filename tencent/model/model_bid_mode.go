/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// BidMode : 出价方式，<br/> 1. bid_mode为billing_event升级字段，不可同时输入，升级后直接写入bid_mode字段即可；<br/> 2. 当投放智能出价广告，可写入BID_MODE_OCPC/BID_MODE_OCPM。此时，optimization_goal优化目标字段必填；<br/> 3. 当投放非智能出价广告，可写入BID_MODE_CPC/BID_MODE_CPM/BID_MODE_CPA。此时，optimization_goal优化目标字段不可填；<br/> 4. 针对非微信流量，BID_MODE_CPC可编辑修改为BID_MODE_OCPC，BID_MODE_CPM可编辑修改为BID_MODE_OCPM，其他修改不可操作。针对微信流量，bid_mode字段不可修改；<br/> 5. 可通过adcreative_templates/get接口查询不同情况下支持的出价方式;
type BidMode string

// List of BidMode
const (
	BidMode_CPC  BidMode = "BID_MODE_CPC"
	BidMode_CPA  BidMode = "BID_MODE_CPA"
	BidMode_CPM  BidMode = "BID_MODE_CPM"
	BidMode_CPT  BidMode = "BID_MODE_CPT"
	BidMode_CPO  BidMode = "BID_MODE_CPO"
	BidMode_OCPC BidMode = "BID_MODE_OCPC"
	BidMode_OCPA BidMode = "BID_MODE_OCPA"
	BidMode_OCPM BidMode = "BID_MODE_OCPM"
)