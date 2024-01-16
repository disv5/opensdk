package target

import (
	"github.com/disv5/opensdk/kwai-marketing-api/core"
	"github.com/disv5/opensdk/kwai-marketing-api/model/target"
)

// TemplateList 查询定向模板接口
func TemplateList(clt *core.SDKClient, accessToken string, req *target.TemplateListRequest) (*target.TemplateListResponse, error) {
	var resp target.TemplateListResponse
	err := clt.Get(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
