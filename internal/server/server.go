package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type (
	LaunchParams struct {
		Port uint16
	}
)

func DefaultLaunchParam() *LaunchParams {
	return &LaunchParams{
		Port: 8080,
	}
}

func Launch(params *LaunchParams) error {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html", []byte(PageDocument))
	})
	router.POST("/submit", func(c *gin.Context) {
		// parse file form
		fileHeader, err := c.FormFile("file")
		if err != nil {
			log.Printf("get file header error: %v", err)
			c.Status(http.StatusBadRequest)
			return
		}
		// store file into tmp path
		if err := c.SaveUploadedFile(fileHeader, "/tmp/pprof-web-file"); err != nil {
			log.Printf("save file error: %v", err)
			c.Status(http.StatusInternalServerError)
			return
		}
		// use pprof to analyze
		c.Redirect(http.StatusTemporaryRedirect, "/")
	})

	return router.Run(fmt.Sprintf(":%d", params.Port))
}
