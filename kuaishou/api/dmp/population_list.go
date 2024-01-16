package dmp

import (
	"github.com/disv5/opensdk/kuaishou/core"
	"github.com/disv5/opensdk/kuaishou/model/dmp"
)

// PopulationList 人群列表查询接口
func PopulationList(clt *core.SDKClient, accessToken string, req *dmp.PopulationListRequest) (*dmp.PopulationListResponse, error) {
	var resp *dmp.PopulationListResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
