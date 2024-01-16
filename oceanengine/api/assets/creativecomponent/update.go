package creativecomponent

import (
	"donson.com.cn/draining/internal/pkg/oceanengine/core"
	"donson.com.cn/draining/internal/pkg/oceanengine/model/assets/creativecomponent"
)

// Update 更新组件 https://open.oceanengine.com/doc/index.html?key=ad&type=api&id=1696710673083407
func Update(clt *core.SDKClient, accessToken string, req *creativecomponent.UpdateRequest) (*creativecomponent.CreateResponseData, error) {
	var resp creativecomponent.CreateResponse

	if err := clt.Post("2/assets/creative_component/update/", req, &resp, accessToken); err != nil {
		return nil, err
	}

	return resp.Data, nil
}
