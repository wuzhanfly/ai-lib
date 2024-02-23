// Copyright 2024 wuzhanfly <wuzhanfly@gmail.com>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package midjourney

import (
	"sync"

	"github.com/gogf/gf/os/glog"
	"github.com/wuzhanfly/ai-lib/app/dao"
	"github.com/wuzhanfly/ai-lib/app/model/entity"
)

// MidjourneyClient Midjourney客户端
type MidjourneyClient struct {
}

var midjourneyClient *MidjourneyClient
var midjourneyClientOnce sync.Once

// Instance 单例
func Instance() *MidjourneyClient {
	midjourneyClientOnce.Do(func() {
		midjourneyClient = &MidjourneyClient{}
		if err := midjourneyClient.init(); err != nil {
			glog.Line().Println("MidjourneyClient init失败:" + err.Error())
		}
	})
	return midjourneyClient
}

// init 初始化
func (c *MidjourneyClient) init() (err error) {
	return
}

// GetConfig 获取配置
func (c *MidjourneyClient) GetConfig() (configData *entity.ConfigMidjourney, err error) {
	configData = &entity.ConfigMidjourney{}
	err = dao.ConfigMidjourney.Where("status=1").Limit(1).Order("call_num ASC").Scan(configData)
	if err != nil {
		return nil, err
	}
	return configData, nil
}
