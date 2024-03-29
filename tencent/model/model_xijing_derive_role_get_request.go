/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type XijingDeriveRoleGetRequest struct {
	AccountId   *int64  `json:"account_id,omitempty"`
	RoleId      *string `json:"role_id,omitempty"`
	Copywriting *string `json:"copywriting,omitempty"`
	Layout      *Layout `json:"layout,omitempty"`
	IsLoop      *bool   `json:"is_loop,omitempty"`
	IsClose     *bool   `json:"is_close,omitempty"`
	IsDownload  *bool   `json:"is_download,omitempty"`
}
