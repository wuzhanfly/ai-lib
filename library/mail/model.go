// Copyright 2024 wuzhanfly <wuzhanfly@gmail.com>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package mail

type SendMailParams struct {
	ReceiverEmails []string
	MsgSubject     string
	MsgContentType string
	MsgBody        string
}
