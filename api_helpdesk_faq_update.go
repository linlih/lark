// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// UpdateFAQ 该接口用于修改知识库。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/faq/patch
func (r *HelpdeskAPI) UpdateFAQ(ctx context.Context, request *UpdateFAQReq, options ...MethodOptionFunc) (*UpdateFAQResp, *Response, error) {
	r.cli.logInfo(ctx, "[lark] Helpdesk#UpdateFAQ call api")
	r.cli.logDebug(ctx, "[lark] Helpdesk#UpdateFAQ request: %s", jsonString(request))

	if r.cli.mock.mockHelpdeskUpdateFAQ != nil {
		r.cli.logDebug(ctx, "[lark] Helpdesk#UpdateFAQ mock enable")
		return r.cli.mock.mockHelpdeskUpdateFAQ(ctx, request, options...)
	}

	req := &RawRequestReq{
		Method:              "PATCH",
		URL:                 "https://open.feishu.cn/open-apis/helpdesk/v1/faqs/:id",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
		NeedHelpdeskAuth:    true,
	}
	resp := new(updateFAQResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark] Helpdesk#UpdateFAQ PATCH https://open.feishu.cn/open-apis/helpdesk/v1/faqs/:id failed: %s", err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark] Helpdesk#UpdateFAQ PATCH https://open.feishu.cn/open-apis/helpdesk/v1/faqs/:id failed, code: %d, msg: %s", resp.Code, resp.Msg)
		return nil, response, NewError("Helpdesk", "UpdateFAQ", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark] Helpdesk#UpdateFAQ request_id: %s, response: %s", response.RequestID, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockHelpdeskUpdateFAQ(f func(ctx context.Context, request *UpdateFAQReq, options ...MethodOptionFunc) (*UpdateFAQResp, *Response, error)) {
	r.mockHelpdeskUpdateFAQ = f
}

func (r *Mock) UnMockHelpdeskUpdateFAQ() {
	r.mockHelpdeskUpdateFAQ = nil
}

type UpdateFAQReq struct {
	ID  string           `path:"id" json:"-"`   // 知识库ID, 示例值："6856395634652479491"
	Faq *UpdateFAQReqFaq `json:"faq,omitempty"` // 修改的知识库内容
}

type UpdateFAQReqFaq struct {
	CategoryID     *string  `json:"category_id,omitempty"`     // 知识库分类ID, 示例值："6836004780707807251"
	Question       string   `json:"question,omitempty"`        // 问题, 示例值："问题"
	Answer         *string  `json:"answer,omitempty"`          // 答案, 示例值："答案"
	AnswerRichtext *string  `json:"answer_richtext,omitempty"` // 富文本答案和答案必须有一个必填。Json Array格式，富文本结构请见[了解更多: 富文本](/ssl:ttdoc/ukTMukTMukTM/uITM0YjLyEDN24iMxQjN), 示例值："[{,                        "content": "这只是一个测试，医保问题",,                        "type": "text",                    }]"
	Tags           []string `json:"tags,omitempty"`            // 关联词
}

type updateFAQResp struct {
	Code int            `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string         `json:"msg,omitempty"`  // 错误描述
	Data *UpdateFAQResp `json:"data,omitempty"`
}

type UpdateFAQResp struct{}
