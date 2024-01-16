package qlorderplan

import (
	"donson.com.cn/draining/internal/pkg/iqiyi/core"
	"donson.com.cn/draining/internal/pkg/iqiyi/model/qlorderplan"
)

// List 获取广告计划信息
func List(clt *core.SDKClient, accessToken string, req *qlorderplan.ListRequest) (*qlorderplan.ListResponse, error) {
	var resp qlorderplan.ListResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
