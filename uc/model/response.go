package model

import "encoding/json"

// Response api response interface
type Response interface {
	IsError() bool
	Error() string
}

// BaseResponse shared response data struct
type BaseResponse struct {
	Header ResponseHeader  `json:"header"`
	Body   json.RawMessage `json:"body"`
}

type ResponseHeader struct {
	Desc      string `json:"desc"`
	Status    uint   `json:"status"`
	Quota     int    `json:"quota"`
	LeftQuota int    `json:"left_quota"`
}

// IsError detect if the response is an error
func (r BaseResponse) IsError() bool {
	return r.Header.Status != 0
}

// Error implement error interface
func (r BaseResponse) Error() string {
	return r.Header.Desc
}
