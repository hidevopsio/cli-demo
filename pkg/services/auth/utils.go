package auth

import (
	"fmt"
	"strings"
	"encoding/json"
	"net/http"
	"bytes"
	"io/ioutil"
	"errors"
	"github.com/manifoldco/promptui"
)

//定义HTTP获取后的对象
type MyToken struct {
	Code 	int  	`json:"code"`
	Message string	`json:"message"`
	Data    string	`json:"data"`
}

//定义用户用以HTTP登陆的JSON对象
type LoginAuth struct {
	Username string 	`json:"username"`
	Password string 	`json:"password"`
}



//检查URL是否合法
func CheckUrl(url string) bool  {
	return strings.HasPrefix(url,"http://") || strings.HasPrefix(url,"https://")
}


//登陆函数，返回Token
func Login(url,username,password string)  (token string,err error) {
	myAuth := LoginAuth{Username:username,Password:password}
	jsonByte,err := json.Marshal(myAuth)
	if err != nil {
		fmt.Println("Error ",err)
		return  token,err
	}
	myToken := MyToken{}
	resp,err := http.Post(url,"application/json",bytes.NewBuffer(jsonByte))
	if err == nil {
		defer resp.Body.Close()
		byteResp,_:= ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(byteResp, &myToken)
		if err == nil {
			if myToken.Code == 200 {
				token = myToken.Data
			} else {
				err = errors.New(myToken.Message)
			}
		}
	}

	return token,err
}

//收集用户终端输入的Username或者URL.通过label指定类型
func GetInput(label string) string  {
	u := promptui.Prompt{
		Label:    label,
	}
	username,_ := u.Run()
	return username
}

//获取终端用户的Password
func GetInputPassword() string {
	p := promptui.Prompt{
		Label:  "password",
		Mask:   '*',
	}
	password,_ := p.Run()
	return  password
}