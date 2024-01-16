package unit

import (
	"github.com/disv5/opensdk/kuaishou/core"
	"github.com/disv5/opensdk/kuaishou/model/unit"
)

// DspBatUpdate 监测链接批量更新接口
func DspBatUpdate(clt *core.SDKClient, accessToken string, req *unit.DspBatUpdateMonitorRequest) ([]*unit.UnitMonitorUrlsResp, error) {
	var resp unit.DspBatUpdateMonitorResp
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp.UnitMonitorUrlsResp, nil
}
