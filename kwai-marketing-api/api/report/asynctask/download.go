package asynctask

import (
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/core"
	"donson.com.cn/draining/internal/pkg/kwai-marketing-api/model/report/asynctask"
)

func Download(clt *core.SDKClient, accessToken string, req *asynctask.DownloadRequest) ([]byte, error) {
	return clt.GetBytes(accessToken, req)
}
