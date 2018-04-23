package auth

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestGetHomeDir(t *testing.T) {
	homeDir,_ := GetHomeDir()
	fmt.Println(homeDir)
	assert.Equal(t,homeDir,`C:\Users\vpclu`)
}

func TestPathExists(t *testing.T) {
	exists:= PathExists(`C:\Users\vpclu\.hicli`)
	assert.Equal(t,exists,true)
}