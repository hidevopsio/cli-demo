package auth

import (
	"testing"
	"github.com/magiconair/properties/assert"
	"fmt"
)

func TestReadYAML(t *testing.T) {
	conf := ReadYAML()
	if len(conf.Hicli.Clusters) != 0 {
		lastIndex := conf.Hicli.LastIndex
		username := conf.Hicli.Clusters[lastIndex].Username
		fmt.Println(lastIndex, username)
	}
	fmt.Println(conf)
}

func TestUpdateYAML(t *testing.T) {
	conf := ReadYAML()
	err := UpdateYAML(conf, "http://www.unknowname.com", "burtte", "Token")
	assert.Equal(t, err, nil)
}
