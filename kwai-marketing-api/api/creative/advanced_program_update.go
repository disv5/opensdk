package creative

import (
	"github.com/disv5/opensdk/kwai-marketing-api/core"
	"github.com/disv5/opensdk/kwai-marketing-api/model/creative"
)

// AdvancedProgramUpdate 修改程序化2.0创意
func AdvancedProgramUpdate(clt *core.SDKClient, accessToken string, req *creative.AdvancedProgramUpdateRequest) error {
	return clt.Post(accessToken, req, nil)
}
