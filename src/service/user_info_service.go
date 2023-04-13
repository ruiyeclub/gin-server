package service

import (
	"encoding/json"
	"fmt"
	"gin-server/src/dao"
	"gin-server/src/enum"
	"gin-server/src/global_const"
	"gin-server/src/model"
	"gin-server/src/utils/errors"
	"gin-server/src/utils/redis"
	"github.com/gin-gonic/gin"
)

var UserInfoService *userInfoService

func init() {
	UserInfoService = &userInfoService{}
}

type userInfoService struct {
}

// GetUserInfo 根据钱包地址查询用户信息
func (s *userInfoService) GetUserInfo(walletAddress string) *model.DexUserInfo {
	userInfo := dao.UserInfoDao.GetUserInfo(walletAddress)
	return userInfo
}

// SaveUserInfo 保存用户信息，并且返回保存后的用户信息
func (s *userInfoService) SaveUserInfo(userInfo *model.DexUserInfo) *model.DexUserInfo {

	walletAddress := userInfo.WalletAddress
	dexUserInfo := s.GetUserInfo(walletAddress)
	if dexUserInfo != nil {
		errors.NewErrorByEnum(enum.UserIsExist)
	}
	dao.UserInfoDao.SaveUserInfo(userInfo)
	return s.GetUserInfo(walletAddress)
}

// UpdateUserInfo 更新用户信息,并且返回更新后的用户信息
func (s *userInfoService) UpdateUserInfo(userInfo *model.DexUserInfo) *model.DexUserInfo {
	walletAddress := userInfo.WalletAddress
	dexUserInfo := s.GetUserInfo(walletAddress)
	if dexUserInfo == nil {
		errors.NewErrorByEnum(enum.NotUser)
	}
	dao.UserInfoDao.UpdateUserInfo(userInfo)
	return s.getUserInfoAndSaveRedis(walletAddress)
}

// QueryOrSave 查询不到就保存
func (s *userInfoService) QueryOrSave(walletAddress string) *model.DexUserInfo {
	userInfo := dao.UserInfoDao.GetUserInfo(walletAddress)
	if userInfo == nil {
		newUserInfo := new(model.DexUserInfo)
		newUserInfo.WalletAddress = walletAddress
		dao.UserInfoDao.SaveUserInfo(newUserInfo)
		return newUserInfo
	}
	return userInfo
}

// GetUserInfoByToken 根据token获取用户信息
func (s *userInfoService) GetUserInfoByToken(walletAddress string) *model.DexUserInfo {
	userInfoKey := fmt.Sprintf(global_const.UserInfoKey, walletAddress)
	var userInfo *model.DexUserInfo
	redis.Server.GetStruct(userInfoKey, &userInfo)
	if userInfo == nil {
		userInfo = s.getUserInfoAndSaveRedis(walletAddress)
	}
	return userInfo
}

func (s *userInfoService) getUserInfoAndSaveRedis(walletAddress string) *model.DexUserInfo {
	userInfo := s.GetUserInfo(walletAddress)
	if userInfo == nil {
		errors.NewErrorByEnum(enum.NotUser)
	}
	userInfoByte, err := json.Marshal(userInfo)
	if err == nil {
		userInfoKey := fmt.Sprintf(global_const.UserInfoKey, walletAddress)
		redis.Server.Set(userInfoKey, string(userInfoByte), global_const.UserInfoExpire)
	}

	return userInfo
}

// GetUserInfoByTokenOfContext 根据token获取用户信息 通过上下文
func (s *userInfoService) GetUserInfoByTokenOfContext(c *gin.Context) *model.DexUserInfo {
	walletAddress := c.GetString(global_const.WalletAddress)
	if walletAddress == "" {
		errors.NewErrorByEnum(enum.AuthErr)
	}
	return s.GetUserInfoByToken(walletAddress)
}
