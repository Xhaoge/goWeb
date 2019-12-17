package main

import (
	"fmt"
	"Pro_golang/Golang/config"
	//"Pro_golang/Golang/TestCase/SearchCase"
	"Pro_golang/Golang/utils"
)

func main(){
	fmt.Println("hello world.......")
	fmt.Printf("CaseConfig main:%+v \n",config.CaseConfig)
	fmt.Printf("BaseConfig main: %+v \n",config.BaseConfig)
	//searchCase.RunCaseProcess()
	utils.Utils()
}

