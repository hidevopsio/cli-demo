package auth

import (
	"testing"
	"github.com/magiconair/properties/assert"
	"fmt"
)

func TestReadYaml(t *testing.T) {
	conf := ReadYaml()
	//servers := conf.Hicli.Clusters
	//username := conf.Hicli.Clusters[0].Username
	fmt.Println(conf)
	//lastIndex := conf.Hicli.LastIndex
	//assert.Equal(t, username,"")
	//assert.Equal(t, lastIndex, 1)
}

func TestUpdateYAML(t *testing.T) {
	conf := ReadYaml()
	err := UpdateYAML(conf, "http://www.unknowname.cn", "burtte", "Token")
	assert.Equal(t, err, nil)
}

func TestInitYAML(t *testing.T) {
	err := InitYAML()
	assert.Equal(t, err, nil)
}
