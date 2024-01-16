package conversion

import (
	"donson.com.cn/draining/internal/pkg/oceanengine/core"
	"donson.com.cn/draining/internal/pkg/oceanengine/model/conversion"
)

// Attribution 电话转化回传
func Attribution(clt *core.SDKClient, req *conversion.Request) error {
	var resp conversion.Response
	return clt.AnalyticsV1Post("attribution", req, &resp)
}
