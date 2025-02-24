package main

import (
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"github.com/neverTanking/TiktokByGo/router"
)

func main() {
	go service.RunMessageServer()

	r := gin.Default()

	router.InitRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
