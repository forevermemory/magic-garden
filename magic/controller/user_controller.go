package controller

/*
date:2020-07-29 14:00:32
*/

import (
	"image/png"
	"magic/db"
	"magic/global"
	"magic/service"
	"magic/utils/middleware"
	"net/http"
	"strconv"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

// UserLogin lgin
func UserLogin(c *gin.Context) {
	var u = global.RegisteruserParams{}
	err := c.ShouldBind(&u)

	user, err := service.UserLogin(&u, c.Request.Header.Get("User-Agent"))
	if err != nil {
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	// 设置cookie token
	generateToken(c, user)
	// return OKResponse{0, user}
}

// RegisterUserSendMsg 注册
func RegisterUserSendMsg(c *gin.Context) interface{} {
	phone, ok := c.GetQuery("phone")
	if !ok {
		return ErrorResponse{-1, "请传入手机号phone参数"}
	}
	code, err := service.RegisterUserSendMsg(phone)
	if err != nil {
		return ErrorResponse{-1, err.Error()}
	}
	res := Response{"ok", gin.H{"code": code}}
	return OKResponse{0, res}
}

// IsUsernameExist IsUsernameExist
func IsUsernameExist(c *gin.Context) interface{} {
	username, ok := c.GetQuery("username")
	if !ok {
		return ErrorResponse{-1, "请传入username"}
	}
	data, err := service.IsUsernameExist(username)
	if err != nil {
		return ErrorResponse{-1, err.Error()}
	}
	return OKResponse{0, data}
}

// AddUsers add
func AddUsers(c *gin.Context) interface{} {

	var u = global.RegisteruserParams{}
	err := c.ShouldBind(&u)
	if err != nil {
		return ErrorResponse{-1, err.Error()}
	}
	err = service.RegisterUser(&u, c.Request.Header.Get("User-Agent"))
	if err != nil {
		return ErrorResponse{-1, err.Error()}
	}
	return OKResponse{0, "ok"}
}

// UpdateUsers update
func UpdateUsers(c *gin.Context) interface{} {
	var u = db.Users{}
	err := c.ShouldBind(&u)
	if err != nil {
		return ErrorResponse{-1, err.Error()}
	}
	err = service.UpdateUsers(&u)
	if err != nil {
		return ErrorResponse{-1, err.Error()}
	}
	return OKResponse{0, "ok"}
}

// UpdateUsersPassword 重置密码
func UpdateUsersPassword(c *gin.Context) interface{} {
	var u = global.RegisteruserParams{}
	err := c.ShouldBind(&u)
	if err != nil {
		return ErrorResponse{-1, err.Error()}
	}
	err = service.UpdateUsersPassword(&u)
	if err != nil {
		return ErrorResponse{-1, err.Error()}
	}
	return OKResponse{0, "ok"}
}

// GetUsersByID  get xxx by id
func GetUsersByID(c *gin.Context) interface{} {
	var u = db.Users{}
	err := c.ShouldBind(&u)
	if err != nil {
		return ErrorResponse{-1, err.Error()}
	}
	data, err := service.GetUsersByID(u.ID)
	if err != nil {
		return ErrorResponse{-1, err.Error()}
	}
	return OKResponse{0, data}
}

// ListUsers // list by page condition
func ListUsers(c *gin.Context) interface{} {
	var u = db.Users{PageSize: 10, PageNo: 1}
	err := c.ShouldBind(&u)
	if err != nil {
		return ErrorResponse{-1, err.Error()}
	}
	data, err := service.ListUsers(&u)
	if err != nil {
		return ErrorResponse{-1, err.Error()}
	}
	return OKResponse{0, data}
}

// DeleteUsers Delete
func DeleteUsers(c *gin.Context) interface{} {
	var u = db.Users{}
	err := c.ShouldBind(&u)
	if err != nil {
		return ErrorResponse{-1, err.Error()}
	}
	err = service.DeleteUsers(u.ID)
	if err != nil {
		return ErrorResponse{-1, err.Error()}
	}
	return OKResponse{0, "ok"}
}

// GenerateCaptcha 验证码
func GenerateCaptcha(c *gin.Context) {
	ca, ok := c.GetQuery("captcha")
	if !ok {
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "请传入参数",
		})
		return
	}
	ua := c.Request.Header.Get("User-Agent")
	img, err := service.HandleGenerateCaptcha(ca, ua)
	if err != nil {
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	png.Encode(c.Writer, img)
}

func generateToken(c *gin.Context, user *db.Users) {
	// 初始化jwt
	j := &middleware.JWT{
		SigningKey: []byte("magic-garden"),
	}
	claims := middleware.CustomClaims{
		UserID:   user.ID,
		Phone:    user.Phone,
		Username: user.Username,
		Nickname: user.Nickname,
		IsVip:    user.IsVip,
		StandardClaims: jwtgo.StandardClaims{
			//NotBefore: int64(time.Now().Unix() + 10),   // 签名生效时间
			ExpiresAt: time.Now().Unix() + 86400, // 过期时间 1天
			Issuer:    "magic-garden",            //签名的发行者
		},
	}
	// 生成token
	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusOK, ErrorResponse{
			Errcode: -1,
			Message: err.Error(),
		})
		return
	}
	// 存入redis
	if err := setToken(token, strconv.Itoa(user.ID)+"__"+user.Phone); err != nil {
		c.JSON(http.StatusOK, ErrorResponse{
			Errcode: -1,
			Message: err.Error(),
		})
		return
	}
	// 设置cookie
	//c.SetCookie("token", token, 604800, "/", "", false, false)
	http.SetCookie(c.Writer, &http.Cookie{
		Name:   "token",
		Value:  token,
		Path:   "/",
		Domain: "",
		//Expires:time.Now().Add(time.Minute*10),
		MaxAge:   86400,
		HttpOnly: false,
		SameSite: http.SameSiteLaxMode,
		Secure:   false,
	})
	// 返回结果
	result := map[string]interface{}{
		"_id":      user.ID,
		"username": user.Username,
		"phone":    user.Phone,
		"nickname": user.Nickname,
		"is_vip":   user.IsVip,
	}
	c.JSON(http.StatusOK, OKResponse{
		Errcode: 0,
		Data:    result,
	})
	return
}

func setToken(token string, key string) error {
	conn := global.REDIS.Get()
	defer conn.Close()
	_, err := conn.Do("set", key, token)
	return err
}
