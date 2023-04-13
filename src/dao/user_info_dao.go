package dao

import (
	"gin-server/src/enum"
	"gin-server/src/model"
	"gin-server/src/utils/database"
	"gin-server/src/utils/errors"

	"gorm.io/gorm"
)

var UserInfoDao *userInfoDao

func init() {
	UserInfoDao = &userInfoDao{database.Orm.DB()}
}

type userInfoDao struct {
	db *gorm.DB
}

// GetUserInfo 根据用户钱包地址查询用户信息
func (dao *userInfoDao) GetUserInfo(walletAddress string) *model.DexUserInfo {
	userInfo := new(model.DexUserInfo)
	result := dao.db.Where("wallet_address", walletAddress)
	result = result.First(userInfo)
	if result.Error == nil {
		return userInfo
	} else if result.Error == gorm.ErrRecordNotFound {
		return nil
	} else {
		errors.NewErrorByEnum(enum.DataErr)
		return nil
	}
}

// UpdateUserInfo 更新用户信息,并且返回更新后的用户信息
func (dao *userInfoDao) UpdateUserInfo(userInfo *model.DexUserInfo) {
	result := dao.db.Model(userInfo)
	result.Omit("id", "wallet_address", "created", "updated")
	result.Where("wallet_address", userInfo.WalletAddress)
	result.Updates(userInfo)
	if result.Error != nil {
		errors.NewErrorByEnum(enum.DataErr)
	}
}

// SaveUserInfo 保存用户信息，并且返回保存后的用户信息
func (dao *userInfoDao) SaveUserInfo(userInfo *model.DexUserInfo) {
	result := dao.db.Omit("Created", "Updated").Create(userInfo)
	if result.Error != nil {
		errors.NewErrorByEnum(enum.DataErr)
	}
}
