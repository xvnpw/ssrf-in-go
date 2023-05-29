package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/internal", func(context *gin.Context) {
		fmt.Printf("-> /debug request: user-agent=%s\n", context.Request.UserAgent())

		context.String(http.StatusOK, "This is internal sensitive endpoint\n")
	})
	r.Run()
}
