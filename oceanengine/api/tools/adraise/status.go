package adraise

import (
	"regexp"
	"strconv"

	"github.com/disv5/opensdk/oceanengine/core"
	"github.com/disv5/opensdk/oceanengine/enum"
	"github.com/disv5/opensdk/oceanengine/model/tools/adraise"
)

// Status 获取当前起量状态: 获取当前起量状态
func Status(clt *core.SDKClient, accessToken string, req *adraise.StatusRequest) (map[uint64]enum.AdRaiseStatus, error) {
	var resp adraise.StatusResponse
	if err := clt.Get("2/tools/ad_raise_status/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	ret := make(map[uint64]enum.AdRaiseStatus, len(req.AdIDs))
	re := regexp.MustCompile(`(\d+):\s+'(.*?)'`)
	matches := re.FindAllStringSubmatch(resp.Data.Status, -1)
	for _, match := range matches {
		if len(match) != 3 {
			continue
		}
		adID, _ := strconv.ParseUint(match[1], 10, 64)
		status := match[2]
		ret[adID] = enum.AdRaiseStatus(status)
	}
	return ret, nil
}
