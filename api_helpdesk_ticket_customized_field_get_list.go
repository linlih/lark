// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetTicketCustomizedFieldList
//
// 该接口用于获取全部工单自定义字段。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket_customized_field/list-ticket-customized-fields
func (r *HelpdeskAPI) GetTicketCustomizedFieldList(ctx context.Context, request *GetTicketCustomizedFieldListReq, options ...MethodOptionFunc) (*GetTicketCustomizedFieldListResp, *Response, error) {
	r.cli.logInfo(ctx, "[lark] Helpdesk#GetTicketCustomizedFieldList call api")
	r.cli.logDebug(ctx, "[lark] Helpdesk#GetTicketCustomizedFieldList request: %s", jsonString(request))

	if r.cli.mock.mockHelpdeskGetTicketCustomizedFieldList != nil {
		r.cli.logDebug(ctx, "[lark] Helpdesk#GetTicketCustomizedFieldList mock enable")
		return r.cli.mock.mockHelpdeskGetTicketCustomizedFieldList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/helpdesk/v1/ticket_customized_fields",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedHelpdeskAuth:      true,
	}
	resp := new(getTicketCustomizedFieldListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark] Helpdesk#GetTicketCustomizedFieldList GET https://open.feishu.cn/open-apis/helpdesk/v1/ticket_customized_fields failed: %s", err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark] Helpdesk#GetTicketCustomizedFieldList GET https://open.feishu.cn/open-apis/helpdesk/v1/ticket_customized_fields failed, code: %d, msg: %s", resp.Code, resp.Msg)
		return nil, response, NewError("Helpdesk", "GetTicketCustomizedFieldList", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark] Helpdesk#GetTicketCustomizedFieldList request_id: %s, response: %s", response.RequestID, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockHelpdeskGetTicketCustomizedFieldList(f func(ctx context.Context, request *GetTicketCustomizedFieldListReq, options ...MethodOptionFunc) (*GetTicketCustomizedFieldListResp, *Response, error)) {
	r.mockHelpdeskGetTicketCustomizedFieldList = f
}

func (r *Mock) UnMockHelpdeskGetTicketCustomizedFieldList() {
	r.mockHelpdeskGetTicketCustomizedFieldList = nil
}

type GetTicketCustomizedFieldListReq struct {
	PageToken *string `query:"page_token" json:"-"` // 分页标记，第一次请求不填，表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token，下次遍历可采用该 page_token 获取查询结果, 示例值："6948728206392295444"
	PageSize  *int    `query:"page_size" json:"-"`  // 分页大小, 示例值：10；默认为20, 最大值：`100`
	Visible   *bool   `json:"visible,omitempty"`    // 是否可见, 示例值：true
}

type getTicketCustomizedFieldListResp struct {
	Code int                               `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                            `json:"msg,omitempty"`  // 错误描述
	Data *GetTicketCustomizedFieldListResp `json:"data,omitempty"` //
}

type GetTicketCustomizedFieldListResp struct {
	HasMore       bool                                    `json:"has_more,omitempty"`        // 是否还有更多项
	NextPageToken string                                  `json:"next_page_token,omitempty"` // 下一分页标识
	Items         []*GetTicketCustomizedFieldListRespItem `json:"items,omitempty"`           // 工单自定义字段列表
}

type GetTicketCustomizedFieldListRespItem struct {
	TicketCustomizedFieldID string                                         `json:"ticket_customized_field_id,omitempty"` // 工单自定义字段ID
	HelpdeskID              string                                         `json:"helpdesk_id,omitempty"`                // 服务台ID
	KeyName                 string                                         `json:"key_name,omitempty"`                   // 键名
	DisplayName             string                                         `json:"display_name,omitempty"`               // 名称
	Position                string                                         `json:"position,omitempty"`                   // 字段在列表后台管理列表中的位置
	FieldType               string                                         `json:"field_type,omitempty"`                 // 类型
	Description             string                                         `json:"description,omitempty"`                // 描述
	Visible                 bool                                           `json:"visible,omitempty"`                    // 是否可见
	Editable                bool                                           `json:"editable,omitempty"`                   // 是否可以修改
	Required                bool                                           `json:"required,omitempty"`                   // 是否必填
	CreatedAt               string                                         `json:"created_at,omitempty"`                 // 创建时间
	UpdatedAt               string                                         `json:"updated_at,omitempty"`                 // 更新时间
	CreatedBy               *GetTicketCustomizedFieldListRespItemCreatedBy `json:"created_by,omitempty"`                 // 创建用户
	UpdatedBy               *GetTicketCustomizedFieldListRespItemUpdatedBy `json:"updated_by,omitempty"`                 // 更新用户
	DropdownOptions         *HelpdeskDropdownOption                        `json:"dropdown_options,omitempty"`           // 下拉列表选项
	DropdownAllowMultiple   bool                                           `json:"dropdown_allow_multiple,omitempty"`    // 是否支持多选，仅在字段类型是dropdown的时候有效
}

type GetTicketCustomizedFieldListRespItemCreatedBy struct {
	ID        string `json:"id,omitempty"`         // 用户ID
	AvatarUrl string `json:"avatar_url,omitempty"` // 用户头像url
	Name      string `json:"name,omitempty"`       // 用户名
	Email     string `json:"email,omitempty"`      // 用户邮箱
}

type GetTicketCustomizedFieldListRespItemUpdatedBy struct {
	ID        string `json:"id,omitempty"`         // 用户ID
	AvatarUrl string `json:"avatar_url,omitempty"` // 用户头像url
	Name      string `json:"name,omitempty"`       // 用户名
	Email     string `json:"email,omitempty"`      // 用户邮箱
}
