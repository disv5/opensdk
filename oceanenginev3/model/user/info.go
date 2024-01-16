package user

import (
	"github.com/disv5/opensdk/oceanengine/model"
)

// GetUserInfoRes 获取授权User信息
type GetUserInfoRes struct {
	model.BaseResponse
	Data *UserItemData `json:"data"` // json返回值
}

// UserItemData
type UserItemData struct {
	Id          uint64 `json:"id"`
	Email       string `json:"email"`
	DisplayName string `json:"display_name"`
}
