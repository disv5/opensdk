package coupon

import (
	"github.com/disv5/opensdk/oceanengine/core"
	"github.com/disv5/opensdk/oceanengine/model/clue/coupon"
)

// Create 创建卡券
// 创建青鸟线索通卡券接口
func Create(clt *core.SDKClient, accessToken string, req *coupon.CreateRequest) (uint64, error) {
	var resp coupon.CreateResponse
	if err := clt.Post("2/clue/coupon/create/", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.ActivityID, nil
}
