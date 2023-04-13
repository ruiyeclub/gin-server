package global_const

import "time"

const Authorization = "Authorization" // token
const Refresh = "Refresh"             // Refresh_token
const TokenExpireDuration = time.Minute * 30
const TokenRedisKey = "token:dex:%s"
const UserInfoKey = "userinfo:%s"
const UserInfoExpire = time.Minute * 5

const ATokenExpiredDuration = 2 * time.Hour
const RTokenExpiredDuration = 2 * 24 * time.Hour
const WalletAddress = "walletAddress"
const NewToken = "newToken"
