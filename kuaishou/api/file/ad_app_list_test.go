package file

import (
	"fmt"
	"github.com/disv5/opensdk/kuaishou/core"
	"github.com/disv5/opensdk/kuaishou/model/file"
	"testing"
)

func TestAdAppList(t *testing.T) {
	var (
		sdk = core.NewSDKClient(0, "xxxxx")
		req = &file.AdAppListRequest{
			AdvertiserID: 11044594,
		}
		list *file.AdAppListResponse
	)

	list, err := AdAppList(sdk, "36f7b5519e1f4437e382c6a34ed3aec3", req)
	fmt.Println(err)
	fmt.Println(list)
}
