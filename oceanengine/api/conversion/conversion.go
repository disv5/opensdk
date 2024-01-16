package conversion

import (
	"github.com/disv5/opensdk/oceanengine/core"
	"github.com/disv5/opensdk/oceanengine/model/conversion"
)

// Conversion 新版转化回传
func Conversion(clt *core.SDKClient, req *conversion.Request) error {
	var resp conversion.Response
	return clt.AnalyticsPost("conversion", req, &resp)
}
