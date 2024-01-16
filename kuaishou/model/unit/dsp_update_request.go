package unit

import (
	"encoding/json"
)

// DspUpdateRequest 修改广告组APIRequest
type DspUpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// UnitID 广告组 ID
	UnitID uint64 `json:"unit_id,omitempty"`
	// UnitName 广告组名称; 长度为1-100个字符，同一个计划下的广告组名称不能重复
	UnitName string `json:"unit_name,omitempty"`
	// SchemaUri
	SchemaUri string `json:"schema_uri,omitempty"`
	// PutStatus
	PutStatus int `json:"put_status,omitempty"`
}

// Url implement PostRequest interface
func (r DspUpdateRequest) Url() string {
	return "gw/dsp/unit/update"
}

// Encode implement PostRequest interface
func (r DspUpdateRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
