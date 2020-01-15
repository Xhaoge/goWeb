package Case

import (
	"fmt"
	// _"Pro_golang/Yuetu/Case/search"
)
type CaseStruct struct {
	CaseName	string
	Belong		string	
}

var SearchList []CaseStruct


// type CaseList struct{
// 	SearchList 		[]CaseStruct
// 	VerifyList		[]CaseStruct
// 	OrderLise		[]CaseStruct
// 	OthersList		[]CaseStruct
// }


func init(){
	fmt.Println("this is case init")
	fmt.Println("SearchList:",SearchList)
}

func CaseTest(){
	fmt.Println("this is testcase package")
}