package conversion

import (
	"github.com/disv5/opensdk/oceanengine/core"
	"github.com/disv5/opensdk/oceanengine/model/conversion"
)

// Attribution 电话转化回传
func Attribution(clt *core.SDKClient, req *conversion.Request) error {
	var resp conversion.Response
	return clt.AnalyticsV1Post("attribution", req, &resp)
}
