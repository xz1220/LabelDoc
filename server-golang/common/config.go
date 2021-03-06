package common

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

// 用户角色配置
var RoleAdmin string = "ROLE_ADMIN"
var RoleUser string = "ROLE_USER"
var RoleReviewer string = "ROLE_REVIEWER"

// 目录限制
var prefix = "/Users/xingzheng"
var IMAGE_DIC = prefix + "/labelprojectdata/image/"
var IMAGE_S_DIC = prefix + "/labelprojectdata/images/"
var IMAGE_L_DIC = prefix + "/labelprojectdata/imagel/"
var IMAGE_DELETE_DIC = prefix + "/labelprojectdata/imaged/"
var VIDEO_DIC = prefix + "/labelprojectdata/video/"
var VIDEO_D_DIC = prefix + "/labelprojectdata/videod/"
var VIDEO_S_DIC = prefix + "/labelprojectdata/videos/"
var LIMITED_LENGTH = 4194304 // 4MB

//init config file
func InitConfig(workDir string) {
	if workDir == "main" {
		Dir, err := os.Getwd()
		if err != nil {
			panic("路径读取失败")
		}
		workDir = Dir
	}
	viper.AddConfigPath(workDir + "/config")
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	err := viper.ReadInConfig()
	if err != nil {
		panic("读取配置文件失败")
	}
	log.Print("读取成功")
}
