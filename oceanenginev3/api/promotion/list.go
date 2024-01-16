package promotion

import (
	"donson.com.cn/draining/internal/pkg/oceanenginev3/core"
	"donson.com.cn/draining/internal/pkg/oceanenginev3/model/promotion"
)

func List(clt *core.SDKClient, accessToken string, req *promotion.ListRequest) (*promotion.ListResponseData, error) {
	var resp promotion.ListResponse
	err := clt.Get("v3.0/promotion/list/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func Update(clt *core.SDKClient, accessToken string, req *promotion.ReqUpdateAd) (*promotion.UpdateDataRes, error) {
	var resp promotion.UpdateAdRes
	err := clt.Post("v3.0/promotion/update/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func UpdateTwo(clt *core.SDKClient, accessToken string, req *promotion.ReqUpdateAdTwo) (*promotion.UpdateDataRes, error) {
	var resp promotion.UpdateAdRes
	err := clt.Post("v3.0/promotion/update/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func Status(clt *core.SDKClient, accessToken string, req *promotion.ReqStatusUpdateAd) (*promotion.UpdateStatusDataRes, error) {
	var resp promotion.UpdateStatusAdRes
	err := clt.Post("v3.0/promotion/status/update/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
