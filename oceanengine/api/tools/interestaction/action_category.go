package interestaction

import (
	"github.com/disv5/opensdk/oceanengine/core"
	"github.com/disv5/opensdk/oceanengine/model/tools/interestaction"
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
