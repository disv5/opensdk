/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type BrandAddResponseData struct {
	AccountId   *int64  `json:"account_id,omitempty"`
	Name        *string `json:"name,omitempty"`
	ImageId     *string `json:"image_id,omitempty"`
	Width       *int64  `json:"width,omitempty"`
	Height      *int64  `json:"height,omitempty"`
	ImageUrl    *string `json:"image_url,omitempty"`
	CreatedTime *int64  `json:"created_time,omitempty"`
}
