package asynctask

import (
	"github.com/disv5/opensdk/kwai-marketing-api/core"
	"github.com/disv5/opensdk/kwai-marketing-api/model/report/asynctask"
)

func Download(clt *core.SDKClient, accessToken string, req *asynctask.DownloadRequest) ([]byte, error) {
	return clt.GetBytes(accessToken, req)
}
