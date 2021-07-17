package service

/*
date:2020-07-29 14:00:32
*/

import (
	"errors"
	"magic/db"
	"magic/global"
	utils "magic/utils"

	"github.com/afocus/captcha"
	"github.com/jinzhu/gorm"

	"github.com/gomodule/redigo/redis"
)

// UserLogin 用户登录
func UserLogin(req *global.RegisteruserParams) (*db.Users, error) {

	// 1.1 是否出发了图片验证码
	// if req.IsCaptcha == 1 {
	// 	conn := global.REDIS.Get()
	// 	defer conn.Close()
	// 	capt, err := redis.String(conn.Do("get", ua))
	// 	if err != nil {
	// 		return nil, errors.New("验证码过期")
	// 	}
	// 	if capt != req.Captcha {
	// 		return nil, errors.New("验证码不正确")
	// 	}
	// }

	// 2. 登陆用户
	user, err := db.GetUsersByUsernameAndPassword(req.Username, req.Password)
	if err != nil {
		return nil, errors.New("用户名不存在,或者密码错误,请重试")
	}
	user.Password = "******"

	// 3. 查询相关信息
	// 查询等级
	ll, err := db.GetUserLevelByLevel(user.Level)
	if err != nil {
		// return nil, err
	}
	user.LevelObject = ll

	return user, nil
}

// RegisterUser 注册用户
func RegisterUser(req *global.RegisteruserParams, ua string) error {
	// conn := global.REDIS.Get()
	// defer conn.Close()
	// // 1.1 是否出发了验证码
	// if req.IsCaptcha == 1 {
	// 	capt, err := redis.String(conn.Do("get", ua))
	// 	if err != nil {
	// 		return errors.New("验证码过期")
	// 	}
	// 	if capt != req.Captcha {
	// 		return errors.New("验证码不正确")
	// 	}
	// }
	// // 1.2校验手机验证码是否正确
	// code, err := redis.String(conn.Do("get", req.Username))
	// if err != nil {
	// 	fmt.Println("redis::err", err)
	// 	return err
	// }
	// if code != req.Code {
	// 	return errors.New("手机验证码不正确")
	// }
	// 2. 注册用户
	uuidd := utils.GetUUID()
	var err error
	user := &db.Users{
		ID:       uuidd,
		Username: req.Username,
		Nickname: req.Nickname,
		Password: req.Password,
		Phone:    req.Username,
		GBMoney:  1000,
		Yuanbao:  "10",
		IsVip:    1,
	}
	if err = db.AddUsers(user); err != nil {
		return err
	}
	// 3.删除redis的code
	// if _, err = conn.Do("del", req.Phone); err != nil {
	// 	return err
	// }
	return nil
}

// RegisterUserSendMsg 发短信
func RegisterUserSendMsg(phone string) (string, error) {
	// 校验手机号吧
	if ok := utils.CheckIsPhoneVaild(phone); !ok {
		return "", errors.New("手机号不合理")
	}
	code := utils.GenValidateCode(4)
	utils.SendQiniuCode(code, phone)
	// 将验证码放到redis TODO key 手机号 val code
	conn := global.REDIS.Get()
	defer conn.Close()
	if _, err := conn.Do("set", phone, code); err != nil {
		return "", err
	}
	// 不设置过期时间
	return code, nil
}

// AddUsers add
func AddUsers(b *db.Users) error {
	return db.AddUsers(b)
}

// IsUsernameExist 是否用户名存在
func IsUsernameExist(username string) (interface{}, error) {
	// 首先判断用户名是否存在
	res := make(map[string]interface{})
	_, err := db.GetUserByUsername(username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			res["msg"] = "用户名不存在"
			res["is_exits"] = "2"
			return res, nil
		}
		return res, err
	}
	// 可以查询到这个username
	res["msg"] = "用户名存在"
	res["is_exits"] = "1"
	return res, nil
}

// UpdateUsersPassword 重置密码
func UpdateUsersPassword(req *global.RegisteruserParams) (*db.Users, error) {
	// 1.用户
	user, err := db.GetUsersByID(req.ID)
	if err != nil {
		return nil, err
	}
	// 2。 code是否正确
	conn := global.REDIS.Get()
	defer conn.Close()
	code, err := redis.String(conn.Do("get", req.Phone))
	if err != nil {
		return nil, err
	}
	if code != req.Code {
		return nil, errors.New("验证码不正确")
	}
	// 3.更新
	user.Password = req.NewPassword
	user.ChangePassTime = utils.GetNowTimeString()
	return db.UpdateUsers(user)
}

// UpdateUsers update
func UpdateUsers(b *db.Users) (*db.Users, error) {
	return db.UpdateUsers(b)
}

// GetUsersByID get by id
func GetUsersByID(id string) (*db.Users, error) {
	return db.GetUsersByID(id)
}

// ListUsers  page by condition
func ListUsers(b *db.Users) (*db.DataStore, error) {
	list, err := db.ListUsers(b)
	if err != nil {
		return nil, err
	}
	total, err := db.CountUsers(b)
	if err != nil {
		return nil, err
	}
	return &db.DataStore{Total: total, Data: list, TotalPage: (int(total) + b.PageSize - 1) / b.PageSize}, nil
}

// DeleteUsers delete
func DeleteUsers(id string) error {
	return db.DeleteUsers(id)
}

// HandleGenerateCaptcha 生成验证码
func HandleGenerateCaptcha(ca, ua string) (*captcha.Image, error) {
	// 存入到redis 防止前端出错
	conn := global.REDIS.Get()
	defer conn.Close()
	_, err := redis.String(conn.Do("set", ua, ca))
	if err != nil {
		return nil, err
	}
	// 设置过期时间 300s
	_, err = conn.Do("expire", ua, 300)
	if err != nil {
		return nil, err
	}
	// 生成验证码
	img := utils.GenerateCaptcha(ca)
	return img, nil
}
