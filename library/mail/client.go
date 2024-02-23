// Copyright 2024 wuzhanfly <wuzhanfly@gmail.com>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package mail

import (
	"sync"

	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"github.com/wuzhanfly/ai-lib/library/helper"
)

// MailClient 邮箱
type MailClient struct {
}

var mailClient *MailClient
var mailClientOnce sync.Once

// Instance 单例
func Instance() *MailClient {
	mailClientOnce.Do(func() {
		mailClient = &MailClient{}
		if err := mailClient.init(); err != nil {
			glog.Line().Println("MailClient init失败:" + err.Error())
		}
	})
	return mailClient
}

// init 初始化
func (c *MailClient) init() (err error) {
	return
}

// GetSmtpHost 获取 GetSmtpHost
func (c *MailClient) GetSmtpHost() (string, error) {
	return helper.GetConfig("smtpHost")
}

// GetSmtpPort 获取 GetSmtpPort
func (c *MailClient) GetSmtpPort() (int, error) {
	port, err := helper.GetConfig("smtpPort")
	if err != nil {
		return 0, err
	}
	return gconv.Int(port), nil
}

// GetSmtpEmail 获取 GetSmtpEmail
func (c *MailClient) GetSmtpEmail() (string, error) {
	return helper.GetConfig("smtpEmail")
}

// GetSmtpEmailPassword 获取 GetSmtpEmailPassword
func (c *MailClient) GetSmtpEmailPassword() (string, error) {
	return helper.GetConfig("smtpEmailPassword")
}
