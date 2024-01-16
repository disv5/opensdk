package file

import (
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/core"
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/model/file"
)

// AdVideoList 查询视频信息list接口
func AdVideoList(clt *core.SDKClient, accessToken string, req *file.AdVideoListRequest) (*file.AdVideoListResponse, error) {
	var resp file.AdVideoListResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
