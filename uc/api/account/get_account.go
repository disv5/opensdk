package account

import (
	"github.com/disv5/opensdk/uc/core"
	"github.com/disv5/opensdk/uc/model/account"
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
