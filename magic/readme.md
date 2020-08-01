## 魔法花园后台

### 说明

-   返回的数据均为 json
-   当 code =0 时说明操作成功
-   code = -1 操作失败

### 用户

#### 获取图片验证码

```
请求方式
    get
请求路径
    /captcha
请求参数
    captcha
说明
    captcha 参数后面的内容是验证码的内容,由前台来控制
返回结果
    图片
```

#### 获取手机验证码

```
请求方式
    get
请求路径
   /api/v1/user/sendsms
请求参数
    phone 手机号
返回结果
    返回验证码
```

#### 检查用户名是否存在

```
请求方式
    get
请求路径
   /api/v1/user/registe/isUsernameExist
请求参数
    username  用户名 用于后面登陆
返回结果
{
    "code": 0,
    "data": {
        "is_exits": "2",
        "msg": "用户名不存在"
    }
}

{
    "code": 0,
    "data": {
        "is_exits": "1",
        "msg": "用户名存在"
    }
}
```

#### 用户注册

```
请求方式
    post
请求路径
    /api/v1/user/registe
请求参数
	username 用于登陆 建议和手机号一样 不要中文
	password 密码 > 六位
	phone 手机号
	code 手机验证码
    is_captcha  下面两个参数在用户非法操作后会触发 值传1 is_captcha=1
    captcha 验证码内容
说明
    判断是否在注册时候存在不合理操作,存在就出现验证码验证

返回结果
{
    "code": 0,
    "data": "ok"
}
```

#### 用户重置密码

```
请求方式
    post
请求路径
    /api/v1/user/reset/password
请求参数
    user_id 用户id
	phone 手机号
	code 手机验证码
    new_password 新的密码

说明
    判断是否在注册时候存在不合理操作,存在就出现验证码验证

返回结果
{
    "code": 0,
    "data": "ok"
}
```

#### 用户登陆

```
请求方式
    post
请求路径
    /api/v1/user/login
请求参数
	username
	password

    is_captcha 下面两个参数在用户非法操作后会触发 值传1 is_captcha=1
    captcha
说明
    判断是否在注册时候存在不合理操作,存在就出现验证码验证

返回结果
{
    "code": 0,
    "data": {
        "_id": 2,
        "is_vip": 0,
        "nickname": "张三",
        "phone": "18862806080",
        "username": "admin"
    }
}
```

#### 修改昵称

```
请求方式
    post
请求路径
    /api/v1/user/update
请求参数
	_id 用户id
	nickname 昵称

返回结果
{
    "code": 0,
    "data": "ok"
}
```

#### 根据 id 查询用户的详情信息

```
请求方式
    get
请求路径
    /api/v1/user/user/get
请求参数
	_id 登陆成功后返回的字段

返回结果
{
    "code": 0,
    "data": {
        "_id": 2,
        "username": "admin",
        "nickname": "张三",
        "password": "123123",
        "phone": "18862806080",
        "status": 0,
        "is_vip": 0,
        "desc": "",
        "gb_money": "10000",
        "yuanbao": "0",
        "change_pass_time": ""
    }
}
```

### 花园

#### 初始化花园

```
请求方式
    get
请求路径
    /api/v1/garden/init
请求参数
    user_id 用户的id

返回结果
{
    "code": 0,
    "data": "ok"
}
```

#### 修改花园名称、公告

```
请求方式
    post
请求路径
    /api/v1/garden/update
请求参数
    _id 花园的id
    g_name 花园名称
    g_info 花园公告

返回结果
{
    "code": 0,
    "data": "ok"
}
```

#### 查看背包 道具和种子

```
请求方式
    get
请求路径
    /api/v1/garden/knapsack/list
请求参数
    cate 分类 1 代表种子 2 代表道具
    page 初始传1 根据返回的 totalPage 判断有多少页
    garden_id 花园的id即 _id

返回结果
{
    "code": 0,
    "data": {
        "total": 1,
        "totalPage": 1,
        "data": [
            {
                "_id": 14,
                "garden_id": 3, 花园id
                "seed_id": 0, 种子id
                "seed_num": "0", 种子数量
                "cate": 2,
                "prop_id": 6,  道具ID
                "prop_num": "33",  道具数量
                "page": 0,
                "seed_name": "", 种子名称
                "prop_name": "终极魔力营养液" 道具名称
            }
        ]
    }
}
```

#### 花园帮助文档

-   帮助列表

```
请求方式
    get
请求路径
    /api/v1/garden/help/list
请求参数
    无

返回结果
    {
    "code": 0,
    "data": [
        {
            "_id": 1,
            "h_title": "如何播种花朵?",
            "h_content": ""
        },
    ]
    }
```

-   帮助详情

```
请求方式
    get
请求路径
    /api/v1/garden/help/detail
请求参数
    _id  列表中返回的_id的值
返回结果
{
    "code": 0,
    "data": {
        "_id": 1,
        "h_title": "如何播种花朵?",
        "h_content": "如何播种花朵?<br/>\r\n首先您需要有一个可播种的花盆,种子可以去商店购买获得.等级越高,可以买到的种子就越多,有机会种出的花种也就越多。随着等级的上升,您还可以添置新的花盆。<br/>\r\n----------<br/>"
    }
}
```
