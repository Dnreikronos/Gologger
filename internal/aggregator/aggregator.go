package aggregator

import (
	"Gologger/pkg/types"
	"context"
	"log"
	"sync"
	"time"
)

type Aggregator struct {
	mu     sync.Mutex
	counts map[string]map[string]int // file -> level -> count
}

func NewAggregator() *Aggregator {
	return &Aggregator{
		counts: make(map[string]map[string]int),
	}
}

func (a *Aggregator) Start(ctx context.Context, wg *sync.WaitGroup, in <-chan types.ParsedLog) {
	wg.Add(1)
	go func() {
		defer wg.Done()

		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				a.printStats()
				return
			case parsed := <-in:
				a.add(parsed)
			case <-ticker.C:
				a.printStats()
			}
		}
	}()
}

func (a *Aggregator) add(log types.ParsedLog) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if _, ok := a.counts[log.Filename]; !ok {
		a.counts[log.Filename] = make(map[string]int)
	}
	a.counts[log.Filename][log.Level]++
}

func (a *Aggregator) printStats() {
	a.mu.Lock()
	defer a.mu.Unlock()

	log.Println("=== Log Summary ===")
	for file, levels := range a.counts {
		log.Printf("File: %s", file)
		for level, count := range levels {
			log.Printf("  %s: %d", level, count)
		}
	}
}

