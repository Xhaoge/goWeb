package SearchCase

import (
	"fmt"
	// "net/http"
	// "bytes"
	// "io/ioutil"
	// "reflect"
)


type SearchCase_0002 struct {}


func  (n *SearchCase_0002)RunCaseProcess(){
	fmt.Println("this is case2 runcaseprocess")
	// url := `http://dev-api.gloryholiday.com/yuetu/search`
	// data := `{
	// 		    "Cid": "ctrip",
	// 		    "TripType": "1",
	// 		    "FromCity": "HKG",
	// 		    "ToCity": "LAX",
	// 		    "FromDate": "20200523",
	// 		    "RetDate": "20200821",
	// 		    "AdultNumber": 1,
	// 		    "ChildNumber": 0,
	// 		    "InfantNumber":0,
	// 		    "Currency":"CNY",
	// 		    "BypassCache": false,
	// 		    "GodPerspective":false }`

	// fmt.Println(url,data)
	// urlPost := "http://test-restful-api.gloryholiday.com/currencyservice/getCurrency"
	// body := `{
	// 	    "originalCode":"USD",
	// 		"targetCode":"CNY",
	// 		"publish_timestamp": "2019-11-28T00:00:00Z"}`
	// fmt.Println("body type:",reflect.TypeOf(body))

	// res,err := http.Post(url,"application/json;charset=utf-8",bytes.NewBuffer([]byte(data)))
	// fmt.Println("res body:",res.Body)
	// if err != nil {
	// 	fmt.Println("err:",err)
	// }
	// defer res.Body.Close()
	// content,err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("err ",err)
	// }
	// fmt.Println(string(content))
	// fmt.Println("content type:",reflect.TypeOf(content))
	// fmt.Println("content value:",reflect.ValueOf(content))

}
