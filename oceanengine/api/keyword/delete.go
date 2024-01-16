package keyword

import (
	"donson.com.cn/draining/internal/pkg/oceanengine/core"
	"donson.com.cn/draining/internal/pkg/oceanengine/model/keyword"
)

// Delete 删除关键词
// 删除指定keyword_id的搜索词，可批量删除。
func Delete(clt *core.SDKClient, accessToken string, req *keyword.DeleteRequest) (*keyword.DeleteResponseData, error) {
	var resp keyword.DeleteResponse
	err := clt.Post("2/keyword/delete_v2/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
