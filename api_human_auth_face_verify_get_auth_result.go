// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetFaceVerifyAuthResult
//
// ::: note
// 无源人脸比对接入需申请白名单，接入前请联系飞书开放平台工作人员，邮箱：openplatform@bytedance.com。
// :::
// 无源人脸比对流程，开发者后台通过调用此接口请求飞书后台，对本次活体比对结果做校验。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/human_authentication-v1/face/query-recognition-result
func (r *HumanAuthAPI) GetFaceVerifyAuthResult(ctx context.Context, request *GetFaceVerifyAuthResultReq, options ...MethodOptionFunc) (*GetFaceVerifyAuthResultResp, *Response, error) {
	r.cli.logInfo(ctx, "[lark] HumanAuth#GetFaceVerifyAuthResult call api")
	r.cli.logDebug(ctx, "[lark] HumanAuth#GetFaceVerifyAuthResult request: %s", jsonString(request))

	if r.cli.mock.mockHumanAuthGetFaceVerifyAuthResult != nil {
		r.cli.logDebug(ctx, "[lark] HumanAuth#GetFaceVerifyAuthResult mock enable")
		return r.cli.mock.mockHumanAuthGetFaceVerifyAuthResult(ctx, request, options...)
	}

	req := &RawRequestReq{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/face_verify/v1/query_auth_result",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getFaceVerifyAuthResultResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark] HumanAuth#GetFaceVerifyAuthResult GET https://open.feishu.cn/open-apis/face_verify/v1/query_auth_result failed: %s", err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark] HumanAuth#GetFaceVerifyAuthResult GET https://open.feishu.cn/open-apis/face_verify/v1/query_auth_result failed, code: %d, msg: %s", resp.Code, resp.Msg)
		return nil, response, NewError("HumanAuth", "GetFaceVerifyAuthResult", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark] HumanAuth#GetFaceVerifyAuthResult request_id: %s, response: %s", response.RequestID, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockHumanAuthGetFaceVerifyAuthResult(f func(ctx context.Context, request *GetFaceVerifyAuthResultReq, options ...MethodOptionFunc) (*GetFaceVerifyAuthResultResp, *Response, error)) {
	r.mockHumanAuthGetFaceVerifyAuthResult = f
}

func (r *Mock) UnMockHumanAuthGetFaceVerifyAuthResult() {
	r.mockHumanAuthGetFaceVerifyAuthResult = nil
}

type GetFaceVerifyAuthResultReq struct {
	ReqOrderNo string  `query:"req_order_no" json:"-"` // 人脸识别单次唯一标识，由`tt.startFaceVerify`接口返回
	OpenID     *string `query:"open_id" json:"-"`      // 用户应用标识, 与employee_id二选其一
	EmployeeID *string `query:"employee_id" json:"-"`  // 用户租户标识, 与open_id二选其一
}

type getFaceVerifyAuthResultResp struct {
	Code int                          `json:"code,omitempty"` // 返回码，非0为失败
	Msg  string                       `json:"msg,omitempty"`  // 返回信息，返回码的描述
	Data *GetFaceVerifyAuthResultResp `json:"data,omitempty"` // 业务数据
}

type GetFaceVerifyAuthResultResp struct {
	AuthState     int `json:"auth_state,omitempty"`     // 认证结果, 0: 认证中, 1: 认证成功, -1: 认证失败
	AuthTimpstamp int `json:"auth_timpstamp,omitempty"` // 认证时间，unix 时间戳
}
