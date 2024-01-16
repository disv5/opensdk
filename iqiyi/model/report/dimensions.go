package report

// Dimensions 维度数据
type Dimensions struct {
	// StatDatetime 数据起始时间，分组条件包含 STAT_GROUP_BY_FIELD_STAT_TIME 时返回，格式为：yyyy-MM-dd HH:mm:ss
	Date string `json:"date,omitempty"`
	Hour int `json:"hour"`
	// 广告主
	AdvertiserID  uint64 `json:"advertiserId,omitempty"`
	AdvertiserName  string `json:"advertiserName,omitempty"`
	// 广告系列
	OrderGroupId   uint64 `json:"orderGroupId,omitempty"`
	OrderGroupName string `json:"orderGroupName,omitempty"`
	// 广告
	OrderPlanId   uint64 `json:"orderPlanId,omitempty"`
	OrderPlanName string `json:"orderPlanName,omitempty"`
	// 创意
	CreativeId   uint64 `json:"creativeId,omitempty"`
	CreativeName string `json:"creativeName,omitempty"`
}
