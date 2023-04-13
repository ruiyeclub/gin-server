package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var Config *config

type config struct {
	App   App
	Db    Db
	Redis Redis
	Auth  Auth
}

func NewConfig() *config {
	cf := new(config)
	cf.loadConfigYaml()
	return cf
}

type App struct {
	Port string
	Name string
}

type Db struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	Charset  string
}

type Redis struct {
	Host     string
	Port     string
	Password string
	Database int
}

type Auth struct {
	Token bool
}

type Listener struct {
}

func init() {
	Config = NewConfig()
}

func (cf *config) loadConfigYaml() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./conf")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Println("conf路径未找到config.yaml，开始在根目录查询文件~")
		viper.AddConfigPath(".")
		err = viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("配置文件读取失败: %v \n", err))
		}
	}
	if err := viper.Unmarshal(cf); nil != err {
		log.Fatalf("赋值配置对象失败，异常信息：%v", err)
	}
}
