/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 处理操作信息
type CropOperationSpec struct {
	CropCustomizedSpec          *CropCustomizedSpec          `json:"crop_customized_spec,omitempty"`
	CropSmartSpec               *CropSmartSpec               `json:"crop_smart_spec,omitempty"`
	ResizeSpec                  *ResizeSpec                  `json:"resize_spec,omitempty"`
	CropCustomizedSpecAndResize *CropCustomizedSpecAndResize `json:"crop_customized_spec_and_resize,omitempty"`
}
