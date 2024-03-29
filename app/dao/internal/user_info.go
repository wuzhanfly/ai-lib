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

// UserInfoDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type UserInfoDao struct {
	gmvc.M                  // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	DB      gdb.DB          // DB is the raw underlying database management object.
	Table   string          // Table is the table name of the DAO.
	Columns userInfoColumns // Columns contains all the columns of Table that for convenient usage.
}

// UserInfoColumns defines and stores column names for table c_user_info.
type userInfoColumns struct {
	UserId   string // ID
	Nickname string // 昵称
	Avatar   string // 头像
}

func NewUserInfoDao() *UserInfoDao {
	return &UserInfoDao{
		M:     g.DB("default").Model("c_user_info").Safe(),
		DB:    g.DB("default"),
		Table: "c_user_info",
		Columns: userInfoColumns{
			UserId:   "user_id",
			Nickname: "nickname",
			Avatar:   "avatar",
		},
	}
}
