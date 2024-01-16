package customaudience

import (
	"donson.com.cn/draining/internal/pkg/oceanengine/core"
	"donson.com.cn/draining/internal/pkg/oceanengine/model"
	"donson.com.cn/draining/internal/pkg/oceanengine/model/dmp/customaudience"
)

// Delete 删除人群包
// 通过此接口可做人群包删除操作。已经在计划中使用的人群包不能被删除，只有该计划被删除后，人群包才可以删除。
func Delete(clt *core.SDKClient, accessToken string, req *customaudience.DeleteRequest) error {
	var resp model.BaseResponse
	return clt.Post("2/dmp/custom_audience/delete/", req, &resp, accessToken)
}
