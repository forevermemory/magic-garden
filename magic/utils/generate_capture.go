package utils

import (
	"github.com/afocus/captcha"
)

// GenerateCaptcha 返回 m image.Image
func GenerateCaptcha(s string) *captcha.Image {
	cap := captcha.New()
	// 设置字体
	cap.SetFont("./comic.ttf")

	cap.SetSize(128, 64)
	// cap.SetDisturbance(captcha.MEDIUM)
	// cap.SetFrontColor(color.RGBA{255, 255, 255, 255})
	// cap.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})

	img := cap.CreateCustom(s)
	return img
}
