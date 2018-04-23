package auth

import (
	"testing"
	"github.com/magiconair/properties/assert"
	"fmt"
)

func TestReadYaml(t *testing.T) {
	conf := ReadYaml()
	//servers := conf.Hicli.Clusters
	username := conf.Hicli.Clusters[0].Username
	fmt.Println(username)
	lastIndex := conf.Hicli.LastIndex
	assert.Equal(t, username,"foo")
	assert.Equal(t, lastIndex, "120")
}

func TestUpdateYAML(t *testing.T) {
	conf := ReadYaml()
	err := UpdateYAML(conf,"http://www.unknowname.win","burtte","Token")
	assert.Equal(t, err, nil)
}

func TestCheckConf(t *testing.T) {
	url := "http://www.biadu.com"
	assert.Equal(t,CheckConf(url,"cheng"),true)
}

func TestGetToken(t *testing.T) {
	token := GetToken("http://www.unknowname.cn", "burtte")
	assert.Equal(t,token,"Token")
}


func TestGetLastIndex(t *testing.T) {
	index,exists := GetLastIndex()
	assert.Equal(t,exists,true)
	assert.Equal(t,index,0)
}