package taskraise

import (
	"donson.com.cn/draining/internal/pkg/oceanengine/core"
	"donson.com.cn/draining/internal/pkg/oceanengine/model/tools/taskraise"
)

// Create 账户优选起量
func Create(clt *core.SDKClient, accessToken string, req *taskraise.CreateRequest) (uint64, error) {
	var resp taskraise.CreateResponse
	if err := clt.Post("2/tools/task_raise/create/", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.ReportID, nil
}
