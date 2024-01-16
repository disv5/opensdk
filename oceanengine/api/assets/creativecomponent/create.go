package creativecomponent

import (
	"github.com/disv5/opensdk/oceanengine/core"
	"github.com/disv5/opensdk/oceanengine/model/assets/creativecomponent"
)

// Create 创建组件 https://open.oceanengine.com/doc/index.html?key=ad&type=api&id=1696710672391183
func Create(clt *core.SDKClient, accessToken string, req *creativecomponent.CreateRequest) (*creativecomponent.CreateResponseData, error) {
	var resp creativecomponent.CreateResponse

	if err := clt.Post("2/assets/creative_component/create/", req, &resp, accessToken); err != nil {
		return nil, err
	}

	return resp.Data, nil
}
