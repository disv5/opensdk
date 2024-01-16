/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type LocalStoresUpdateRequest struct {
	AccountId      *int64                    `json:"account_id,omitempty"`
	LocalStoreList *[]UpdateLocalStoreStruct `json:"local_store_list,omitempty"`
}
