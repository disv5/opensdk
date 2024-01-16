package common

// OceanEngineApiCommonResultParam 巨量引擎Api返回的通用参数
type OceanEngineApiCommonResultParam struct {
	Code      int    `form:"code" json:"code,omitempty"`             // 业务响应状态码
	Message   string `form:"message" json:"message,omitempty"`       // 提示信息
	MessageCN string `form:"message_cn" json:"message_cn,omitempty"` // 中文提示信息
	RequestId string `form:"request_id" json:"request_id,omitempty"` // 请求日志id
}
