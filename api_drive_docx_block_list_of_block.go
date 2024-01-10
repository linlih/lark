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

// GetDocxBlockListOfBlock 获取文档中指定块的所有子块的富文本内容并分页返回。文档版本号可选。
//
// 应用频率限制: 单个应用调用频率上限为每秒 5 次, 超过该频率限制, 接口将返回 HTTP 状态码 400 及错误码 99991400。当请求被限频, 应用需要处理限频状态码, 并使用指数退避算法或其它一些频控策略降低对 API 的调用速率。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/document-docx/docx-v1/document-block-children/get
// new doc: https://open.feishu.cn/document/server-docs/docs/docs/docx-v1/document-block/get-2
func (r *DriveService) GetDocxBlockListOfBlock(ctx context.Context, request *GetDocxBlockListOfBlockReq, options ...MethodOptionFunc) (*GetDocxBlockListOfBlockResp, *Response, error) {
	if r.cli.mock.mockDriveGetDocxBlockListOfBlock != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#GetDocxBlockListOfBlock mock enable")
		return r.cli.mock.mockDriveGetDocxBlockListOfBlock(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "GetDocxBlockListOfBlock",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/docx/v1/documents/:document_id/blocks/:block_id/children",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getDocxBlockListOfBlockResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveGetDocxBlockListOfBlock mock DriveGetDocxBlockListOfBlock method
func (r *Mock) MockDriveGetDocxBlockListOfBlock(f func(ctx context.Context, request *GetDocxBlockListOfBlockReq, options ...MethodOptionFunc) (*GetDocxBlockListOfBlockResp, *Response, error)) {
	r.mockDriveGetDocxBlockListOfBlock = f
}

// UnMockDriveGetDocxBlockListOfBlock un-mock DriveGetDocxBlockListOfBlock method
func (r *Mock) UnMockDriveGetDocxBlockListOfBlock() {
	r.mockDriveGetDocxBlockListOfBlock = nil
}

// GetDocxBlockListOfBlockReq ...
type GetDocxBlockListOfBlockReq struct {
	DocumentID         string  `path:"document_id" json:"-"`           // 文档的唯一标识。点击[这里](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/document-docx/docx-overview)了解如何获取文档的 `document_id`, 示例值: "doxcnePuYufKa49ISjhD8Iabcef"
	BlockID            string  `path:"block_id" json:"-"`              // Block 的唯一标识。你可调用[获取文档所有块](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/document-docx/docx-v1/document-block/list)获取文档中块的 `block_id`, 示例值: "doxcnO6UW6wAw2qIcYf4hZabcef"
	DocumentRevisionID *int64  `query:"document_revision_id" json:"-"` // 查询的文档版本, 1 表示文档最新版本。文档创建后, 版本为 1。若查询的版本为文档最新版本, 则需要持有文档的阅读权限；若查询的版本为文档的历史版本, 则需要持有文档的编辑权限。你可通过调用[获取文档基本信息](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/document-docx/docx-v1/document/get)获取文档的 `revision_id`, 示例值:1, 默认值: `-1`, 最小值: `-1`
	PageToken          *string `query:"page_token" json:"-"`           // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: "aw7DoMKBFMOGwqHCrcO8w6jCmMOvw6ILeADCvsKNw57Di8O5XGV3LG4_w5HCqhFxSnDCrCzCn0BgZcOYUg85EMOYcEAcwqYOw4ojw5QFwofCu8KoIMO3K8Ktw4IuNMOBBHNYw4bCgCV3U1zDu8K-J8KSR8Kgw7Y0fsKZdsKvW3d9w53DnkHDrcO5bDkYwrvDisOEPcOtVFJ-I03CnsOILMOoAmLDknd6dsKqG1bClAjDuS3CvcOTwo7Dg8OrwovDsRdqIcKxw5HDohTDtXN9w5rCkWo"
	PageSize           *int64  `query:"page_size" json:"-"`            // 分页大小, 示例值: 500, 默认值: `500`, 最大值: `500`
	UserIDType         *IDType `query:"user_id_type" json:"-"`         // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
}

// GetDocxBlockListOfBlockResp ...
type GetDocxBlockListOfBlockResp struct {
	Items     []*DocxBlock `json:"items,omitempty"`      // block 的 children 列表
	PageToken string       `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	HasMore   bool         `json:"has_more,omitempty"`   // 是否还有更多项
}

// getDocxBlockListOfBlockResp ...
type getDocxBlockListOfBlockResp struct {
	Code int64                        `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                       `json:"msg,omitempty"`  // 错误描述
	Data *GetDocxBlockListOfBlockResp `json:"data,omitempty"`
}
