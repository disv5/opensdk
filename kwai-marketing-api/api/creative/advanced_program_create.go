package creative

import (
	"github.com/disv5/opensdk/kwai-marketing-api/core"
	"github.com/disv5/opensdk/kwai-marketing-api/model/creative"
)

// AdvancedProgramCreate 创建程序化2.0创意
func AdvancedProgramCreate(clt *core.SDKClient, accessToken string, req *creative.AdvancedProgramCreateRequest) (uint64, error) {
	var resp creative.AdvancedProgramCreateResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return 0, err
	}
	return resp.UnitID, nil
}
