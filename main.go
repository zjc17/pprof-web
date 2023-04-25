package main

import (
	"fmt"
	"github.com/zjc17/pprof-web/internal/server"
	"github.com/zjc17/pprof-web/internal/version"
	"log"
)

func main() {
	fmt.Printf("pprof web v%s\n\n", version.Version)

	launchParam := server.DefaultLaunchParam()
	if err := server.Launch(launchParam); err != nil {
		log.Fatal(err)
	}
}
