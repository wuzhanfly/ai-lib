// Copyright 2024 wuzhanfly <wuzhanfly@gmail.com>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// ConfigDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type ConfigDao struct {
	gmvc.M                // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	DB      gdb.DB        // DB is the raw underlying database management object.
	Table   string        // Table is the table name of the DAO.
	Columns configColumns // Columns contains all the columns of Table that for convenient usage.
}

// ConfigColumns defines and stores column names for table c_config.
type configColumns struct {
	ConfigName string // 配置参数名
	Title      string // 标题
	Unit       string // 单位
	InputType  string // 表单类型
	Options    string // 参数配置的选项
	Value      string // 配置值
	Type       string // 类型
	Sort       string // 排序
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
}

func NewConfigDao() *ConfigDao {
	return &ConfigDao{
		M:     g.DB("default").Model("c_config").Safe(),
		DB:    g.DB("default"),
		Table: "c_config",
		Columns: configColumns{
			ConfigName: "config_name",
			Title:      "title",
			Unit:       "unit",
			InputType:  "input_type",
			Options:    "options",
			Value:      "value",
			Type:       "type",
			Sort:       "sort",
			CreatedAt:  "created_at",
			UpdatedAt:  "updated_at",
		},
	}
}
