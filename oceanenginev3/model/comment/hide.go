package comment

import (
	"donson.com.cn/draining/internal/pkg/oceanengine/model"
	"donson.com.cn/draining/internal/pkg/oceanengine/util"
)

type HideReq struct {
	AdvertiserId int64   `json:"advertiser_id"`
	CommentIds   []int64 `json:"comment_ids"`
}

// Encode implement GetRequest interface
func (r HideReq) Encode() []byte {
	return util.JSONMarshal(r)
}

// HideCommentRes 返回的数据
type HideCommentRes struct {
	model.BaseResponse
	//common.OceanEngineApiCommonResultParam
	Data *CommentHideItemData `json:"data"` // json返回值
}

type CommentHideItemData struct {
	SuccessCommentIds []int64 `json:"success_comment_ids"` //项目列表
}

//************************************

type ReplyReq struct {
	AdvertiserId int64   `json:"advertiser_id"`
	CommentIds   []int64 `json:"comment_ids"`
	ReplyText    string  `json:"reply_text"`
}

// Encode implement GetRequest interface
func (r ReplyReq) Encode() []byte {
	return util.JSONMarshal(r)
}

// ReplyCommentRes 返回的数据
type ReplyCommentRes struct {
	model.BaseResponse
	//common.OceanEngineApiCommonResultParam
	Data *CommentReplyItemData `json:"data"` // json返回值
}

type CommentReplyItemData struct {
	SuccessCommentIds []int64 `json:"success_comment_ids"` //项目列表
}

// ****************************
type TopReq struct {
	AdvertiserId int64   `json:"advertiser_id"`
	CommentIds   []int64 `json:"comment_ids"`
	StickType    string  `json:"stick_type"`
}

// Encode implement GetRequest interface
func (r TopReq) Encode() []byte {
	return util.JSONMarshal(r)
}

// TopRes DelTermsRes
type TopRes struct {
	model.BaseResponse
	//common.OceanEngineApiCommonResultParam
	Data struct{} `json:"data"` // json返回值
}
