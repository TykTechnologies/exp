package main

import (
	"context"
	"flag"
	"log"
	"time"
)

func main() {
	config := &flags{}
	config.Bind()
	flag.Parse()

	t := time.Now()

	// os.Notify; CTRL+C;

	ctx := context.Background()
	if err := start(ctx, config); err != nil {
		log.Fatal(err)
		return
	}

	d := time.Since(t)
	seconds := d.Seconds()

	log.Printf("Rendered dir index in %.4f seconds", seconds)
}
