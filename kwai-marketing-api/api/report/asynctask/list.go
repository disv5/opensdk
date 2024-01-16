package asynctask

import (
	"github.com/disv5/opensdk/kwai-marketing-api/core"
	"github.com/disv5/opensdk/kwai-marketing-api/model/report/asynctask"
)

func List(clt *core.SDKClient, accessToken string, req *asynctask.ListRequest) (*asynctask.ListResponse, error) {
	var resp asynctask.ListResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
