package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	yaml "gopkg.in/yaml.v2"
)


func init() {
	ReadYamlConfig()
	ReadCaseXML()
}


func ReadYamlConfig(){
	conf := new(Yaml)
	dir,_ := os.Getwd()
	str := dir+`\config\config.yaml`
	fmt.Println("str:",str)
	yamlFile, err := ioutil.ReadFile(str)
	log.Println("yamlFile:\n",string(yamlFile))
	if err != nil {
		log.Printf("%v:",err)
	}
	err = yaml.Unmarshal(yamlFile,conf)
	fmt.Println("conf:",conf.Mysql)
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
