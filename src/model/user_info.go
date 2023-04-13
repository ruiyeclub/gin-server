package model

import (
	"gin-server/src/utils/localtime"
)

type DexUserInfo struct {
	Id            int64               `gorm:"column:id" json:"id" form:"id"`                                    //type:bigint      comment:主键       version:2023-02-22 16:14
	Avatar        string              `gorm:"column:avatar" json:"avatar" form:"avatar"`                        //type:string      comment:头像       version:2023-02-22 16:14
	Nickname      string              `gorm:"column:nickname" json:"nickname" form:"nickname"`                  //type:string      comment:昵称       version:2023-02-22 16:14
	Email         string              `gorm:"column:email" json:"email" form:"email"`                           //type:string      comment:邮箱       version:2023-02-22 16:14
	Twitter       string              `gorm:"column:twitter" json:"twitter" form:"twitter"`                     //type:string      comment:推特       version:2023-02-22 16:14
	Discord       string              `gorm:"column:discord" json:"discord" form:"discord"`                     //type:string      comment:discord    version:2023-02-22 16:14
	Telegram      string              `gorm:"column:telegram" json:"telegram" form:"telegram"`                  //type:string      comment:tg         version:2023-02-22 16:14
	WalletAddress string              `gorm:"column:wallet_address" json:"walletAddress" form:"walletAddress" ` //type:string      comment:用户地址   version:2023-02-22 16:14
	Introduction  string              `gorm:"column:introduction" json:"introduction" form:"introduction"`      //type:text        comment:简介       version:2023-02-22 16:14
	Created       localtime.LocalTime `gorm:"column:created" json:"created" form:"created"`                     //type:timestamp   comment:创建时间   version:2023-02-22 16:14
	Updated       localtime.LocalTime `gorm:"column:updated" json:"updated" form:"updated"`                     //type:timestamp   comment:更新时间   version:2023-02-22 16:14
}

func (dex DexUserInfo) TableName() string {
	return "dex_user_info"
}
