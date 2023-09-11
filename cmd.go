package main

import (
	"gin-plus-app/logic"
	"gin-plus-app/logic/enterprise/contract"
	"gin-plus-app/logic/graphql"
	"gin-plus-app/middler"
	ginplus "github.com/aide-cloud/gin-plus"
	"github.com/gin-gonic/gin"
)

var httpMethodPrefixes = []ginplus.HttpMethod{
	{
		Prefix: "Create",
		Method: ginplus.Post,
	}, {
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
	// 初始化gin
	r := gin.Default()
	// 注册中间件
	r.Use(
		middler.RecoverMiddleware(),
		middler.LoggerMiddleware(),
		middler.MetricSecondsMiddleware(),
		middler.IpMetricMiddleware(),
		middler.MetricRequestsMiddleware(),
	)

	// 初始化gin-plus
	ginEngine := ginplus.New(r,
		ginplus.AppendHttpMethodPrefixes(httpMethodPrefixes...),
		ginplus.WithControllers(logic.NewApi()),
		ginplus.WithGraphqlConfig(ginplus.GraphqlConfig{
			Enable:     true,
			HandlePath: "graphql",
			ViewPath:   "graphql",
			Root:       graphql.NewResolver(),
			Content:    graphql.Sdl,
		}),
	)

	ginEngine.GET("/say", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	// 注册公共无权限路由
	ginEngine.RegisterSwaggerUI().RegisterPing().RegisterMetrics().RegisterGraphql()
	// 注册公共有权限路由
	authGroup := r.Group("/auth", middler.MustLogin())
	// 注册企业路由
	enterpriseGroup := authGroup.Group("/enterprise/:eid")
	ginEngine.GenRoute(enterpriseGroup.Group("/api/v2/enterprise/:eid"), contract.New())
	// 启动gin-plus
	ginplus.NewCtrlC(ginEngine).Start()
}
