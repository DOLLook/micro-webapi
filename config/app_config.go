// 配置文件包
package config

import (
	"fmt"
	"go-webapi-fw/common/mongoutils"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// 应用程序配置
var AppConfig = &appConfig

// 程序配置
var appConfig = struct {
	Server struct {
		Name string `yaml:"name"`
		Port string `yaml:"port"`
	} `yaml:"server"`

	MySql      string `yaml:"mysql"`
	MongodbUri string `yaml:"mongodb_uri"`

	Redis struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Password string `yaml:"password"`
		// 超时时间 毫秒
		Timeout int32 `yaml:"timeout"`
	} `yaml:"redis"`

	Rabbitmq struct {
		Address string `yaml:"address"`
	} `yaml:"rabbitmq"`

	Consul struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"consul"`
}{}

func init() {
	bytes, err := ioutil.ReadFile("application.yml")
	if err != nil {
		mongoutils.Error(err)
		fmt.Println(err)
		return
	}

	if err := yaml.Unmarshal(bytes, AppConfig); err != nil {
		mongoutils.Error(err)
		fmt.Println(err)
	}
}
