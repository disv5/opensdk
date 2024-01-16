package site

import (
	"donson.com.cn/draining/internal/pkg/oceanengine/core"
	"donson.com.cn/draining/internal/pkg/oceanengine/model/tools/site"
)

// Read 获取橙子建站站点详细信息
// 通过此接口，用户可以获取站点的详细信息，包括新建或更新时传递的全量数据。
func Read(clt *core.SDKClient, accessToken string, req *site.ReadRequest) (*site.SiteDetail, error) {
	var resp site.ReadResponse
	if err := clt.Get("2/tools/site/read/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
