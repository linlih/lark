// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GenerateCaldavConf 用于为当前用户生成一个CalDAV账号密码。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/setting/generate_caldav_conf
func (r *CalendarAPI) GenerateCaldavConf(ctx context.Context, request *GenerateCaldavConfReq, options ...MethodOptionFunc) (*GenerateCaldavConfResp, *Response, error) {
	r.cli.logInfo(ctx, "[lark] Calendar#GenerateCaldavConf call api")
	r.cli.logDebug(ctx, "[lark] Calendar#GenerateCaldavConf request: %s", jsonString(request))

	if r.cli.mock.mockCalendarGenerateCaldavConf != nil {
		r.cli.logDebug(ctx, "[lark] Calendar#GenerateCaldavConf mock enable")
		return r.cli.mock.mockCalendarGenerateCaldavConf(ctx, request, options...)
	}

	req := &RawRequestReq{
		Method:              "POST",
		URL:                 "https://open.feishu.cn/open-apis/calendar/v4/settings/generate_caldav_conf",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
	}
	resp := new(generateCaldavConfResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark] Calendar#GenerateCaldavConf POST https://open.feishu.cn/open-apis/calendar/v4/settings/generate_caldav_conf failed: %s", err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark] Calendar#GenerateCaldavConf POST https://open.feishu.cn/open-apis/calendar/v4/settings/generate_caldav_conf failed, code: %d, msg: %s", resp.Code, resp.Msg)
		return nil, response, NewError("Calendar", "GenerateCaldavConf", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark] Calendar#GenerateCaldavConf request_id: %s, response: %s", response.RequestID, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockCalendarGenerateCaldavConf(f func(ctx context.Context, request *GenerateCaldavConfReq, options ...MethodOptionFunc) (*GenerateCaldavConfResp, *Response, error)) {
	r.mockCalendarGenerateCaldavConf = f
}

func (r *Mock) UnMockCalendarGenerateCaldavConf() {
	r.mockCalendarGenerateCaldavConf = nil
}

type GenerateCaldavConfReq struct {
	DeviceName *string `json:"device_name,omitempty"` // 需要同步日历的设备名，在日历中展示用来管理密码, 示例值："iPhone", 最大长度：`100` 字符
}

type generateCaldavConfResp struct {
	Code int                     `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 错误描述
	Data *GenerateCaldavConfResp `json:"data,omitempty"` //
}

type GenerateCaldavConfResp struct {
	Password      string `json:"password,omitempty"`       // caldav密码
	UserName      string `json:"user_name,omitempty"`      // caldav用户名
	ServerAddress string `json:"server_address,omitempty"` // 服务器地址
	DeviceName    string `json:"device_name,omitempty"`    // 设备名
}
