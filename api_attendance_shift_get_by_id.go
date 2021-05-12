// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetShiftByID
//
// 通过班次 ID 获取班次详情。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//shift_by_id
func (r *AttendanceAPI) GetShiftByID(ctx context.Context, request *GetShiftByIDReq, options ...MethodOptionFunc) (*GetShiftByIDResp, *Response, error) {
	r.cli.logInfo(ctx, "[lark] Attendance#GetShiftByID call api")
	r.cli.logDebug(ctx, "[lark] Attendance#GetShiftByID request: %s", jsonString(request))

	if r.cli.mock.mockAttendanceGetShiftByID != nil {
		r.cli.logDebug(ctx, "[lark] Attendance#GetShiftByID mock enable")
		return r.cli.mock.mockAttendanceGetShiftByID(ctx, request, options...)
	}

	req := &RawRequestReq{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/attendance/v1/shifts/:shift_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getShiftByIDResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark] Attendance#GetShiftByID GET https://open.feishu.cn/open-apis/attendance/v1/shifts/:shift_id failed: %s", err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark] Attendance#GetShiftByID GET https://open.feishu.cn/open-apis/attendance/v1/shifts/:shift_id failed, code: %d, msg: %s", resp.Code, resp.Msg)
		return nil, response, NewError("Attendance", "GetShiftByID", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark] Attendance#GetShiftByID request_id: %s, response: %s", response.RequestID, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockAttendanceGetShiftByID(f func(ctx context.Context, request *GetShiftByIDReq, options ...MethodOptionFunc) (*GetShiftByIDResp, *Response, error)) {
	r.mockAttendanceGetShiftByID = f
}

func (r *Mock) UnMockAttendanceGetShiftByID() {
	r.mockAttendanceGetShiftByID = nil
}

type GetShiftByIDReq struct {
	ShiftID string `path:"shift_id" json:"-"` // 班次 ID，示例值："6919358778597097404"
}

type getShiftByIDResp struct {
	Code int               `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string            `json:"msg,omitempty"`  // 错误描述
	Data *GetShiftByIDResp `json:"data,omitempty"` // -
}

type GetShiftByIDResp struct {
	ShiftID           string                             `json:"shift_id,omitempty"`              // 班次 ID
	ShiftName         string                             `json:"shift_name,omitempty"`            // 班次名称
	PunchTimes        int                                `json:"punch_times,omitempty"`           // 打卡次数
	IsFlexible        bool                               `json:"is_flexible,omitempty"`           // 是否弹性打卡
	FlexibleMinutes   int                                `json:"flexible_minutes,omitempty"`      // 弹性打卡时间
	NoNeedOff         bool                               `json:"no_need_off,omitempty"`           // 是否下班免打卡
	PunchTimeRule     *GetShiftByIDRespPunchTimeRule     `json:"punch_time_rule,omitempty"`       // 打卡规则
	LateOffLateOnRule *GetShiftByIDRespLateOffLateOnRule `json:"late_off_late_on_rule,omitempty"` // 晚走晚到规则
	RestTimeRule      *GetShiftByIDRespRestTimeRule      `json:"rest_time_rule,omitempty"`        // 休息时间规则
}

type GetShiftByIDRespPunchTimeRule struct {
	OnTime              string `json:"on_time,omitempty"`                // 上班时间
	OffTime             string `json:"off_time,omitempty"`               // 下班时间
	LateMinutesAsLate   int    `json:"late_minutes_as_late,omitempty"`   // 晚到多长时间记为迟到
	LateMinutesAsLack   int    `json:"late_minutes_as_lack,omitempty"`   // 晚到多长时间记为缺卡
	OnAdvanceMinutes    int    `json:"on_advance_minutes,omitempty"`     // 最早可提前多长时间打上班卡
	EarlyMinutesAsEarly int    `json:"early_minutes_as_early,omitempty"` // 早走多长时间记为早退
	EarlyMinutesAsLack  int    `json:"early_minutes_as_lack,omitempty"`  // 早走多长时间记为缺卡
	OffDelayMinutes     int    `json:"off_delay_minutes,omitempty"`      // 最晚可延后多长时间打下班卡
}

type GetShiftByIDRespLateOffLateOnRule struct {
	LateOffMinutes int `json:"late_off_minutes,omitempty"` // 晚走多长时间
	LateOnMinutes  int `json:"late_on_minutes,omitempty"`  // 晚到多长时间
}

type GetShiftByIDRespRestTimeRule struct {
	RestBeginTime string `json:"rest_begin_time,omitempty"` // 休息开始时间
	RestEndTime   string `json:"rest_end_time,omitempty"`   // 休息结束时间
}
