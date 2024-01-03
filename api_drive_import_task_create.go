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

// CreateDriveImportTask 创建导入任务。支持导入为新版文档、电子表格、多维表格以及旧版文档。该接口为异步接口, 需要通过[查询导入结果](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/import_task/get)接口获取导入结果, 调用方式可参考[导入使用指南](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/import_task/import-user-guide)。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/import_task/create
// new doc: https://open.feishu.cn/document/server-docs/docs/drive-v1/import_task/create
func (r *DriveService) CreateDriveImportTask(ctx context.Context, request *CreateDriveImportTaskReq, options ...MethodOptionFunc) (*CreateDriveImportTaskResp, *Response, error) {
	if r.cli.mock.mockDriveCreateDriveImportTask != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#CreateDriveImportTask mock enable")
		return r.cli.mock.mockDriveCreateDriveImportTask(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "CreateDriveImportTask",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/drive/v1/import_tasks",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(createDriveImportTaskResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveCreateDriveImportTask mock DriveCreateDriveImportTask method
func (r *Mock) MockDriveCreateDriveImportTask(f func(ctx context.Context, request *CreateDriveImportTaskReq, options ...MethodOptionFunc) (*CreateDriveImportTaskResp, *Response, error)) {
	r.mockDriveCreateDriveImportTask = f
}

// UnMockDriveCreateDriveImportTask un-mock DriveCreateDriveImportTask method
func (r *Mock) UnMockDriveCreateDriveImportTask() {
	r.mockDriveCreateDriveImportTask = nil
}

// CreateDriveImportTaskReq ...
type CreateDriveImportTaskReq struct {
	FileExtension string                         `json:"file_extension,omitempty"` // 导入文件格式后缀, 示例值: "xlsx"
	FileToken     string                         `json:"file_token,omitempty"`     // 导入文件Drive FileToken, 示例值: "boxcnxe5OxxxxxxxSNdsJviENsk", 最大长度: `27` 字符
	Type          string                         `json:"type,omitempty"`           // 导入目标云文档格式, 示例值: "sheet"
	FileName      *string                        `json:"file_name,omitempty"`      // 导入目标云文档文件名, 若为空使用Drive文件名, 示例值: "test"
	Point         *CreateDriveImportTaskReqPoint `json:"point,omitempty"`          // 挂载点
}

// CreateDriveImportTaskReqPoint ...
type CreateDriveImportTaskReqPoint struct {
	MountType int64  `json:"mount_type,omitempty"` // 挂载类型, 示例值: 1, 可选值有: 1: 挂载到云空间
	MountKey  string `json:"mount_key,omitempty"`  // 挂载位置, 对于mount_type=1, 云空间目录token, 空表示根目录, 示例值: "fldxxxxxxxx"
}

// CreateDriveImportTaskResp ...
type CreateDriveImportTaskResp struct {
	Ticket string `json:"ticket,omitempty"` // 导入任务ID
}

// createDriveImportTaskResp ...
type createDriveImportTaskResp struct {
	Code int64                      `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 错误描述
	Data *CreateDriveImportTaskResp `json:"data,omitempty"`
}
