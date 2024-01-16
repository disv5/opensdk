package terms

import (
	"github.com/disv5/opensdk/oceanengine/model"
	"github.com/disv5/opensdk/oceanengine/util"
	"github.com/disv5/opensdk/oceanenginev3/model/promotion"
	"net/url"
	"strconv"
)

// AddTermsReq 添加屏蔽词
type AddTermsReq struct {
	AdvertiserId int64    `json:"advertiser_id"`
	AwemeId      string   `json:"aweme_id"`
	IsApplyToAdv bool     `json:"is_apply_to_adv"`
	Terms        []string `json:"terms"`
}

// Encode implement GetRequest interface
func (r AddTermsReq) Encode() []byte {
	return util.JSONMarshal(r)
}

// AddTermsRes AddTermsRes
type AddTermsRes struct {
	model.BaseResponse
	//common.OceanEngineApiCommonResultParam
	Data struct{} `json:"data"` // json返回值
}

// DelTermsReq 删除屏蔽词
type DelTermsReq struct {
	AdvertiserId int64    `json:"advertiser_id"`
	AwemeId      string   `json:"aweme_id"`
	IsApplyToAdv bool     `json:"is_apply_to_adv"`
	Terms        []string `json:"terms"`
}

// Encode implement GetRequest interface
func (r DelTermsReq) Encode() []byte {
	return util.JSONMarshal(r)
}

// DelTermsRes DelTermsRes
type DelTermsRes struct {
	model.BaseResponse
	//common.OceanEngineApiCommonResultParam
	Data struct{} `json:"data"` // json返回值
}

// UpdateTermsReq 更新屏蔽词
type UpdateTermsReq struct {
	AdvertiserId int64  `json:"advertiser_id"`
	AwemeId      string `json:"aweme_id,omitempty"`
	IsApplyToAdv bool   `json:"is_apply_to_adv"`
	OriginTerm   string `json:"origin_term"`
	NewTerm      string `json:"new_term"`
}

// Encode implement GetRequest interface
func (r UpdateTermsReq) Encode() []byte {
	return util.JSONMarshal(r)
}

// UpdateTermsRes DelTermsRes
type UpdateTermsRes struct {
	model.BaseResponse
	//common.OceanEngineApiCommonResultParam
	Data struct{} `json:"data"` // json返回值
}

// TermsListReq 获取屏蔽词
type TermsListReq struct {
	AdvertiserId int64  `json:"advertiser_id"`
	AwemeId      string `json:"aweme_id"`
	IsApplyToAdv bool   `json:"is_apply_to_adv"`
	Page         int    `json:"page"`
	PageSize     int    `json:"page_size"`
}

// Encode implement GetRequest interface
func (r TermsListReq) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatInt(r.AdvertiserId, 10))

	values.Set("aweme_id", r.AwemeId)
	values.Set("is_apply_to_adv", strconv.FormatBool(r.IsApplyToAdv))
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	return values.Encode()
}

// TermsListRes
type TermsListRes struct {
	model.BaseResponse
	Data *TermsListItermData `json:"data"` // json返回值
}

type TermsListItermData struct {
	TermsListIterms []string           `json:"terms"`               //
	PageInfo        promotion.PageInfo `json:"page_info,omitempty"` //分页信息
}

// AddUserTermsReq 添加屏蔽用户
type AddUserTermsReq struct {
	AdvertiserId     int64    `json:"advertiser_id"`
	BannedType       string   `json:"banned_type"`
	AwemeId          string   `json:"aweme_id,omitempty"`
	IsApplyToAdv     bool     `json:"is_apply_to_adv"`
	AwemeUserIds     []string `json:"aweme_user_ids,omitempty"`
	NicknameKeywords []string `json:"nickname_keywords,omitempty"`
}

// Encode implement GetRequest interface
func (r AddUserTermsReq) Encode() []byte {
	return util.JSONMarshal(r)
}

// AddUserTermsRes
type AddUserTermsRes struct {
	model.BaseResponse
	Data *AddUserTermData `json:"data"` // json返回值
}

type AddUserTermData struct {
	Success []string `json:"success"` //
	Fail    []*Fail  `json:"fail"`    //
}

type Fail struct {
	ResultKey string `json:"result_key"` //
	Message   string `json:"message"`    //
}

// DelUserTermsReq 删除屏蔽用户
type DelUserTermsReq struct {
	AdvertiserId     int64    `json:"advertiser_id"`
	BannedType       string   `json:"banned_type"`
	AwemeId          string   `json:"aweme_id,omitempty"`
	IsApplyToAdv     bool     `json:"is_apply_to_adv"`
	AwemeUserIds     []string `json:"aweme_user_ids,omitempty"`
	NicknameKeywords []string `json:"nickname_keywords,omitempty"`
}

// Encode implement GetRequest interface
func (r DelUserTermsReq) Encode() []byte {
	return util.JSONMarshal(r)
}

// DelUserTermsRes
type DelUserTermsRes struct {
	model.BaseResponse
	Data *DelUserTermData `json:"data"` // json返回值
}

type DelUserTermData struct {
	Success []string `json:"success"` //
	Fail    []*Fail  `json:"fail"`    //
}

// TermsUserListReq 获取屏蔽用户
type TermsUserListReq struct {
	AdvertiserId    int64  `json:"advertiser_id"`
	AwemeId         string `json:"aweme_id,omitempty"`
	IsApplyToAdv    bool   `json:"is_apply_to_adv,omitempty"`
	BannedType      string `json:"banned_type,omitempty"`
	AwemeUserId     string `json:"aweme_user_id,omitempty"`
	NicknameKeyword string `json:"nickname_keyword,omitempty"`
	Page            int    `json:"page"`
	PageSize        int    `json:"page_size"`
}

// Encode implement GetRequest interface
func (r TermsUserListReq) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatInt(r.AdvertiserId, 10))

	values.Set("is_apply_to_adv", strconv.FormatBool(r.IsApplyToAdv))
	if r.AwemeId != "" {
		values.Set("aweme_id", r.AwemeId)
	}
	if r.BannedType != "" {
		values.Set("banned_type", r.BannedType)
	}
	if r.AwemeUserId != "" {
		values.Set("aweme_user_id", r.AwemeUserId)
	}
	if r.NicknameKeyword != "" {
		values.Set("nickname_keyword", r.NicknameKeyword)
	}
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	return values.Encode()
}

// TermsUserListRes
type TermsUserListRes struct {
	model.BaseResponse
	Data *TermsUserListItermData `json:"data"` // json返回值
}

type TermsUserListItermData struct {
	List     []TermsUserListIterm `json:"list"`                //
	PageInfo promotion.PageInfo   `json:"page_info,omitempty"` //分页信息
}

type TermsUserListIterm struct {
	AwemeName       string `json:"aweme_name"`
	BannedType      string `json:"banned_type"`
	AwemeUserId     string `json:"aweme_user_id"`
	NicknameKeyword string `json:"nickname_keyword"`
}
