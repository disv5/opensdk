package creative

import (
	"github.com/disv5/opensdk/kuaishou/core"
	"github.com/disv5/opensdk/kuaishou/model/creative"
)

// Preview 创意体验
func Preview(clt *core.SDKClient, accessToken string, req *creative.PreviewRequest) error {
	return clt.Post(accessToken, req, nil)
}
