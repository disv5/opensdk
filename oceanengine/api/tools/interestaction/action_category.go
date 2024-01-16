package interestaction

import (
	"donson.com.cn/draining/internal/pkg/oceanengine/core"
	"donson.com.cn/draining/internal/pkg/oceanengine/model/tools/interestaction"
)

// ActionCategory 行为类目查询
func ActionCategory(clt *core.SDKClient, accessToken string, req *interestaction.ActionCategoryRequest) ([]interestaction.Object, error) {
	var resp interestaction.ActionCategoryResponse
	err := clt.Get("2/tools/interest_action/action/category", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
