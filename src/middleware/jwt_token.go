package middleware

import (
	"errors"
	"gin-server/src/global_const"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var Secret = []byte("gin-server")

type MyClaims struct {
	WalletAddress string
	jwt.StandardClaims
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	return Secret, nil
}

// GenToken 颁发token access token 和 refresh token
func GenToken(walletAddress string) (aToken, rToken string, err error) {
	rc := jwt.StandardClaims{
		Issuer:    "ruiyeclub", // 签发人
		ExpiresAt: time.Now().Add(global_const.ATokenExpiredDuration).Unix(),
	}
	at := MyClaims{
		walletAddress,
		rc,
	}
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, at).SignedString(Secret)

	// refresh token 不需要保存任何用户信息
	rt := rc
	rt.ExpiresAt = time.Now().Add(global_const.RTokenExpiredDuration).Unix()
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, rt).SignedString(Secret)
	return
}

// ParseToken 验证Token
func ParseToken(tokenID string) (*MyClaims, error) {
	var myc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenID, myc, keyFunc)
	if err != nil {
		// 验证错误这里返回
		return myc, err
	}
	if !token.Valid {
		return nil, errors.New("verify Token Failed")
	}
	return myc, nil
}

// RefreshToken 通过 refresh token 刷新 newToken
func RefreshToken(aToken, rToken string) (newAToken string, err error) {
	// rToken 无效直接返回
	if _, err = jwt.Parse(rToken, keyFunc); err != nil {
		return
	}
	// 从旧access token 中解析出claims数据
	var claim MyClaims
	_, err = jwt.ParseWithClaims(aToken, &claim, keyFunc)
	if err == nil {
		// 未过期
		return aToken, err
	}
	// 判断错误是不是因为access token 正常过期导致的
	v, _ := err.(*jwt.ValidationError)
	if v.Errors == jwt.ValidationErrorExpired {
		at := MyClaims{
			claim.WalletAddress,
			jwt.StandardClaims{
				Issuer:    "ruiyeclub", // 签发人
				ExpiresAt: time.Now().Add(global_const.ATokenExpiredDuration).Unix(),
			},
		}
		return jwt.NewWithClaims(jwt.SigningMethodHS256, at).SignedString(Secret)
	}
	return
}
