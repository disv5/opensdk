package creative

import (
	"github.com/disv5/opensdk/oceanengine/model"
)

// GetResponse 获取创意列表 API Response
type GetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *GetResponseData `json:"data,omitempty"`
}

// GetResponseData json返回值
type GetResponseData struct {
	// List Creative数组
	List []Creative `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
