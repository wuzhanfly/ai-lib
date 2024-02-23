// Copyright 2024 wuzhanfly <wuzhanfly@gmail.com>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

// ShopGoods is the golang structure for table c_shop_goods.
type ShopGoods struct {
	Id                int64  `orm:"id,primary"          json:"id"`                // ID
	Title             string `orm:"title"               json:"title"`             // 商品标题
	Content           string `orm:"content"             json:"content"`           // 商品内容
	FeatItems         string `orm:"feat_items"          json:"featItems"`         // 商品特色条目JSON
	BuyType           int    `orm:"buy_type"            json:"buyType"`           // 购买类型 1购买级别 2购买balance 3购买gpt3 4购买gpt4 5购买midjourney
	ActiveLevelId     int    `orm:"active_level_id"     json:"activeLevelId"`     // 购买的级别
	ActiveExpireType  int    `orm:"active_expire_type"  json:"activeExpireType"`  // 激活有效期类型 0无 1一天 2一月 3一年
	ActiveExpireValue int    `orm:"active_expire_value" json:"activeExpireValue"` // 激活有效期值
	BuyValue          int    `orm:"buy_value"           json:"buyValue"`          // 购买的提问次数的值分单位
	MarketPrice       int    `orm:"market_price"        json:"marketPrice"`       // 市场价
	RealPrice         int    `orm:"real_price"          json:"realPrice"`         // 实际价格
	Status            int    `orm:"status"              json:"status"`            // 是否上架
	Sort              int    `orm:"sort"                json:"sort"`              // 排序
	CreatedAt         int    `orm:"created_at"          json:"createdAt"`         // 创建时间
	UpdatedAt         int    `orm:"updated_at"          json:"updatedAt"`         // 更新时间
}
