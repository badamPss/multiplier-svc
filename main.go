package main

import (
	"flag"
	"log"
	"math/rand"
	"time"
)

var rtp float64

func main() {
	flag.Float64Var(&rtp, "rtp", 0, "rtp")
	flag.Parse()

	if rtp <= 0 || rtp > 1 {
		log.Fatal("error: rtp should be between 0 and 1")
	}

	rand.Seed(time.Now().UnixNano())
	startServer(rtp)
}
