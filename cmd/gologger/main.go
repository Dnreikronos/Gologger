package main

import (
	"Gologger/internal/aggregator"
	"Gologger/internal/parser"
	"Gologger/internal/tailer"
	"Gologger/pkg/types"
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

func main() {
	flag.Usage = func() {
		log.Println("Usage: gologger -files=/path/to/log1,/path/to/log2")
	}
	filesArg := flag.String("files", "", "Comma-separated list of log file paths to monitor")
	flag.Parse()

	if *filesArg == "" {
		flag.Usage()
		os.Exit(1)
	}

	files := splitAndTrim(*filesArg)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go handleInterrupt(cancel)

	var wg sync.WaitGroup
	rawChan := make(chan types.LogLine, 100)
	parsedChan := make(chan types.ParsedLog, 100)

	for _, file := range files {
		err := tailer.TailFile(ctx, &wg, file, rawChan)
		if err != nil {
			log.Printf("Failed to tail file %s: %v", file, err)
		}
	}

	parser.StartParser(ctx, &wg, rawChan, parsedChan)

	agg := aggregator.NewAggregator()
	agg.Start(ctx, &wg, parsedChan)

	wg.Wait()
}

func splitAndTrim(s string) []string {
	var out []string
	for _, part := range strings.Split(s, ",") {
		trimed := strings.TrimSpace(part)
		if trimed != "" {
			out = append(out, trimed)
		}
	}
	return out
}

func handleInterrupt(cancelFunc context.CancelFunc) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
	log.Println("Interrupt received, shutting down....")
	cancelFunc()
}
