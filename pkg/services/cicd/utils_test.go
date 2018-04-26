package cicd

import (
	"testing"
	"fmt"
	"github.com/hidevopsio/hicli/pkg/common"
	"github.com/stretchr/testify/assert"
)

func TestToHttpData(t *testing.T) {
	env := common.EnvOptions{App: "name", Project: "cheng", Name: "TEST", Profile: "dev"}
	jStr, _ := ToHttpData(env)
	fmt.Println(jStr)
}

func TestInitEnvOpt(t *testing.T) {
	name := ""
	app := ""
	project := ""
	profile := ""
	env, err := InitEnvOpt(name, profile, app, project)
	if err == nil {
		fmt.Println("No Erro,Env is", env)
	} else {
		fmt.Println("err is", err, env)
	}
}

func TestGetToken(t *testing.T) {
	token, url := GetTokenUrl()
	assert.Equal(t, token, "")
	assert.Equal(t, url, "http://127.0.0.1:8080")
}

func TestCICDRun(t *testing.T) {
	token, url := GetTokenUrl()
	env := common.EnvOptions{
		App:     "hello-world",
		Name:    "java",
		Profile: "dev",
		Project: "demo",
	}
	err := CICDRun(url, token, env)
	if err == nil {
		fmt.Println("Deploy Sucessfuly")
	} else {
		fmt.Println("Deploy Failed", err)
	}
}
