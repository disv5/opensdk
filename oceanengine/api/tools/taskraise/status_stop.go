package taskraise

import (
	"donson.com.cn/draining/internal/pkg/oceanengine/core"
	"donson.com.cn/draining/internal/pkg/oceanengine/model/tools/taskraise"
)

// StatusStop 关闭优选起量任务
func StatusStop(clt *core.SDKClient, accessToken string, req *taskraise.StatusStopRequest) error {
	return clt.Post("2/tools/task_raise/status/stop/", req, nil, accessToken)
}
