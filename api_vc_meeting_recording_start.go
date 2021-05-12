// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// StartMeetingRecording 在会议中开始录制。
//
// 会议正在进行中，且操作者具有相应权限（如果操作者为用户，必须是会中当前主持人）
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting-recording/start
func (r *VCAPI) StartMeetingRecording(ctx context.Context, request *StartMeetingRecordingReq, options ...MethodOptionFunc) (*StartMeetingRecordingResp, *Response, error) {
	r.cli.logInfo(ctx, "[lark] VC#StartMeetingRecording call api")
	r.cli.logDebug(ctx, "[lark] VC#StartMeetingRecording request: %s", jsonString(request))

	if r.cli.mock.mockVCStartMeetingRecording != nil {
		r.cli.logDebug(ctx, "[lark] VC#StartMeetingRecording mock enable")
		return r.cli.mock.mockVCStartMeetingRecording(ctx, request, options...)
	}

	req := &RawRequestReq{
		Method:              "PATCH",
		URL:                 "https://open.feishu.cn/open-apis/vc/v1/meetings/:meeting_id/recording/start",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
	}
	resp := new(startMeetingRecordingResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark] VC#StartMeetingRecording PATCH https://open.feishu.cn/open-apis/vc/v1/meetings/:meeting_id/recording/start failed: %s", err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark] VC#StartMeetingRecording PATCH https://open.feishu.cn/open-apis/vc/v1/meetings/:meeting_id/recording/start failed, code: %d, msg: %s", resp.Code, resp.Msg)
		return nil, response, NewError("VC", "StartMeetingRecording", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark] VC#StartMeetingRecording request_id: %s, response: %s", response.RequestID, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockVCStartMeetingRecording(f func(ctx context.Context, request *StartMeetingRecordingReq, options ...MethodOptionFunc) (*StartMeetingRecordingResp, *Response, error)) {
	r.mockVCStartMeetingRecording = f
}

func (r *Mock) UnMockVCStartMeetingRecording() {
	r.mockVCStartMeetingRecording = nil
}

type StartMeetingRecordingReq struct {
	MeetingID string `path:"meeting_id" json:"-"` // 会议ID, 示例值: "6911188411932033028"
	Timezone  *int   `json:"timezone,omitempty"`  // 录制文件时间显示使用的时区[-12,12], 示例值: 8
}

type startMeetingRecordingResp struct {
	Code int                        `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 错误描述
	Data *StartMeetingRecordingResp `json:"data,omitempty"`
}

type StartMeetingRecordingResp struct{}
