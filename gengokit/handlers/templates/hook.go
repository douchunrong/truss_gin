package templates

const Hook = `
package handlers

import (
	"context"

	"fmt"
	"{{.ImportPath -}} /svc"
	"os"
	"os/signal"
	"syscall"
)

`

const HookInterruptHandler = `
func InterruptHandler(errc chan<- error) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	terminateError := fmt.Errorf("%s", <-c)

	// Place whatever shutdown handling you want here

	errc <- terminateError
}
`
const HookSetConfig = `
func SetConfig(cfg svc.Config) svc.Config {
	// Place add custom ctx value 
	ctxMap := make(map[string]interface{})
	svc.SetCustomCtx(ctxMap)

	return cfg
}
`
