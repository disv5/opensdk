package enum

// AdConvertSource 转化工具来源
type AdConvertSource string

const (
	// AD_CONVERT_SOURCE_TYPE_INTERNAL 平台提供的转化工具
	AD_CONVERT_SOURCE_TYPE_INTERNAL AdConvertSource = "AD_CONVERT_SOURCE_TYPE_INTERNAL"
	// AD_CONVERT_SOURCE_TYPE_JS  代码检测转化
	AD_CONVERT_SOURCE_TYPE_JS AdConvertSource = "AD_CONVERT_SOURCE_TYPE_JS"
	// AD_CONVERT_SOURCE_TYPE_XPATH 路径转化
	AD_CONVERT_SOURCE_TYPE_XPATH AdConvertSource = "AD_CONVERT_SOURCE_TYPE_XPATH"
	// AD_CONVERT_SOURCE_TYPE_APP_DOWNLOAD 应用下载API
	AD_CONVERT_SOURCE_TYPE_APP_DOWNLOAD AdConvertSource = "AD_CONVERT_SOURCE_TYPE_APP_DOWNLOAD"
	// AD_CONVERT_SOURCE_TYPE_API 落地页API（废弃，不能新建，历史数据保留）
	AD_CONVERT_SOURCE_TYPE_API AdConvertSource = "AD_CONVERT_SOURCE_TYPE_API"
	// AD_CONVERT_SOURCE_TYPE_H5_API 落地页API（H5）
	AD_CONVERT_SOURCE_TYPE_H5_API AdConvertSource = "AD_CONVERT_SOURCE_TYPE_H5_API"
	// AD_CONVERT_SOURCE_TYPE_SDK 应用下载SDK
	AD_CONVERT_SOURCE_TYPE_SDK AdConvertSource = "AD_CONVERT_SOURCE_TYPE_SDK"
	// AD_CONVERT_SOURCE_TYPE_OPEN_URL 应用直达API（应用直达链接）
	AD_CONVERT_SOURCE_TYPE_OPEN_URL AdConvertSource = "AD_CONVERT_SOURCE_TYPE_OPEN_URL"
	// AD_CONVERT_SOURCE_TYPE_NORMAL_APP_DOWNLOAD 普通应用下载
	AD_CONVERT_SOURCE_TYPE_NORMAL_APP_DOWNLOAD AdConvertSource = "AD_CONVERT_SOURCE_TYPE_NORMAL_APP_DOWNLOAD"
	// AD_CONVERT_SOURCE_TYPE_CONFIG 无转化跟踪
	AD_CONVERT_SOURCE_TYPE_CONFIG AdConvertSource = "AD_CONVERT_SOURCE_TYPE_CONFIG"
	// AD_CONVERT_SOURCE_TYPE_CPS_GAME  内部游戏cps应用下载
	AD_CONVERT_SOURCE_TYPE_CPS_GAME AdConvertSource = "AD_CONVERT_SOURCE_TYPE_CPS_GAME"
	// AD_CONVERT_SOURCE_TYPE_APP_API_TEMAI 落地页API(支持特卖)
	AD_CONVERT_SOURCE_TYPE_APP_API_TEMAI AdConvertSource = "AD_CONVERT_SOURCE_TYPE_APP_API_TEMAI"
)
