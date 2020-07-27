package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
	"unsafe"

	"github.com/qiniu/api.v7/v7/auth"
	"github.com/qiniu/api.v7/v7/sms"
)

/**
 * @author @liuqt
 * @date 2020/3/7
 */

// Manager send msg obj
var Manager *sms.Manager

func init() {
	accessKey := os.Getenv("accessKey")
	secretKey := os.Getenv("secretKey")

	mac := auth.New(accessKey, secretKey)
	Manager = sms.NewManager(mac)

}

// SendMessage 创蓝
func SendMessage(msg, phone string) {
	params := make(map[string]interface{})

	//请登录zz.253.com获取API账号、密码以及短信发送的URL
	params["account"] = "N6046102"        //创蓝API账号
	params["password"] = "1QpIL2Y5oC6065" //创蓝API密码
	params["phone"] = phone               //手机号码

	//设置您要发送的内容：其中“【】”中括号为运营商签名符号，多签名内容前置添加提交
	params["msg"] = url.QueryEscape("【南通应急指挥信息系统】" + msg)
	params["report"] = "true"
	bytesData, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	reader := bytes.NewReader(bytesData)
	url := "http://smssh1.253.com/msg/send/json" //短信发送URL
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	str := (*string)(unsafe.Pointer(&respBytes))
	fmt.Println(*str)
}

// SendCode 创蓝
func SendCode(msg, phone string) {
	params := make(map[string]interface{})

	//请登录zz.253.com获取API账号、密码以及短信发送的URL
	params["account"] = "YZM3313036"      //创蓝API账号
	params["password"] = "0xAPywQCYR3064" //创蓝API密码
	params["phone"] = phone               //手机号码

	//设置您要发送的内容：其中“【】”中括号为运营商签名符号，多签名内容前置添加提交
	params["msg"] = url.QueryEscape("【南通应急指挥信息系统】" + msg)
	params["report"] = "true"
	bytesData, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	reader := bytes.NewReader(bytesData)
	url := "http://smssh1.253.com/msg/send/json" //短信发送URL
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	str := (*string)(unsafe.Pointer(&respBytes))
	fmt.Println(*str)
}

// GenValidateCode 生成指定长度的数字验证码
func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

const accessKey = "tdVsm0u5zbaCcG2IvTyaL7Ret6VqCWWfAQ14QquG"
const secretKey = "5i4WNAdca3g-mpidakOolh5gHs7-JahmADw2gGxD"

// SendQiniuMessage 固定消息
func SendQiniuMessage(phone string) {
	//fmt.Println(accessKey,secretKey)
	mac := auth.New(accessKey, secretKey)
	Manager := sms.NewManager(mac)
	mob := make([]string, 0)
	mob = append(mob, phone)
	// mp := make(map[string]interface{}, 0)
	mes := sms.MessagesRequest{
		SignatureID: "1236537723758583808",
		TemplateID:  "1236847965105037312",
		Mobiles:     mob,
		// Parameters:  mp,
	}
	//fmt.Println(mob)
	Manager.SendMessage(mes)

}

// SendQiniuMessage2 可发送任意消息
func SendQiniuMessage2(msg, phone string) {
	//fmt.Println(accessKey,secretKey)
	mac := auth.New(accessKey, secretKey)
	Manager := sms.NewManager(mac)
	mob := make([]string, 0)
	mob = append(mob, phone)
	mp := make(map[string]interface{}, 0)
	mp["code"] = msg
	mes := sms.MessagesRequest{
		SignatureID: "1236537723758583808",
		TemplateID:  "1267620483851890688",
		Mobiles:     mob,
		Parameters:  mp,
	}
	//fmt.Println(mob)
	Manager.SendMessage(mes)

}

// SendQiniuCode 发送验证码
func SendQiniuCode(code, phone string) {

	mac := auth.New(accessKey, secretKey)
	Manager := sms.NewManager(mac)
	mob := make([]string, 0)
	mob = append(mob, phone)
	mp := make(map[string]interface{}, 0)
	mp["code"] = code
	mes := sms.MessagesRequest{
		SignatureID: "1236537723758583808",
		TemplateID:  "1236848264251187200",
		Mobiles:     mob,
		Parameters:  mp,
	}
	//fmt.Println(mob)
	Manager.SendMessage(mes)
	//fmt.Println(res,err)
}
