package target

import (
	"github.com/disv5/opensdk/kuaishou/core"
	"github.com/disv5/opensdk/kuaishou/model/target"
)

// TemplateUpdate 修改定向模板
func TemplateUpdate(clt *core.SDKClient, accessToken string, req *target.TemplateUpdateRequest) (*target.Template, error) {
	var resp target.Template
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
