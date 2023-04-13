package middleware

import (
	"gin-server/src/enum"
	"gin-server/src/global_const"
	"gin-server/src/utils/config"
	"gin-server/src/utils/errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		isAuth := config.Config.Auth.Token
		if isAuth {
			// 这里根据自己的业务进行更改
			authorizationStr := c.Request.Header.Get(global_const.Authorization)
			if authorizationStr == "" {
				errors.NewErrorByEnum(enum.LoginErr)
				c.Abort()
			}

			// parse authorization，是否为我们关心的token
			claim, err := ParseToken(authorizationStr)
			if err != nil {
				v, _ := err.(*jwt.ValidationError)
				if v.Errors == jwt.ValidationErrorExpired {
					refreshToken := c.Request.Header.Get(global_const.Refresh)
					token, err := RefreshToken(authorizationStr, refreshToken)
					if err != nil {
						errors.NewErrorByEnum(enum.TokenExpires)
						c.Abort()
					}
					c.Writer.Header().Set(global_const.NewToken, token)
				} else {
					errors.NewErrorByEnum(enum.TokenInvalid)
					c.Abort()
				}
			}
			// 将当前请求的username信息保存到请求的上下文c上
			c.Set(global_const.WalletAddress, claim.WalletAddress)
			c.Next()
		}
	}
}
