package config

import (
	"fmt"
	"os"
	"encoding/xml"
)

//https://blog.csdn.net/wade3015/article/details/83351776

type Configuration struct {
	Enabled		bool 	`xml:"enabled"`
	Path 		string 	`xml:"path"`
}

func ReadCaseXML(){

	xmlFile, err := os.Open(dir)
}