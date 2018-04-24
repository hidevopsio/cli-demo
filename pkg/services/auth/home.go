package auth

import (
	"os/user"
	"runtime"
	"os"
	"bytes"
	"os/exec"
	"strings"
	"errors"
	"fmt"
)

//获取用户HOME目录
func GetHomeDir() (string, error) {
	user, err := user.Current()
	if nil == err {
		//fmt.Println("get dir")
		return user.HomeDir, nil
	}

	if "windows" == runtime.GOOS {
		fmt.Println("windows")
		return homeWindows()
	}
	fmt.Println("linux")
	return homeUnix()
}

//获取*unix系统家目录，不对外提供服务。给GetHomeDir调用
func homeUnix() (string, error) {
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}
	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output when reading home directory")
	}
	fmt.Println("linux")
	return result, nil
}

//获取Windows系统家目录，不对外提供服务。给GetHomeDir调用
func homeWindows() (string, error) {
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
	}
	fmt.Println("windwos ")
	return home, nil
}

//检查指定目录或者文件是否存在
func PathExists(filePath string) (bool) {
	_, err := os.Stat(filePath)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
