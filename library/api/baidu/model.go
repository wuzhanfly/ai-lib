// Copyright 2024 wuzhanfly <wuzhanfly@gmail.com>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package baidu

import "github.com/wuzhanfly/ai-lib/app/model/entity"

type Config struct {
	entity.ConfigBaidu
	AccessToken         string
	AccessTokenExpireIn int64
}

type TransTextParams struct {
	From string `json:"from"`
	To   string `json:"to"`
	Q    string `json:"q"`
}

type TransTextResponse struct {
	LogId     string                   `json:"log_id"`
	Result    *TransTextResponseResult `json:"result"`
	ErrorMsg  string                   `json:"error_msg"`
	ErrorCode string                   `json:"error_code"`
}

type TransTextResponseResult struct {
	From        string                          `json:"from"`
	To          string                          `json:"to"`
	TransResult []*TransTextResponseTransResult `json:"trans_result"`
}

type TransTextResponseTransResult struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
}
