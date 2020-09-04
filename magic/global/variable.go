package global

// Port 服务启动端口
var Port string

// PageSize 分页大小
var PageSize int = 10

// 2干旱(浇水) 3有虫(除虫) 4有草(除草)

// PotStatus 花盆状态
var PotStatus = map[int]string{2: "浇水", 3: "除虫", 4: "除草"}
