// Copyright 2015 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package core

import (
	"fmt"
)

type Tree struct {
	Docs    map[string][]*doc // 按apiGroup分组的文档结构。
	Version string            // 程序的版本号
}

// 表示一个api文档。
type doc struct {
	Group       string    // 所属分组
	Version     string    // 版本号
	Methods     string    // 请求的方法，GET，POST等
	URL         string    // 请求地址
	Summary     string    // 简要描述
	Description string    // 详细描述
	Queries     []*param  // 查询参数
	Params      []*param  // URL参数
	Request     *request  // 若是GET，则使用此描述请求的具体数据
	Status      []*status // 各种状态码下返回的数据描述
}

type request struct {
	Type     string            // 请求的类型，xml或是json
	Headers  map[string]string // 请求必须携带的头
	Params   []*param          //提交的各个字段的描述
	Examples []*example
}

// 表示一次请求或是返回的数据。
type status struct {
	Code     string            // 状态码
	Summary  string            // 该状态下的简要描述
	Headers  map[string]string // 必须提交的头信息或是返回的头信息。
	Params   []*param          // 提交或是返回数据的各个字段描述
	Examples []*example        // 提交或是返回数据的示例，键名为数据类型，键值为示例代码
}

// 用于描述提交和返回的参数信息。
type param struct {
	Name        string // 参数名称
	Type        string // 类型
	Description string // 参数介绍
}

// 示例代码
type example struct {
	Type string // 示例代码的类型，小写，xml或是json
	Code string // 示例代码
}

// 语法法错误，语法比较简单，仅包含了在第几行发生错误。
type SyntaxError struct {
	line int
	file string
}

func (err *SyntaxError) Error() string {
	return fmt.Sprint("在%v文件的第%v行发生语法错误", err.file, err.line)
}

func NewTree() *Tree {
	return &Tree{
		Docs: map[string][]*doc{},
	}
}
