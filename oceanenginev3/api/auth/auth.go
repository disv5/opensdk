package auth

import (
	"github.com/disv5/opensdk/oceanenginev3/core"
	auth "github.com/disv5/opensdk/oceanenginev3/model/status"
)

func AuthStatus(clt *core.SDKClient, accessToken string, req *auth.AuthGetRequest) (*auth.AuthGetResponseData, error) {
	var resp auth.AuthGetResponse
	err := clt.Get("2/event_manager/auth/get_auth_status", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func AuthDisable(clt *core.SDKClient, accessToken string, req *auth.AuthOprRequest) (*auth.AuthResData, error) {
	var resp auth.AuthRes
	err := clt.Post("2/event_manager/auth/disable", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func AuthEnable(clt *core.SDKClient, accessToken string, req *auth.AuthOprRequest) (*auth.AuthResData, error) {
	var resp auth.AuthRes
	err := clt.Post("2/event_manager/auth/enable", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
