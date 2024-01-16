package project

import (
	"github.com/disv5/opensdk/oceanenginev3/core"
	promotion "github.com/disv5/opensdk/oceanenginev3/model/project"
)

func List(clt *core.SDKClient, accessToken string, req *promotion.GetProjectListReq) (*promotion.ProjectListItemData, error) {
	var resp promotion.GetProjectListRes
	err := clt.Get("v3.0/project/list/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func Update(clt *core.SDKClient, accessToken string, req *promotion.UpdateProjectReq) (*promotion.UpdateProjectResData, error) {
	var resp promotion.UpdateProjectRes
	err := clt.Post("v3.0/project/update/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
