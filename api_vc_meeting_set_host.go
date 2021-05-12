// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// SetHostMeeting 设置会议的主持人
//
// 发起设置主持人的操作者必须具有相应的权限（如果操作者为用户，必须是会中当前主持人）；该操作使用CAS并发安全机制，需传入会中当前主持人，如果操作失败可使用返回的最新数据重试
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting/set_host
func (r *VCAPI) SetHostMeeting(ctx context.Context, request *SetHostMeetingReq, options ...MethodOptionFunc) (*SetHostMeetingResp, *Response, error) {
	r.cli.logInfo(ctx, "[lark] VC#SetHostMeeting call api")
	r.cli.logDebug(ctx, "[lark] VC#SetHostMeeting request: %s", jsonString(request))

	if r.cli.mock.mockVCSetHostMeeting != nil {
		r.cli.logDebug(ctx, "[lark] VC#SetHostMeeting mock enable")
		return r.cli.mock.mockVCSetHostMeeting(ctx, request, options...)
	}

	req := &RawRequestReq{
		Method:              "PATCH",
		URL:                 "https://open.feishu.cn/open-apis/vc/v1/meetings/:meeting_id/set_host",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
	}
	resp := new(setHostMeetingResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark] VC#SetHostMeeting PATCH https://open.feishu.cn/open-apis/vc/v1/meetings/:meeting_id/set_host failed: %s", err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark] VC#SetHostMeeting PATCH https://open.feishu.cn/open-apis/vc/v1/meetings/:meeting_id/set_host failed, code: %d, msg: %s", resp.Code, resp.Msg)
		return nil, response, NewError("VC", "SetHostMeeting", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark] VC#SetHostMeeting request_id: %s, response: %s", response.RequestID, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockVCSetHostMeeting(f func(ctx context.Context, request *SetHostMeetingReq, options ...MethodOptionFunc) (*SetHostMeetingResp, *Response, error)) {
	r.mockVCSetHostMeeting = f
}

func (r *Mock) UnMockVCSetHostMeeting() {
	r.mockVCSetHostMeeting = nil
}

type SetHostMeetingReq struct {
	UserIDType  *IDType                       `query:"user_id_type" json:"-"`  // 用户 ID 类型, 示例值："open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 userid
	MeetingID   string                        `path:"meeting_id" json:"-"`     // 会议ID, 示例值："6911188411932033028"
	HostUser    *SetHostMeetingReqHostUser    `json:"host_user,omitempty"`     // 将要设置的主持人
	OldHostUser *SetHostMeetingReqOldHostUser `json:"old_host_user,omitempty"` // 当前主持人（CAS并发安全：如果和会中当前主持人不符则会设置失败，可使用返回的最新数据需重新设置）
}

type SetHostMeetingReqHostUser struct {
	ID       *string `json:"id,omitempty"`        // 用户ID, 示例值："ou_3ec3f6a28a0d08c45d895276e8e5e19b"
	UserType *int    `json:"user_type,omitempty"` // 用户类型, 示例值：1, 可选值有: `1`：lark用户, `2`：rooms用户, `3`：文档用户, `4`：neo单品用户, `5`：neo单品游客用户, `6`：pstn用户, `7`：sip用户
}

type SetHostMeetingReqOldHostUser struct {
	ID       *string `json:"id,omitempty"`        // 用户ID, 示例值："ou_3ec3f6a28a0d08c45d895276e8e5e19b"
	UserType *int    `json:"user_type,omitempty"` // 用户类型, 示例值：1, 可选值有: `1`：lark用户, `2`：rooms用户, `3`：文档用户, `4`：neo单品用户, `5`：neo单品游客用户, `6`：pstn用户, `7`：sip用户
}

type setHostMeetingResp struct {
	Code int                 `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string              `json:"msg,omitempty"`  // 错误描述
	Data *SetHostMeetingResp `json:"data,omitempty"` //
}

type SetHostMeetingResp struct {
	HostUser *SetHostMeetingRespHostUser `json:"host_user,omitempty"` // 会中当前主持人
}

type SetHostMeetingRespHostUser struct {
	ID       string `json:"id,omitempty"`        // 用户ID
	UserType int    `json:"user_type,omitempty"` // 用户类型, 可选值有: `1`：lark用户, `2`：rooms用户, `3`：文档用户, `4`：neo单品用户, `5`：neo单品游客用户, `6`：pstn用户, `7`：sip用户
}
