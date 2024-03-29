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

// UserDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type UserDao struct {
	gmvc.M              // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	DB      gdb.DB      // DB is the raw underlying database management object.
	Table   string      // Table is the table name of the DAO.
	Columns userColumns // Columns contains all the columns of Table that for convenient usage.
}

// UserColumns defines and stores column names for table c_user.
type userColumns struct {
	Id               string // ID
	Username         string // 用户名
	Password         string // 密码
	LevelId          string // 级别ID
	LevelExpireDate  string // 级别到期日期
	LevelExpireYear  string // 级别到期日期年
	LevelExpireMonth string // 级别到期日期月
	LevelExpireDay   string // 级别到期日期日
	LastLoginAt      string // 最后一次登录时间
	IsBan            string // 是否被禁用
	CreatedAt        string // 创建时间
	UpdatedAt        string // 更新时间
}

func NewUserDao() *UserDao {
	return &UserDao{
		M:     g.DB("default").Model("c_user").Safe(),
		DB:    g.DB("default"),
		Table: "c_user",
		Columns: userColumns{
			Id:               "id",
			Username:         "username",
			Password:         "password",
			LevelId:          "level_id",
			LevelExpireDate:  "level_expire_date",
			LevelExpireYear:  "level_expire_year",
			LevelExpireMonth: "level_expire_month",
			LevelExpireDay:   "level_expire_day",
			LastLoginAt:      "last_login_at",
			IsBan:            "is_ban",
			CreatedAt:        "created_at",
			UpdatedAt:        "updated_at",
		},
	}
}
