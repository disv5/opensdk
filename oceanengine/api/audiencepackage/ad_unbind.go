package audiencepackage

import (
	"donson.com.cn/draining/internal/pkg/oceanengine/core"
	"donson.com.cn/draining/internal/pkg/oceanengine/model/audiencepackage"
)

// AdUnbind 定向包解绑
func AdUnbind(clt *core.SDKClient, accessToken string, req *audiencepackage.AdBindRequest) (uint64, error) {
	var resp audiencepackage.AdBindResponse
	err := clt.Post("2/audience_package/ad/unbind/", req, &resp, accessToken)
	if err != nil {
		return 0, err
	}
	return resp.Data.AudiencePackageID, nil
}
