package campaign

// CreateResponse 创建广告计划 API Response
type CreateResponse struct {
	// CampaignID 广告计划ID
	CampaignID uint64 `json:"campaign_id,omitempty"`
}
