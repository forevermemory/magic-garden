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
        "is_vip": 1, 是否是vip 1否 2是
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
        "is_vip": 1, 是否是vip 1否 2是
        "desc": "",
        "gb_money": "10000",
        "yuanbao": "0",
        "change_pass_time": ""
    }
}
```

### 花园

#### 签到

```
请求方式
    get
请求路径
    /api/v1/garden/signin
请求参数
    _id 花园id(用户的id)

返回结果
{
    "code": 0,
    "msg": "ok",
    "data": [
        {
            "o_id": 66,
            "o_name": "风华正茂",
            "o_num": 3
        },
        {
            "o_name": "GB",
            "o_num": 7000
        }
    ]
}

{
    "code": 0,
    "msg": "ok",
    "data": [
        {
            "o_name": "你今天签到过了,明天记得来领取签到奖励哦!"
        }
    ]
}
```

#### 签到历史

```
请求方式
    get
请求路径
    /api/v1/garden/signin/list
请求参数
    page 初始传1

```

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

#### 查询一个花园

```
请求方式
    get
请求路径
    /api/v1/garden/list/:oid
请求参数
    oid 花园的id(用户的id)

返回结果
{
    "code": 0,
    "msg": "ok",
    "data": {
        "_id": 2,
        "g_name": "哈哈哈",
        "g_info": "劳动可耻,偷窃光荣!",
        "g_level": 1,
        "is_signin": 0,
        "sign_days": "0",
        "g_atlas": "0",
        "g_current_ex": "0"
    }
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
    "msg": "ok"
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

#### 花盆

-   查询花盆列表

```
请求方式
    get
请求路径
    /api/v1/garden/flowerpot/list
请求参数
    garden_id 花园的id

返回结果
参数说明:


{
    "code": 0,
    "msg": "ok",
    "data": [
        {
            "_id": 1, // 花盆id
            "user_id": 2, // 用户id 花园id
            "garden_id": 2,
            "number": 1, // 该用户的花盆编号
            "is_lock": 2, // 是否解锁 1未解锁 2解锁
            "is_sow": 2, // 是否播种 空还是播种 1 空 2 播种
            "seed_id": 1, // 种的种子 id
            "status": 0,
            "seed_result": 3, // 种子开花结果 开出的花是图谱id
            "seed_result_str": "", // 种子开花结果 开出的花 字符串
            "flower_num": 7, // 开花数量
            "flower_num_haldle": 3,
            "current_stage": "花蕾期", // 当前阶段
            "next_stage": "开花", // 下个阶段
            "next_stage_str": "1小时39分钟后开花",
            "sow_time": "2020-08-07 11:03:50",
            "is_change_color": 0, // 否使用染色剂 1 未使用 2使用了
            "change_result": "", // 染色结果 string
            "is_harvest": 1 // 是否可以收获 1不可以 2 可以
            "disaster": 3 // 自然灾害类型 1健康 2干旱(浇水) 3有虫(除虫) 4有草(除草)
        },
        {
            "_id": 2,
            "user_id": 2,
            "garden_id": 2,
            "number": 2,
            "is_lock": 2,
            "is_sow": 2,
            "seed_id": 2,
            "status": 0,
            "seed_result": 5,
            "seed_result_str": "红色野花",
            "flower_num": 9,
            "flower_num_haldle": 3,
            "current_stage": "花蕾期",
            "next_stage": "开花",
            "next_stage_str": "4秒后开花",
            "sow_time": "2020-08-07 10:24:10",
            "is_change_color": 0,
            "change_result": "",
            "is_harvest": 2,
            "disaster": 2
        }
    ]
}
```

-   花盆播种

```
请求方式
    post
请求路径
    /api/v1/garden/flowerpot/sow
请求参数
    garden_id 花园的id
    number 花盆编号
    seed_id 播种的种子id

    入参示例
        普通播种入参

            {
                "garden_id":2, // 花园id
                "number":1, // 花盆编号
                "seed_id":3 // 种子id
            }
        vip 播种入参
            {
                "garden_id":2,
                "seed_id":3,
                "is_vip":2
            }
    ps: 普通用户只能一个种子一个种子进行播种
        vip用户可以点击种子进行一键播种
返回结果


```

-   查看已经播种花盆状态

```
请求方式
    get
请求路径
    /api/v1/garden/flowerpot/detail
请求参数
    garden_id 花园的id
    number 花盆编号

    入参示例
        "garden_id":2, // 花园id
        "number":1, // 花盆编号
返回结果
{
    "code": 0,
    "msg": "ok",
    "data": {
        "_id": 1,
        "user_id": 2,
        "garden_id": 2,
        "number": 1,
        "is_lock": 2,
        "is_sow": 2,
        "seed_id": 66,
        "status": 0,
        "seed_result": 290,
        "seed_result_str": "摩羯花",
        "flower_num": 1,
        "flower_num_haldle": 3,
        "current_stage": "花种期",
        "next_stage": "花苗期",
        "next_stage_str": "2小时39分钟后进入花苗期",
        "sow_time": "2020-08-10 14:42:21",
        "is_change_color": 1,
        "change_result": "",
        "is_harvest": 1,
        "disaster": 4
    }
}

```

-   浇水除虫除草操作

```
请求方式
    post
请求路径
    /api/v1/garden/flowerpot/lookafter
请求参数
    入参示例
    {
        "garden_id":2,
        "handles":[
            {"number":1,"kind":3},
            {"number":2,"kind":4}
        ],
        "is_vip":2
    }
    // 普通用户 只能一个一个操作 即 handles 只能传一个参数
    // vip用户 需要携带 "is_vip":2 handles可以传多个参数
返回结果
{
    "code": 0,
    "msg": "ok",
    "data": {
        "result": "一键操作成功,您获得2点经验值",
        "total": 2
    }
}
```

-   移除花盆中成长的花朵

```
请求方式
    post
请求路径
    /api/v1/garden/flowerpot/lookafter
请求参数
    入参示例
    {
        "garden_id": 2,
        "number": 1
    }
返回结果
{
    "code": 0,
    "msg": "ok",
    "data": {
        "result": "移除成功"
    }
}
```

-   染色 TODO

```
请求方式
    post
请求路径
    /api/v1/garden/flowerpot/dyeing
请求参数
    入参示例
        {
            "garden_id":"887ec461385e4ec69cec4fcabdabbaf8",
            "number":1,
        }
返回结果
{
    "code": 0,
    "msg": "ok",
    "data": {
        "result": "当前花朵暂不支持染色,请正确操作"
    }
}
{
    "code": 0,
    "msg": "ok",
    "data": {
        "result": "染色成功,您获得了蓝色雪花纷飞"
    }
}
```

-   施肥

```
请求方式
    post
请求路径
    /api/v1/garden/flowerpot/fertilizer
请求参数
    入参示例
    {
        "garden_id": "887ec461385e4ec69cec4fcabdabbaf8",
        "number": 2,
        "prop_id": 4
    }
    // TODO 5 6 会直接报错
返回结果
{
    "code": 0,
    "msg": "ok",
    "data": {
        "result": "操作成功"
    }
}
```

-   收获

```
请求方式
    post
请求路径
    /api/v1/garden/flowerpot/harvest
请求参数
    入参示例
    普通一个一个收获
        {
            "garden_id": "887ec461385e4ec69cec4fcabdabbaf8",
            "number": 2
        }
    vip一键收获
        {
            "garden_id": "887ec461385e4ec69cec4fcabdabbaf8",
            "is_vip": 2
        }
返回结果
{
    "code": 0,
    "msg": "收获成功",
    "data": {
        "result": [
            "收获成功,经验值+2，GB+9000",
            "恭喜您种出了一种新的花朵[紫色三色旋花]!!!经验值+600",
            "您一共收获了4朵紫色三色旋花！",
            "收获成功,经验值+2，GB+9000",
            "恭喜您种出了一种新的花朵[粉色羽叶茑萝]!!!经验值+600",
            "您一共收获了9朵粉色羽叶茑萝！",
            "收获成功,经验值+2，GB+9000",
            "恭喜您种出了一种新的花朵[蓝色金鱼草]!!!经验值+600",
            "您一共收获了5朵蓝色金鱼草！"
        ],
        "result2": true,
        "total": 0
    }
}

{
    "code": 0,
    "msg": "收获成功",
    "data": {
        "result": [
            "收获成功,经验值+2，GB+9000",
            "您一共收获了4朵紫色三色旋花！",
            "收获成功,经验值+2，GB+9000",
            "您一共收获了9朵粉色羽叶茑萝！",
            "收获成功,经验值+2，GB+9000",
            "您一共收获了5朵蓝色金鱼草！"
        ],
        "result2": true,
        "total": 0
    }
}
```

-   购买种子

```
请求方式
    post
请求路径
    /api/v1/garden/gb/shop/buy/seed
请求参数
    入参示例
        {
            "garden_id": "887ec461385e4ec69cec4fcabdabbaf8",
            "seed_id": 3,
            "seed_num":2
        }
返回结果
{
    "code": 0,
    "msg": "ok",
    "data": {
        "result": "购买成功,您获得2颗梅花,花费400GB"
    }
}
```

-   购买道具

```
请求方式
    post
请求路径
    /api/v1/garden/gb/shop/buy/prop
请求参数
    入参示例
        {
            "garden_id": "887ec461385e4ec69cec4fcabdabbaf8",
            "prop_id": 3,
            "prop_num":2
        }
返回结果
{
    "code": 0,
    "msg": "ok",
    "data": {
        "result": "购买成功,您获得2个中魔力营养液,花费5000GB"
    }
}
```

-   魔法屋可合成列表

```
请求方式
    get
请求路径
    /api/v1/garden/magician/list?page=1
请求参数
    入参示例
    page = 1
返回结果
{
    "code": 0,
    "msg": "ok",
    "data": {
        "total": 217,
        "totalPage": 22,
        "data": [
            {
                "_id": 121,
                "raw_id": 121,
                "img_url": "s_l_121.gif",
                "level_num": 10,
                "raw_price": 0,
                "vip_price": 0,
                "level_str": "见习魔法学徒",
                "seed_name": "独孤求败",
                "rarity": 2,
                "is_change_color": 1,
                "forecast_num": 1,
                "forecast_time": 12,
                "meaning": "不是光棍要种,是光棍更要种的光棍节非卖活动花!",
                "page": 0,
                "page_size": 0
            }
        ]
    }
}
```

-   魔法屋查询一个种子合成所需材料

```
请求方式
    post
请求路径
    /api/v1/garden/magician/detail
请求参数
    入参示例
    {
        "garden_id":"887ec461385e4ec69cec4fcabdabbaf8",
        "seed_id":115
    }
返回结果
{
    "code": 0,
    "msg": "ok",
    "data": [
        {
            "seed_id": 115,
            "atlas_id": 120,
            "num": 12,  // 合成所需数量
            "flower_cate_name": "白色烈焰焚情",
            "flower_image": "m_l_120.gif",
            "rarity": 1,
            "seed_name": "金色烈焰焚情",
            "garden_id": "887ec461385e4ec69cec4fcabdabbaf8",
            "total_num": 81 // 持有数量
        },
        {
            "seed_id": 115,
            "atlas_id": 121,
            "num": 12,
            "flower_cate_name": "红色烈焰焚情",
            "flower_image": "m_l_121.gif",
            "rarity": 1,
            "seed_name": "金色烈焰焚情",
            "garden_id": "887ec461385e4ec69cec4fcabdabbaf8",
            "total_num": 45
        },
        {
            "seed_id": 115,
            "atlas_id": 122,
            "num": 12,
            "flower_cate_name": "黄色烈焰焚情",
            "flower_image": "m_l_122.gif",
            "rarity": 1,
            "seed_name": "金色烈焰焚情",
            "garden_id": "",
            "total_num": 0
        },
        {
            "seed_id": 115,
            "atlas_id": 123,
            "num": 12,
            "flower_cate_name": "粉色烈焰焚情",
            "flower_image": "m_l_123.gif",
            "rarity": 1,
            "seed_name": "金色烈焰焚情",
            "garden_id": "",
            "total_num": 0
        },
        {
            "seed_id": 115,
            "atlas_id": 319,
            "num": 4,
            "flower_cate_name": "银色烈焰焚情",
            "flower_image": "m_l_319.gif",
            "rarity": 2,
            "seed_name": "金色烈焰焚情",
            "garden_id": "",
            "total_num": 0
        }
    ]
}
```

-   魔法屋合成种子

```
请求方式
    post
请求路径
    /api/v1/garden/magician/synthesis
请求参数
    入参示例
    {
        "garden_id":"887ec461385e4ec69cec4fcabdabbaf8",
        "seed_id":115
    }
返回结果
{
    "code": 0,
    "msg": "ok",
    "data": {
        "result": ["不满足合成条件"]
    }
}

{
    "code": 0,
    "msg": "ok",
    "data": {
        "result": [
            "合成成功,您获得了一颗金色烈焰焚情",
            "您获得100经验",
            "您获得1000GB"
        ]
    }
}
```

-   查看花房里面花朵

```
请求方式
    get
请求路径
    /api/v1/garden/house/list
请求参数
    入参示例
        garden_id=887ec461385e4ec69cec4fcabdabbaf8
        cate=1  // 花篮 cate=2 花瓶
        page default=1
返回结果
{
    "code": 0,
    "msg": "ok",
    "data": {
        "total": 7,
        "totalPage": 1,
        "data": [
            {
                "_id": 26,
                "garden_id": "887ec461385e4ec69cec4fcabdabbaf8",
                "atlas_id": 114,
                "num": 36, // 该种花朵持有数量
                "cate": 1,
                "seed_id": 35,
                "flower_cate_name": "紫色三色旋花", // 名称
                "flower_image": "m_l_114.gif",
                "rarity": 0
            }
        ]
    }
}
```

-   查看花房里面花朵 -- 统计

```
请求方式
    get
请求路径
    /api/v1/garden/house/statistics
请求参数
    入参示例
        garden_id=887ec461385e4ec69cec4fcabdabbaf8
        cate=1  // 花篮 cate=2 花瓶
返回结果
    // 您收获和采摘的花均存放于您的花篮中,共有鲜花8种,21朵
    // 您收到的鲜花会存放在您的花瓶中,共有鲜花1种,1朵.
{
    "code": 0,
    "msg": "ok",
    "data": [
        {
            "total_duo": 216, // 所有花朵总数
            "total_zhong": 7  // 种类
        }
    ]
}
```

### 好友系统

#### 新增好友

docker network create link-network
docker run -d --network link-network --network-alias redis redis:latest
docker run -d --network link-network --network-alias mysql mysql:latest

docker run -it --network link-network magic:v1 /bin/bash
