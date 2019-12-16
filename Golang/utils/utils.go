package utils


import (
	"fmt"
	"Pro_golang/Golang/config"
	//"Pro_golang/Golang/TestCase/SearchCase"
)

func utils() {
	fmt.Println("this is utils package")
}

func init(){
	for i ,v range config.CaseConfig {
		Println(i,v)
	}
}