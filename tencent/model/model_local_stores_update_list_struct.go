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
type LocalStoresUpdateListStruct struct {
	PoiId                     *string                     `json:"poi_id,omitempty"`
	LocalStoreName            *string                     `json:"local_store_name,omitempty"`
	LocalStoreProvince        *string                     `json:"local_store_province,omitempty"`
	LocalStoreCity            *string                     `json:"local_store_city,omitempty"`
	LocalStoreAddress         *string                     `json:"local_store_address,omitempty"`
	LocalStoreBizInfo         *LocalStoreBizInfoStructRsp `json:"local_store_biz_info,omitempty"`
	WechatEcosystemAccounts   *WechatEcosystemAccounts    `json:"wechat_ecosystem_accounts,omitempty"`
	LocalStoreStreet          *string                     `json:"local_store_street,omitempty"`
	LocalStoreBusinessArea    *string                     `json:"local_store_business_area,omitempty"`
	LocalStoreDistrict        *string                     `json:"local_store_district,omitempty"`
	LocalStoreLocation        *LocalStoreLocation         `json:"local_store_location,omitempty"`
	WechatWorkCorpId          *int64                      `json:"wechat_work_corp_id,omitempty"`
	WechatCustomerServiceLink *string                     `json:"wechat_customer_service_link,omitempty"`
	ImageSet                  *[]ImageSetDataStruct       `json:"image_set,omitempty"`
	RecommendWords            *[]RecommendWordStruct      `json:"recommend_words,omitempty"`
	IsUseStandardizedName     *bool                       `json:"is_use_standardized_name,omitempty"`
}
