package report

import (
	"donson.com.cn/draining/internal/pkg/uc/model"
	"encoding/json"
)

type CampaignRequest struct {
	Header model.HeaderTarget `json:"header"`
	Body   CampaignBody       `json:"body"`
}

type CampaignBody struct {
	StartDate         string             `json:"startDate"`
	EndDate           string             `json:"endDate"`
	UnitOfTime        int                `json:"unitOfTime"`
	CampaignIds       []int64            `json:"campaignIds,omitempty"`
	CampaignCondition *CampaignCondition `json:"campaignCondition,omitempty"`
	PerformanceData   []string           `json:"performanceData"`
}

type CampaignCondition struct {
	OptTarget    int `json:"optTarget"`
	Delivery     int `json:"delivery"`
	DeliveryMode int `json:"deliveryMode"`
	ChargeType   int `json:"chargeType"`
}

// Url implement GetRequest interface
func (r CampaignRequest) Url() string {
	return "api/report/campaign"
}

// Encode implement GetRequest interface
func (r CampaignRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
