package auth

import (
	"testing"
	"fmt"
)


func TestLogin(t *testing.T) {
	token,err := Login("http://www.baidu.com","tkg","123456")
	fmt.Println("err is ",err)
	fmt.Println("token is ",token)
}


func TestGetInput(t *testing.T) {
	u := GetInput("username")
	fmt.Println("username is ",u)
}
