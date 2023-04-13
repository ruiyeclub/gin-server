package api

import (
	"gin-server/src/enum"
	"gin-server/src/middleware"
	"gin-server/src/model"
	"gin-server/src/service"
	"gin-server/src/utils/errors"
	"gin-server/src/utils/format"
	"github.com/gin-gonic/gin"
)

func UserAccountApi(rg *gin.Engine) {
	r := rg.Group("/v1/user")

	r.POST("login", login)
	r.POST("refresh", RefreshToken)
	r.GET("address", address)
	// 鉴权获取token
	r.Use(middleware.JWTAuthMiddleware())
	// 根据token获取用户信息
	r.GET("info", getUserInfo)
	// 修改用户信息
	r.POST("update", updateUserInfo)
}

// 获取用户信息接口
func getUserInfo(c *gin.Context) {
	userInfo := service.UserInfoService.GetUserInfoByTokenOfContext(c)
	format.NewApiResult(c).Success(userInfo)
}

// 获取用户信息接口
func address(c *gin.Context) {
	walletAddress := c.Query("walletAddress")
	userInfo := service.UserInfoService.GetUserInfo(walletAddress)
	format.NewApiResult(c).Success(userInfo)
}

// 更新用户信息接口
func updateUserInfo(c *gin.Context) {
	p := new(model.DexUserInfo)
	_ = c.Bind(p)
	userInfo := service.UserInfoService.GetUserInfoByTokenOfContext(c)
	if userInfo.WalletAddress != p.WalletAddress {
		errors.NewErrorByEnum(enum.PermissionsErr)
	}
	resultUserInfo := service.UserInfoService.UpdateUserInfo(p)
	format.NewApiResult(c).Success(resultUserInfo)
}

func login(c *gin.Context) {
	walletAddress := c.PostForm("walletAddress")
	if walletAddress == "" {
		errors.NewErrorByEnum(enum.ArgumentErr)
	}
	_ = service.UserInfoService.QueryOrSave(walletAddress)
	token, refreshToken, err := middleware.GenToken(walletAddress)
	if err != nil {
		errors.NewErrorByEnum(enum.AuthErr)
	}
	format.NewApiResult(c).Success(gin.H{
		"token":        token,
		"refreshToken": refreshToken,
	})
}

func RefreshToken(c *gin.Context) {
	token := c.PostForm("token")
	refreshToken := c.PostForm("refreshToken")
	newToken, err := middleware.RefreshToken(token, refreshToken)
	if err != nil {
		errors.NewErrorByEnum(enum.AuthErr)
	}
	format.NewApiResult(c).Success(newToken)
}
