/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// ConversionScene : 转化场景，网页及微信小程序转化场景无需输入推广目标id
type ConversionScene string

// List of ConversionScene
const (
	ConversionScene_ANDROID             ConversionScene = "CONVERSION_SCENE_ANDROID"
	ConversionScene_IOS                 ConversionScene = "CONVERSION_SCENE_IOS"
	ConversionScene_WEB                 ConversionScene = "CONVERSION_SCENE_WEB"
	ConversionScene_WECHAT_MINI_PROGRAM ConversionScene = "CONVERSION_SCENE_WECHAT_MINI_PROGRAM"
	ConversionScene_WECHAT_MINI_GAME    ConversionScene = "CONVERSION_SCENE_WECHAT_MINI_GAME"
	ConversionScene_QQ_MINI_GAME        ConversionScene = "CONVERSION_SCENE_QQ_MINI_GAME"
	ConversionScene_QUICK_APP           ConversionScene = "CONVERSION_SCENE_QUICK_APP"
	ConversionScene_WE_COM              ConversionScene = "CONVERSION_SCENE_WE_COM"
)