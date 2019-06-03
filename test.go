package main

import (
	"communicate/wx"
	"encoding/json"
	"io/ioutil"
	"log"
)

func main() {
	data,err := ioutil.ReadFile("init_user_data")
	if err != nil{
		log.Fatal(err)
	}

	var initD wx.InitData

	if err := json.Unmarshal(data, &initD); err != nil{
		log.Fatal(err)
	}
	log.Print(initD)
}
