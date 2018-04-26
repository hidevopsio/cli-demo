package auth

import (
	"path/filepath"
	"github.com/hidevopsio/hiboot/pkg/system"
	"fmt"
	"github.com/hidevopsio/hicli/pkg/config"
)

const (
	name     = "client"
	fileType = "yml"
)

var userHomeDir, _ = GetHomeDir()
var configDir = filepath.Join(userHomeDir, ".hicli")

//读取client.yml
func ReadYAML() *config.Configuration {
	builder := &system.Builder{
		Path:       configDir,
		Name:       name,
		FileType:   fileType,
		ConfigType: config.Configuration{},
	}
	err := builder.Init()
	if err != nil {
		fmt.Println("Read config file failed", err)
	}
	cp, err := builder.Build()
	if err != nil {
		fmt.Println("Read config file failed", err)
	}
	c := cp.(*config.Configuration)
	return c
}

//更新或添加client.yml
func UpdateYAML(conf *config.Configuration, url, username, token string) error {
	exists := false
	var servers []config.Cluster
	//检查相关URL与用户名是否存在，存在更新
	for index, v := range conf.Hicli.Clusters {
		if v.Cluster == url && v.Username == username {
			v.Token = token
			conf.Hicli.LastIndex = index
			exists = true
		}
		servers = append(servers, v)
	}
	//不存在则添加
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
	}
	conf.Hicli.Clusters = servers
	//初始化build
	b := &system.Builder{
		Path:       configDir,
		Name:       name,
		FileType:   fileType,
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
