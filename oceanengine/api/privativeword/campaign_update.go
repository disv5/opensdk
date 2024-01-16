package privativeword

import (
	"donson.com.cn/draining/internal/pkg/oceanengine/core"
	"donson.com.cn/draining/internal/pkg/oceanengine/model/privativeword"
)

// CampaignUpdate 设置组否定词
func CampaignUpdate(clt *core.SDKClient, accessToken string, req *privativeword.AdUpdateRequest) (uint64, error) {
	var resp privativeword.AdUpdateResponse
	err := clt.Post("2/privative_word/campaign/update", req, &resp, accessToken)
	if err != nil {
		return 0, err
	}
	return resp.Data.CampaignID, nil
}
