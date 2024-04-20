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

// CreateTaskCustomField 创建一个自定义字段, 并将其加入一个资源上（目前资源只支持清单）。创建自定义字段必须提供字段名称, 类型和相应类型的设置。
//
// 目前任务自定义字段支持数字(number), 成员(member), 日期(datetime), 单选(single_select), 多选(multi_select), 文本(text)几种类型。分别使用"number_setting", "member_setting", "datetime_setting", "single_select_setting", "multi_select_setting", "text_setting"来设置。
// 例如创建一个数字类型的自定义字段, 并添加到guid为"ec5ed63d-a4a9-44de-a935-7ba243471c0a"的清单, 可以这样发请求。
// ```
// POST /task/v2/custom_fields
// {
// "name": "价格",
// "type": "number",
// "resource_type": "tasklist",
// "resource_id": "ec5ed63d-a4a9-44de-a935-7ba243471c0a",
// "number_setting": {
// "format": "cny",
// "decimal_count": 2,
// "separator": "thousand"
// }
// }
// ```
// 表示创建一个叫做“价格”的自定义字段, 保留两位小数。在界面上显示时采用人民币的格式, 并显示千分位分割符。
// 类似的, 创建一个单选字段, 可以这样调用接口:
// ```
// POST /task/v2/custom_fields
// {
// "name": "优先级",
// "type": "single_select",
// "resource_type": "tasklist",
// "resource_id": "ec5ed63d-a4a9-44de-a935-7ba243471c0a",
// "single_select_setting": {
// "options": [
// {
// "name": "高",
// "color_index": 1
// },
// {
// "name": "中",
// "color_index": 11
// },
// {
// "name": "低",
// "color_index": 16
// }
// ]
// }
// }
// ```
// 表示创建一个叫“优先级”的单选, 包含“高”, “中”, “低”三个选项, 每个选项设置一个颜色值。
// 在一个资源上创建自定义字段需要该资源的可编辑权限。
//
// doc: https://open.larkoffice.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/custom_field/create
func (r *TaskService) CreateTaskCustomField(ctx context.Context, request *CreateTaskCustomFieldReq, options ...MethodOptionFunc) (*CreateTaskCustomFieldResp, *Response, error) {
	if r.cli.mock.mockTaskCreateTaskCustomField != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Task#CreateTaskCustomField mock enable")
		return r.cli.mock.mockTaskCreateTaskCustomField(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Task",
		API:                   "CreateTaskCustomField",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/task/v2/custom_fields",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(createTaskCustomFieldResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockTaskCreateTaskCustomField mock TaskCreateTaskCustomField method
func (r *Mock) MockTaskCreateTaskCustomField(f func(ctx context.Context, request *CreateTaskCustomFieldReq, options ...MethodOptionFunc) (*CreateTaskCustomFieldResp, *Response, error)) {
	r.mockTaskCreateTaskCustomField = f
}

// UnMockTaskCreateTaskCustomField un-mock TaskCreateTaskCustomField method
func (r *Mock) UnMockTaskCreateTaskCustomField() {
	r.mockTaskCreateTaskCustomField = nil
}

// CreateTaskCustomFieldReq ...
type CreateTaskCustomFieldReq struct {
	UserIDType          *IDType                                      `query:"user_id_type" json:"-"`          // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	ResourceType        string                                       `json:"resource_type,omitempty"`         // 自定义字段要归属的资源类型, 支持"tasklist", 示例值: "tasklist", 默认值: `tasklist`
	ResourceID          string                                       `json:"resource_id,omitempty"`           // 自定义字段要归属的资源ID, 当`resource_type`为"tasklist"时必须填写清单的GUID, 示例值: "ec5ed63d-a4a9-44de-a935-7ba243471c0a", 最大长度: `100` 字符
	Name                string                                       `json:"name,omitempty"`                  // 字段名称, 最大50个字符, 示例值: "优先级"
	Type                string                                       `json:"type,omitempty"`                  // 自定义字段类型, 示例值: "number", 可选值有: number: 数字, datetime: 日期, member: 成员, single_select: 单选, multi_select: 多选, text: 文本
	NumberSetting       *CreateTaskCustomFieldReqNumberSetting       `json:"number_setting,omitempty"`        // 数字类型的字段设置
	MemberSetting       *CreateTaskCustomFieldReqMemberSetting       `json:"member_setting,omitempty"`        // 人员类型的字段设置
	DatetimeSetting     *CreateTaskCustomFieldReqDatetimeSetting     `json:"datetime_setting,omitempty"`      // 时间日期类型的字段设置
	SingleSelectSetting *CreateTaskCustomFieldReqSingleSelectSetting `json:"single_select_setting,omitempty"` // 单选设置
	MultiSelectSetting  *CreateTaskCustomFieldReqMultiSelectSetting  `json:"multi_select_setting,omitempty"`  // 多选设置
	TextSetting         *CreateTaskCustomFieldReqTextSetting         `json:"text_setting,omitempty"`          // 文本类型设置
}

// CreateTaskCustomFieldReqDatetimeSetting ...
type CreateTaskCustomFieldReqDatetimeSetting struct {
	Format *string `json:"format,omitempty"` // 日期时间格式, 支持, yyyy-mm-dd: 以短横分隔的年月日, 例如2023-08-24, yyyy/mm/dd: 以斜杠分隔的年月日, 例如2023/08/04, mm/dd/yyyy: 以斜杠分隔的月日年, 例如08/24/2023, dd/mm/yyyy: 以斜杠分隔的日月年, 例如24/08/2023, 默认为"yyyy-mm-dd", 注意本设置仅影响App中的时间日期类型字段的字段值的显示格式, 并不会影响openapi输入/输出的字段值的格式, 示例值: "yyyy/mm/dd"
}

// CreateTaskCustomFieldReqMemberSetting ...
type CreateTaskCustomFieldReqMemberSetting struct {
	Multi *bool `json:"multi,omitempty"` // 是否支持多选, 默认为false, 示例值: true, 默认值: `false`
}

// CreateTaskCustomFieldReqMultiSelectSetting ...
type CreateTaskCustomFieldReqMultiSelectSetting struct {
	Options []*CreateTaskCustomFieldReqMultiSelectSettingOption `json:"options,omitempty"` // 多选选项, 长度范围: `0` ～ `100`
}

// CreateTaskCustomFieldReqMultiSelectSettingOption ...
type CreateTaskCustomFieldReqMultiSelectSettingOption struct {
	Name       string `json:"name,omitempty"`        // 选项名称, 不能为空, 最大50个字符, 示例值: "高优"
	ColorIndex *int64 `json:"color_index,omitempty"` // 选项的颜色索引值, 可以是0～54中的一个数字。如果不填写则会随机选一个, 示例值: 1, 取值范围: `0` ～ `54`
	IsHidden   *bool  `json:"is_hidden,omitempty"`   // 选项是否隐藏。隐藏后的选项在App界面不可见, 也不可以通过oapi将字段值设为该选项, 示例值: false
}

// CreateTaskCustomFieldReqNumberSetting ...
type CreateTaskCustomFieldReqNumberSetting struct {
	Format               *string `json:"format,omitempty"`                 // 数字类型的自定义字段的值在App展示的格式, 注意本设置仅影响App中的数字类型字段的字段值的显示格式, 并不会影响openapi中输入/输出的字段值的格式, 示例值: "normal", 可选值有: normal: 常规数字, percentage: 百分比格式, cny: 人民币格式, usd: 美元格式, custom: 自定义符号, 默认值: `normal`
	CustomSymbol         *string `json:"custom_symbol,omitempty"`          // 当`format`设为"custom"时, 设置具体的自定义符号, 注意本设置仅影响App中的数字类型字段的字段值的显示格式, 并不会影响openapi输入/输出的字段值的格式, 示例值: "自定义符号"
	CustomSymbolPosition *string `json:"custom_symbol_position,omitempty"` // 当`format`设为"custom"时, 自定义符号相对于数字的显示位置, 注意本设置仅影响App中的数字类型字段的字段值的显示格式, 并不会影响openapi输入/输出的字段值的格式, 示例值: "left", 可选值有: left: 自定义符号显示在数字左边, right: 自定义符号显示在数字右边, 默认值: `right`
	Separator            *string `json:"separator,omitempty"`              // 数字类型自定义字段整数部分的分隔符样式, 注意本设置仅影响App中的数字类型字段的字段值的显示格式, 并不会影响openapi输入/输出的字段值的格式, 示例值: "thousand", 可选值有: none: 无分隔符, thousand: 千分位分隔符, 默认值: `none`
	DecimalCount         *int64  `json:"decimal_count,omitempty"`          // 数字类型自定义字段的值保留的小数位数。多余的位数将被四舍五入, 默认为0, 示例值: 2, 默认值: `0`, 取值范围: `0` ～ `6`
}

// CreateTaskCustomFieldReqSingleSelectSetting ...
type CreateTaskCustomFieldReqSingleSelectSetting struct {
	Options []*CreateTaskCustomFieldReqSingleSelectSettingOption `json:"options,omitempty"` // 单选选项, 长度范围: `0` ～ `100`
}

// CreateTaskCustomFieldReqSingleSelectSettingOption ...
type CreateTaskCustomFieldReqSingleSelectSettingOption struct {
	Name       string `json:"name,omitempty"`        // 选项名称, 不能为空, 最大50个字符, 示例值: "高优"
	ColorIndex *int64 `json:"color_index,omitempty"` // 选项的颜色索引值, 取值0～54。如不填写会自动从未使用的颜色索引值中随机选一个, 示例值: 1, 取值范围: `0` ～ `54`
	IsHidden   *bool  `json:"is_hidden,omitempty"`   // 选项是否隐藏。隐藏后的选项在界面不可见, 也不可以再通过openapi将字段值设为该选项, 示例值: false
}

// CreateTaskCustomFieldResp ...
type CreateTaskCustomFieldResp struct {
	CustomField *CreateTaskCustomFieldRespCustomField `json:"custom_field,omitempty"` // 创建的自定义字段
}

// CreateTaskCustomFieldRespCustomField ...
type CreateTaskCustomFieldRespCustomField struct {
	Guid                string                                                   `json:"guid,omitempty"`                  // 自定义字段的GUID
	Name                string                                                   `json:"name,omitempty"`                  // 自定义字段名称
	Type                string                                                   `json:"type,omitempty"`                  // 自定义字段类型
	NumberSetting       *CreateTaskCustomFieldRespCustomFieldNumberSetting       `json:"number_setting,omitempty"`        // 数字类型的字段设置
	MemberSetting       *CreateTaskCustomFieldRespCustomFieldMemberSetting       `json:"member_setting,omitempty"`        // 人员类型的字段设置
	DatetimeSetting     *CreateTaskCustomFieldRespCustomFieldDatetimeSetting     `json:"datetime_setting,omitempty"`      // 时间日期类型的字段设置
	SingleSelectSetting *CreateTaskCustomFieldRespCustomFieldSingleSelectSetting `json:"single_select_setting,omitempty"` // 单选类型的字段设置
	MultiSelectSetting  *CreateTaskCustomFieldRespCustomFieldMultiSelectSetting  `json:"multi_select_setting,omitempty"`  // 多选类型的字段设置
	Creator             *CreateTaskCustomFieldRespCustomFieldCreator             `json:"creator,omitempty"`               // 创建人
	CreatedAt           string                                                   `json:"created_at,omitempty"`            // 自定义字段创建的时间戳(ms)
	UpdatedAt           string                                                   `json:"updated_at,omitempty"`            // 自定义字段的更新时间戳(ms)
	TextSetting         *CreateTaskCustomFieldRespCustomFieldTextSetting         `json:"text_setting,omitempty"`          // 文本类型设置
}

// CreateTaskCustomFieldRespCustomFieldCreator ...
type CreateTaskCustomFieldRespCustomFieldCreator struct {
	ID   string `json:"id,omitempty"`   // 表示member的id
	Type string `json:"type,omitempty"` // 成员的类型
	Role string `json:"role,omitempty"` // 成员角色
}

// CreateTaskCustomFieldRespCustomFieldDatetimeSetting ...
type CreateTaskCustomFieldRespCustomFieldDatetimeSetting struct {
	Format string `json:"format,omitempty"` // 日期显示格式
}

// CreateTaskCustomFieldRespCustomFieldMemberSetting ...
type CreateTaskCustomFieldRespCustomFieldMemberSetting struct {
	Multi bool `json:"multi,omitempty"` // 是否支持多选
}

// CreateTaskCustomFieldRespCustomFieldMultiSelectSetting ...
type CreateTaskCustomFieldRespCustomFieldMultiSelectSetting struct {
	Options []*CreateTaskCustomFieldRespCustomFieldMultiSelectSettingOption `json:"options,omitempty"` // 选项
}

// CreateTaskCustomFieldRespCustomFieldMultiSelectSettingOption ...
type CreateTaskCustomFieldRespCustomFieldMultiSelectSettingOption struct {
	Guid       string `json:"guid,omitempty"`        // 选项的GUID
	Name       string `json:"name,omitempty"`        // 选项名称, 不能为空, 最大50个字符。
	ColorIndex int64  `json:"color_index,omitempty"` // 选项的颜色索引值。
	IsHidden   bool   `json:"is_hidden,omitempty"`   // 选项是否隐藏。隐藏后的选项在界面不可见, 也不可以再通过openapi将字段值设为该选项。
}

// CreateTaskCustomFieldRespCustomFieldNumberSetting ...
type CreateTaskCustomFieldRespCustomFieldNumberSetting struct {
	Format               string `json:"format,omitempty"`                 // 数字展示的格式, 可选值有: normal: 常规数字, percentage: 百分比格式, cny: 人民币格式, usd: 美元格式, custom: 自定义符号
	CustomSymbol         string `json:"custom_symbol,omitempty"`          // 自定义符号。只有`format`设为custom时才会生效。
	CustomSymbolPosition string `json:"custom_symbol_position,omitempty"` // 自定义符号的位置。默认为right, 可选值有: left: 自定义符号放在数字左边, right: 自定义符号放在数字右边
	Separator            string `json:"separator,omitempty"`              // 分隔符样式, 可选值有: none: 无分隔符, thousand: 千分位分隔符
	DecimalCount         int64  `json:"decimal_count,omitempty"`          // 保留小数位数。输入的数字值的小数位数如果比该设置多, 多余的位数将被四舍五入后舍弃。如果`format`为"percentage", 表示变为百分数之后的小数位数。
}

// CreateTaskCustomFieldRespCustomFieldSingleSelectSetting ...
type CreateTaskCustomFieldRespCustomFieldSingleSelectSetting struct {
	Options []*CreateTaskCustomFieldRespCustomFieldSingleSelectSettingOption `json:"options,omitempty"` // 选项
}

// CreateTaskCustomFieldRespCustomFieldSingleSelectSettingOption ...
type CreateTaskCustomFieldRespCustomFieldSingleSelectSettingOption struct {
	Guid       string `json:"guid,omitempty"`        // 选项的GUID
	Name       string `json:"name,omitempty"`        // 选项名称, 不能为空, 最大50个字符
	ColorIndex int64  `json:"color_index,omitempty"` // 选项的颜色索引值, 取值0～54。如不填写会自动从未使用的颜色索引值中随机选一个。
	IsHidden   bool   `json:"is_hidden,omitempty"`   // 选项是否隐藏。隐藏后的选项在界面不可见, 也不可以再通过openapi将字段值设为该选项。
}

// createTaskCustomFieldResp ...
type createTaskCustomFieldResp struct {
	Code  int64                      `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                     `json:"msg,omitempty"`  // 错误描述
	Data  *CreateTaskCustomFieldResp `json:"data,omitempty"`
	Error *ErrorDetail               `json:"error,omitempty"`
}
