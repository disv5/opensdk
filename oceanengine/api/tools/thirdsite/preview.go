package thirdsite

import (
	"github.com/disv5/opensdk/oceanengine/core"
	"github.com/disv5/opensdk/oceanengine/model/tools/thirdsite"
)

// Preview 获取第三方落地页预览地址
// 通过此接口，用户可以获取第三方落地页预览地址。
func Preview(clt *core.SDKClient, accessToken string, req *thirdsite.PreviewRequest) (*thirdsite.PreviewResponseData, error) {
	var resp thirdsite.PreviewResponse
	err := clt.Get("2/tools/third_site/preview/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
