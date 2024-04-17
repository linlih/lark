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

// UpdateCalendarEvent 调用该接口以当前身份（应用或用户）更新指定日历上的一个日程, 包括日程标题、描述、开始与结束时间、视频会议以及日程地点等信息。
//
// - 当前身份由 Header Authorization 的 Token 类型决定。tenant_access_token 指应用身份, user_access_token 指用户身份。
// - 如果使用应用身份调用该接口, 则需要确保应用开启了[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)。
// - 当前身份必须对日历有 writer 或 owner 权限, 并且日历的类型只能为 primary 或 shared。你可以调用[查询日历信息](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/get)接口, 获取日历类型以及当前身份对该日历的访问权限。
// - 当前身份为日程组织者时, 可修改该接口内的所有可编辑字段。
// - 当前身份为日程参与者时, 仅可编辑部分字段（包括 visibility、free_busy_status、color、reminders）。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/patch
// new doc: https://open.feishu.cn/document/server-docs/calendar-v4/calendar-event/patch
func (r *CalendarService) UpdateCalendarEvent(ctx context.Context, request *UpdateCalendarEventReq, options ...MethodOptionFunc) (*UpdateCalendarEventResp, *Response, error) {
	if r.cli.mock.mockCalendarUpdateCalendarEvent != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Calendar#UpdateCalendarEvent mock enable")
		return r.cli.mock.mockCalendarUpdateCalendarEvent(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Calendar",
		API:                   "UpdateCalendarEvent",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(updateCalendarEventResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCalendarUpdateCalendarEvent mock CalendarUpdateCalendarEvent method
func (r *Mock) MockCalendarUpdateCalendarEvent(f func(ctx context.Context, request *UpdateCalendarEventReq, options ...MethodOptionFunc) (*UpdateCalendarEventResp, *Response, error)) {
	r.mockCalendarUpdateCalendarEvent = f
}

// UnMockCalendarUpdateCalendarEvent un-mock CalendarUpdateCalendarEvent method
func (r *Mock) UnMockCalendarUpdateCalendarEvent() {
	r.mockCalendarUpdateCalendarEvent = nil
}

// UpdateCalendarEventReq ...
type UpdateCalendarEventReq struct {
	CalendarID       string                            `path:"calendar_id" json:"-"`        // 日程所在的日历 ID。了解更多, 参见[日历 ID 说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/introduction), 示例值: "feishu.cn_xxxxxxxxxx@group.calendar.feishu.cn"
	EventID          string                            `path:"event_id" json:"-"`           // 日程 ID, 创建日程时会返回日程 ID。你也可以调用以下接口获取某一日历的 ID, [获取日程列表](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/list), [搜索日程](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/search), 示例值: "00592a0e-7edf-4678-bc9d-1b77383ef08e_0"
	UserIDType       *IDType                           `query:"user_id_type" json:"-"`      // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	Summary          *string                           `json:"summary,omitempty"`           // 日程标题, 默认值: 空, 表示不更新该字段, 示例值: "日程标题", 最大长度: `1000` 字符
	Description      *string                           `json:"description,omitempty"`       // 日程描述, 注意: 目前 API 方式不支持编辑富文本描述。如果日程描述通过客户端编辑为富文本内容, 则使用 API 更新描述会导致富文本格式丢失, 默认值: 空, 表示不更新该字段, 示例值: "日程描述", 最大长度: `40960` 字符
	NeedNotification *bool                             `json:"need_notification,omitempty"` // 更新日程时, 是否给日程参与人发送 Bot 通知, 默认值: 空, 表示不更新该字段, 可选值有: true: 发送通知, false: 不发送通知, 示例值: false
	StartTime        *UpdateCalendarEventReqStartTime  `json:"start_time,omitempty"`        // 日程开始时间。不传值则表示不更新该字段。
	EndTime          *UpdateCalendarEventReqEndTime    `json:"end_time,omitempty"`          // 日程结束时间。不传值则表示不更新该字段。
	Vchat            *UpdateCalendarEventReqVchat      `json:"vchat,omitempty"`             // 视频会议信息。不传值则表示不更新该字段。
	Visibility       *string                           `json:"visibility,omitempty"`        // 日程公开范围, 注意: 更新日程时如果修改了该参数值, 则仅对当前身份生效, 默认值: 空, 表示不更新该字段, 示例值: "default", 可选值有: default: 默认权限, 即跟随日历权限, 默认仅向他人显示是否忙碌, public: 公开, 显示日程详情, private: 私密, 仅自己可见详情
	AttendeeAbility  *string                           `json:"attendee_ability,omitempty"`  // 参与人权限, 默认值: 空, 表示不更新该字段, 示例值: "can_see_others", 可选值有: none: 无法编辑日程、无法邀请其他参与人、无法查看参与人列表, can_see_others: 无法编辑日程、无法邀请其他参与人、可以查看参与人列表, can_invite_others: 无法编辑日程、可以邀请其他参与人、可以查看参与人列表, can_modify_event: 可以编辑日程、可以邀请其他参与人、可以查看参与人列表
	FreeBusyStatus   *string                           `json:"free_busy_status,omitempty"`  // 日程占用的忙闲状态, 新建日程默认为 `busy`, 注意: 更新日程时如果修改了该参数值, 则仅对当前身份生效, 默认值: 空, 表示不更新该字段, 示例值: "busy", 可选值有: busy: 忙碌, free: 空闲
	Location         *UpdateCalendarEventReqLocation   `json:"location,omitempty"`          // 日程地点。不传值则表示不更新该字段。
	Color            *int64                            `json:"color,omitempty"`             // 日程颜色, 取值通过颜色 RGB 值的 int32 表示, 注意: 该参数仅对当前身份生效, 客户端展示时会映射到色板上最接近的一种颜色, 取值为 0 或 -1 时, 默认跟随日历颜色, 默认值: 空, 表示不更新该字段, 示例值:1
	Reminders        []*UpdateCalendarEventReqReminder `json:"reminders,omitempty"`         // 日程提醒列表。不传值则表示不更新该字段。
	Recurrence       *string                           `json:"recurrence,omitempty"`        // 重复日程的重复性规则, 规则设置方式参考[rfc5545](https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10), 注意: COUNT 和 UNTIL 不支持同时出现, 预定会议室重复日程长度不得超过两年, 默认值: 空, 表示不更新该字段, 示例值: "FREQ=DAILY;INTERVAL=1", 最大长度: `2000` 字符
	Schemas          []*UpdateCalendarEventReqSchema   `json:"schemas,omitempty"`           // 日程自定义信息, 控制日程详情页的 UI 展示。不传值则表示不更新该字段。
}

// UpdateCalendarEventReqEndTime ...
type UpdateCalendarEventReqEndTime struct {
	Date      *string `json:"date,omitempty"`      // 结束时间, 仅全天日程使用该字段, [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339) 格式, 例如, 2018-09-01, 注意: 该参数不能与 `timestamp` 同时指定, 示例值: "2018-09-01"
	Timestamp *string `json:"timestamp,omitempty"` // 秒级时间戳, 用于设置具体的结束时间。例如, 1602504000 表示 2020/10/12 20:00:00（UTC +8 时区）, 注意: 该参数不能与 `date` 同时指定, 示例值: "1602504000"
	Timezone  *string `json:"timezone,omitempty"`  // 时区。使用 IANA Time Zone Database 标准, 例如 Asia/Shanghai, 全天日程时区固定为UTC +0, 非全天日程时区默认为 Asia/Shanghai, 示例值: "Asia/Shanghai"
}

// UpdateCalendarEventReqLocation ...
type UpdateCalendarEventReqLocation struct {
	Name      *string  `json:"name,omitempty"`      // 地点名称, 示例值: "地点名称", 长度范围: `1` ～ `512` 字符
	Address   *string  `json:"address,omitempty"`   // 地点地址, 示例值: "地点地址", 长度范围: `1` ～ `255` 字符
	Latitude  *float64 `json:"latitude,omitempty"`  // 地点坐标纬度信息, 对于国内的地点, 采用 GCJ-02 标准, 对于海外的地点, 采用 WGS84 标准, 示例值: 1.100000023841858
	Longitude *float64 `json:"longitude,omitempty"` // 地点坐标经度信息, 对于国内的地点, 采用 GCJ-02 标准, 对于海外的地点, 采用 WGS84 标准, 示例值: 2.200000047683716
}

// UpdateCalendarEventReqReminder ...
type UpdateCalendarEventReqReminder struct {
	Minutes *int64 `json:"minutes,omitempty"` // 日程提醒时间的偏移量, 正数时表示在日程开始前 X 分钟提醒, 负数时表示在日程开始后 X 分钟提醒, 注意: 更新日程时修改该字段仅对当前身份生效, 示例值: 5, 取值范围: `-20160` ～ `20160`
}

// UpdateCalendarEventReqSchema ...
type UpdateCalendarEventReqSchema struct {
	UiName   *string `json:"ui_name,omitempty"`   // UI 名称, 可选值有: ForwardIcon: 日程转发按钮, MeetingChatIcon: 会议群聊按钮, MeetingMinutesIcon: 会议纪要按钮, MeetingVideo: 视频会议区域, RSVP: 接受、拒绝、待定区域, Attendee: 参与者区域, OrganizerOrCreator: 组织者或创建者区域, 示例值: "ForwardIcon"
	UiStatus *string `json:"ui_status,omitempty"` // UI 项的状态。目前只支持选择 `hide`, 示例值: "hide", 可选值有: hide: 隐藏显示, readonly: 只读, editable: 可编辑, unknown: 未知 UI 项自定义状态。该参数仅用于读取时兼容, 不支持作为请求参数值传入
	AppLink  *string `json:"app_link,omitempty"`  // 按钮点击后跳转的链接, 注意: 兼容性参数, 只读, 因此暂不支持传入该请求参数, 示例值: "https://applink.feishu.cn/client/calendar/event/detail?calendarId=xxxxxx&key=xxxxxx&originalTime=xxxxxx&startTime=xxxxxx", 最大长度: `2000` 字符
}

// UpdateCalendarEventReqStartTime ...
type UpdateCalendarEventReqStartTime struct {
	Date      *string `json:"date,omitempty"`      // 开始时间, 仅全天日程使用该字段, [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339) 格式, 例如, 2018-09-01, 注意: 该参数不能与 `timestamp` 同时指定, 示例值: "2018-09-01"
	Timestamp *string `json:"timestamp,omitempty"` // 秒级时间戳, 用于设置具体的开始时间。例如, 1602504000 表示 2020/10/12 20:00:00（UTC +8 时区）, 注意: 该参数不能与 `date` 同时指定, 示例值: "1602504000"
	Timezone  *string `json:"timezone,omitempty"`  // 时区。使用 IANA Time Zone Database 标准, 例如 Asia/Shanghai, 全天日程时区固定为UTC +0, 非全天日程时区默认为 Asia/Shanghai, 示例值: "Asia/Shanghai"
}

// UpdateCalendarEventReqVchat ...
type UpdateCalendarEventReqVchat struct {
	VCType          *string                                     `json:"vc_type,omitempty"`          // 视频会议类型。如果无需视频会议, 则必须传入 `no_meeting`, 示例值: "third_party", 可选值有: vc: 飞书视频会议。取该类型时, vchat 内的其他字段均无效。, third_party: 第三方链接视频会议。取该类型时, 仅生效 vchat 内的 icon_type、description、meeting_url 字段。, no_meeting: 无视频会议。取该类型时, vchat 内的其他字段均无效。, lark_live: 飞书直播。该值用于客户端, 不支持通过 API 调用, 只读。, unknown: 未知类型。该值用于客户端做兼容使用, 不支持通过 API 调用, 只读。
	IconType        *string                                     `json:"icon_type,omitempty"`        // 第三方视频会议的 icon 类型, 默认值: 空, 表示不更新该字段, 示例值: "vc", 可选值有: vc: 飞书视频会议 icon, live: 直播视频会议 icon, default: 默认 icon
	Description     *string                                     `json:"description,omitempty"`      // 第三方视频会议文案, 默认值: 空, 表示不更新该字段, 示例值: "发起视频会议", 长度范围: `0` ～ `500` 字符
	MeetingURL      *string                                     `json:"meeting_url,omitempty"`      // 视频会议 URL, 默认值: 空, 表示不更新该字段, 示例值: "https://example.com", 长度范围: `1` ～ `2000` 字符
	MeetingSettings *UpdateCalendarEventReqVchatMeetingSettings `json:"meeting_settings,omitempty"` // 飞书视频会议（VC）的会前设置, 需满足以下全部条件: 当 `vc_type` 为 `vc` 时生效, 需要有日程的编辑权限, 不传值则表示不更新该字段。
}

// UpdateCalendarEventReqVchatMeetingSettings ...
type UpdateCalendarEventReqVchatMeetingSettings struct {
	OwnerID               *string  `json:"owner_id,omitempty"`                // 设置会议 owner 的用户 ID, ID 类型需和 user_id_type 保持一致, 该参数需满足以下全部条件才会生效: 应用身份（tenant_access_token）请求, 且在应用日历上操作日程, 首次将日程设置为 VC 会议时, 才能设置owner, owner 不能为非用户身份, owner 不能为外部租户用户身份, 示例值: "ou_7d8a6e6df7621556ce0d21922b676706ccs"
	JoinMeetingPermission *string  `json:"join_meeting_permission,omitempty"` // 设置入会范围, 示例值: "only_organization_employees", 可选值有: anyone_can_join: 所有人可以加入会议, only_organization_employees: 仅企业内用户可以加入会议, only_event_attendees: 仅日程参与者可以加入会议
	AssignHosts           []string `json:"assign_hosts,omitempty"`            // 通过用户 ID 指定主持人, ID 类型需和 user_id, _type 保持一致, 注意: 仅日程组织者可以指定主持人, 主持人不能是非用户身份, 主持人不能是外部租户用户身份, 在应用日历上操作日程时, 不允许指定主持人, 示例值: ["ou_7d8a6e6df7621556ce0d21922b676706ccs"], 最大长度: `10`
	AutoRecord            *bool    `json:"auto_record,omitempty"`             // 是否开启自动录制, 可选值有: true: 开启, false: 不开启, 默认值: 空, 表示不更新该字段, 示例值: false
	OpenLobby             *bool    `json:"open_lobby,omitempty"`              // 是否开启等候室, 可选值有: true: 开启, false: 不开启, 默认值: 空, 表示不更新该字段, 示例值: true
	AllowAttendeesStart   *bool    `json:"allow_attendees_start,omitempty"`   // 是否允许日程参与者发起会议, 注意: 应用日历上操作日程时, 该字段必须为 true, 否则没有人能发起会议, 可选值有: true: 允许, false: 不允许, 默认值: 空, 表示不更新该字段, 示例值: true
}

// UpdateCalendarEventResp ...
type UpdateCalendarEventResp struct {
	Event *UpdateCalendarEventRespEvent `json:"event,omitempty"` // 更新后的日程实体信息。
}

// UpdateCalendarEventRespEvent ...
type UpdateCalendarEventRespEvent struct {
	EventID             string                                      `json:"event_id,omitempty"`              // 日程 ID。后续可通过该 ID 查询、更新或删除日程信息。更多信息参见[日程 ID 说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/introduction)。
	OrganizerCalendarID string                                      `json:"organizer_calendar_id,omitempty"` // 该日程组织者的日历 ID。关于日历 ID 的说明可参见[日历 ID 说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/introduction)。
	Summary             string                                      `json:"summary,omitempty"`               // 日程标题。
	Description         string                                      `json:"description,omitempty"`           // 日程描述。
	NeedNotification    bool                                        `json:"need_notification,omitempty"`     // 更新日程是否给日程参与人发送 Bot 通知。
	StartTime           *UpdateCalendarEventRespEventStartTime      `json:"start_time,omitempty"`            // 日程开始时间。
	EndTime             *UpdateCalendarEventRespEventEndTime        `json:"end_time,omitempty"`              // 日程结束时间。
	Vchat               *UpdateCalendarEventRespEventVchat          `json:"vchat,omitempty"`                 // 视频会议信息。
	Visibility          string                                      `json:"visibility,omitempty"`            // 日程公开范围, 仅对当前身份生效, 可选值有: default: 默认权限, 跟随日历权限, 即默认仅向他人显示是否忙碌, public: 公开, 显示日程详情, private: 私密, 仅自己可见详情
	AttendeeAbility     string                                      `json:"attendee_ability,omitempty"`      // 参与人权限, 可选值有: none: 无法编辑日程、无法邀请其他参与人、无法查看参与人列表, can_see_others: 无法编辑日程、无法邀请其他参与人、可以查看参与人列表, can_invite_others: 无法编辑日程、可以邀请其他参与人、可以查看参与人列表, can_modify_event: 可以编辑日程、可以邀请其他参与人、可以查看参与人列表
	FreeBusyStatus      string                                      `json:"free_busy_status,omitempty"`      // 日程占用的忙闲状态, 仅对当前身份生效, 可选值有: busy: 忙碌, free: 空闲
	Location            *UpdateCalendarEventRespEventLocation       `json:"location,omitempty"`              // 日程地点。
	Color               int64                                       `json:"color,omitempty"`                 // 日程颜色, 由颜色 RGB 值的 int32 表示, 说明: 仅对当前身份生效, 取值为 0 或 -1 时, 表示默认跟随日历颜色, 客户端展示时会映射到色板上最接近的一种颜色。
	Reminders           []*UpdateCalendarEventRespEventReminder     `json:"reminders,omitempty"`             // 日程提醒列表。
	Recurrence          string                                      `json:"recurrence,omitempty"`            // 重复日程的重复性规则, 规则格式可参见 [rfc5545](https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10)。
	Status              string                                      `json:"status,omitempty"`                // 日程状态, 可选值有: tentative: 未回应, confirmed: 已确认, cancelled: 日程已取消
	IsException         bool                                        `json:"is_exception,omitempty"`          // 日程是否是一个重复日程的例外日程。了解例外日程, 可参见[例外日程](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/introduction#71c5ec78)。
	RecurringEventID    string                                      `json:"recurring_event_id,omitempty"`    // 例外日程对应的原重复日程的 event_id。
	CreateTime          string                                      `json:"create_time,omitempty"`           // 日程的创建时间（秒级时间戳）。
	Schemas             []*UpdateCalendarEventRespEventSchema       `json:"schemas,omitempty"`               // 日程自定义信息, 控制日程详情页的 UI 展示。
	EventOrganizer      *UpdateCalendarEventRespEventEventOrganizer `json:"event_organizer,omitempty"`       // 日程组织者信息。
	AppLink             string                                      `json:"app_link,omitempty"`              // 日程的 app_link, 用于跳转到具体的某个日程。
}

// UpdateCalendarEventRespEventEndTime ...
type UpdateCalendarEventRespEventEndTime struct {
	Date      string `json:"date,omitempty"`      // 结束时间, 仅全天日程使用该字段, [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339) 格式, 例如, 2018-09-01。
	Timestamp string `json:"timestamp,omitempty"` // 秒级时间戳, 指日程具体的结束时间。例如, 1602504000 表示 2020/10/12 20:00:00（UTC +8 时区）。
	Timezone  string `json:"timezone,omitempty"`  // 时区。使用 IANA Time Zone Database 标准。
}

// UpdateCalendarEventRespEventEventOrganizer ...
type UpdateCalendarEventRespEventEventOrganizer struct {
	UserID      string `json:"user_id,omitempty"`      // 日程组织者 user ID。
	DisplayName string `json:"display_name,omitempty"` // 日程组织者姓名。
}

// UpdateCalendarEventRespEventLocation ...
type UpdateCalendarEventRespEventLocation struct {
	Name      string  `json:"name,omitempty"`      // 地点名称。
	Address   string  `json:"address,omitempty"`   // 地点地址。
	Latitude  float64 `json:"latitude,omitempty"`  // 地点坐标纬度信息, 对于国内的地点, 采用 GCJ-02 标准, 对于海外的地点, 采用 WGS84 标准
	Longitude float64 `json:"longitude,omitempty"` // 地点坐标经度信息, 对于国内的地点, 采用 GCJ-02 标准, 对于海外的地点, 采用 WGS84 标准
}

// UpdateCalendarEventRespEventReminder ...
type UpdateCalendarEventRespEventReminder struct {
	Minutes int64 `json:"minutes,omitempty"` // 日程提醒时间的偏移量。该参数仅对当前身份生效, 正数时表示在日程开始前 X 分钟提醒, 负数时表示在日程开始后 X 分钟提醒。
}

// UpdateCalendarEventRespEventSchema ...
type UpdateCalendarEventRespEventSchema struct {
	UiName   string `json:"ui_name,omitempty"`   // UI 名称。可能值: ForwardIcon: 日程转发按钮, MeetingChatIcon: 会议群聊按钮, MeetingMinutesIcon: 会议纪要按钮, MeetingVideo: 视频会议区域, RSVP: 接受、拒绝、待定区域, Attendee: 参与者区域, OrganizerOrCreator: 组织者或创建者区域
	UiStatus string `json:"ui_status,omitempty"` // UI项自定义状态, 可选值有: hide: 隐藏显示, readonly: 只读, editable: 可编辑, unknown: 未知 UI 项自定义状态, 仅用于读取时兼容
	AppLink  string `json:"app_link,omitempty"`  // 按钮点击后跳转的链接。
}

// UpdateCalendarEventRespEventStartTime ...
type UpdateCalendarEventRespEventStartTime struct {
	Date      string `json:"date,omitempty"`      // 开始时间, 仅全天日程使用该字段, [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339) 格式, 例如, 2018-09-01。
	Timestamp string `json:"timestamp,omitempty"` // 秒级时间戳, 指日程具体的开始时间。例如, 1602504000 表示 2020/10/12 20:00:00（UTC +8 时区）。
	Timezone  string `json:"timezone,omitempty"`  // 时区。使用 IANA Time Zone Database 标准。
}

// UpdateCalendarEventRespEventVchat ...
type UpdateCalendarEventRespEventVchat struct {
	VCType          string                                            `json:"vc_type,omitempty"`          // 视频会议类型, 可以为空, 表示在首次添加日程参与人时, 会自动生成飞书视频会议 URL, 可选值有: vc: 飞书视频会议。, third_party: 第三方链接视频会议。, no_meeting: 无视频会议。, lark_live: 飞书直播, 只读参数。, unknown: 未知类型, 用于兼容的只读参数。
	IconType        string                                            `json:"icon_type,omitempty"`        // 第三方视频会议 icon 类型, 可以为空, 表示展示默认 icon, 可选值有: vc: 飞书视频会议 icon。, live: 直播视频会议 icon。, default: 默认 icon。
	Description     string                                            `json:"description,omitempty"`      // 第三方视频会议文案。
	MeetingURL      string                                            `json:"meeting_url,omitempty"`      // 视频会议 URL。
	MeetingSettings *UpdateCalendarEventRespEventVchatMeetingSettings `json:"meeting_settings,omitempty"` // 飞书视频会议（VC）的会前设置。
}

// UpdateCalendarEventRespEventVchatMeetingSettings ...
type UpdateCalendarEventRespEventVchatMeetingSettings struct {
	OwnerID               string   `json:"owner_id,omitempty"`                // 会议 owner 的用户 ID 信息。
	JoinMeetingPermission string   `json:"join_meeting_permission,omitempty"` // 设置入会范围, 可选值有: anyone_can_join: 所有人可以加入会议, only_organization_employees: 仅企业内用户可以加入会议, only_event_attendees: 仅日程参与者可以加入会议
	AssignHosts           []string `json:"assign_hosts,omitempty"`            // 主持人的用户 ID 信息。
	AutoRecord            bool     `json:"auto_record,omitempty"`             // 是否开启自动录制。
	OpenLobby             bool     `json:"open_lobby,omitempty"`              // 是否开启等候室。
	AllowAttendeesStart   bool     `json:"allow_attendees_start,omitempty"`   // 是否允许日程参与者发起会议。
}

// updateCalendarEventResp ...
type updateCalendarEventResp struct {
	Code  int64                    `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                   `json:"msg,omitempty"`  // 错误描述
	Data  *UpdateCalendarEventResp `json:"data,omitempty"`
	Error *ErrorDetail             `json:"error,omitempty"`
}
