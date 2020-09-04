package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func route(f func(ctx *gin.Context) interface{}) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, f(context))
	}
}

// PaymentOrderInParams 支付下单 前台 --> go 入参
type PaymentOrderInParams struct {
	MerchantNo     string `json:"merchantNo"`      // 商户号
	PayType        int    `json:"payType"`         // 1:支付宝H5支付 2：支付宝扫码支付 3：微信H5支付 4：微信公众号支付 5：微信公众号扫码支付
	OutTradeNo     string `json:"outTradeNo"`      // 商户订单号
	TotalAmount    int    `json:"totalAmount" `    // 支付总金额 单位:分
	Subject        string `json:"subject"`         // 支付标题
	Body           string `json:"body"`            // 订单描述
	TimeoutExpress int    `json:"timeoutExpress" ` // 支付有效分钟 default 5 min
	ReturnURL      string `json:"returnUrl"`       // 支付宝WAP支付需要传
	ClientIP       string `json:"clientIp" `       // 微信WAP支付传
	OpenID         string `json:"openId" `         // 微信公众号支付必传用户openid
}

func handle(c *gin.Context) {
	fmt.Println("post -----")
	fmt.Println("header -----", c.Request.Header.Get("Test"))
	var u = PaymentOrderInParams{}
	err := c.ShouldBind(&u)
	if err != nil {
		c.JSON(200, gin.H{"code": 0})
		return
	}
	fmt.Println(u.OutTradeNo)
	fmt.Println(u.Body)
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
