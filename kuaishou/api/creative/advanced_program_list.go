package creative

import (
	"github.com/disv5/opensdk/kuaishou/core"
	"github.com/disv5/opensdk/kuaishou/model/creative"
)

// AdvancedProgramList 获取程序化创意2.0信息
func AdvancedProgramList(clt *core.SDKClient, accessToken string, req *creative.AdvancedProgramListRequest) (*creative.AdvancedProgramListResponse, error) {
	var resp creative.AdvancedProgramListResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
