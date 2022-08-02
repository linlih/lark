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

// GetVCExportTask 查看异步导出的进度
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/export/get
func (r *VCService) GetVCExportTask(ctx context.Context, request *GetVCExportTaskReq, options ...MethodOptionFunc) (*GetVCExportTaskResp, *Response, error) {
	if r.cli.mock.mockVCGetVCExportTask != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] VC#GetVCExportTask mock enable")
		return r.cli.mock.mockVCGetVCExportTask(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "VC",
		API:                   "GetVCExportTask",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/vc/v1/exports/:task_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getVCExportTaskResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockVCGetVCExportTask mock VCGetVCExportTask method
func (r *Mock) MockVCGetVCExportTask(f func(ctx context.Context, request *GetVCExportTaskReq, options ...MethodOptionFunc) (*GetVCExportTaskResp, *Response, error)) {
	r.mockVCGetVCExportTask = f
}

// UnMockVCGetVCExportTask un-mock VCGetVCExportTask method
func (r *Mock) UnMockVCGetVCExportTask() {
	r.mockVCGetVCExportTask = nil
}

// GetVCExportTaskReq ...
type GetVCExportTaskReq struct {
	TaskID string `path:"task_id" json:"-"` // 任务id, 示例值: "7108646852144136212"
}

// GetVCExportTaskResp ...
type GetVCExportTaskResp struct {
	Status int64  `json:"status,omitempty"` // 任务状态, 可选值有: 1: 处理中, 2: 失败, 3: 完成
	URL    string `json:"url,omitempty"`    // 文件下载地址
}

// getVCExportTaskResp ...
type getVCExportTaskResp struct {
	Code int64                `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string               `json:"msg,omitempty"`  // 错误描述
	Data *GetVCExportTaskResp `json:"data,omitempty"`
}
