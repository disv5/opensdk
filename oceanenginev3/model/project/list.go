package project

import (
	"encoding/json"
	"github.com/disv5/opensdk/oceanengine/model"
	"github.com/disv5/opensdk/oceanengine/util"
	"github.com/disv5/opensdk/oceanenginev3/model/promotion"
	"net/url"
	"strconv"
)

// GetProjectListReq 获取项目列表请求的数据
type GetProjectListReq struct {
	AdvertiserID uint64     `json:"advertiser_id" binding:"required"` // 广告账户ID
	Filtering    *Filtering `json:"filtering,omitempty"`              // 过滤条件，若此字段不传，或传空则视为无限制条件
	Fields       []string   `json:"fields,omitempty"`                 // 查询字段集合，如果指定则应答结果仅返回指定字段 可参考应答参数返回的指标字段（不支持audience下字段筛选）
	Page         int        `json:"page"`                             // 页数默认值：1，page范围为[1,99999]
	PageSize     int        `json:"page_size"`                        // 页面大小默认值：10，page_size范围为[1,100]
}

type Filtering struct {
	Ids               []int64 `json:"ids,omitempty"`                 // 按广告项目ID过滤，范围为1-100
	DeliveryMode      string  `json:"delivery_mode,omitempty"`       // 投放模式，允许值：MANUAL手动投放、PROCEDURAL自动投放
	LandingType       string  `json:"landing_type,omitempty"`        // 推广目的
	AppPromotionType  string  `json:"app_promotion_type,omitempty"`  // 子目标
	MarketingGoal     string  `json:"marketing_goal,omitempty"`      // 营销场景
	AdType            string  `json:"ad_type,omitempty"`             // 广告类型
	Name              string  `json:"name,omitempty"`                // 项目名称
	Status            string  `json:"status,omitempty"`              // 广告项目状态过滤
	ProjectCreateTime string  `json:"project_create_time,omitempty"` // 项目创建时间
	ProjectModifyTime string  `json:"project_modify_time,omitempty"` // 项目更新时间
	Pricing           string  `json:"pricing,omitempty"`             // 计费方式
	InventoryType     string  `json:"inventory_type,omitempty"`      // 首选位置
	Platform          string  `json:"platform,omitempty"`            // 按平台过滤，允许值：IOS、ANDROID
}

