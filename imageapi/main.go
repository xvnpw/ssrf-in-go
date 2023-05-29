package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/redirect", func(context *gin.Context) {
		fmt.Printf("-> /debug request: user-agent=%s\n", context.Request.UserAgent())

		targetFromUser := context.Query("target")

		context.Redirect(http.StatusMovedPermanently, targetFromUser)
	})
	r.Run()
}
