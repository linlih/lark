// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// DeleteShift
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//shift_delete
func (r *AttendanceAPI) DeleteShift(ctx context.Context, request *DeleteShiftReq, options ...MethodOptionFunc) (*DeleteShiftResp, *Response, error) {
	r.cli.logInfo(ctx, "[lark] Attendance#DeleteShift call api")
	r.cli.logDebug(ctx, "[lark] Attendance#DeleteShift request: %s", jsonString(request))

	if r.cli.mock.mockAttendanceDeleteShift != nil {
		r.cli.logDebug(ctx, "[lark] Attendance#DeleteShift mock enable")
		return r.cli.mock.mockAttendanceDeleteShift(ctx, request, options...)
	}

	req := &RawRequestReq{
		Method:                "DELETE",
		URL:                   "https://open.feishu.cn/open-apis/attendance/v1/shifts/:shift_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(deleteShiftResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark] Attendance#DeleteShift DELETE https://open.feishu.cn/open-apis/attendance/v1/shifts/:shift_id failed: %s", err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark] Attendance#DeleteShift DELETE https://open.feishu.cn/open-apis/attendance/v1/shifts/:shift_id failed, code: %d, msg: %s", resp.Code, resp.Msg)
		return nil, response, NewError("Attendance", "DeleteShift", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark] Attendance#DeleteShift request_id: %s, response: %s", response.RequestID, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockAttendanceDeleteShift(f func(ctx context.Context, request *DeleteShiftReq, options ...MethodOptionFunc) (*DeleteShiftResp, *Response, error)) {
	r.mockAttendanceDeleteShift = f
}

func (r *Mock) UnMockAttendanceDeleteShift() {
	r.mockAttendanceDeleteShift = nil
}

type DeleteShiftReq struct {
	ShiftID string `path:"shift_id" json:"-"` // 班次 ID，示例值："6919358778597097404"
}

type deleteShiftResp struct {
	Code int              `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string           `json:"msg,omitempty"`  // 错误描述
	Data *DeleteShiftResp `json:"data,omitempty"`
}

type DeleteShiftResp struct{}
