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
type LeadCluesGetListStruct struct {
	AccountId                *int64              `json:"account_id,omitempty"`
	LeadsId                  *int64              `json:"leads_id,omitempty"`
	OuterLeadsId             *string             `json:"outer_leads_id,omitempty"`
	ClickId                  *string             `json:"click_id,omitempty"`
	WechatAppid              *string             `json:"wechat_appid,omitempty"`
	AgencyId                 *string             `json:"agency_id,omitempty"`
	AgencyName               *string             `json:"agency_name,omitempty"`
	CampaignId               *int64              `json:"campaign_id,omitempty"`
	CampaignName             *string             `json:"campaign_name,omitempty"`
	AdgroupId                *string             `json:"adgroup_id,omitempty"`
	AdgroupName              *string             `json:"adgroup_name,omitempty"`
	CreativeId               *string             `json:"creative_id,omitempty"`
	CreativeName             *string             `json:"creative_name,omitempty"`
	AdId                     *int64              `json:"ad_id,omitempty"`
	AdName                   *string             `json:"ad_name,omitempty"`
	AdcreativeId             *int64              `json:"adcreative_id,omitempty"`
	AdcreativeName           *string             `json:"adcreative_name,omitempty"`
	ComponentId              *string             `json:"component_id,omitempty"`
	ComponentName            *string             `json:"component_name,omitempty"`
	PageId                   *string             `json:"page_id,omitempty"`
	PageName                 *string             `json:"page_name,omitempty"`
	PageUrl                  *string             `json:"page_url,omitempty"`
	LeadsType                LeadCluesLeadsType  `json:"leads_type,omitempty"`
	LeadsSubType             *string             `json:"leads_sub_type,omitempty"`
	ChatId                   *string             `json:"chat_id,omitempty"`
	LeadsSource              *string             `json:"leads_source,omitempty"`
	LeadsPotentialScore      *string             `json:"leads_potential_score,omitempty"`
	LeadsFollowTag           *string             `json:"leads_follow_tag,omitempty"`
	OuterLeadsConvertType    *string             `json:"outer_leads_convert_type,omitempty"`
	OuterLeadsIneffectReason *string             `json:"outer_leads_ineffect_reason,omitempty"`
	LeadsUserId              *string             `json:"leads_user_id,omitempty"`
	LeadsUserType            LeadsUserType       `json:"leads_user_type,omitempty"`
	LeadsUserWechatAppid     *string             `json:"leads_user_wechat_appid,omitempty"`
	LeadsName                *string             `json:"leads_name,omitempty"`
	LeadsTelephone           *string             `json:"leads_telephone,omitempty"`
	TelephoneLocation        *string             `json:"telephone_location,omitempty"`
	LeadsArea                *string             `json:"leads_area,omitempty"`
	LeadsEmail               *string             `json:"leads_email,omitempty"`
	LeadsQq                  *string             `json:"leads_qq,omitempty"`
	LeadsWechat              *string             `json:"leads_wechat,omitempty"`
	LeadsGender              LeadCluesGenderType `json:"leads_gender,omitempty"`
	Nationality              *string             `json:"nationality,omitempty"`
	WorkingYears             *string             `json:"working_years,omitempty"`
	Age                      *string             `json:"age,omitempty"`
	Profession               *string             `json:"profession,omitempty"`
	IdNumber                 *string             `json:"id_number,omitempty"`
	Address                  *string             `json:"address,omitempty"`
	Bundle                   *string             `json:"bundle,omitempty"`
	PosType                  *int64              `json:"pos_type,omitempty"`
	LeadsCreateTime          *string             `json:"leads_create_time,omitempty"`
	LeadsActionTime          *string             `json:"leads_action_time,omitempty"`
	LeadsTags                *string             `json:"leads_tags,omitempty"`
	ShopName                 *string             `json:"shop_name,omitempty"`
	ShopAddress              *string             `json:"shop_address,omitempty"`
	CallMiddleNum            *string             `json:"call_middle_num,omitempty"`
	CallConsumerHotline      *string             `json:"call_consumer_hotline,omitempty"`
	CallTouchTag             *string             `json:"call_touch_tag,omitempty"`
	CallDuration             *string             `json:"call_duration,omitempty"`
	CallRecordUrl            *string             `json:"call_record_url,omitempty"`
	LayerFormContent         *string             `json:"layer_form_content,omitempty"`
	NickName                 *string             `json:"nick_name,omitempty"`
	IsBroadCastLeads         *string             `json:"is_broad_cast_leads,omitempty"`
	OwnerName                *string             `json:"owner_name,omitempty"`
	OwnerId                  *int64              `json:"owner_id,omitempty"`
	AllFollowRecords         *string             `json:"all_follow_records,omitempty"`
	ClaimAdvertiserId        *int64              `json:"claim_advertiser_id,omitempty"`
}