package tool

import (
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/core"
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/model/tool"
)

// TargetingTagsList 获取可选的定向标签
func TargetingTagsList(clt *core.SDKClient, accessToken string, req *tool.TargetingTagsListRequest) (*tool.TargetingTag, error) {
	var resp tool.TargetingTag
	err := clt.Get(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
