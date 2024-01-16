package unit

import (
	"encoding/json"
)

// DspBatUpdateMonitorRequest 修改监测链接APIRequest
type DspBatUpdateMonitorRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// SchemaUri
	UnitMonitorUrls []UnitMonitorUrls `json:"unit_monitor_urls,omitempty"`
}

type UnitMonitorUrls struct {
	ClickUrl           string `json:"click_url,omitempty"`
	ImpressionUrl      string `json:"impression_url,omitempty"`
	LiveTrackUrl       string `json:"live_track_url,omitempty"`
	ActionbarClickUrl  string `json:"actionbar_click_url,omitempty"`
	ExistValidCreative bool   `json:"exist_valid_creative,omitempty"`
	IsDpa              bool   `json:"is_dpa,omitempty"`
	UnitId             int64  `json:"unit_id,omitempty"`
}

// Url implement PostRequest interface
func (r DspBatUpdateMonitorRequest) Url() string {
	return "gw/dsp/v3/unit/batchUpdateMonitorUrls"
}

// Encode implement PostRequest interface
func (r DspBatUpdateMonitorRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}

// DspBatUpdateMonitorResp 修改监测链接Rep
type DspBatUpdateMonitorResp struct {
	// UnitMonitorUrls UnitMonitorUrls
	UnitMonitorUrlsResp []*UnitMonitorUrlsResp `json:"unit_monitor_urls,omitempty"`
}

type UnitMonitorUrlsResp struct {
	Result              bool        `json:"result"`
	LiveTrackUrl        interface{} `json:"live_track_url"`
	ImpressionUrl       interface{} `json:"impression_url"`
	AdPhotoPlayedT3SUrl interface{} `json:"ad_photo_played_t3s_url"`
	Message             interface{} `json:"message"`
	ClickUrl            string      `json:"click_url"`
	UnitId              int64       `json:"unit_id"`
	ActionbarClickUrl   interface{} `json:"actionbar_click_url"`
}
