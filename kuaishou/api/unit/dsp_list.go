package unit

import (
	"github.com/disv5/opensdk/kuaishou/core"
	"github.com/disv5/opensdk/kuaishou/model/unit"
)

// DspList 获取广告组信息
func DspList(clt *core.SDKClient, accessToken string, req *unit.DspListRequest) (*unit.ListResponse, error) {
	var resp unit.ListResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
