package main

import (
	"fmt"
	"io/ioutil"
	"Pro_golang/Golang/config"
	yaml "gopkg.in/yaml.v2"
)


https://blog.csdn.net/kenkao/article/details/85071404

func main(){
	fmt.Println("hello world.......")
	conf := new(config.Yaml)
	yamlFile, err := ioutil.ReadFile("../case.yaml")
	log.Println("yamlFile:",yamlFile)
	if err != nil {
		log.Printf("%v:",err)
	}
	log.Println("conf:",conf)
	fmt.Println(yamlFile)
}

