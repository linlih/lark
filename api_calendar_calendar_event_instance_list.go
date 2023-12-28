// Code generated by lark_sdk_gen. DO NOT EDIT.
/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lark

import (
	"context"
)

// GetCalendarEventInstanceList 该接口用于以当前身份（应用 / 用户）在获取日历上重复性日程的日程实例信息。
//
// 身份由 Header Authorization 的 Token 类型决定。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/instances
func (r *CalendarService) GetCalendarEventInstanceList(ctx context.Context, request *GetCalendarEventInstanceListReq, options ...MethodOptionFunc) (*GetCalendarEventInstanceListResp, *Response, error) {
	if r.cli.mock.mockCalendarGetCalendarEventInstanceList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Calendar#GetCalendarEventInstanceList mock enable")
		return r.cli.mock.mockCalendarGetCalendarEventInstanceList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Calendar",
		API:                   "GetCalendarEventInstanceList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id/instances",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getCalendarEventInstanceListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCalendarGetCalendarEventInstanceList mock CalendarGetCalendarEventInstanceList method
func (r *Mock) MockCalendarGetCalendarEventInstanceList(f func(ctx context.Context, request *GetCalendarEventInstanceListReq, options ...MethodOptionFunc) (*GetCalendarEventInstanceListResp, *Response, error)) {
	r.mockCalendarGetCalendarEventInstanceList = f
}

// UnMockCalendarGetCalendarEventInstanceList un-mock CalendarGetCalendarEventInstanceList method
func (r *Mock) UnMockCalendarGetCalendarEventInstanceList() {
	r.mockCalendarGetCalendarEventInstanceList = nil
}

// GetCalendarEventInstanceListReq ...
type GetCalendarEventInstanceListReq struct {
	CalendarID string  `path:"calendar_id" json:"-"` // 日历资源ID, 示例值: "feishu.cn_HF9U2MbibE8PPpjro6xjqa@group.calendar.feishu.cn"
	EventID    string  `path:"event_id" json:"-"`    // 日程资源ID, 示例值: "75d28f9b-e35c-4230-8a83-4a661497db54_0"
	StartTime  string  `query:"start_time" json:"-"` // 日程实例开始Unix时间戳, 单位为秒, 返回的日程实例信息的end_time下限（不包含）, 示例值: 1631777271
	EndTime    string  `query:"end_time" json:"-"`   // 日程实例结束Unix时间戳, 单位为秒, 返回的日程实例信息的start_time上限（不包含）, 示例值: 1631777271
	PageSize   *int64  `query:"page_size" json:"-"`  // 分页大小, 示例值: 10, 默认值: `50`, 取值范围: `10` ～ `500`
	PageToken  *string `query:"page_token" json:"-"` // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: eVQrYzJBNDNONlk4VFZBZVlSdzlKdFJ4bVVHVExENDNKVHoxaVdiVnViQT0=
}

// GetCalendarEventInstanceListResp ...
type GetCalendarEventInstanceListResp struct {
	Items     []*GetCalendarEventInstanceListRespItem `json:"items,omitempty"`      // instances实例
	PageToken string                                  `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	HasMore   bool                                    `json:"has_more,omitempty"`   // 是否还有更多项
}

// GetCalendarEventInstanceListRespItem ...
type GetCalendarEventInstanceListRespItem struct {
	EventID     string                                         `json:"event_id,omitempty"`     // 日程实例ID
	Summary     string                                         `json:"summary,omitempty"`      // 日程主题
	Description string                                         `json:"description,omitempty"`  // 日程描述
	StartTime   *GetCalendarEventInstanceListRespItemStartTime `json:"start_time,omitempty"`   // 开始时间
	EndTime     *GetCalendarEventInstanceListRespItemEndTime   `json:"end_time,omitempty"`     // 结束时间
	Status      string                                         `json:"status,omitempty"`       // 日程状态, 可选值有: tentative: 未回应, confirmed: 已确认, cancelled: 日程已取消
	IsException bool                                           `json:"is_exception,omitempty"` // 是否是例外日程实例
	AppLink     string                                         `json:"app_link,omitempty"`     // 日程的app_link, 跳转到具体的某个日程
	Location    *GetCalendarEventInstanceListRespItemLocation  `json:"location,omitempty"`     // 日程地点
}

// GetCalendarEventInstanceListRespItemEndTime ...
type GetCalendarEventInstanceListRespItemEndTime struct {
	Date      string `json:"date,omitempty"`      // 仅全天日程使用该字段, 如2018-09-01。需满足 RFC3339 格式。不能与 time_stamp 同时指定
	Timestamp string `json:"timestamp,omitempty"` // 秒级时间戳, 如1602504000(表示2020/10/12 20:0:00 +8时区)
	Timezone  string `json:"timezone,omitempty"`  // 时区名称, 使用IANA Time Zone Database标准, 如Asia/Shanghai；全天日程时区固定为UTC, 非全天日程时区默认为Asia/Shanghai
}

// GetCalendarEventInstanceListRespItemLocation ...
type GetCalendarEventInstanceListRespItemLocation struct {
	Name      string  `json:"name,omitempty"`      // 地点名称
	Address   string  `json:"address,omitempty"`   // 地点地址
	Latitude  float64 `json:"latitude,omitempty"`  // 地点坐标纬度信息, 对于国内的地点, 采用GCJ-02标准, 海外地点采用WGS84标准
	Longitude float64 `json:"longitude,omitempty"` // 地点坐标经度信息, 对于国内的地点, 采用GCJ-02标准, 海外地点采用WGS84标准
}

// GetCalendarEventInstanceListRespItemStartTime ...
type GetCalendarEventInstanceListRespItemStartTime struct {
	Date      string `json:"date,omitempty"`      // 仅全天日程使用该字段, 如2018-09-01。需满足 RFC3339 格式。不能与 time_stamp 同时指定
	Timestamp string `json:"timestamp,omitempty"` // 秒级时间戳, 如1602504000(表示2020/10/12 20:0:00 +8时区)
	Timezone  string `json:"timezone,omitempty"`  // 时区名称, 使用IANA Time Zone Database标准, 如Asia/Shanghai；全天日程时区固定为UTC, 非全天日程时区默认为Asia/Shanghai
}

// getCalendarEventInstanceListResp ...
type getCalendarEventInstanceListResp struct {
	Code int64                             `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                            `json:"msg,omitempty"`  // 错误描述
	Data *GetCalendarEventInstanceListResp `json:"data,omitempty"`
}
