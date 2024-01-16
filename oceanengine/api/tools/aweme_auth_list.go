package tools

import (
	"donson.com.cn/draining/internal/pkg/oceanengine/core"
	"donson.com.cn/draining/internal/pkg/oceanengine/model/tools"
)

// AwemeAuthList 获取抖音授权关系
// 可以获取账户下抖音号授权关系以及授权视频
func AwemeAuthList(clt *core.SDKClient, accessToken string, req *tools.AwemeAuthListRequest) (*tools.AwemeAuthListData, error) {
	var resp tools.AwemeAuthListResponse
	if err := clt.Get("2/tools/aweme_auth_list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
