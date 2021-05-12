// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// DeleteRoom 该接口用于删除会议室。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUzMxYjL1MTM24SNzEjN
func (r *MeetingRoomAPI) DeleteRoom(ctx context.Context, request *DeleteRoomReq, options ...MethodOptionFunc) (*DeleteRoomResp, *Response, error) {
	r.cli.logInfo(ctx, "[lark] MeetingRoom#DeleteRoom call api")
	r.cli.logDebug(ctx, "[lark] MeetingRoom#DeleteRoom request: %s", jsonString(request))

	if r.cli.mock.mockMeetingRoomDeleteRoom != nil {
		r.cli.logDebug(ctx, "[lark] MeetingRoom#DeleteRoom mock enable")
		return r.cli.mock.mockMeetingRoomDeleteRoom(ctx, request, options...)
	}

	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/meeting_room/room/delete",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(deleteRoomResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark] MeetingRoom#DeleteRoom POST https://open.feishu.cn/open-apis/meeting_room/room/delete failed: %s", err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark] MeetingRoom#DeleteRoom POST https://open.feishu.cn/open-apis/meeting_room/room/delete failed, code: %d, msg: %s", resp.Code, resp.Msg)
		return nil, response, NewError("MeetingRoom", "DeleteRoom", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark] MeetingRoom#DeleteRoom request_id: %s, response: %s", response.RequestID, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockMeetingRoomDeleteRoom(f func(ctx context.Context, request *DeleteRoomReq, options ...MethodOptionFunc) (*DeleteRoomResp, *Response, error)) {
	r.mockMeetingRoomDeleteRoom = f
}

func (r *Mock) UnMockMeetingRoomDeleteRoom() {
	r.mockMeetingRoomDeleteRoom = nil
}

type DeleteRoomReq struct {
	RoomID string `json:"room_id,omitempty"` // 要删除的会议室ID
}

type deleteRoomResp struct {
	Code int             `json:"code,omitempty"` // 返回码，非 0 表示失败
	Msg  string          `json:"msg,omitempty"`  // 返回码的描述，"success" 表示成功，其他为错误提示信息
	Data *DeleteRoomResp `json:"data,omitempty"`
}

type DeleteRoomResp struct{}
