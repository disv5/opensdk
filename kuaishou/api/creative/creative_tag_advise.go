package creative

import (
	"github.com/disv5/opensdk/kuaishou/core"
	"github.com/disv5/opensdk/kuaishou/model/creative"
)

// CreativeTagAdvise 创意标签填写建议
func CreativeTagAdvise(clt *core.SDKClient, accessToken string, req *creative.CreativeTagAdviseRequest) (*creative.CreativeTagAdviseResponse, error) {
	var resp creative.CreativeTagAdviseResponse
	err := clt.Get(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
