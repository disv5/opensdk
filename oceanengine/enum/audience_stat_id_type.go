package enum

// AudienceStatIDType 受众分析查询ID类型
type AudienceStatIDType string

const (
	// AUDIENCE_STAT_ID_TYPE_ADVERTISER 广告主ID
	AUDIENCE_STAT_ID_TYPE_ADVERTISER AudienceStatIDType = "AUDIENCE_STAT_ID_TYPE_ADVERTISER"
	// AUDIENCE_STAT_ID_TYPE_CAMPAIGN 广告组ID
	AUDIENCE_STAT_ID_TYPE_CAMPAIGN AudienceStatIDType = "AUDIENCE_STAT_ID_TYPE_CAMPAIGN"
	// AUDIENCE_STAT_ID_TYPE_AD 广告计划ID
	AUDIENCE_STAT_ID_TYPE_AD AudienceStatIDType = "AUDIENCE_STAT_ID_TYPE_AD"
	// AUDIENCE_STAT_ID_TYPE_CREATIVE 广告创意ID
	AUDIENCE_STAT_ID_TYPE_CREATIVE AudienceStatIDType = "AUDIENCE_STAT_ID_TYPE_CREATIVE"
)