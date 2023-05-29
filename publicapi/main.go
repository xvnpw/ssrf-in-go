package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func validateTargetUrl(input string) bool {
	u, err := url.ParseRequestURI(input)
	if err != nil {
		return false
	}
	if (u.Scheme == "http" || u.Scheme == "https") &&
		(u.Hostname() == "imageapi") &&
		(u.Port() == "" || u.Port() == "80" || u.Port() == "443") {
		return true
	}

	return false
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	/* no protection */
	/*
		router.GET("/debug", func(context *gin.Context) {
			fmt.Printf("-> /debug request: user-agent=%s\n", context.Request.UserAgent())
			urlFromUser := context.Query("url")
			resp, err := http.Get(urlFromUser)
	*/

	/* positive validation */
	/*
		router.GET("/debug", func(context *gin.Context) {
			fmt.Printf("-> /debug request: user-agent=%s\n", context.Request.UserAgent())
			urlFromUser := context.Query("url")
			if !validateTargetUrl(urlFromUser) {
				context.String(http.StatusBadRequest, "Bad url")
				return
			}
			resp, err := http.Get(urlFromUser)
	*/

	/* safeurl */
	/*
		config := safeurl.GetConfigBuilder().
			SetAllowedHosts("imageapi").
			Build()
		router.GET("/debug", func(context *gin.Context) {
			fmt.Printf("-> /debug request: user-agent=%s\n", context.Request.UserAgent())
			urlFromUser := context.Query("url")
			client := safeurl.Client(config)
			resp, err := client.Get(urlFromUser)
	*/

	/* http with turned off redirects */
	/*
		http.DefaultClient.CheckRedirect = func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse }

		router.GET("/debug", func(context *gin.Context) {
			fmt.Printf("-> /debug request: user-agent=%s\n", context.Request.UserAgent())
			urlFromUser := context.Query("url")
			if !validateTargetUrl(urlFromUser) {
				context.String(http.StatusBadRequest, "Bad url")
				return
			}
			resp, err := http.Get(urlFromUser)
	*/

	http.DefaultClient.CheckRedirect =
		func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}

	router.GET("/debug", func(context *gin.Context) {
		fmt.Printf("-> /debug request: user-agent=%s\n", context.Request.UserAgent())
		urlFromUser := context.Query("url")
		resp, err := http.Get(urlFromUser)

		if err != nil {
			stringError := fmt.Errorf("request return error: %v", err)
			fmt.Print(stringError)
			context.String(http.StatusBadRequest, err.Error())
			return
		}

		defer resp.Body.Close()
		bodyString, err := io.ReadAll(resp.Body)

		if err != nil {
			context.String(http.StatusInternalServerError, err.Error())
			return
		}

		context.String(http.StatusOK, string(bodyString))
	})
	router.Run()
}
