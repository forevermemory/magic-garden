package middleware

import (
	"errors"
	"magic/global"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

const (
	// SignKey 签名
	SignKey string = "magic-garden"
)

// ErrTokenExpired 过期
var ErrTokenExpired = errors.New("Token is expired")

// CustomClaims 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Nickname string `json:"nickname"`
	IsVip    int    `json:"is_vip"`
	jwt.StandardClaims
}

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

// NewJWT 新建一个jwt实例
func NewJWT() *JWT {
	return &JWT{
		[]byte(SignKey),
	}
}

// CreateToken 生成一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析Token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	//fmt.Println(err)
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("That's even not  a token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("Token not active yet")
			} else {
				return nil, errors.New("Couldn't handle this token")
			}
		}
		return nil, errors.New("Couldn't handle this token")
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("Couldn't handle this token")
}

// RefreshToken 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now

		claims.StandardClaims.ExpiresAt = time.Now().Add(24 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", errors.New("Couldn't handle this token")
}

// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, _ := c.Cookie("token")
		token2 := token
		// 1. 是否携带token
		if token == "" {
			c.JSON(http.StatusForbidden, gin.H{
				"code": -1,
				"msg":  "请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}

		j := NewJWT()
		// parseToken 解析token包含的信息
		// 2.解析token失败
		claims, err := j.ParseToken(token)
		if claims == nil {
			if err != nil {
				c.JSON(http.StatusForbidden, gin.H{
					"code": -1,
					"msg":  err.Error(),
				})
				c.Abort()
				return
			}

			c.JSON(http.StatusForbidden, gin.H{
				"code": -1,
				"msg":  "解析token时发生未知错误",
			})
			c.Abort()
			return
		}
		// 3.token 是否有效 是否过期 TODO
		if ok, err := isTokenExist(token, claims.UserID+"__"+claims.Phone); err != nil || !ok {
			c.JSON(http.StatusForbidden, gin.H{
				"code": -1,
				"msg":  "token已过期，请重新登录",
			})
			c.Abort()
			return
		}
		//
		c.Request.Header.Set("username", claims.Username) // 是否有用
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("claims", claims)
		//fmt.Println(claims.StandardClaims.ExpiresAt)
		if claims.StandardClaims.ExpiresAt-86400 < time.Now().Unix() {
			newtoken, err := j.RefreshToken(token2)
			if err == nil {
				http.SetCookie(c.Writer, &http.Cookie{
					Name:     "token",
					Value:    newtoken,
					Path:     "/",
					Domain:   "",
					MaxAge:   -1,
					Secure:   false,
					HttpOnly: false,
					SameSite: http.SameSiteLaxMode,
				})
				//刷新历史记录
				setToken(newtoken, claims.UserID+"__"+claims.Phone)
			}
		}

	}
}

func setToken(token string, key string) error {
	conn := global.REDIS.Get()
	defer conn.Close()
	_, err := conn.Do("set", key, token)
	return err
}

func isTokenExist(token string, key string) (bool, error) {
	conn := global.REDIS.Get()
	defer conn.Close()
	if to2, err := redis.String(conn.Do("get", key)); err == nil {
		if to2 == token {
			return true, nil
		}
	} else {
		return false, err
	}
	return false, nil
}
