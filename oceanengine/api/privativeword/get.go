package privativeword

import (
	"donson.com.cn/draining/internal/pkg/oceanengine/core"
	"donson.com.cn/draining/internal/pkg/oceanengine/model/privativeword"
)

// Get 获取否定词列表
// 获取否定词列表，根据广告组或广告计划获取否定词
func Get(clt *core.SDKClient, accessToken string, req *privativeword.GetRequest) (*privativeword.GetResponseData, error) {
	var resp privativeword.GetResponse
	err := clt.Get("2/privative_word/get", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
