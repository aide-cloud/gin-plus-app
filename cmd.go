package main

import (
	"gin-plus-app/logic"
	ginplus "github.com/aide-cloud/gin-plus"
	"github.com/gin-gonic/gin"
)

var httpMethodPrefixes = []ginplus.HttpMethod{
	{
		Prefix: "Create",
		Method: ginplus.Post,
	},
	{
		Prefix: "Update",
		Method: ginplus.Put,
	}, {
		Prefix: "Delete",
		Method: ginplus.Delete,
	}, {
		Prefix: "Detail",
		Method: ginplus.Get,
	}, {
		Prefix: "List",
		Method: ginplus.Get,
	},
}

func main() {
	r := gin.Default()
	ginEngine := ginplus.New(r,
		ginplus.AppendHttpMethodPrefixes(httpMethodPrefixes...),
		ginplus.WithControllers(logic.NewApi()),
	)
	ginEngine.RegisterSwaggerUI()
	ginplus.NewCtrlC(ginEngine).Start()
}
