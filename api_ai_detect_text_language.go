// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// DetectTextLanguage 支持 100 多种语言识别，返回符合 ISO 693-1 标准
//
// 单租户限流：20QPS，同租户下的应用没有限流，共享本租户的 20QPS 限流
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/ai/translation-v1/text/detect
func (r *AIAPI) DetectTextLanguage(ctx context.Context, request *DetectTextLanguageReq, options ...MethodOptionFunc) (*DetectTextLanguageResp, *Response, error) {
	r.cli.logInfo(ctx, "[lark] AI#DetectTextLanguage call api")
	r.cli.logDebug(ctx, "[lark] AI#DetectTextLanguage request: %s", jsonString(request))

	if r.cli.mock.mockAIDetectTextLanguage != nil {
		r.cli.logDebug(ctx, "[lark] AI#DetectTextLanguage mock enable")
		return r.cli.mock.mockAIDetectTextLanguage(ctx, request, options...)
	}

	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/translation/v1/text/detect",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(detectTextLanguageResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark] AI#DetectTextLanguage POST https://open.feishu.cn/open-apis/translation/v1/text/detect failed: %s", err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark] AI#DetectTextLanguage POST https://open.feishu.cn/open-apis/translation/v1/text/detect failed, code: %d, msg: %s", resp.Code, resp.Msg)
		return nil, response, NewError("AI", "DetectTextLanguage", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark] AI#DetectTextLanguage request_id: %s, response: %s", response.RequestID, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockAIDetectTextLanguage(f func(ctx context.Context, request *DetectTextLanguageReq, options ...MethodOptionFunc) (*DetectTextLanguageResp, *Response, error)) {
	r.mockAIDetectTextLanguage = f
}

func (r *Mock) UnMockAIDetectTextLanguage() {
	r.mockAIDetectTextLanguage = nil
}

type DetectTextLanguageReq struct {
	Text string `json:"text,omitempty"` // 需要被识别语种的文本, 示例值："你好"
}

type detectTextLanguageResp struct {
	Code int                     `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 错误描述
	Data *DetectTextLanguageResp `json:"data,omitempty"` //
}

type DetectTextLanguageResp struct {
	Language string `json:"language,omitempty"` // 识别的文本语种，返回符合 ISO 693-1 标准
}
