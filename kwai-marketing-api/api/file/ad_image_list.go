package file

import (
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/core"
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/model/file"
)

// AdImageList 查询图片信息list接口
func AdImageList(clt *core.SDKClient, accessToken string, req *file.AdImageListRequest) (*file.AdImageListResponse, error) {
	var resp file.AdImageListResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
