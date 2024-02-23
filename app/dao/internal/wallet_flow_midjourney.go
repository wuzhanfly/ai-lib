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

// WalletFlowMidjourneyDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type WalletFlowMidjourneyDao struct {
	gmvc.M                              // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	DB      gdb.DB                      // DB is the raw underlying database management object.
	Table   string                      // Table is the table name of the DAO.
	Columns walletFlowMidjourneyColumns // Columns contains all the columns of Table that for convenient usage.
}

// WalletFlowMidjourneyColumns defines and stores column names for table c_wallet_flow_midjourney.
type walletFlowMidjourneyColumns struct {
	Id         string // ID
	UserId     string // 会员ID
	Amount     string // 变动金额
	Total      string // 变动后的余额
	IsIncr     string // 增加减少
	TargetType string // 目标类型
	TargetId   string // 目标ID
	Remark     string // 备注
	AdminName  string // 操作管理员
	Year       string // 年
	Month      string // 月
	Day        string // 日
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
}

func NewWalletFlowMidjourneyDao() *WalletFlowMidjourneyDao {
	return &WalletFlowMidjourneyDao{
		M:     g.DB("default").Model("c_wallet_flow_midjourney").Safe(),
		DB:    g.DB("default"),
		Table: "c_wallet_flow_midjourney",
		Columns: walletFlowMidjourneyColumns{
			Id:         "id",
			UserId:     "user_id",
			Amount:     "amount",
			Total:      "total",
			IsIncr:     "is_incr",
			TargetType: "target_type",
			TargetId:   "target_id",
			Remark:     "remark",
			AdminName:  "admin_name",
			Year:       "year",
			Month:      "month",
			Day:        "day",
			CreatedAt:  "created_at",
			UpdatedAt:  "updated_at",
		},
	}
}
