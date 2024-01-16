package account

import (
	"github.com/disv5/opensdk/uc/core"
	"github.com/disv5/opensdk/uc/model/account"
)

// GetChildrenAccountByAccountId 获取当前账户管家或者代理商下的子账户列表
func GetChildrenAccountByAccountId(clt *core.SDKClient, req *account.GetChildrenAccountByAccountIdRequest) (*account.GetChildrenAccountByAccountIdResponse, error) {
	var resp account.GetChildrenAccountByAccountIdResponse
	err := clt.Post(req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
