package report

import (
	"donson.com.cn/draining/internal/pkg/oceanenginev3/core"
	"donson.com.cn/draining/internal/pkg/oceanenginev3/model/report"
)

func CustomGet(clt *core.SDKClient, accessToken string, req *report.CustomGetRequest) (*report.CustomGetResponseData, error) {
	var resp report.CustomGetResponse
	err := clt.Get("v3.0/report/custom/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
