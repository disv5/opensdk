package comment

import (
	"encoding/json"
	"github.com/disv5/opensdk/oceanengine/model"
	"github.com/disv5/opensdk/oceanenginev3/model/promotion"
	"net/url"
	"strconv"
)

type GetCommentListReq struct {
	AdvertiserId    int64      `json:"advertiser_id"`
	PlatformVersion string     `json:"platform_version"`
	StartTime       string     `json:"start_time"`
	EndTime         string     `json:"end_time"`
	OrderField      string     `json:"order_field"`
	OrderType       string     `json:"order_type"`
	Filtering       *Filtering `json:"filtering"`
	IncludeMetrics  bool       `json:"include_metrics"`
	Page            int        `json:"page"`
	PageSize        int        `json:"page_size"`
}

type Filtering struct {
	AdIds        []int64 `json:"ad_ids"`
	CreativeIds  []int64 `json:"creative_ids"`
	ItemIds      []int64 `json:"item_ids"`
	BindRelation string  `json:"bind_relation"`
	LevelType    string  `json:"level_type"`
	IsReplied    int     `json:"is_replied,omitempty"`
	HideStatus   string  `json:"hide_status"`
	EmotionType  string  `json:"emotion_type,omitempty"`
}

// Encode implement GetRequest interface
func (r GetCommentListReq) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatInt(r.AdvertiserId, 10))
	if r.Filtering != nil {
		filtering, _ := json.Marshal(r.Filtering)
		values.Set("filtering", string(filtering))
	}

	values.Set("platform_version", r.PlatformVersion)
	values.Set("start_time", r.StartTime)
	values.Set("end_time", r.EndTime)
	values.Set("order_type", r.OrderType)
	values.Set("order_field", r.OrderField)
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	return values.Encode()
}

// GetCommentListRes 获取项目列表返回的数据
type GetCommentListRes struct {
	model.BaseResponse
	//common.OceanEngineApiCommonResultParam
	Data *CommentListItemData `json:"data"` // json返回值
}

type CommentListItemData struct {
	CommentList []CommentListItem  `json:"comment_list"`        //项目列表
	PageInfo    promotion.PageInfo `json:"page_info,omitempty"` //分页信息
}
type CommentListItem struct {
	AdId        int    `json:"ad_id"`
	AdName      string `json:"ad_name"`
	AwemeId     string `json:"aweme_id"`
	AwemeName   string `json:"aweme_name"`
	CommentId   int64  `json:"comment_id"`
	CreateTime  string `json:"create_time"`
	CreativeId  int64  `json:"creative_id"`
	EmotionType string `json:"emotion_type"`
	HideStatus  string `json:"hide_status"`
	IsStick     int    `json:"is_stick"`
	ItemId      int64  `json:"item_id"`
	ItemTitle   string `json:"item_title"`
	LevelType   string `json:"level_type"`
	LikeCount   int    `json:"like_count"`
	PromotionId int    `json:"promotion_id"`
	ReplyCount  int    `json:"reply_count"`
	ReplyStatus string `json:"reply_status"`
	Text        string `json:"text"`
}
