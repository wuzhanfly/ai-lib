// Copyright 2024 wuzhanfly <wuzhanfly@gmail.com>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

// LogOperation is the golang structure for table c_log_operation.
type LogOperation struct {
	Id            int64  `orm:"id,primary"     json:"id"`            // ID
	StatusCode    string `orm:"status_code"    json:"statusCode"`    // 状态码
	Router        string `orm:"router"         json:"router"`        // 请求路径
	RequestHeader string `orm:"request_header" json:"requestHeader"` // 请求头
	Content       string `orm:"content"        json:"content"`       // 操作内容
	AdminName     string `orm:"admin_name"     json:"adminName"`     // 管理员名
	CreatedAt     int    `orm:"created_at"     json:"createdAt"`     // 创建时间
}
