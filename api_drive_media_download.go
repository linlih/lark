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
	"io"
)

// DownloadDriveMedia 下载各类云文档中的素材, 例如电子表格中的图片。该接口支持通过在请求头添加`Range` 参数分片下载素材。
//
// * 本接口仅支持下载云文档而非云空间中的资源文件。如要下载云空间中的资源文件, 需调用[下载文件](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/download)接口。
// * 调用此接口之前, 你需确保应用已拥有素材的下载权限。否则接口将返回 403 的 HTTP 状态码。参考[云空间常见问题](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/faq)第五点了解如何分享素材的下载权限给应用。
// * 对于拥有高级权限的多维表格, 在下载素材时, 你需要添加额外的 extra 作为 URL 查询参数, 未填正确填写 extra 接口将返回 403 的 HTTP 状态码。请参考[extra 参数说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/media/introduction#a478a7c3)正确填写 extra 参数。
// 该接口不支持较高并发, 且调用频率上限为 5 QPS。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/media/download
// new doc: https://open.feishu.cn/document/server-docs/docs/drive-v1/media/download
func (r *DriveService) DownloadDriveMedia(ctx context.Context, request *DownloadDriveMediaReq, options ...MethodOptionFunc) (*DownloadDriveMediaResp, *Response, error) {
	if r.cli.mock.mockDriveDownloadDriveMedia != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Drive#DownloadDriveMedia mock enable")
		return r.cli.mock.mockDriveDownloadDriveMedia(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "DownloadDriveMedia",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/drive/v1/medias/:file_token/download",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(downloadDriveMediaResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveDownloadDriveMedia mock DriveDownloadDriveMedia method
func (r *Mock) MockDriveDownloadDriveMedia(f func(ctx context.Context, request *DownloadDriveMediaReq, options ...MethodOptionFunc) (*DownloadDriveMediaResp, *Response, error)) {
	r.mockDriveDownloadDriveMedia = f
}

// UnMockDriveDownloadDriveMedia un-mock DriveDownloadDriveMedia method
func (r *Mock) UnMockDriveDownloadDriveMedia() {
	r.mockDriveDownloadDriveMedia = nil
}

// DownloadDriveMediaReq ...
type DownloadDriveMediaReq struct {
	FileToken string  `path:"file_token" json:"-"` // 素材文件的 token。获取方式如下所示: * 新版文档: 通过[获取文档所有块](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/document-docx/docx-v1/document-block/list)接口获取指定文件块（File Block）或图片块（Image Block）的 token, 即为素材的 token, * 电子表格: 通过[读取多个范围](https://open.feishu.cn/document/ukTMukTMukTM/ukTMzUjL5EzM14SOxMTN)接口获取指定附件的, `fileToken` 参数, 即为素材的 token, * 多维表格: 通过[列出记录](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/list)接口获取指定附件的 `file_token`, 即为素材的 token, 示例值: "boxcnrHpsg1QDqXAAAyachabcef"
	Extra     *string `query:"extra" json:"-"`     // 拥有高级权限的多维表格在下载素材时, 需要添加额外的扩展信息作为 URL 查询参数鉴权。详情参考[extra 参数说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/media/introduction)。未填正确填写该参数的接口将返回 403 的 HTTP 状态码, 示例值: 无
}

// downloadDriveMediaResp ...
type downloadDriveMediaResp struct {
	Code int64                   `json:"code,omitempty"`
	Msg  string                  `json:"msg,omitempty"`
	Data *DownloadDriveMediaResp `json:"data,omitempty"`
}

func (r *downloadDriveMediaResp) SetReader(file io.Reader) {
	if r.Data == nil {
		r.Data = &DownloadDriveMediaResp{}
	}
	r.Data.File = file
}

func (r *downloadDriveMediaResp) SetFilename(filename string) {
	if r.Data == nil {
		r.Data = &DownloadDriveMediaResp{}
	}
	r.Data.Filename = filename
}

// DownloadDriveMediaResp ...
type DownloadDriveMediaResp struct {
	File     io.Reader `json:"file,omitempty"`
	Filename string    `json:"filename,omitempty"`
}
