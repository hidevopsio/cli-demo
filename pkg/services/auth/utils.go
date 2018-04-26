package auth

import (
	"fmt"
	"github.com/hidevopsio/hicli/pkg/common"
	"net/http"
	"bytes"
	"io/ioutil"
	"errors"
	"strings"
	"github.com/manifoldco/promptui"
	"encoding/json"
)

//定义用户用以HTTP登陆的JSON对象
type LoginAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//检查URL是否合法
func CheckUrl(url string) bool {
	return strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")
}

//通过HTTP登陆，返回Token
func Login(url, username, password string) (token string, err error) {
	myAuth := LoginAuth{Username: username, Password: password}
	jsonByte, err := json.Marshal(myAuth)
	if err != nil {
		fmt.Println("Login Failed ", err)
		return token, err
	}
	myToken := common.HicicdResponse{}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonByte))
	if err == nil {
		defer resp.Body.Close()
		byteResp, _ := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(byteResp, &myToken)
		if err == nil {
			if myToken.Code == 200 {
				token = myToken.Data
				err = errors.New(myToken.Message)
			} else {
				err = errors.New(myToken.Message)
			}
		}
	} else {
		//隐藏登陆完整URL信息
		errs := strings.Split(err.Error(), ":")
		err = errors.New(errs[len(errs)-1])
	}

	return token, err
}

//收集用户终端输入
func GetInput(label string) (userInput string) {
	validate := func(input string) error {
		if len(input) < 8 {
			return errors.New("Password must have more than 8 characters")
		}
		return nil
	}
	checkName := func(input string) error {
		if label == "Username" {
			if input == "" {
				return errors.New("Please Input username!")
			}
		}
		return  nil
	}
	if label == "Password" {
		u := promptui.Prompt{
			Label:    label,
			Mask:     '*',
			Validate: validate,
		}
		userInput, _ = u.Run()
	} else {
		u := promptui.Prompt{
			Label: label,
			Validate: checkName,
		}
		userInput, _ = u.Run()
	}
	return userInput
}

//专门用于获取用户输入的SERVER，并做校验
func GetInputServer(server string) (userInput string) {
	checkURL := func(input string) error {
		serverStrs := strings.Split(server, "[")
		defaultServer := serverStrs[len(serverStrs)-1]
		if strings.HasPrefix(defaultServer, "http://") || strings.HasPrefix(defaultServer, "https://") {
			if input == "" {
				return nil
			}
		}
		if ! (strings.HasPrefix(input, "http://") || strings.HasPrefix(input, "https://")) {
			return errors.New("Please Use Like http://SERVER:PORT OR https://SERVER:PORT")
		}
		return nil
	}
	u := promptui.Prompt{
		Label:    server,
		Validate: checkURL,
	}
	userInput, _ = u.Run()
	return userInput
}
