package server

import (
	"github.com/google/pprof/driver"
	"github.com/zjc17/pprof-web/internal/pprof"
	"net/http"
	"path"
)

// startPProfServer start the pprof web handler without browser
func startPProfServer(filePath string) (mux *http.ServeMux, err error) {
	mux = http.NewServeMux()

	options := &driver.Options{
		Flagset: &pprof.Flags{
			Args: []string{"-http=localhost:0", "-no_browser", filePath},
		},
		HTTPServer: func(args *driver.HTTPServerArgs) error {
			for pattern, handler := range args.Handlers {
				if pattern == "/" {
					mux.Handle(pprofWebPath, handler)
				} else {
					mux.Handle(path.Join(pprofWebPath, pattern), handler)
				}
			}
			return nil
		},
	}
	if err = driver.PProf(options); err != nil {
		return
	}
	return
}