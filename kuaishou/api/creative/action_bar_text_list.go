package creative

import (
	"github.com/disv5/opensdk/kuaishou/core"
	"github.com/disv5/opensdk/kuaishou/model/creative"
)

// ActionBarTextList 获取行动号召按钮
func ActionBarTextList(clt *core.SDKClient, accessToken string, req *creative.ActionBarTextListRequest) ([]string, error) {
	var resp creative.ActionBarTextListResponse
	err := clt.Get(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp.ActionBarText, nil
}
