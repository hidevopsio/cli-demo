package auth

import (
	"testing"
	"fmt"
	"os"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	username := os.Getenv("SCM_USERNAME")
	password := os.Getenv("SCM_PASSWORD")
	token, err := Login("http://127.0.0.1:8080/user/login", username, password)
	fmt.Println("token is ", token)
	assert.Equal(t, nil, err)
}


func TestGetInput(t *testing.T) {
	u := GetInput("username")
	fmt.Println("username is ", u)
}


