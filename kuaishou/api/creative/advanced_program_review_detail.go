package creative

import (
	"github.com/disv5/opensdk/kuaishou/core"
	"github.com/disv5/opensdk/kuaishou/model/creative"
)

// AdvancedProgramReviewDetail 获取程序化创意2.0审核信息
func AdvancedProgramReviewDetail(clt *core.SDKClient, accessToken string, req *creative.AdvancedProgramReviewDetailRequest) (*creative.AdvancedProgramReviewDetail, error) {
	var resp creative.AdvancedProgramReviewDetail
	err := clt.Get(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
