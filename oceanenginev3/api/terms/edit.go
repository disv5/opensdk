package terms

import (
	"github.com/disv5/opensdk/oceanenginev3/core"
	"github.com/disv5/opensdk/oceanenginev3/model/terms"
)

func TermsList(clt *core.SDKClient, accessToken string, req *terms.TermsListReq) (*terms.TermsListItermData, error) {
	var resp terms.TermsListRes
	err := clt.Get("v3.0/tools/comment/terms_banned/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func TermsUserList(clt *core.SDKClient, accessToken string, req *terms.TermsUserListReq) (*terms.TermsUserListItermData, error) {
	var resp terms.TermsUserListRes
	err := clt.Get("v3.0/tools/aweme_banned/list/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func AddTerms(clt *core.SDKClient, accessToken string, req *terms.AddTermsReq) (struct{}, error) {
	var resp terms.AddTermsRes
	err := clt.Post("v3.0/tools/comment/terms_banned/add/", req, &resp, accessToken)
	if err != nil {
		return struct{}{}, err
	}
	return resp.Data, nil
}

func DelTerms(clt *core.SDKClient, accessToken string, req *terms.DelTermsReq) (struct{}, error) {
	var resp terms.DelTermsRes
	err := clt.Post("v3.0/tools/comment/terms_banned/delete/", req, &resp, accessToken)
	if err != nil {
		return struct{}{}, err
	}
	return resp.Data, nil
}

func UpdateTerms(clt *core.SDKClient, accessToken string, req *terms.UpdateTermsReq) (struct{}, error) {
	var resp terms.UpdateTermsRes
	err := clt.Post("v3.0/tools/comment/terms_banned/update/", req, &resp, accessToken)
	if err != nil {
		return struct{}{}, err
	}
	return resp.Data, nil
}

func AddUserTerms(clt *core.SDKClient, accessToken string, req *terms.AddUserTermsReq) (*terms.AddUserTermData, error) {
	var resp terms.AddUserTermsRes
	err := clt.Post("v3.0/tools/aweme_banned/create/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func DelUserTerms(clt *core.SDKClient, accessToken string, req *terms.DelUserTermsReq) (*terms.DelUserTermData, error) {
	var resp terms.DelUserTermsRes
	err := clt.Post("v3.0/tools/aweme_banned/delete/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
