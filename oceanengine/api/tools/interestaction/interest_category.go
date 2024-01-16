package interestaction

import (
	"donson.com.cn/draining/internal/pkg/oceanengine/core"
	"donson.com.cn/draining/internal/pkg/oceanengine/model/tools/interestaction"
)

// InterestCategory 兴趣类目查询
func InterestCategory(clt *core.SDKClient, accessToken string, req *interestaction.InterestCategoryRequest) ([]interestaction.Object, error) {
	var resp interestaction.InterestCategoryResponse
	err := clt.Get("2/tools/interest_action/interest/category", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
