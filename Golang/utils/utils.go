package utils


import (
	"fmt"
	"Pro_golang/Golang/config"
	"reflect"
	//"Pro_golang/Golang/TestCase/Search"
	//"Pro_golang/Golang/TestCase/SearchCase"
)

func Utils() {
	fmt.Println("this is utils package")
}



func init(){
	fmt.Println("this is utils package")
	var regStruct map[string]interface{}
	regStruct = make(map[string]interface{})
	fmt.Println("regStruct:",regStruct)
	
	for _,c := range config.CaseConfig.Cases {
		fmt.Println(c.Casename)
		str := c.Casename
	
		fmt.Println("str:",str)
		// str2 := strconv.Itoa(str1)
		// str := string(str2)
		if regStruct[str] != nil {
		 	t := reflect.ValueOf(regStruct[str]).Type()
		 	v := reflect.New(t).Elem()
		 	fmt.Println("t,v:",t,v)
		// 	v.RunCaseProcess()

		//runCase := new(c.Casename{})
	 	}
	}
}