/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 应用直达选项
type DeepLinkEnabled struct {
	IosDeepLinkEnabled     *bool `json:"ios_deep_link_enabled,omitempty"`
	AndroidDeepLinkEnabled *bool `json:"android_deep_link_enabled,omitempty"`
	H5DeepLinkEnabled      *bool `json:"h5_deep_link_enabled,omitempty"`
}
