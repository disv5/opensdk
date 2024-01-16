package tools

import (
	"github.com/disv5/opensdk/oceanengine/core"
	"github.com/disv5/opensdk/oceanengine/model/tools"
)

// CountryInfo 查询国家/区域信息
func CountryInfo(clt *core.SDKClient, accessToken string, req *tools.CountryInfoRequest) (*tools.CountryInfoResponseData, error) {
	var resp tools.CountryInfoResponse
	if err := clt.Get("2/tools/country/info/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
