package unit

// ListResponse 获取广告组信息 API Response
type ListResponse struct {
	// TotalCount 数据总数
	TotalCount int `json:"total_count,omitempty"`
	// Details 返回值详情
	Details []Unit `json:"details,omitempty"`
}
