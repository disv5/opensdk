package audiencepackage

import (
	"donson.com.cn/draining/internal/pkg/oceanengine/core"
	"donson.com.cn/draining/internal/pkg/oceanengine/model/audiencepackage"
)

// Delete 删除定向包
func Delete(clt *core.SDKClient, accessToken string, req *audiencepackage.DeleteRequest) (uint64, error) {
	var resp audiencepackage.DeleteResponse
	err := clt.Post("2/audience_package/delete/", req, &resp, accessToken)
	if err != nil {
		return 0, err
	}
	return resp.Data.AudiencePackageID, nil
}
