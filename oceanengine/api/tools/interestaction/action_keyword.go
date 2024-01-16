package interestaction

import (
	"github.com/disv5/opensdk/oceanengine/core"
	"github.com/disv5/opensdk/oceanengine/model/tools/interestaction"
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
