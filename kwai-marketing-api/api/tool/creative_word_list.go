package tool

import (
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/core"
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/model/tool"
)

// CreativeWordList 获取可选的动态词包
func CreativeWordList(clt *core.SDKClient, accessToken string, advertiserID uint64) ([]tool.CreativeWord, error) {
	req := &tool.CreativeWordListRequest{
		AdvertiserID: advertiserID,
	}
	var resp []tool.CreativeWord
	err := clt.Get(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
