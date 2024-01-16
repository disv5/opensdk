/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// ConversionOptimizationGoal : 深度优化ROI目标，深度优化行为目标、深度优化ROI和深度强化ROI目标仅可填写其中一个
type ConversionOptimizationGoal string

// List of ConversionOptimizationGoal
const (
	ConversionOptimizationGoal_NONE                      ConversionOptimizationGoal = "GOAL_NONE"
	ConversionOptimizationGoal_7DAYPURCHASEROAS          ConversionOptimizationGoal = "GOAL_7DAY_PURCHASE_ROAS"
	ConversionOptimizationGoal_15DAYPURCHASEROAS         ConversionOptimizationGoal = "GOAL_15DAY_PURCHASE_ROAS"
	ConversionOptimizationGoal_30DAYPURCHASEROAS         ConversionOptimizationGoal = "GOAL_30DAY_PURCHASE_ROAS"
	ConversionOptimizationGoal_60DAYPURCHASEROAS         ConversionOptimizationGoal = "GOAL_60DAY_PURCHASE_ROAS"
	ConversionOptimizationGoal_30DAYMONETIZATIONROAS     ConversionOptimizationGoal = "GOAL_30DAY_MONETIZATION_ROAS"
	ConversionOptimizationGoal_30DAYORDERROAS            ConversionOptimizationGoal = "GOAL_30DAY_ORDER_ROAS"
	ConversionOptimizationGoal_1DAYPURCHASEROAS          ConversionOptimizationGoal = "GOAL_1DAY_PURCHASE_ROAS"
	ConversionOptimizationGoal_1DAYMONETIZATIONROAS      ConversionOptimizationGoal = "GOAL_1DAY_MONETIZATION_ROAS"
	ConversionOptimizationGoal_3DAYPURCHASEROAS          ConversionOptimizationGoal = "GOAL_3DAY_PURCHASE_ROAS"
	ConversionOptimizationGoal_3DAYMONETIZATIONROAS      ConversionOptimizationGoal = "GOAL_3DAY_MONETIZATION_ROAS"
	ConversionOptimizationGoal_7DAYMONETIZATIONROAS      ConversionOptimizationGoal = "GOAL_7DAY_MONETIZATION_ROAS"
	ConversionOptimizationGoal_15DAYMONETIZATIONROAS     ConversionOptimizationGoal = "GOAL_15DAY_MONETIZATION_ROAS"
	ConversionOptimizationGoal_7DAYRETENTIONTIMES        ConversionOptimizationGoal = "GOAL_7DAY_RETENTION_TIMES"
	ConversionOptimizationGoal_7DAYLONGTERMPURCHASEROAS  ConversionOptimizationGoal = "GOAL_7DAY_LONGTERM_PURCHASE_ROAS"
	ConversionOptimizationGoal_14DAYLONGTERMPURCHASEROAS ConversionOptimizationGoal = "GOAL_14DAY_LONGTERM_PURCHASE_ROAS"
	ConversionOptimizationGoal_30DAYLONGTERMPURCHASEROAS ConversionOptimizationGoal = "GOAL_30DAY_LONGTERM_PURCHASE_ROAS"
)
