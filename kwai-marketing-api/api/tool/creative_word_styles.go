package tool

import (
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/core"
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/model/tool"
)

// CreativeWordStyles 获取可选的封面贴纸样式
func CreativeWordStyles(clt *core.SDKClient, accessToken string, advertiserID uint64) ([]tool.CreativeWordStyle, error) {
	req := &tool.CreativeWordStylesRequest{
		AdvertiserID: advertiserID,
	}
	var resp []tool.CreativeWordStyle
	err := clt.Get(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
