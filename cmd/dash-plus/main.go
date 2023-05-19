package main

import (
	"github.com/TykTechnologies/tyk-analytics/dashboard"
	"github.com/TykTechnologies/tyk/gateway"
)

func main() {
	var shutdownOnce sync.Once

	ctx, cancel := context.WithCancel(context.Background())
	shutdown := func() {
		shutdownOnce.Do(func() {
			println("Shutting down")

			// dashboard.Stop, gateway.Stop

			if err := recover(); err != nil {
				println("Panic caught:")
				println(err)
			}
			cancel()
		})
	}

	// Start gateway and dashboard
	go func() {
		defer shutdown()
		gateway.Start()
	}()
	go func() {
		defer shutdown()
		dashboard.Start()
	}()
	<-ctx.Done()

}
