package coupon

import (
	"donson.com.cn/draining/internal/pkg/oceanengine/core"
	"donson.com.cn/draining/internal/pkg/oceanengine/model/clue/coupon"
)

// EmployeeDelete 删除核销员
func EmployeeDelete(clt *core.SDKClient, accessToken string, req *coupon.EmployeeDeleteRequest) error {
	return clt.Post("2/clue/coupon/employee/delete/", req, nil, accessToken)
}
