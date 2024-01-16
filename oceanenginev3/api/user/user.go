package user

import (
	"donson.com.cn/draining/internal/pkg/oceanenginev3/core"
	"donson.com.cn/draining/internal/pkg/oceanenginev3/model/user"
)

func Info(clt *core.SDKClient, accessToken string) (*user.UserItemData, error) {
	var resp user.GetUserInfoRes
	err := clt.Get("2/user/info/", nil, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
