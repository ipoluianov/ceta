package main

import (
	"github.com/ipoluianov/ceta/app"
	"github.com/ipoluianov/ceta/system"
	"github.com/ipoluianov/gomisc/logger"
)

func main() {
	logger.Init(logger.CurrentExePath() + "/logs")
	sys := system.NewSystem()
	app.Init("cetuspoolsui", "cetuspoolsui", func() {
		sys.Start()
	}, func() {
		sys.Stop()
	})
	if !app.TryService() {
		app.RunDesktop()
	}
}
