// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// BatchCreateUserFlow
//
// 导入授权内员工的打卡流水记录。导入后，会根据员工所在的考勤组班次规则，计算最终的打卡状态与结果。
// 适用于考勤机数据导入等场景。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//ImportAttendanceFlowRecords
func (r *AttendanceAPI) BatchCreateUserFlow(ctx context.Context, request *BatchCreateUserFlowReq, options ...MethodOptionFunc) (*BatchCreateUserFlowResp, *Response, error) {
	r.cli.logInfo(ctx, "[lark] Attendance#BatchCreateUserFlow call api")
	r.cli.logDebug(ctx, "[lark] Attendance#BatchCreateUserFlow request: %s", jsonString(request))

	if r.cli.mock.mockAttendanceBatchCreateUserFlow != nil {
		r.cli.logDebug(ctx, "[lark] Attendance#BatchCreateUserFlow mock enable")
		return r.cli.mock.mockAttendanceBatchCreateUserFlow(ctx, request, options...)
	}

	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/attendance/v1/user_flows/batch_create",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(batchCreateUserFlowResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark] Attendance#BatchCreateUserFlow POST https://open.feishu.cn/open-apis/attendance/v1/user_flows/batch_create failed: %s", err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark] Attendance#BatchCreateUserFlow POST https://open.feishu.cn/open-apis/attendance/v1/user_flows/batch_create failed, code: %d, msg: %s", resp.Code, resp.Msg)
		return nil, response, NewError("Attendance", "BatchCreateUserFlow", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark] Attendance#BatchCreateUserFlow request_id: %s, response: %s", response.RequestID, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockAttendanceBatchCreateUserFlow(f func(ctx context.Context, request *BatchCreateUserFlowReq, options ...MethodOptionFunc) (*BatchCreateUserFlowResp, *Response, error)) {
	r.mockAttendanceBatchCreateUserFlow = f
}

func (r *Mock) UnMockAttendanceBatchCreateUserFlow() {
	r.mockAttendanceBatchCreateUserFlow = nil
}

type BatchCreateUserFlowReq struct {
	EmployeeType EmployeeType                        `query:"employee_type" json:"-"` // 请求体中的 user_id 和 creator_id 的员工工号类型，可用值：【employee_id（员工的 employeeId），employee_no（员工工号）】，示例值："employee_id"
	FlowRecords  []*BatchCreateUserFlowReqFlowRecord `json:"flow_records,omitempty"`  // 打卡流水记录列表
}

type BatchCreateUserFlowReqFlowRecord struct {
	UserID       string `json:"user_id,omitempty"`       // 员工工号
	CreatorID    string `json:"creator_id,omitempty"`    // 打卡记录创建者的工号
	LocationName string `json:"location_name,omitempty"` // 打卡位置名称信息
	CheckTime    string `json:"check_time,omitempty"`    // 打卡时间，精确到秒的时间戳
	Comment      string `json:"comment,omitempty"`       // 打卡备注
}

type batchCreateUserFlowResp struct {
	Code int                      `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *BatchCreateUserFlowResp `json:"data,omitempty"` // -
}

type BatchCreateUserFlowResp struct {
	FlowRecords []*BatchCreateUserFlowRespFlowRecord `json:"flow_records,omitempty"` // 打卡流水记录列表
}

type BatchCreateUserFlowRespFlowRecord struct {
	UserID       string   `json:"user_id,omitempty"`       // 员工工号
	CreatorID    string   `json:"creator_id,omitempty"`    // 打卡记录创建者的 employee_no
	LocationName string   `json:"location_name,omitempty"` // 打卡位置名称信息
	CheckTime    string   `json:"check_time,omitempty"`    // 打卡时间，精确到秒的时间戳
	Comment      string   `json:"comment,omitempty"`       // 打卡备注
	RecordID     string   `json:"record_id,omitempty"`     // 打卡记录 ID
	Longitude    float64  `json:"longitude,omitempty"`     // 打卡经度
	Latitude     float64  `json:"latitude,omitempty"`      // 打卡纬度
	Ssid         string   `json:"ssid,omitempty"`          // 打卡 Wi-Fi 的 SSID
	Bssid        string   `json:"bssid,omitempty"`         // 打卡 Wi-Fi 的 MAC 地址
	IsField      bool     `json:"is_field,omitempty"`      // 是否为外勤打卡
	IsWifi       bool     `json:"is_wifi,omitempty"`       // 是否为 Wi-Fi 打卡
	Type         int      `json:"type,omitempty"`          // 记录生成方式，可用值：【0（用户自己打卡），1（管理员修改），2（用户补卡），3（系统自动生成），4（下班免打卡），5（考勤机打卡），6（极速打卡），7（考勤开放平台导入）】
	PhotoUrls    []string `json:"photo_urls,omitempty"`    // 打卡照片列表
}
