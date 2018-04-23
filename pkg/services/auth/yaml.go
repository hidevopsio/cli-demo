package auth

import (
	"github.com/hidevopsio/hicli/config"
	"fmt"
	"github.com/hidevopsio/hi/pkg/system"
)

const (
	//application = "application"
	name      = "client"
	yaml        = "yml"
)


type Boot struct {
	config *config.Configuration
}

//读取用户YAML配置文件
func ReadYaml() *config.Configuration {
	userHomeDir,err := GetHomeDir()
	if err != nil {
		fmt.Println("Get Home Dir Failed",err)
	}
	fmt.Println(userHomeDir)
	builder := &system.Builder{
		Path:       userHomeDir,
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
	for _,v := range conf.Hicli.Clusters {
		if v.Cluster == url && v.Username == username {
			fmt.Println("User token is ",token)
			v.Token = token
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
		fmt.Println("Add the server to conf")
	}
	conf.Hicli.Clusters = servers
	//增加更新功能结束

	/*
	//初始化用户提供的相关信息
	newCluster := config.Cluster{
		Cluster:  url,
		Username: username,
		Token:    token,
	}
	//新集群相关信息追加进结构体
	servers := append(conf.Hicli.Clusters,newCluster)
	conf.Hicli.Clusters = servers
	*/

	//初始化build
	userHomeDir,_ := GetHomeDir()
	b := &system.Builder{
		Path:       userHomeDir,
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

//检查用户提供的URL是否存在于YAML文件中，如果存在，返回true。后续读出Token
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

//获取配置文件中的Token，前提保证URL存在于YAML文件中
func GetToken(url string)  string {
	conf := ReadYaml()
	for _,v := range conf.Hicli.Clusters {
		if v.Cluster == url {
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