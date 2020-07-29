package utils

import "regexp"

// CheckIsPhoneVaild 校验手机号是否合理
func CheckIsPhoneVaild(phone string) bool {
	reg := `^[1]([3-9])[0-9]{9}$`
	return regexp.MustCompile(reg).MatchString(phone)

}
