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
	/*
	if resp,err := http.Post(url, "application/json",bytes.NewBuffer(jsonByte));err == nil {
		byteResp,err:= ioutil.ReadAll(resp.Body)
		if err = json.Unmarshal(byteResp, &myToken); err == nil {
			//fmt.Println(myToken.Data)
			if myToken.Code == 200 {
				token = myToken.Data
			}
		}
	}
	*/
	return token,err
}

//收集用户终端输入的用户名与密码。不做错误检查，错误检查在Login函数会有检验
func GetInput() (username,password string)  {
	u := promptui.Prompt{
		Label:    "Username",
	}
	p:= promptui.Prompt{
		Label:    "Password",
		Mask:     '*',
	}
	username,_ = u.Run()
	password,_ = p.Run()
	return username,password
}