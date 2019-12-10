package config

import (
	"fmt"
)

func config(){
	fmt.Println("config")
}

type Yaml struct {
	Mysql struct {
		User 	string `yaml:"user"`
		Host	string `yaml:"host"`
		Password string `yaml:"password"`
		Port	string  `yaml:"port"`
		Name	string  `yaml:"name"`
	}
	Cache struct {
		Enable  bool `yaml:"enable"`
		List	[]string `yaml:"list,flow"`
	}
}


type Yaml1 struct {
	SQLConf Mysql `yaml:"mysql"`
	CacheConf Cache `yaml:"cache"`
}


type Yaml12 struct {
	Mysql 	`yaml:"mysql,inline"`
	Cache 	`yaml:"cache,inline"`
}

type Mysql struct {
	User 	string  	`yaml:"user"`
	Hose    string 		`yaml:"host"`
	Password  string 	`yaml:"password"`
	Port 	string 		`yaml:"port"`
	Name 	string  	`yaml:"name"`
}

type Cache struct {
	Enable 	bool 	`yaml:"enable"`
	List 	[]string `yaml:"list,flow"`
}