// Encode implement GetRequest interface
func (r GetProjectListReq) Encode() string {
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

// GetProjectListRes 获取项目列表返回的数据
type GetProjectListRes struct {
	model.BaseResponse
	//common.OceanEngineApiCommonResultParam
	Data *ProjectListItemData `json:"data"` // json返回值
}

type ProjectListItemData struct {
	List     []ProjectListItem  `json:"list"`                //项目列表
	PageInfo promotion.PageInfo `json:"page_info,omitempty"` //分页信息
}

// ProjectListItem 项目列表项
type ProjectListItem struct {
	ProjectId           int64           `json:"project_id"`             // 项目ID
	AdvertiserId        int64           `json:"advertiser_id"`          // 广告账户id
	DeliveryMode        string          `json:"delivery_mode"`          // 投放模式，允许值：MANUAL手动投放(默认值）、PROCEDURAL自动投放(仅支持landing_type=APP或MICRO_GAME ）
	LandingType         string          `json:"landing_type"`           // 推广目的，允许值：APP 应用推广、LINK 销售线索推广、MICRO_GAME 小游戏、SHOP 电商店铺推广、QUICK_APP快应用, 必填
	AppPromotionType    string          `json:"app_promotion_type"`     // 子目标，当 landing_type = APP 有效且必填
	MarketingGoal       string          `json:"marketing_goal"`         // 营销场景，允许值：VIDEO_AND_IMAGE 短视频/图片, LIVE直播, 必填
	AdType              string          `json:"ad_type"`                // 广告类型，允许值：ALL，
	OptStatus           string          `json:"opt_status"`             // 目标操作，枚举值：ENABLE 启用项目、DISABLE暂停项目。
	Name                string          `json:"name"`                   // 项目名称，长度是1-50个字（两个英文字符占1个字），必填
	ProjectCreateTime   string          `json:"project_create_time"`    // 项目创建时间
	ProjectModifyTime   string          `json:"project_modify_time"`    // 项目更新时间
	Status              string          `json:"status"`                 // 广告项目状态过滤
	Pricing             string          `json:"pricing"`                // 计费方式
	PackageName         string          `json:"package_name"`           // 应用包名
	AppName             string          `json:"app_name"`               // 应用名
	FeedDeliverySearch  string          `json:"feed_delivery_search"`   // 搜索快投关键词功能，HAS_OPEN:启用，DISABLED:未启用
	SearchBidRatio      float64         `json:"search_bid_ratio"`       // 出价系数
	AudienceExtend      string          `json:"audience_extend"`        // 定向拓展
	Keywords            []Keyword       `json:"keywords"`               // 搜索关键词列表
	RelatedProduct      RelatedProduct  `json:"related_product"`        // 关联产品投放相关
	AssetType           string          `json:"asset_type"`             // 资产类型，landing_type = LINK 或SHOP时有效且必填 允许值： ORANGE 橙子落地页、THIRDPARTY 自研落地页
	QuickAppId          int             `json:"quick_app_id"`           // 快应用资产id ，仅支持已通过审核的快应用资产
	MicroPromotionType  string          `json:"micro_promotion_type"`   // 小程序类型，landing_type = MICRO_GAME 时有效且必填 允许值： WECHAT_GAME 微信小游戏、WECHAT_APP微信小程序、BYTE_GAME字节小游戏、BYTE_APP字节小程序
	DownloadUrl         string          `json:"download_url"`           // 下载链接，landing_type=APP 子目标为 DOWNLOAD 或者LAUNCH 时有效且必填
	DownloadType        string          `json:"download_type"`          // 下载方式，landing_type=APP 子目标为 DOWNLOAD 时有效，子目标为LAUNCH时无需传入
	DownloadMode        string          `json:"download_mode"`          // 优先从系统应用商店下载（下载模式），landing_type=APP 子目标为 DOWNLOAD 时有效，子目标为LAUNCH时无需传入
	LaunchType          string          `json:"launch_type"`            // 调起方式， landing_type = APP 且子目标为LAUNCH有效
	OpenUrl             string          `json:"open_url"`               // Deeplink直达链接，landing_type = APP 且子目标为 LAUNCH 时有效且必填
	UlinkUrl            string          `json:"ulink_url"`              // ulink直达链接，landing_type = APP 且子目标为LAUNCH 时有效
	SubscribeUrl        string          `json:"subscribe_url"`          // 预约下载链接，landing_type = APP 且子目标为LAUNCH 时有效且必填
	OptimizeGoal        OptimizeGoal    `json:"optimize_goal"`          // 优化目标
	LandingPageStayTime int64           `json:"landing_page_stay_time"` // 店铺停留时长，单位为毫秒，当external_action为AD_CONVERT_TYPE_STAY_TIME时有效且必填 允许值：1000 、5000 、12000 、20000、30000、40000、50000、60000
	DeliveryRange       DeliveryRange   `json:"delivery_range"`         // 广告版位
	Audience            Audience        `json:"audience"`               // 人群定向
	DeliverySetting     DeliverySetting `json:"delivery_setting"`       // 排期、预算与出价
	TrackUrlSetting     TrackUrlSetting `json:"track_url_setting"`      // 监测链接
}

/**********************************
*********************************/

// UpdateProjectReq 更新项目请求参数
type UpdateProjectReq struct {
	AdvertiserId      uint64          `json:"advertiser_id" binding:"required"` // 广告账户id, 必填
	ProjectId         string          `json:"project_id" binding:"required"`    // 项目id
	Name              string          `json:"name,omitempty"`                   // 项目名称，长度是1-50个字（两个英文字符占1个字）
	SearchBidRatioKey                 // 搜索快投
	Keywords          []Keyword       `json:"keywords,omitempty"`                                   // 关键词
	DownloadMode      string          `json:"download_mode,omitempty"`                              // 优先从系统应用商店下载（下载模式），landing_type=APP 子目标为 DOWNLOAD 时有效，子目标为LAUNCH时无需传入
	OpenURL           string          `json:"open_url,omitempty"`                                   // Deeplink直达链接，landing_type = APP 且子目标为 LAUNCH 时有效且必填
	UlinkURL          string          `json:"ulink_url,omitempty"`                                  // ulink直达链接，landing_type = APP 且子目标为LAUNCH 时有效
	Audience          Audience        `form:"audience" json:"audience,omitempty"`                   // 人群定向
	DeliverySetting   DeliverySetting `form:"delivery_setting" json:"delivery_setting,omitempty"`   // 排期、预算与出价
	TrackUrlSetting   TrackUrlSetting `form:"track_url_setting" json:"track_url_setting,omitempty"` // 监测链接
}

// Encode implement PostRequest interface
func (r UpdateProjectReq) Encode() []byte {
	return util.JSONMarshal(r)
}

// UpdateProjectRes 更新项目返回参数
type UpdateProjectRes struct {
	//common.OceanEngineApiCommonResultParam
	model.BaseResponse
	Data *UpdateProjectResData `json:"data"` // json返回值
}

type UpdateProjectResData struct {
	ProjectId int64 `json:"project_id" binding:"required"` // 项目id
	ErrorList []struct {
		ObjectType   string `json:"object_type"`   // 错误对象类型，返回值：BUDGET 预算、PROJECT_SETTING 项目设置
		ErrorCode    int    `json:"error_code"`    // 错误码
		ErrorMessage string `json:"error_message"` // 错误信息
	} `json:"error_list"` // 错误list，广告为分块更新，存在部分内容更新失败，部分内容更新成功的的情况 若广告更新成功，则返回为空数组 若更新失败，则返回错误的部分及原因
}

// SearchBidRatioKey 搜索快投
type SearchBidRatioKey struct {
	// 出价系数，默认系数为1，出价系数可通过【获取快投推荐出价系数】查询，小数点后最多两位,取值范围 [1,2]
	// 当符合以下所有条件时填写有效
	// 1. bid_type != NO_BID && pricing = PRICING_OCPM
	// 2. deep_bid_type = DEEP_BID_DEFAULT 无深度优化方式 /BID_PER_ACTION 每次付费
	SearchBidRatio float64 `form:"search_bid_ratio" json:"search_bid_ratio,omitempty"` //出价系数，默认系数为1，小数点后最多两位,取值范围 [1,2]
	AudienceExtend string  `form:"audience_extend" json:"audience_extend,omitempty"`   //定向拓展，允许值：ON:开启，OFF:关闭，默认值：ON开启
}

// Keyword 关键词
type Keyword struct {
	Word      string `form:"word" json:"word,omitempty" binding:"required"`             // 关键词
	BidType   string `form:"bid_type" json:"bid_type,omitempty" binding:"required"`     // 出价类型, 允许值: FEED_TO_SEARCH 搜索快投（默认值）
	MatchType string `form:"match_type" json:"match_type,omitempty" binding:"required"` // 匹配类型, 允许值: PHRASE短语匹配，EXTENSIVE广泛匹配，PRECISION精准匹配
}

// RelatedProduct 关联产品投放
type RelatedProduct struct {
	ProductSetting    string `form:"product_setting" json:"product_setting,omitempty"`         // 商品库设置，允许值：SINGLE 启用SDPA、NO_MAP不启用（默认值）
	ProductPlatformId int    `form:"product_platform_id" json:"product_platform_id,omitempty"` // 商品库ID，当启用商品库时必填，可通过【商品广告-查询商品库】查询，创建后不可修改
	ProductId         string `form:"product_id" json:"product_id,omitempty"`                   // 产品ID，当启用商品库时必填，可通过【商品广告-获取商品列表】 查询，创建后不可修改
}

// AppLaunchContent 投放内容与目标
type AppLaunchContent struct {
	DownloadType       string `form:"download_type" json:"download_type,omitempty"`               // 下载方式，landing_type=APP 子目标为 DOWNLOAD 时有效，子目标为LAUNCH时无需传入
	DownloadURL        string `form:"download_url" json:"download_url,omitempty"`                 // 下载链接，landing_type=APP 子目标为 DOWNLOAD 或者LAUNCH 时有效且必填
	DownloadMode       string `form:"download_mode" json:"download_mode,omitempty"`               // 优先从系统应用商店下载（下载模式），landing_type=APP 子目标为 DOWNLOAD 时有效，子目标为LAUNCH时无需传入
	QuickAppID         int    `form:"quick_app_id" json:"quick_app_id,omitempty"`                 // 快应用资产id ，仅支持已通过审核的快应用资产
	LaunchType         string `form:"launch_type" json:"launch_type,omitempty"`                   // 调起方式， landing_type = APP 且子目标为LAUNCH有效
	OpenURL            string `form:"open_url" json:"open_url,omitempty"`                         // Deeplink直达链接，landing_type = APP 且子目标为 LAUNCH 时有效且必填
	UlinkURL           string `form:"ulink_url" json:"ulink_url,omitempty"`                       // ulink直达链接，landing_type = APP 且子目标为LAUNCH 时有效
	SubscribeURL       string `form:"subscribe_url" json:"subscribe_url,omitempty"`               // 预约下载链接，landing_type = APP 且子目标为LAUNCH 时有效且必填
	AssetType          string `form:"asset_type" json:"asset_type,omitempty"`                     // 资产类型，landing_type = LINK 或SHOP时有效且必填 允许值： ORANGE 橙子落地页、THIRDPARTY 自研落地页
	MicroPromotionType string `form:"micro_promotion_type" json:"micro_promotion_type,omitempty"` // 小程序类型，landing_type = MICRO_GAME 时有效且必填 允许值： WECHAT_GAME 微信小游戏、WECHAT_APP微信小程序、BYTE_GAME字节小游戏、BYTE_APP字节小程序
	MiniProgramID      string `form:"mini_program_id" json:"mini_program_id,omitempty"`           // 小程序ID，landing_type = APP 且子目标为LAUNCH 时有效且必填
	MiniProgramURL     string `form:"mini_program_url" json:"mini_program_url,omitempty"`         // 小程序落地页链接，landing_type = APP 且子目标为LAUNCH 时有效且必填
}

// OptimizeGoal 优化目标
type OptimizeGoal struct {
	AssetIds           []int64 `json:"asset_ids,omitempty"`            // 事件管理资产 id，list长度上限为 1
	ConvertId          int64   `json:"convert_id,omitempty"`           // 转化跟踪id
	ExternalAction     string  `json:"external_action,omitempty"`      // 优化目标
	DeepExternalAction string  `json:"deep_external_action,omitempty"` // 深度转化目标
}

// DeliveryRange 广告版位
type DeliveryRange struct {
	InventoryCatalog string `json:"inventory_catalog,omitempty" binding:"required"` // 必填，广告位大类，允许值：MANUAL 首选媒体，UNIVERSAL_SMART 通投智选
	// 当 delivery_mode = PROCEDURAL 时仅支持通投智选，当 marketing_goal = LIVE 时，仅支持 MANUAL 首选媒体
	InventoryType []string `json:"inventory_type,omitempty"` // 条件必填，广告投放位置（首选媒体），inventory_catalog = MANUAL 有效且必填，允许值：
	// INVENTORY_FEED 今日头条
	// INVENTORY_VIDEO_FEED 西瓜视频
	// INVENTORY_AWEME_FEED 抖音短视频
	// INVENTORY_TOMATO_NOVEL 番茄小说
	// INVENTORY_UNION_SLOT 穿山甲
	// UNION_BOUTIQUE_GAME ohayoo 精品游戏
	// 当 marketing_goal = LIVE 时，仅支持 INVENTORY_AWEME_FEED 抖音短视频，INVENTORY_UNION_SLOT 穿山甲
	UnionVideoType string `json:"union_video_type,omitempty"` // 条件必填，投放形式（穿山甲视频创意类型），
	// inventory_catalog = MANUAL && inventory_type 仅有 INVENTORY_UNION_SLOT 时有效且必填
	// 当 marketing_goal = LIVE 时，不支持该字段
	// 允许值：ORIGINAL_VIDEO 原生视频、REWARDED_VIDEO 激励视频、SPLASH_VIDEO 开屏视频
}

// Audience 人群定向
type Audience struct {
	AudiencePackageID int        `json:"audience_package_id,omitempty"` // 定向包ID，由【工具-定向包管理-获取定向包】获取
	District          string     `json:"district,omitempty"`            // 地域类型，允许值： CITY 省市、 COUNTY 区县、BUSINESS_DISTRICT商圈、REGION 行政区域、OVERSEA 海外区域、NONE 不限
	Geolocation       []struct { // 从地图添加(地图位置)
		Name   string  `json:"name,omitempty"`   // 地点名称
		Long   float64 `json:"long,omitempty"`   // 经度
		Lat    float64 `json:"lat,omitempty"`    // 纬度
		Radius int     `json:"radius,omitempty"` // 半径，单位为m，允许值为：3000～30000，默认值为6000
	} `json:"geolocation,omitempty"`
	RegionVersion             string   `json:"region_version,omitempty"`               // 行政区域版本号，district =REGION/OVERSEA时必填，通过【获取行政信息】接口获取
	City                      []int    `json:"city,omitempty"`                         // 城市或区县编号，使用省市示例：{"district": "CITY","city": [12]}，使用区县示例：{"district": "COUNTY","city": [130102]}
	LocationType              string   `json:"location_type,omitempty"`                // 位置类型，允许值：CURRENT正在该地区的用户，HOME居住在该地区的用户，TRAVEL到该地区旅行的用户，ALL该地区内的所有用户
	Gender                    string   `json:"gender,omitempty"`                       // 性别
	Age                       []string `json:"age,omitempty"`                          // 年龄
	RetargetingTagsInclude    []int    `json:"retargeting_tags_include,omitempty"`     // 定向人群包列表
	RetargetingTagsExclude    []int    `json:"retargeting_tags_exclude,omitempty"`     // 排除人群包列表
	InterestActionMode        string   `json:"interest_action_mode,omitempty"`         // 行为兴趣
	ActionScene               []string `json:"action_scene,omitempty"`                 // 行为场景
	ActionDays                int      `json:"action_days,omitempty"`                  // 用户发生行为天数
	ActionCategories          []int    `json:"action_categories,omitempty"`            // 行为类目词，interest_action_mode = CUSTOM时有效
	ActionWords               []int    `json:"action_words,omitempty"`                 // 行为关键词，interest_action_mode = CUSTOM时有效
	InterestCategories        []int    `json:"interest_categories,omitempty"`          // 兴趣类目词，interest_action_mode = CUSTOM时有效兴趣类目词
	InterestWords             []int    `json:"interest_words,omitempty"`               // 兴趣关键词，interest_action_mode = CUSTOM时有效，传入具体的词id
	AwemeFanBehaviors         []string `json:"aweme_fan_behaviors,omitempty"`          // 抖音达人互动用户行为类型, 详见【附录-抖音用户行为类型】
	AwemeFanTimeScope         string   `json:"aweme_fan_time_scope,omitempty"`         // 抖音达人互动行为时间范围，默认时间范围为 FIFTEEN_DAYS 15天
	AwemeFanCategories        []int    `json:"aweme_fan_categories,omitempty"`         // 抖音达人分类ID列表
	AwemeFanAccounts          []int    `json:"aweme_fan_accounts,omitempty"`           // 抖音达人ID列表
	SuperiorPopularityType    string   `json:"superior_popularity_type,omitempty"`     // 媒体定向
	FlowPackage               []int    `json:"flow_package,omitempty"`                 // 定向逻辑，可通过【工具-穿山甲流量包-获取穿山甲流量包】
	ExcludeFlowPackage        []int    `json:"exclude_flow_package,omitempty"`         // 排除定向逻辑，可通过【工具-穿山甲流量包-获取穿山甲流量包】
	Platform                  []string `json:"platform,omitempty"`                     // 投放平台列表，不传值为全选
	AndroidOsv                string   `json:"android_osv,omitempty"`                  // 最低安卓版本，当app_type为APP_ANDROID时有效
	IosOsv                    string   `json:"ios_osv,omitempty"`                      // 最低IOS版本，当app_type为APP_IOS 时有效
	DeviceType                []string `json:"device_type,omitempty"`                  // 设备类型
	Ac                        []string `json:"ac,omitempty"`                           // 网络类型
	Carrier                   []string `json:"carrier,omitempty"`                      // 运营商
	HideIfExists              string   `json:"hide_if_exists,omitempty"`               // 过滤已安装，仅对Android链接生效，允许值：UNLIMITED、FILTER、TARGETING
	HideIfConverted           string   `json:"hide_if_converted,omitempty"`            // 过滤已转化用户，允许值：NO_EXCLUDE、PROMOTION、PROJECT、ADVERTISER、APP、CUSTOMER、ORGANIZATION
	ConvertedTimeDuration     string   `json:"converted_time_duration,omitempty"`      // 过滤时间范围，当 hide_if_converted = APP、CUSTOMER、ORGANIZATION 时有效，默认为THREE_MONTH
	FilterAwemeAbnormalActive int      `json:"filter_aweme_abnormal_active,omitempty"` // 过滤高活跃用户，即过滤关注、点赞、评论行为高活跃的用户允许值：ON 过滤 OFF不过滤（默认值）当marketing_goal= LIVE 且inventory_type非仅穿山甲时，支持该字段
	FilterAwemeFansCount      int      `json:"filter_aweme_fans_count,omitempty"`      // 过滤高关注数用户，例如"filter_aweme_fans_count": 1000表示过滤粉丝数在1000以上的用户 允许值：1000、500、200 当marketing_goal= Live 且inventory_type非仅穿山甲时，支持该字段
	FilterOwnAwemeFans        int      `json:"filter_own_aweme_fans,omitempty"`        // 过滤自己的粉丝，允许值：ON 过滤 OFF不过滤（默认值）当marketing_goal= Live 且inventory_type非仅穿山甲时，支持该字段
	DeviceBrand               []string `json:"device_brand,omitempty"`                 // 手机品牌
	LaunchPrice               []int    `json:"launch_price,omitempty"`                 // 手机价格
	AutoExtendTargets         []string `json:"auto_extend_targets,omitempty"`          // 可放开定向，允许值：AGE、REGION、GENDER、CUSTOM_AUDIENCE
}

// DeliverySetting 排期、预算与出价
type DeliverySetting struct {
	ProjectId    int64  `json:"project_id,omitempty"`
	ScheduleType string `json:"schedule_type,omitempty"` // 必填，投放时间类型，允许值：SCHEDULE_FROM_NOW 从今天起长期投放、SCHEDULE_START_END 设置开始和结束日期
	StartTime    string `json:"start_time,omitempty"`    // 条件必填，投放起始时间，当schedule_type = SCHEDULE_START_END时必填，如：2017-01-01 精确到天，广告投放起始时间不允许修改
	EndTime      string `json:"end_time,omitempty"`      // 条件必填，投放结束时间，当schedule_type = SCHEDULE_START_END时必填，如：2017-01-01 精确到天
	ScheduleTime string `json:"schedule_time,omitempty"` // 投放时段，默认全时段投放，格式是48*7位字符串，且都是0或1。也就是以半个小时为最小粒度，周一至周日每天分为48个区段，0为不投放，1为投放，不传、全传0、全传1均代表全时段投放。例如：000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000，则投放时段为周一到周日的11:30~13:30

	DeepBidType string  `json:"deep_bid_type,omitempty"` // 深度优化方式，必填字段，允许值：DEEP_BID_MIN, ROI_COEFFICIENT, BID_PER_ACTION, SOCIAL_ROI
	BidType     string  `json:"bid_type,omitempty"`      // 竞价策略，必填字段，允许值：CUSTOM, NO_BID, CONSERVATIVE, UPPER_CONTROL
	BidSpeed    string  `json:"bid_speed,omitempty"`     // 投放速度，当 bid_type=CUSTOM 时必填
	BudgetMode  string  `json:"budget_mode,omitempty"`   // 项目预算类型，必填字段，允许值：BUDGET_MODE_INFINITE, BUDGET_MODE_DAY
	Budget      float64 `json:"budget,omitempty"`        // 项目预算，当 budget_mode=BUDGET_MODE_DAY 时有效且必填
	Pricing     string  `json:"pricing,omitempty"`       // 计费方式，允许值：PRICING_CPM, PRICING_CPC, PRICING_OCPM, PRICING_OCPC

	CpaBid               float64 `json:"cpa_bid,omitempty"`                // 目标转化出价/预期成本，取值范围：0.1-10000元，当bid_type=NO_BID时不填写该字段，当delivery_mode=PROCEDURAL时必填，当delivery_mode=MANUAL时填写报错
	RoiGoal              float64 `json:"roi_goal,omitempty"`               // 深度转化ROI系数，范围(0,5]，当delivery_mode=PROCEDURAL且深度优化方式为付费ROI时必填，当delivery_mode=MANUAL时填写报错
	BudgetOptimizeSwitch string  `json:"budget_optimize_switch,omitempty"` // 支持预算择优分配，允许值：ON开启，OFF不开启（默认值），当bid_type=NO_BID且BUDGET=BUDGET_MODE_DAY时，“预算择优分配”必填，当delivery_mode=PROCEDURAL时不支持OFF
}

// TrackUrlSetting 监测链接
type TrackUrlSetting struct {
	TrackUrlType               string   `json:"track_url_type,omitempty"`                 // 监测链接类型
	TrackUrlGroupID            int64    `json:"track_url_group_id,omitempty"`             // 监测链接组ID
	TrackURL                   []string `json:"track_url,omitempty"`                      // 展示监测链接
	ActionTrackURL             []string `json:"action_track_url,omitempty"`               // 点击监测链接
	VideoPlayEffectiveTrackURL []string `json:"video_play_effective_track_url,omitempty"` // 视频有效播放监测链接
	VideoPlayDoneTrackURL      []string `json:"video_play_done_track_url,omitempty"`      // 视频播完监测链接
	VideoPlayFirstTrackURL     []string `json:"video_play_first_track_url,omitempty"`     // 视频开始播放监测链接
	SendType                   string   `json:"send_type,omitempty"`                      // 数据发送方式
}

// BatchUpdateProjectBudgetReq 批量更新项目预算请求数据
type BatchUpdateProjectBudgetReq struct {
	AdvertiserID int64 `json:"advertiser_id" binding:"required"` // 广告账户id
	Data         []struct {
		ProjectID  int64   `json:"project_id"`  // 项目ID
		BudgetMode string  `json:"budget_mode"` // 允许值: BUDGET_MODE_DAY日预算 BUDGET_MODE_INFINITE不限
		Budget     float64 `json:"budget"`      // 项目预算，取值范围: ≥ 300 当budget_mode = BUDGET_MODE_DAY时有效且必填
	} `json:"data"` // 批量更新项目状态，包含项目ID和目标操作，list长度限制1～10
}

// BatchUpdateProjectStatusReq 批量更新项目状态请求数据
type BatchUpdateProjectStatusReq struct {
	AdvertiserID int64 `json:"advertiser_id" binding:"required"` // 广告账户id
	Data         []struct {
		ProjectID int64  `json:"project_id"` // 项目ID
		OptStatus string `json:"opt_status"` // 目标操作，允许值：ENABLE 启用项目、DISABLE 暂停项目
	} `json:"data"` // 批量更新项目状态，包含项目ID和目标操作，list长度限制1～10
}
