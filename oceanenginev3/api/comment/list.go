package comment

import (
	"github.com/disv5/opensdk/oceanenginev3/core"
	"github.com/disv5/opensdk/oceanenginev3/model/comment"
)

func List(clt *core.SDKClient, accessToken string, req *comment.GetCommentListReq) (*comment.CommentListItemData, error) {
	var resp comment.GetCommentListRes
	err := clt.Get("v3.0/tools/comment/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func Hide(clt *core.SDKClient, accessToken string, req *comment.HideReq) (*comment.CommentHideItemData, error) {
	var resp comment.HideCommentRes
	err := clt.Post("v3.0/tools/comment/hide/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func Reply(clt *core.SDKClient, accessToken string, req *comment.ReplyReq) (*comment.CommentReplyItemData, error) {
	var resp comment.ReplyCommentRes
	err := clt.Post("v3.0/tools/comment/reply/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func Top(clt *core.SDKClient, accessToken string, req *comment.TopReq) (struct{}, error) {
	var resp comment.TopRes
	err := clt.Post("v3.0/tools/comment/reply/", req, &resp, accessToken)
	if err != nil {
		return struct{}{}, err
	}
	return resp.Data, nil
}
