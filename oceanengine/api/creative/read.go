package creative

import (
	"github.com/disv5/opensdk/oceanengine/core"
	"github.com/disv5/opensdk/oceanengine/model/creative"
)

// Read 创意详细信息
func Read(clt *core.SDKClient, accessToken string, req *creative.ReadRequest) (*creative.CreativeDetail, error) {
	var resp creative.ReadResponse
	err := clt.Get("2/creative/read_v2/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
