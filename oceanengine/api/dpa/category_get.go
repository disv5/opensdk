package dpa

import (
	"github.com/disv5/opensdk/oceanengine/core"
	"github.com/disv5/opensdk/oceanengine/model/dpa"
)

// CategoryGet 获取DPA分类
func CategoryGet(clt *core.SDKClient, accessToken string, req *dpa.CategoryGetRequest) ([]dpa.Category, error) {
	var resp dpa.CategoryGetResponse
	if err := clt.Get("2/dpa/category/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
