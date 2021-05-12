// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetFAQ 该接口用于获取服务台知识库详情。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/faq/get
func (r *HelpdeskAPI) GetFAQ(ctx context.Context, request *GetFAQReq, options ...MethodOptionFunc) (*GetFAQResp, *Response, error) {
	r.cli.logInfo(ctx, "[lark] Helpdesk#GetFAQ call api")
	r.cli.logDebug(ctx, "[lark] Helpdesk#GetFAQ request: %s", jsonString(request))

	if r.cli.mock.mockHelpdeskGetFAQ != nil {
		r.cli.logDebug(ctx, "[lark] Helpdesk#GetFAQ mock enable")
		return r.cli.mock.mockHelpdeskGetFAQ(ctx, request, options...)
	}

	req := &RawRequestReq{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/helpdesk/v1/faqs/:id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedHelpdeskAuth:      true,
	}
	resp := new(getFAQResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark] Helpdesk#GetFAQ GET https://open.feishu.cn/open-apis/helpdesk/v1/faqs/:id failed: %s", err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark] Helpdesk#GetFAQ GET https://open.feishu.cn/open-apis/helpdesk/v1/faqs/:id failed, code: %d, msg: %s", resp.Code, resp.Msg)
		return nil, response, NewError("Helpdesk", "GetFAQ", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark] Helpdesk#GetFAQ request_id: %s, response: %s", response.RequestID, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockHelpdeskGetFAQ(f func(ctx context.Context, request *GetFAQReq, options ...MethodOptionFunc) (*GetFAQResp, *Response, error)) {
	r.mockHelpdeskGetFAQ = f
}

func (r *Mock) UnMockHelpdeskGetFAQ() {
	r.mockHelpdeskGetFAQ = nil
}

type GetFAQReq struct {
	ID string `path:"id" json:"-"` // 知识库ID, 示例值："6856395634652479491"
}

type getFAQResp struct {
	Code int         `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string      `json:"msg,omitempty"`  // 错误描述
	Data *GetFAQResp `json:"data,omitempty"` //
}

type GetFAQResp struct {
	Faq *GetFAQRespFaq `json:"faq,omitempty"` // 知识库详情
}

type GetFAQRespFaq struct {
	FaqID          string                   `json:"faq_id,omitempty"`          // 知识库ID
	ID             string                   `json:"id,omitempty"`              // 知识库旧版ID，请使用faq_id
	HelpdeskID     string                   `json:"helpdesk_id,omitempty"`     // 服务台ID
	Question       string                   `json:"question,omitempty"`        // 问题
	Answer         string                   `json:"answer,omitempty"`          // 答案
	AnswerRichtext string                   `json:"answer_richtext,omitempty"` // 富文本答案
	CreateTime     int                      `json:"create_time,omitempty"`     // 创建时间
	UpdateTime     int                      `json:"update_time,omitempty"`     // 修改时间
	Categories     []*HelpdeskCategory      `json:"categories,omitempty"`      // 分类
	Tags           []string                 `json:"tags,omitempty"`            // 关联词列表
	ExpireTime     int                      `json:"expire_time,omitempty"`     // 失效时间
	UpdateUser     *GetFAQRespFaqUpdateUser `json:"update_user,omitempty"`     // 更新用户
	CreateUser     *GetFAQRespFaqCreateUser `json:"create_user,omitempty"`     // 创建用户
}

type GetFAQRespFaqUpdateUser struct {
	ID        string `json:"id,omitempty"`         // 用户ID
	AvatarUrl string `json:"avatar_url,omitempty"` // 用户头像url
	Name      string `json:"name,omitempty"`       // 用户名
}

type GetFAQRespFaqCreateUser struct {
	ID        string `json:"id,omitempty"`         // 用户ID
	AvatarUrl string `json:"avatar_url,omitempty"` // 用户头像url
	Name      string `json:"name,omitempty"`       // 用户名
}
