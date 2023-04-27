package server

import (
	"fmt"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/zjc17/pprof-web/internal/files"
	"log"
	"net/http"
	"net/http/pprof"
)

const (
	pprofWebPath = "/pprof/"
	// maxUploadSize 32 MiB
	maxUploadSize = 32 << 20
)

type (
	// LaunchParams for gin server
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
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	router.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html", []byte(PageDocument))
	})

	router.POST("/submit", func(c *gin.Context) {
		// check upload size
		if err := c.Request.ParseMultipartForm(maxUploadSize); err != nil {
			c.Data(http.StatusBadRequest, "text/plain; charset=UTF-8", []byte("file too large"))
			return
		}
		// parse file form
		fileHeader, err := c.FormFile("file")
		if err != nil {
			log.Printf("get file header error: %v", err)
			c.Status(http.StatusBadRequest)
			return
		}
		// store file into tmp path
		fileID := files.GenerateFileID()
		filePath := files.GetFilePathByFileID(fileID)
		if err := c.SaveUploadedFile(fileHeader, filePath); err != nil {
			log.Printf("save file error: %v", err)
			c.Status(http.StatusInternalServerError)
			return
		}
		// use pprof to analyze
		c.Redirect(http.StatusTemporaryRedirect, pprofWebPath+"?file_id="+fileID)
	})
	router.Any(pprofWebPath+"*pprof-path", func(c *gin.Context) {
		fileID := c.Query("file_id")
		filePath := files.GetFilePathByFileID(fileID)
		if fileID == "" {
			c.Redirect(http.StatusTemporaryRedirect, "/")
			return
		}
		mux, err := startPProfServer(filePath)
		if err != nil {
			log.Printf("start pprof server error: %v", err)
			c.Status(http.StatusInternalServerError)
			return
		}
		mux.ServeHTTP(c.Writer, c.Request)
	})
	RegisterDebugRouter(router)
	return router.Run(fmt.Sprintf(":%d", params.Port))
}

func RegisterDebugRouter(router *gin.Engine) {
	wrapHandler := func(handler func(w http.ResponseWriter, r *http.Request)) gin.HandlerFunc {
		return func(c *gin.Context) {
			handler(c.Writer, c.Request)
		}
	}
	router.Any("/debug/pprof/", wrapHandler(pprof.Index))
	router.Any("/debug/pprof/cmdline", wrapHandler(pprof.Cmdline))
	router.Any("/debug/pprof/profile", wrapHandler(pprof.Profile))
	router.Any("/debug/pprof/symbol", wrapHandler(pprof.Symbol))
	router.Any("/debug/pprof/trace", wrapHandler(pprof.Trace))
}
