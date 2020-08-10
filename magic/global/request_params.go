package global

// RegisteruserParams 注册传入的参数
type RegisteruserParams struct {
	ID          int    `json:"user_id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
	Phone       string `json:"phone"`
	Code        string `json:"code"`
	// Captcha 触发验证码后需要判断
	Captcha   string `json:"captcha"`
	IsCaptcha int    `json:"is_captcha"`
}

// UserAddGamesParams 用户添加游戏传入的参数
type UserAddGamesParams struct {
	UserID int    `json:"user_id" form:"user_id"`
	GameID int    `json:"game_id" form:"game_id"`
	Ginfo  string `json:"g_info" form:"g_info"`
	Gname  string `json:"g_name" form:"g_name"`
}

// GardenParams 花园参数
type GardenParams struct {
	UserID      int `json:"user_id" form:"user_id"`
	GardenID    int `json:"_id" form:"_id"`
	FlowerpotID int `json:"number" form:"number"`
	SeedID      int `json:"seed_id" form:"seed_id"`
	Page        int `json:"page" form:"page"`
	// Cate 背包的内容分类 分类 1种子 2道具
	Cate int `json:"cate" form:"cate"`
}

// GardenPotParams 花盆参数 [{},{}]
type GardenPotParams struct {
	GardenID    int `json:"garden_id" form:"garden_id"`
	FlowerpotID int `json:"number" form:"number"`
	SeedID      int `json:"seed_id" form:"seed_id"`
	IsVip       int `json:"is_vip" form:"is_vip"`
	Page        int `json:"page" form:"page"`
	// 2干旱(浇水) 3有虫(除虫) 4有草(除草)[{"number":1,"kind":"2"}]
	Handle []map[string]int `json:"handles" form:"handles"`
}
