package promotion

import (
	"encoding/json"
	"github.com/disv5/opensdk/oceanengine/util"
	"net/url"
	"strconv"

	"github.com/disv5/opensdk/oceanenginev3/enum"
	"github.com/disv5/opensdk/oceanenginev3/model"
)

// GetRequest 获取广告计划 API Request
type ListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filtering 过滤条件
	Filtering *ListFiltering `json:"filtering,omitempty"`
	// Fields 查询字段集合, 如果指定, 则返回结果数组中, 每个元素是包含所查询字段的字典允许值: "id", "name", "budget", "budget_mode", "status", "opt_status","open_url", "modify_time", "start_time", "end_time", "bid","advertiser_id", "pricing", "flow_control_mode", "download_url", quick_app_url, "inventory_type", "schedule_type", "app_type", "cpa_bid","cpa_skip_first_phrase", "audience", "external_url", "package","campaign_id", "ad_modify_time", "ad_create_time","audit_reject_reason", "retargeting_type", "retargeting_tags","convert_id", "interest_tags", "hide_if_converted","external_actions", "device_type","auto_extend_enabled", "auto_extend_targets", "dpa_lbs", "dpa_city", "dpa_province", "dpa_recommend_type", "roi_goal","subscribe_url","form_id","form_index","app_desc","app_thumbnails"
	Fields []string `json:"fields,omitempty"`
	// Page 页数默认值: 1，page必须大于0
	Page int `json:"page,omitempty"`
	// PageSize 页面大小默认值:10，page_size范围为1-1000
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r ListRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Filtering != nil {
		filtering, _ := json.Marshal(r.Filtering)
		values.Set("filtering", string(filtering))
	}
	if len(r.Fields) > 0 {
		fields, _ := json.Marshal(r.Fields)
		values.Set("fields", string(fields))
	}
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	return values.Encode()
}

// GetFiltering 过滤条件
type ListFiltering struct {
	// IDs 按广告计划ID过滤，范围为1-100
	IDs []uint64 `json:"ids,omitempty"`
	// AdName 按广告计划name过滤，长度为1-30个字符
	Name string `json:"name,omitempty"`
	// Status 按计划状态过滤，默认为返回“所有不包含已删除”，如果要返回所有包含已删除有对应枚举表示
	Status enum.PromotionStatus `json:"status,omitempty"`
	// AdCreateTime 广告计划创建时间，格式"yyyy-mm-dd"，表示过滤出当天创建的广告计划
	PromotionCreateTime string `json:"promotion_create_time,omitempty"`
	// AdModifyTime 广告计划更新时间，格式"yyyy-mm-dd"，表示过滤出当天更新的广告计划
	PromotionModifyTime string `json:"promotion_modify_time,omitempty"`
	// 新增的
	StatusFirst  enum.PromotionStatus `json:"status_first,omitempty"`
	ProjectId    int64                `json:"project_id,omitempty"`
	DeliveryMode string               `json:"delivery_mode,omitempty"`
}

// GetResponse 获取广告计划 API Response
type ListResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *ListResponseData `json:"data,omitempty"`
}

