package unit

import (
	"github.com/disv5/opensdk/kwai-marketing-api/core"
	"github.com/disv5/opensdk/kwai-marketing-api/model/unit"
)

// DspUpdate 修改广告组
func DspUpdate(clt *core.SDKClient, accessToken string, req *unit.DspUpdateRequest) (uint64, error) {
	var resp unit.UpdateResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return 0, err
	}
	return resp.UnitID, nil
}
