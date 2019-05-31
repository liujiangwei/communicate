package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
)


func main() {
	data, err := ioutil.ReadFile("login_page")
	if err != nil{
		log.Fatal(err)
	}

	var ld LoginData
	if err = xml.Unmarshal(data, &ld); err != nil{
		log.Fatal(err)
	}

	log.Print(ld)
}
