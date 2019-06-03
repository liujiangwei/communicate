package controller

import (
	"communicate/wx"
	"encoding/base64"
	"github.com/kataras/iris"
	"github.com/labstack/gommon/log"
)

func Login(ctx iris.Context) {
	client, err := wx.NewClient()
	if err != nil{
		_,_ = ctx.JSON(map[string]interface{}{
			"code" : 1,
			"message" : err.Error(),
		})
		return
	}

	code, err :=  client.GetCode()
	if err != nil{
		_,_ = ctx.JSON(map[string]interface{}{
			"code" : 1,
			"message" : err.Error(),
		})
		return
	}

	log.Print(base64.StdEncoding.EncodeToString(code))
	_,_ = ctx.JSON(map[string]interface{}{
		"code" : 0,
		"message" : 0,
		"data" : map[string]interface{}{
			"code" : base64.StdEncoding.EncodeToString(code),
			"uuid" : client.UUid,
		},
	})
}