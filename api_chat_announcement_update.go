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

// UpdateChatAnnouncement 更新会话中的群公告信息, 更新公告信息的格式和更新[旧版云文档](https://open.feishu.cn/document/ukTMukTMukTM/uAzM5YjLwMTO24CMzkjN)格式相同, 不支持新版文档格式。
//
// 注意事项:
// - 应用需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)
// - 机器人或授权用户必须在群里
// - 操作者需要拥有群公告文档的阅读权限
// - 获取内部群信息时, 操作者须与群组在同一租户下
// - 若群开启了 [仅群主和群管理员可编辑群信息] 配置, 群主/群管理员 或 创建群组且具备 [更新应用所创建群的群信息] 权限的机器人, 可更新群公告
// - 若群未开启 [仅群主和群管理员可编辑群信息] 配置, 所有成员可以更新群公告
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-announcement/patch
// new doc: https://open.feishu.cn/document/server-docs/group/chat-announcement/patch
func (r *ChatService) UpdateChatAnnouncement(ctx context.Context, request *UpdateChatAnnouncementReq, options ...MethodOptionFunc) (*UpdateChatAnnouncementResp, *Response, error) {
	if r.cli.mock.mockChatUpdateChatAnnouncement != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Chat#UpdateChatAnnouncement mock enable")
		return r.cli.mock.mockChatUpdateChatAnnouncement(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Chat",
		API:                   "UpdateChatAnnouncement",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v1/chats/:chat_id/announcement",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(updateChatAnnouncementResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockChatUpdateChatAnnouncement mock ChatUpdateChatAnnouncement method
func (r *Mock) MockChatUpdateChatAnnouncement(f func(ctx context.Context, request *UpdateChatAnnouncementReq, options ...MethodOptionFunc) (*UpdateChatAnnouncementResp, *Response, error)) {
	r.mockChatUpdateChatAnnouncement = f
}

// UnMockChatUpdateChatAnnouncement un-mock ChatUpdateChatAnnouncement method
func (r *Mock) UnMockChatUpdateChatAnnouncement() {
	r.mockChatUpdateChatAnnouncement = nil
}

// UpdateChatAnnouncementReq ...
type UpdateChatAnnouncementReq struct {
	ChatID   string   `path:"chat_id" json:"-"`   // 待修改公告的群 ID, 详情参见[群ID 说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-id-description), 注意: 不支持P2P单聊, 示例值: "oc_5ad11d72b830411d72b836c20"
	Revision string   `json:"revision,omitempty"` // 文档当前版本号 int64 类型, [获取群公告信息](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-announcement/get)接口会返回, 注意: 传入的版本号和最新版本号的差距不能超过100, 示例值: "12"
	Requests []string `json:"requests,omitempty"` // 修改文档请求的序列化字段, 更新公告信息的格式和更新[云文档](https://open.feishu.cn/document/ukTMukTMukTM/uYDM2YjL2AjN24iNwYjN)格式相同, 示例值: ["{\"requestType\":\"InsertBlocksRequestType\", \"insertBlocksRequest\":{\"payload\":\"{\\\"blocks\\\":[{\\\"type\\\":\\\"paragraph\\\", \\\"paragraph\\\":{\\\"elements\\\":[{\\\"type\\\":\\\"textRun\\\", \\\"textRun\\\":{\\\"text\\\":\\\"Docs API Sample Content\\\", \\\"style\\\":{}}}], \\\"style\\\":{}}}]}\", \"location\":{\"zoneId\":\"0\", \"index\":0, \"endOfZone\": true}}}"]
}

// UpdateChatAnnouncementResp ...
type UpdateChatAnnouncementResp struct {
}

// updateChatAnnouncementResp ...
type updateChatAnnouncementResp struct {
	Code int64                       `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                      `json:"msg,omitempty"`  // 错误描述
	Data *UpdateChatAnnouncementResp `json:"data,omitempty"`
}
