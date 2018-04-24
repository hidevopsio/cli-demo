package auth

import (
	"testing"
	"github.com/magiconair/properties/assert"
)

func TestReadYaml(t *testing.T) {
	conf := ReadYaml()
	username := conf.Hicli.Clusters[0].Username
	lastIndex := conf.Hicli.LastIndex
	assert.Equal(t, username,"")
	assert.Equal(t, lastIndex, 0)
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
