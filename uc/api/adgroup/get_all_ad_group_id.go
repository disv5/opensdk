package adgroup

import (
	"donson.com.cn/draining/internal/pkg/uc/core"
	"donson.com.cn/draining/internal/pkg/uc/model/adgroup"
)

func GetAllAdGroupId(clt *core.SDKClient, req *adgroup.GetAllAdGroupIdRequest) (*adgroup.GetAllAdGroupIdResponse, error) {
	var resp adgroup.GetAllAdGroupIdResponse
	err := clt.Post(req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
