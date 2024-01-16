/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type ExtendPackageAddResponseData struct {
	PackageId      *int64              `json:"package_id,omitempty"`
	SuccessResults *[]ResultDataStruct `json:"success_results,omitempty"`
	FailedResults  *[]ResultDataStruct `json:"failed_results,omitempty"`
}
