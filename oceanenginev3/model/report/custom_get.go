package report

import (
	"encoding/json"
	"github.com/disv5/opensdk/oceanenginev3/enum"
	"net/url"
	"strconv"

	"github.com/disv5/opensdk/oceanenginev3/model"
)

// GetRequest 获取广告计划 API Request
type CustomGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// 数据主题，默认值：BASIC_DATA广告基础数据 枚举值：BASIC_DATA广告基础数据、QUERY_DATA搜索词数据、BIDWORD_DATA关键词数据、MATERIAL_DATA素材数据
	DataTopic enum.DataTopicType `json:"data_topic,omitempty"`
	// 维度列表
	Dimensions []string `json:"dimensions"`
	// 指标列表
	Metrics []string `json:"metrics"`
	// Filtering 过滤条件
	Filters   []CustomGetFiltering `json:"filters"`
	StartTime string               `json:"start_time"`
	EndTime   string               `json:"end_time"`
	// 排序
	OrderBy []OrderBy `json:"order_by"`
	// Page 页数默认值: 1，page必须大于0
	Page int `json:"page,omitempty"`
	// PageSize 页面大小默认值:10，page_size范围为1-1000
	PageSize int `json:"page_size,omitempty"`
}

type OrderBy struct {
	Field string           `json:"field"`
	Type  enum.OrderByType `json:"type"`
}

// Encode implement GetRequest interface
func (r CustomGetRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("start_time", r.StartTime)
	values.Set("end_time", r.EndTime)
	if r.Filters != nil {
		filters, _ := json.Marshal(r.Filters)
		values.Set("filters", string(filters))
	}
	if len(r.Dimensions) > 0 {
		dimensions, _ := json.Marshal(r.Dimensions)
		values.Set("dimensions", string(dimensions))
	}
	if len(r.Metrics) > 0 {
		metrics, _ := json.Marshal(r.Metrics)
		values.Set("metrics", string(metrics))
	}
	if len(r.OrderBy) > 0 {
		orderBy, _ := json.Marshal(r.OrderBy)
		values.Set("order_by", string(orderBy))
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
type CustomGetFiltering struct {
	Field    string   `json:"field,omitempty"`  // 过滤的消耗指标字段
	Type     int      `json:"type,omitempty"`   // 字段类型
	Operator int      `json:"operator"`         // 处理方式
	Values   []string `json:"values,omitempty"` // 过滤字段具体值
}

// GetResponse 获取广告计划 API Response
type CustomGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *CustomGetResponseData `json:"data,omitempty"`
}

type PageInfo struct {
	Page        int `json:"page"`
	PageSize    int `json:"page_size"`
	TotalNumber int `json:"total_number"`
	TotalPage   int `json:"total_page"`
}
type Dimensions struct {
	CdpPromotionID string `json:"cdp_promotion_id"`
	StatTimeDay    string `json:"stat_time_day"`
	StatTimeHour   string `json:"stat_time_hour"`
}

type DimensionsHour struct {
	CdpPromotionID string `json:"cdp_promotion_id"`
	StatTimeHour   string `json:"stat_time_hour"`
}
type Metrics struct {
	ClickCnt              string `json:"click_cnt"`
	ConversionCost        string `json:"conversion_cost"`
	ConversionRate        string `json:"conversion_rate"`
	ConvertCnt            string `json:"convert_cnt"`
	CpcPlatform           string `json:"cpc_platform"`
	CpmPlatform           string `json:"cpm_platform"`
	Ctr                   string `json:"ctr"`
	ShowCnt               string `json:"show_cnt"`
	StatCost              string `json:"stat_cost"`
	AttributionConvertCnt string `json:"attribution_convert_cnt"`
	// v2.44 新增字段
	TotalPlay       string `json:"total_play"`
	ValidPlay       string `json:"valid_play"`
	PlayDuration3s  string `json:"play_duration_3s"`
	Play25FeedBreak string `json:"play_25_feed_break"`
	Play50FeedBreak string `json:"play_50_feed_break"`
	Play75FeedBreak string `json:"play_75_feed_break"`
	Play99FeedBreak string `json:"play_99_feed_break"`
}
type Row struct {
	Dimensions Dimensions `json:"dimensions"`
	Metrics    Metrics    `json:"metrics"`
}
type TotalMetrics struct {
	ClickCnt       string `json:"click_cnt"`
	ConversionCost string `json:"conversion_cost"`
	ConversionRate string `json:"conversion_rate"`
	ConvertCnt     string `json:"convert_cnt"`
	CpcPlatform    string `json:"cpc_platform"`
	CpmPlatform    string `json:"cpm_platform"`
	Ctr            string `json:"ctr"`
	ShowCnt        string `json:"show_cnt"`
	StatCost       string `json:"stat_cost"`
}
type CustomGetResponseData struct {
	PageInfo     *PageInfo    `json:"page_info"`
	Rows         []Row        `json:"rows"`
	TotalMetrics TotalMetrics `json:"total_metrics"`
}
