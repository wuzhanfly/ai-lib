// Copyright 2024 wuzhanfly <wuzhanfly@gmail.com>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

// Wallet is the golang structure for table c_wallet.
type Wallet struct {
	UserId     int64 `orm:"user_id,primary" json:"userId"`     // 会员ID
	Balance    uint  `orm:"balance"         json:"balance"`    // 余额
	Gpt3       uint  `orm:"gpt3"            json:"gpt3"`       // gpt3提问次数
	Gpt4       uint  `orm:"gpt4"            json:"gpt4"`       // gpt4提问次数
	Midjourney uint  `orm:"midjourney"      json:"midjourney"` // midjourney提问次数
}