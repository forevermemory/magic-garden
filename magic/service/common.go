package service

// Result 花盆返回参数
type Result map[string]interface{}

// Set set
func (r Result) Set(key string, value interface{}) {
	r[key] = value
}

// Get get
func (r Result) Get(key string) bool {
	_, ok := r[key]
	return ok
}
