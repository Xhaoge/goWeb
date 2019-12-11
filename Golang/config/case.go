package config

import (
	"fmt"
	"os"
	"encoding/xml"
	"io/ioutil"
)

type Cases struct {
	Case	[]Case
}


type Case struct {
	Module		string `xml:"module"`
	Casename    string `xml:"casename"`
}


type Configuration struct {
	Enabled		bool 	`xml:"enabled"`
	Path 		string 	`xml:"path"`
}

func ReadCaseXML(){
	dir1,_ := os.Getwd()
	fmt.Println("dir1:",dir1)
	dir := dir1 + `\config\case.xml`
	fmt.Println("dir1:",dir)
	File, err := os.Open(dir)
	fmt.Println(File)
	if err != nil {
		fmt.Println("error opening file:",err)
		return
	}
	defer File.Close()

	// var conf Configuration
	// if err := xml.NewDecoder(xmlFile).Decode(&conf); err != nil {
	// 	fmt.Println("error decode file:",err)
	// 	return
	// }
	// fmt.Println(conf)
	// fmt.Println(conf.Enabled)
	// fmt.Println(conf.Path)

	xmlFile, err := ioutil.ReadAll(File)
	if err != nil {
		fmt.Println(err)
		return 
	}

	test := Cases{}
	err = xml.Unmarshal(xmlFile,&test)
	if err != nil {
		fmt.Println(err)
		return 
	}
	fmt.Println("test:",test)
	for _,c := range test.Case {
		fmt.Println("Casename:",c.Casename)
	}


}