package auth

import (
	"github.com/hidevopsio/hicli/config"
	"fmt"
	"github.com/hidevopsio/hi/pkg/system"
	"path/filepath"
	"os"
)

const (
	name      = "client"
	yaml        = "yml"
)

type Boot struct {
	config *config.Configuration
}

//读取用户YAML配置文件
func ReadYaml() *config.Configuration {
	userHomeDir,err := GetHomeDir()
	yamlDir := filepath.Join(userHomeDir, ".hicli")
	if err != nil {
		fmt.Println("Get Home Dir Failed",err)
	}
	fmt.Println(userHomeDir)
	builder := &system.Builder{
		Path:       yamlDir,
		Name:       name,
		FileType:   yaml,
		ConfigType: config.Configuration{},
	}
	cp, err := builder.Build()
	if err != nil {
		fmt.Println("error",err)
	}
	c  := cp.(*config.Configuration)
	return c
}

//更新或添加YAML
func UpdateYAML(conf *config.Configuration, url,username,token string)  error {
	//增加更新功能开始
	exists := false
	var servers  []config.Cluster
	for index,v := range conf.Hicli.Clusters {
		if v.Cluster == url && v.Username == username {
			fmt.Println("User token is ",token)
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
		fmt.Println("Add the server to conf")
	}
	conf.Hicli.Clusters = servers
	//初始化build
	userHomeDir,_ := GetHomeDir()
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
	if err == nil {
		fmt.Println("Save Suceess")
	}
	return err
}

//检查用户提供的URL与用户名是否存在于YAML文件中，如果存在，返回true。后续读出Token
func CheckConf(url,username string)  bool {
	fileConf := ReadYaml()
	for _,v := range fileConf.Hicli.Clusters {
		if url == v.Cluster && username == v.Username {
			return true
		}
		fmt.Println(v.Cluster)
	}
	return false
}

//根据用户URL与用户名，返回YAML中相对应的Token
func GetToken(url,username string)  string {
	conf := ReadYaml()
	for _,v := range conf.Hicli.Clusters {
		if v.Cluster == url && v.Username == username {
			return v.Token
		}
	}
	return ""
}

//更新Token,主要应对Token失效的情况。
func UpdateToken(url,token string) {
	conf := ReadYaml()
	for _,v := range conf.Hicli.Clusters {
		if v.Cluster == url {
			v.Token = token
		}
	}
}

//获取YAML中的LastIndex。如果文件与文件夹都不存放，则创建它们并返回空
func GetLastIndex() (int,bool) {
	userHomeDir, _ := GetHomeDir()
	yamlDir := filepath.Join(userHomeDir,".hicli")
	if PathExists(yamlDir) {
		if PathExists(filepath.Join(yamlDir,"client.yml")) {
			conf := ReadYaml()
			//fmt.Println(conf.Hicli.LastIndex)
			return conf.Hicli.LastIndex,true
		} else {
			_, err := os.Create(filepath.Join(yamlDir,"client.yml"))
			if err != nil {
				fmt.Println("Create client.yml failed ",err)
			}
		}
	} else {
		err := os.Mkdir(yamlDir, 755)
		if err != nil {
			fmt.Println("Mkidr .hicli failed ",err)
		}
		if _, err := os.Create(filepath.Join(yamlDir,"client.yml"));err != nil {
			fmt.Println("Create client.yml failed ",err)
		}
	}
	return 0,false
}