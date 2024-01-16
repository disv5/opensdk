package coupon

import (
	"github.com/disv5/opensdk/oceanengine/core"
	"github.com/disv5/opensdk/oceanengine/model/clue/coupon"
)

// EmployeeDelete 删除核销员
func EmployeeDelete(clt *core.SDKClient, accessToken string, req *coupon.EmployeeDeleteRequest) error {
	return clt.Post("2/clue/coupon/employee/delete/", req, nil, accessToken)
}
