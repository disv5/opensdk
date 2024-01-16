package account

import (
	"donson.com.cn/draining/internal/pkg/uc/core"
	"donson.com.cn/draining/internal/pkg/uc/model/account"
)

// GetAccount 获取username所对应的账户信息
func GetAccount(clt *core.SDKClient, req *account.GetAccountRequest) (*account.GetAccountResponse, error) {
	var resp account.GetAccountResponse
	err := clt.Post(req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
