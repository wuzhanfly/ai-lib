// Copyright 2024 wuzhanfly <wuzhanfly@gmail.com>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/wuzhanfly/ai-lib/app/dao/internal"
)

// walletFlowBalanceDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type walletFlowBalanceDao struct {
	*internal.WalletFlowBalanceDao
}

var (
	// WalletFlowBalance is globally public accessible object for table c_wallet_flow_balance operations.
	WalletFlowBalance walletFlowBalanceDao
)

func init() {
	WalletFlowBalance = walletFlowBalanceDao{
		internal.NewWalletFlowBalanceDao(),
	}
}

// Fill with you ideas below.
