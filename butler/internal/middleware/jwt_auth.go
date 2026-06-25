package middleware

import (
    "net/http"
    "strings"
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
)

// 用于签名的密钥（生产环境请从环境变量读取）
var jwtSecret = []byte("your-secret-key")

// 自定义Claims结构体，可根据业务需求添加字段
type UserClaims struct {
    UserID   string `json:"user_id"`
    Username string `json:"username"`
    jwt.RegisteredClaims
}

// JWTAuthMiddleware 返回一个Gin中间件，用于验证JWT
func JWTAuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // 1. 从Authorization头获取Token，格式为 "Bearer <token>"
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
            c.Abort()
            return
        }

        // 2. 检查并解析Bearer前缀
        parts := strings.SplitN(authHeader, " ", 2)
        if !(len(parts) == 2 && parts[0] == "Bearer") {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
            c.Abort()
            return
        }
        tokenString := parts[1]

        // 3. 解析并验证Token
        claims := &UserClaims{}
        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            // 确保使用的签名算法是预期的HS256
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, jwt.ErrSignatureInvalid
            }
            return jwtSecret, nil
        })

        // 4. 处理解析错误（Token无效或过期）
        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
            c.Abort()
            return
        }

        // 5. 将用户信息存入Gin Context，供后续处理器使用
        c.Set("user_id", claims.UserID)
        c.Set("username", claims.Username)
        // 也可以直接存入整个claims对象
        c.Set("claims", claims)

        c.Next() // 继续执行下一个中间件或路由处理函数
    }
}