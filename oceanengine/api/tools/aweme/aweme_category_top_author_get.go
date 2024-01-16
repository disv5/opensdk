package aweme

import (
	"donson.com.cn/draining/internal/pkg/oceanengine/core"
	"donson.com.cn/draining/internal/pkg/oceanengine/model/tools/aweme"
)

// AwemeCategoryTopAuthorGet 查询抖音类目下的推荐达人
// 查询抖音类目下的推荐达人，根据类目id查询抖音推荐达人。
func AwemeCategoryTopAuthorGet(clt *core.SDKClient, accessToken string, req *aweme.AwemeCategoryTopAuthorGetRequest) ([]aweme.Author, error) {
	var resp aweme.AwemeCategoryTopAuthorGetResponse
	err := clt.Get("2/tools/aweme_category_top_author/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
