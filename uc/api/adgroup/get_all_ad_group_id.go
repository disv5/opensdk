package adgroup

import (
	"github.com/disv5/opensdk/uc/core"
	"github.com/disv5/opensdk/uc/model/adgroup"
)

func GetAllAdGroupId(clt *core.SDKClient, req *adgroup.GetAllAdGroupIdRequest) (*adgroup.GetAllAdGroupIdResponse, error) {
	var resp adgroup.GetAllAdGroupIdResponse
	err := clt.Post(req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
