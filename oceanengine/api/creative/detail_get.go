package creative

import (
	"github.com/disv5/opensdk/oceanengine/core"
	"github.com/disv5/opensdk/oceanengine/model/creative"
)

// DetailGet 创意详细信息(新)
func DetailGet(clt *core.SDKClient, accessToken string, req *creative.DetailGetRequest) (*creative.CreativeDetailV2, error) {
	var resp creative.DetailGetResponse
	err := clt.Get("v3.0/creative/detail/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
