package conf

import (
	"fmt"
	"gopkg.in/ini.v1"
	"strings"
	"user/model"
)

var (
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
)

func Init() {
	//读取config.ini配置文件
	file, err := ini.Load("conf/config.ini")
	if err != nil {
		fmt.Println("Fail to read file: ", err)
	}

	LoadMysqlData(file)
	path := strings.Join([]string{DbUser, ":", DbPassword, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}, "")
	print(path)
	model.Database(path)
}
func LoadMysqlData(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassword = file.Section("mysql").Key("DbPassword").String()
	DbName = file.Section("mysql").Key("DbName").String()

}
