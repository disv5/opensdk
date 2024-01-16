package taskraise

import (
	"donson.com.cn/draining/internal/pkg/oceanengine/core"
	"donson.com.cn/draining/internal/pkg/oceanengine/model/tools/taskraise"
)

// Get 查询优选起量任务
func Get(clt *core.SDKClient, accessToken string, req *taskraise.GetRequest) (*taskraise.GetResponseData, error) {
	var resp taskraise.GetResponse
	if err := clt.Get("2/tools/task_raise/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
