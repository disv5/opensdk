package unit

import "encoding/json"

// DspMonitorListRequest 批量获取监测链接接口
type DspMonitorListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// UnitIDs 广告组ID集
	UnitIDs []uint64 `json:"unit_ids,omitempty"`
	// Page 请求的页码数 默认为1
	Page int `json:"page,omitempty"`
	// PageSize 请求的每页行数; 默认为20
	PageSize int `json:"page_size,omitempty"`
}

// Url implement PostRequest interface
func (r DspMonitorListRequest) Url() string {
	return "gw/dsp/v3/unit/getMonitorUrls"
}

// Encode implement PostRequest interface
func (r DspMonitorListRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
