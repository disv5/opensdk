package taskraise

import (
	"donson.com.cn/draining/internal/pkg/oceanengine/core"
	"donson.com.cn/draining/internal/pkg/oceanengine/enum"
	"donson.com.cn/draining/internal/pkg/oceanengine/model/tools/taskraise"
)

// OptimizationIDsGet 查询优选起量状态
func OptimizationIDsGet(clt *core.SDKClient, accessToken string, advertiserID uint64) (enum.TaskRaiseStatus, error) {
	var resp taskraise.OptimizationIDsGetResponse
	if err := clt.Get("2/tools/task_raise/optimization_ids/get/", &taskraise.OptimizationIDsGetRequest{AdvertiserID: advertiserID}, &resp, accessToken); err != nil {
		return enum.TaskRaiseStatus_UNKNOWN, err
	}
	return resp.Data.Status, nil
}
