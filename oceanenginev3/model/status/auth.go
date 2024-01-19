package auth

import (
	"donson.com.cn/draining/internal/pkg/oceanengine/model"
	"donson.com.cn/draining/internal/pkg/oceanengine/util"
	"net/url"
	"strconv"
)

// AuthGetRequest
type AuthGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID int64 `json:"advertiser_id,omitempty"`
}

// Encode implement GetRequest interface
func (r AuthGetRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatInt(r.AdvertiserID, 10))
	return values.Encode()
}

// AuthGetResponse
type AuthGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *AuthGetResponseData `json:"data,omitempty"`
}

type AuthGetResponseData struct {
	Enabled bool `json:"enabled"`
}

// AuthOprRequest
type AuthOprRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID int64 `json:"advertiser_id,omitempty"`
}

// Encode implement GetRequest interface
func (r AuthOprRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// AuthRes 返回的数据
type AuthRes struct {
	model.BaseResponse
	//common.OceanEngineApiCommonResultParam
	Data *AuthResData `json:"data"` // json返回值
}

type AuthResData struct {
}
