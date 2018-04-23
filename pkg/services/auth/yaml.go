package auth

import (
	"github.com/hidevopsio/hicli/config"
	"fmt"
	"github.com/hidevopsio/hi/pkg/system"
	"path/filepath"
	"os"
)

const (
	name = "client"
	yaml = "yml"
)

type Boot struct {
	config *config.Configuration
}

//读取用户YAML配置文件
func ReadYaml() *config.Configuration {
	userHomeDir, err := GetHomeDir()
	yamlDir := filepath.Join(userHomeDir, ".hicli")
	InitYAML()
	if err != nil {
		fmt.Println("Get Home Dir Failed", err)
	}
	//fmt.Println(userHomeDir)
	builder := &system.Builder{
		Path:       yamlDir,
		Name:       name,
		FileType:   yaml,
		ConfigType: config.Configuration{},
	}
	cp, err := builder.Build()
	if err != nil {
		fmt.Println("error", err)
	}
	c := cp.(*config.Configuration)
	return c
}

//更新或添加YAML
func UpdateYAML(conf *config.Configuration, url, username, token string) error {
	//增加更新功能开始
	exists := false
	var servers []config.Cluster
	for index, v := range conf.Hicli.Clusters {
		if v.Cluster == url && v.Username == username {
			v.Token = token
			conf.Hicli.LastIndex = index
			exists = true
		}
		servers = append(servers, v)
	}
	//追加
	if ! exists {
		newCluster := config.Cluster{
			Cluster:  url,
			Username: username,
			Token:    token,
		}
		//新集群相关信息追加进结构体
		servers = append(conf.Hicli.Clusters, newCluster)
		lastIndex := len(servers) - 1
		conf.Hicli.LastIndex = lastIndex
		//fmt.Println("Add the server to conf")
	}
	conf.Hicli.Clusters = servers
	//初始化build
	userHomeDir, _ := GetHomeDir()
	yamlPath := filepath.Join(userHomeDir, ".hicli")
	b := &system.Builder{
		Path:       yamlPath,
		Name:       name,
		FileType:   yaml,
		ConfigType: config.Cluster{},
	}
	err := b.Init()
	if err != nil {
		fmt.Println(err);
		return err
	}
	err = b.Save(conf)
	return err
}

//当client.yml文件不存在时，创建一个空白文件
func InitYAML() (err error) {
	userHomeDir, err := GetHomeDir()
	yamlDir := filepath.Join(userHomeDir, ".hicli")
	yamlFile := filepath.Join(yamlDir, "client.yml")
	if PathExists(yamlDir) {
		if ! PathExists(yamlFile) {
			_, err = os.Create(yamlFile)
		}
	} else {
		err = os.Mkdir(yamlDir, 755)
		_, err = os.Create(yamlFile)
	}
	return err
}
