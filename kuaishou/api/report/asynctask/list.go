package asynctask

import (
	"github.com/disv5/opensdk/kuaishou/core"
	"github.com/disv5/opensdk/kuaishou/model/report/asynctask"
)

func List(clt *core.SDKClient, accessToken string, req *asynctask.ListRequest) (*asynctask.ListResponse, error) {
	var resp asynctask.ListResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
