package conversion

import (
	"donson.com.cn/draining/internal/pkg/oceanengine/core"
	"donson.com.cn/draining/internal/pkg/oceanengine/model/conversion"
)

// Conversion 新版转化回传
func Conversion(clt *core.SDKClient, req *conversion.Request) error {
	var resp conversion.Response
	return clt.AnalyticsPost("conversion", req, &resp)
}
