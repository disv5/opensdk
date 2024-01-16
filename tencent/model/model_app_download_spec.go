/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 应用下载信息
type AppDownloadSpec struct {
	Title          *string         `json:"title,omitempty"`
	AppIosSpec     *AppIosSpec     `json:"app_ios_spec,omitempty"`
	AppAndroidSpec *AppAndroidSpec `json:"app_android_spec,omitempty"`
}
