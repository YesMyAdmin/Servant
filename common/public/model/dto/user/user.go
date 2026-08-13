package dto

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

// LoginSession 登录会话
type LoginSession struct {
	//用户id
	UserId uint64 `json:"userId"`
	//登录时间
	LoginTime time.Time `json:"loginTime"`
	//登录过期时间
	LoginExpireTime time.Time `json:"loginExpireTime"`
	jwt.RegisteredClaims
}

//登录请求
type LoginRequest struct {
	UserName string `json:"userName"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Password string `json:"password"`
}

// LoginSuccessResponse 登录成功返回
type LoginSuccessResponse struct {
	// BearerToken 登录成功后返回的token
	BearerToken string `json:"bearerToken"`
}