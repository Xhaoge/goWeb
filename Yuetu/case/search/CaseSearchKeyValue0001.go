package searchCase

import (
	"fmt"
	"Pro_golang/Yuetu/case"
)


var CaseSearchKeyValue0001 = case.CaseStruct{
	"CaseSearchKeyValue0001","searchcase"
}
case.SearchList = append(SearchList,CaseSearchKeyValue0001)

func (s CaseSearchKeyValue0001)TestInit(){
	fmt.Println("this is CaseSearchKeyValue0001 TestInit")
}

func (s CaseSearchKeyValue0001)TestProcess(){
	fmt.Println("this is CaseSearchKeyValue0001 TestProcess")
}

func (s CaseSearchKeyValue0001)TestResult(){
	fmt.Println("this is CaseSearchKeyValue0001 TestResult")
}

