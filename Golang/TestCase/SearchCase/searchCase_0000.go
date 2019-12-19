package SearchCase


import (
	"fmt"
	"reflect"
)

var SearchCaseList []interface{}

func init(){
	fmt.Println("this is searchCase init")
	// var SearchCase_0002 SearchCase_0002
	SearchCaseList = append(SearchCaseList,SearchCase_0002{})
	fmt.Println("SearchCaseList:",SearchCaseList)
	p := reflect.TypeOf(SearchCaseList[0])
	SearchCaseList[0].RunCaseProcess()
	c := reflect.ValueOf(SearchCaseList[0])
	fmt.Println("p:",p)
	fmt.Println("c:",c)
}