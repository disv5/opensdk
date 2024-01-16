package creative

import (
	"github.com/disv5/opensdk/oceanengine/util"
)

// UpdateRequest 修改创意信息 API Request
type UpdateRequest struct {
	CreativeDetail
}

// Encode implement PostRequest interface
func (r UpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
