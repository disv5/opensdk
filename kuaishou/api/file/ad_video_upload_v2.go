package file

import (
	"github.com/disv5/opensdk/kuaishou/core"
	"github.com/disv5/opensdk/kuaishou/model/file"
)

// AdVideoUploadV2 上传视频v2接口
func AdVideoUploadV2(clt *core.SDKClient, accessToken string, req *file.AdVideoUploadRequestV2) (*file.Video, error) {
	var resp file.Video
	err := clt.Upload(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
