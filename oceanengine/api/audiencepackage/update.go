package audiencepackage

import (
	"donson.com.cn/draining/internal/pkg/oceanengine/core"
	"donson.com.cn/draining/internal/pkg/oceanengine/model/audiencepackage"
)

// Update 更新定向包
func Update(clt *core.SDKClient, accessToken string, req *audiencepackage.UpdateRequest) (uint64, error) {
	var resp audiencepackage.UpdateResponse
	err := clt.Post("2/audience_package/update/", req, &resp, accessToken)
	if err != nil {
		return 0, err
	}
	return resp.Data.AudiencePackageID, nil
}
