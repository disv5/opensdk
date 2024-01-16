package file

import (
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/core"
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/model/file"
)

// AdImageGet 查询图片信息get接口
func AdImageGet(clt *core.SDKClient, accessToken string, req *file.AdImageGetRequest) (*file.Image, error) {
	var resp file.Image
	err := clt.Get(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
