package main

// The intent of this is to expose a httpbin API as implemented by
// mccutchen/go-httpbin/v2, only adding the http server around it,
// along with logging the request details to stdout as JSON.
//
// The request details are used to plot a histogram / graph of the
// requests as they hit the service. This is particularly applicable
// to testing rate limiters and seeing the request count / rate.
//
// Logging graceful shutdown and program exit is done with `log` package,
// which logs the output into the standard error, so it doesn't interfere
// with the JSON output in standard output.

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/SentimensRG/ctx/sigctx"
	"github.com/mccutchen/go-httpbin/v2/httpbin"
)

func main() {
	if err := start(); err != nil {
		// Don't log expected server shutdown error
		if errors.Is(err, http.ErrServerClosed) {
			return
		}

		log.Fatal(err)
	}
}

type logEntry struct {
	Time time.Time
	httpbin.Result
}

func start() error {
	var (
		addr   = ":8085"
		output = "stdout"
	)
	flag.StringVar(&addr, "addr", addr, "address to listen to")
	flag.StringVar(&output, "output", output, "output, 'stdout' or file")
	flag.Parse()

	ctx := sigctx.New()

	var encoder *json.Encoder
	switch output {
	case "stdout":
		encoder = json.NewEncoder(os.Stdout)
	default:
		log.Println("Writing out log to", output)
		f, err := os.Create(output)
		if err != nil {
			return err
		}
		defer f.Close()

		encoder = json.NewEncoder(f)
	}

	var observerMu sync.Mutex
	observer := func(res httpbin.Result) {
		observerMu.Lock()
		defer observerMu.Unlock()

		encoder.Encode(logEntry{
			Time:   time.Now(),
			Result: res,
		})
	}

	srv := &http.Server{
		Addr:    addr,
		Handler: httpbin.New(httpbin.WithObserver(observer)),
	}

	log.Println("Starting server on", addr)

	go func() {
		<-ctx.Done()
		log.Println("Graceful shutdown")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		srv.Shutdown(shutdownCtx)
	}()

	return srv.ListenAndServe()
}
