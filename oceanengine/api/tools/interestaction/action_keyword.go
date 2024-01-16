package interestaction

import (
	"donson.com.cn/draining/internal/pkg/oceanengine/core"
	"donson.com.cn/draining/internal/pkg/oceanengine/model/tools/interestaction"
)

// ActionKeyword 行为关键词查询
func ActionKeyword(clt *core.SDKClient, accessToken string, req *interestaction.ActionKeywordRequest) ([]interestaction.Object, error) {
	var resp interestaction.ActionKeywordResponse
	err := clt.Get("2/tools/interest_action/action/keyword", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
