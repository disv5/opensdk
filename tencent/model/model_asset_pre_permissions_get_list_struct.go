/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 返回结构
type AssetPrePermissionsGetListStruct struct {
	AccountId      *int64       `json:"account_id,omitempty"`
	OwnerAccountId *int64       `json:"owner_account_id,omitempty"`
	AssetId        *int64       `json:"asset_id,omitempty"`
	AssetName      *string      `json:"asset_name,omitempty"`
	AssetType      AssetType    `json:"asset_type,omitempty"`
	PathId         *int64       `json:"path_id,omitempty"`
	PathType       PathType     `json:"path_type,omitempty"`
	IsGrantedAll   GrantAllType `json:"is_granted_all,omitempty"`
	GrantTime      *int64       `json:"grant_time,omitempty"`
	PermissionList *[]string    `json:"permission_list,omitempty"`
}
