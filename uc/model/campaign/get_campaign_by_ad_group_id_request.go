package campaign

import (
	"donson.com.cn/draining/internal/pkg/uc/model"
	"encoding/json"
)

type GetCampaignByAdGroupIdRequest struct {
	Header model.HeaderTarget         `json:"header"`
	Body   GetCampaignByAdGroupIdBody `json:"body"`
}

type GetCampaignByAdGroupIdBody struct {
	AdGroupIds []int64 `json:"adGroupIds"`
}

// Url implement GetRequest interface
func (r GetCampaignByAdGroupIdRequest) Url() string {
	return "api/campaign/getCampaignByAdGroupId"
}

// Encode implement GetRequest interface
func (r GetCampaignByAdGroupIdRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