// GetResponseData json返回值
type ListResponseData struct {
	// List 广告数组
	List []Ad `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}

type LowQualityMaterialList struct {
	LowQualityImageIds interface{} `json:"low_quality_image_ids"`
	LowQualityVideoIds interface{} `json:"low_quality_video_ids"`
}
type MaterialScoreInfo struct {
	LowQualityMaterialList LowQualityMaterialList `json:"low_quality_material_list"`
	MaterialAdvice         []string               `json:"material_advice"`
	ScoreNumOfMaterial     string                 `json:"score_num_of_material"`
	ScoreTypeOfMaterial    string                 `json:"score_type_of_material"`
	ScoreValueOfMaterial   string                 `json:"score_value_of_material"`
}
type NativeSetting struct {
	AnchorRelatedType interface{} `json:"anchor_related_type,omitempty"`
	AwemeID           interface{} `json:"aweme_id,omitempty"`
	IsFeedAndFavSee   interface{} `json:"is_feed_and_fav_see,omitempty"`
}
type Images struct {
	ImageID        string `json:"image_id"`
	MaterialID     int64  `json:"material_id"`
	MaterialStatus string `json:"material_status"`
}
type ImageMaterialList struct {
	ImageMode string   `json:"image_mode,omitempty"`
	Images    []Images `json:"images,omitempty"`
}
type VideoMaterialList struct {
	ImageMode         string      `json:"image_mode,omitempty"`
	ItemId            interface{} `json:"item_id,omitempty"`
	MaterialId        int64       `json:"material_id,omitempty"`
	MaterialStatus    string      `json:"material_status,omitempty"`
	VideoCoverId      string      `json:"video_cover_id,omitempty"`
	VideoId           string      `json:"video_id,omitempty"`
	VideoTaskIds      interface{} `json:"video_task_ids,omitempty"`
	VideoTemplateType interface{} `json:"video_template_type,omitempty"`
}

type ProductInfo struct {
	ImageIds      []string `json:"image_ids,omitempty"`
	SellingPoints []string `json:"selling_points,omitempty"`
	Titles        []string `json:"titles,omitempty"`
}
type TitleMaterialList struct {
	Title    string      `json:"title,omitempty"`
	WordList interface{} `json:"word_list,omitempty"`
}
type PromotionMaterials struct {
	OpenUrl                 string               `json:"open_url,omitempty"`
	CallToActionButtons     []string             `json:"call_to_action_buttons,omitempty"`
	ExternalURLMaterialList []string             `json:"external_url_material_list,omitempty"`
	ImageMaterialList       []ImageMaterialList  `json:"image_material_list,omitempty"`
	VideoMaterialList       []VideoMaterialList  `json:"video_material_list,omitempty"`
	ProductInfo             ProductInfo          `json:"product_info,omitempty"`
	TitleMaterialList       []TitleMaterialList  `json:"title_material_list,omitempty"`
	AnchorMaterialList      []AnchorMaterialList `json:"anchor_material_list,omitempty"`
}

type AnchorMaterialList struct {
	AnchorId   string `json:"anchor_id"`
	AnchorType string `json:"anchor_type"`
}

type Ad struct {
	AdDownloadStatus           interface{}        `json:"ad_download_status"`
	AdvertiserID               int64              `json:"advertiser_id"`
	Budget                     float64            `json:"budget"`
	CpaBid                     float64            `json:"cpa_bid"`
	DeepCpabid                 float64            `json:"deep_cpabid"`
	IsCommentDisable           interface{}        `json:"is_comment_disable"`
	LearningPhase              string             `json:"learning_phase"`
	MaterialScoreInfo          MaterialScoreInfo  `json:"material_score_info"`
	MaterialsType              interface{}        `json:"materials_type"`
	NativeSetting              NativeSetting      `json:"native_setting"`
	OptStatus                  string             `json:"opt_status"`
	ProjectID                  int64              `json:"project_id"`
	PromotionCreateTime        string             `json:"promotion_create_time"`
	PromotionID                int64              `json:"promotion_id"`
	PromotionMaterials         PromotionMaterials `json:"promotion_materials"`
	PromotionModifyTime        string             `json:"promotion_modify_time"`
	PromotionName              string             `json:"promotion_name"`
	RoiGoal                    interface{}        `json:"roi_goal"`
	Source                     string             `json:"source"`
	Status                     string             `json:"status"`
	CreativeAutoGenerateSwitch string             `json:"creative_auto_generate_switch"`
	ConfigId                   int64              `json:"config_id"`
	BudgetMode                 string             `json:"budget_mode"`
}
type PageInfo struct {
	Page        int `json:"page"`
	PageSize    int `json:"page_size"`
	TotalNumber int `json:"total_number"`
	TotalPage   int `json:"total_page"`
}

// ReqUpdateAd 修改广告
type ReqUpdateAd struct {
	AdvertiserId       int64              `form:"advertiser_id" json:"advertiser_id" binding:"required"`    // 广告账户id
	PromotionId        int64              `form:"promotion_id" json:"promotion_id" binding:"required"`      // 广告ID
	Name               string             `form:"name" json:"name" binding:"required"`                      // 广告名称
	Operation          string             `form:"operation" json:"operation,omitempty"`                     // 广告状态
	MaterialsType      string             `form:"materials_type" json:"materials_type,omitempty"`           // 素材类型，直播场景必填，允许值：LIVE_MATERIALS直播素材，PROMOTION_MATERIALS广告素材
	PromotionMaterials PromotionMaterials `form:"promotion_materials" json:"promotion_materials,omitempty"` // 直播素材与广告素材组合
	NativeSetting      *NativeSetting     `form:"native_setting" json:"native_setting,omitempty"`           //原生广告设置(支持直播链路）
	AdSource           string             `form:"source" json:"source,omitempty"`                           // 广告来源，字数限制：[2-10]（英文字符限制：[4-20]）。当landing_type = LINK、MICRO_GAME、SHOP、QUICK_APP时必填。
	IsCommentDisable   string             `form:"is_comment_disable" json:"is_comment_disable,omitempty"`   // 广告评论，ON为启用，OFF为不启用， 默认值：ON。允许值: ON、OFF。注：inventory_type 仅包含INVENTORY_UNION_SLOT（穿山甲）时，不支持广告评论。
	AdDownloadStatus   string             `form:"ad_download_status" json:"ad_download_status,omitempty"`   // 客户端下载视频功能，ON为启用，OFF为不启用，默认值：OFF。允许值: ON、OFF。若关闭，用户则无法下载您本地上传的视频，该设置在广告投放后无法修改。注：仅支持账户信息推广身份（aweme_id传入时，ad_download_status传入无效）；inventory_type 仅包含INVENTORY_UNION_SLOT（穿山甲）时，不支持广告评论。

	//智能创意生成
	CreativeAutoGenerateSwitch string  `form:"creative_auto_generate_switch" json:"creative_auto_generate_switch,omitempty"` //是否开启自动生成素材 默认值：OFF 枚举值：ON开启、OFF不开启
	ConfigId                   int64   `form:"config_id" json:"config_id,omitempty"`                                         //配置ID，开关打开，不传为黑盒明投派生
	Budget                     float64 `json:"budget,omitempty"`
	CpaBid                     float64 `json:"cpa_bid,omitempty"`
	DeepCpabid                 float64 `json:"deep_cpabid,omitempty"`
	RoiGoal                    float64 `json:"roi_goal,omitempty"`
}

// Encode implement PostRequest interface
func (r ReqUpdateAd) Encode() []byte {
	//str := util.JSONMarshal(r)
	//strr := string(str)
	//fmt.Println(strr)
	//if r.NativeSetting.AnchorRelatedType == nil {
	//	r.NativeSetting = struct{}{}
	//}
	return util.JSONMarshal(r)
}

// UpdateAdRes 修改广告返回数据
type UpdateAdRes struct {
	model.BaseResponse
	Data *UpdateDataRes `form:"data" json:"data,omitempty"`
}
type UpdateDataRes struct {
	PromotionId int `form:"promotion_id" json:"promotion_id,omitempty"` // 广告ID
	ErrorList   []struct {
		ErrorCode    int    `json:"error_code"`
		ErrorMessage string `json:"error_message"`
		ObjectType   string `json:"object_type"`
	} `json:"error_list"`
}

// ReqStatusUpdateAd 暂停
type ReqStatusUpdateAd struct {
	AdvertiserId int64 `form:"advertiser_id" json:"advertiser_id" binding:"required"` // 广告账户id
	Data         []struct {
		PromotionId int64  `json:"promotion_id"`
		OptStatus   string `json:"opt_status"`
	} `json:"data"`
}

// Encode implement PostRequest interface
func (r ReqStatusUpdateAd) Encode() []byte {
	//str := util.JSONMarshal(r)
	//strr := string(str)
	//fmt.Println(strr)
	//if r.NativeSetting.AnchorRelatedType == nil {
	//	r.NativeSetting = struct{}{}
	//}
	return util.JSONMarshal(r)
}

// UpdateStatusAdRes 修改广告状态返回数据
type UpdateStatusAdRes struct {
	model.BaseResponse
	Data *UpdateStatusDataRes `form:"data" json:"data,omitempty"`
}
type UpdateStatusDataRes struct {
	PromotionIds []int          `form:"promotion_ids" json:"promotion_ids,omitempty"` // 更新成功的广告ID集合
	StatusErrors []StatusErrors `form:"errors" json:"errors,omitempty"`
}

type StatusErrors struct {
	PromotionId  int    `form:"promotion_id" json:"promotion_id,omitempty"`   // 广告ID
	ErrorMessage string `form:"error_message" json:"error_message,omitempty"` // 错误信息
}

//****************二跳用的

// ReqUpdateAdTwo 二跳用的修改广告
type ReqUpdateAdTwo struct {
	AdvertiserId       int64                 `form:"advertiser_id" json:"advertiser_id" binding:"required"`    // 广告账户id
	PromotionId        int64                 `form:"promotion_id" json:"promotion_id" binding:"required"`      // 广告ID
	Name               string                `form:"name" json:"name" binding:"required"`                      // 广告名称
	Operation          string                `form:"operation" json:"operation,omitempty"`                     // 广告状态
	MaterialsType      string                `form:"materials_type" json:"materials_type,omitempty"`           // 素材类型，直播场景必填，允许值：LIVE_MATERIALS直播素材，PROMOTION_MATERIALS广告素材
	PromotionMaterials PromotionMaterialsTwo `form:"promotion_materials" json:"promotion_materials,omitempty"` // 直播素材与广告素材组合
	NativeSetting      *NativeSetting        `form:"native_setting" json:"native_setting,omitempty"`           //原生广告设置(支持直播链路）
	AdSource           string                `form:"source" json:"source,omitempty"`                           // 广告来源，字数限制：[2-10]（英文字符限制：[4-20]）。当landing_type = LINK、MICRO_GAME、SHOP、QUICK_APP时必填。
	IsCommentDisable   string                `form:"is_comment_disable" json:"is_comment_disable,omitempty"`   // 广告评论，ON为启用，OFF为不启用， 默认值：ON。允许值: ON、OFF。注：inventory_type 仅包含INVENTORY_UNION_SLOT（穿山甲）时，不支持广告评论。
	AdDownloadStatus   string                `form:"ad_download_status" json:"ad_download_status,omitempty"`   // 客户端下载视频功能，ON为启用，OFF为不启用，默认值：OFF。允许值: ON、OFF。若关闭，用户则无法下载您本地上传的视频，该设置在广告投放后无法修改。注：仅支持账户信息推广身份（aweme_id传入时，ad_download_status传入无效）；inventory_type 仅包含INVENTORY_UNION_SLOT（穿山甲）时，不支持广告评论。

	//智能创意生成
	CreativeAutoGenerateSwitch string  `form:"creative_auto_generate_switch" json:"creative_auto_generate_switch,omitempty"` //是否开启自动生成素材 默认值：OFF 枚举值：ON开启、OFF不开启
	ConfigId                   int64   `form:"config_id" json:"config_id,omitempty"`                                         //配置ID，开关打开，不传为黑盒明投派生
	Budget                     float64 `json:"budget,omitempty"`
	CpaBid                     float64 `json:"cpa_bid,omitempty"`
	DeepCpabid                 float64 `json:"deep_cpabid,omitempty"`
	RoiGoal                    float64 `json:"roi_goal,omitempty"`
}

// Encode implement PostRequest interface
func (r ReqUpdateAdTwo) Encode() []byte {
	//str := util.JSONMarshal(r)
	//strr := string(str)
	//fmt.Println(strr)
	//if r.NativeSetting.AnchorRelatedType == nil {
	//	r.NativeSetting = struct{}{}
	//}
	return util.JSONMarshal(r)
}

type PromotionMaterialsTwo struct {
	OpenUrl                 string               `json:"open_url"`
	CallToActionButtons     []string             `json:"call_to_action_buttons,omitempty"`
	ExternalURLMaterialList []string             `json:"external_url_material_list,omitempty"`
	ImageMaterialList       []ImageMaterialList  `json:"image_material_list,omitempty"`
	VideoMaterialList       []VideoMaterialList  `json:"video_material_list,omitempty"`
	ProductInfo             ProductInfo          `json:"product_info,omitempty"`
	TitleMaterialList       []TitleMaterialList  `json:"title_material_list,omitempty"`
	AnchorMaterialList      []AnchorMaterialList `json:"anchor_material_list,omitempty"`
}
