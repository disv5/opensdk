package target

import (
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/core"
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/model/target"
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
