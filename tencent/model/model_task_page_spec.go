/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 落地页信息
type TaskPageSpec struct {
	PageId                   *int64                   `json:"page_id,omitempty"`
	PageUrl                  *string                  `json:"page_url,omitempty"`
	OverrideCanvasHeadOption OverrideCanvasHeadOption `json:"override_canvas_head_option,omitempty"`
}
