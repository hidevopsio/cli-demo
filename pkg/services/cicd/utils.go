package cicd

import (
	"os"
	"strings"
	"encoding/json"
	"bytes"
	"github.com/hidevopsio/hicli/pkg/common"
	"github.com/hidevopsio/hicli/pkg/services/auth"
	"net/http"
	"io/ioutil"
	"errors"
	"time"
)

const cicdPath = "/cicd/run"

//初始化envOptions
func InitEnvOpt(name, profile, app, project string) (env common.EnvOptions, err error) {
	if currDir, err := os.Getwd(); err == nil {
		dirString := strings.Replace(currDir, `/`, `\`, -1)
		dirs := strings.Split(dirString, `\`)
		baseIndex := len(dirs)
		if app == "" {
			app = dirs[baseIndex-1]
		}
		if project == "" {
			project = dirs[baseIndex-2]
		}
		if profile == "" {
			profile = "dev"
		}
	}
	env.Name = name
	env.App = app
	env.Project = project
	env.Profile = profile
	return env, err
}

//传入待JSON化的Struce数据，转换成HTTP POST Data
func ToHttpData(i interface{}) (buffer *bytes.Buffer, err error) {
	jsonByte, err := json.Marshal(i)
	if err == nil {
		buffer = bytes.NewBuffer(jsonByte)
	}
	return buffer, err
}

//获取缓存的Token与URL。如果不存在，返回空
func GetTokenUrl() (token, url string) {
	conf := auth.ReadYAML()
	if len(conf.Hicli.Clusters) != 0 {
		lastIndex := conf.Hicli.LastIndex
		token = conf.Hicli.Clusters[lastIndex].Token
		url = conf.Hicli.Clusters[lastIndex].Cluster
	}
	return token, url
}

//收集校验完所有必需参数后，最终执行的动作
func CICDRun(url, token string, env common.EnvOptions) (err error) {
	serverResp := common.HicicdResponse{}
	client := &http.Client{Timeout: 600 * time.Second}
	postData, err := ToHttpData(env)
	req, err := http.NewRequest("POST", url+cicdPath, postData)
	if err == nil {
		req.Header.Add("Authorization", "Bearer "+token)
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	respByte, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(respByte, &serverResp)
	if err == nil {
		if serverResp.Code == 200 && serverResp.Message == "Successful." {
			err = nil
		} else {
			err = errors.New(serverResp.Message)
		}
	}
	defer resp.Body.Close()
	return err
}
