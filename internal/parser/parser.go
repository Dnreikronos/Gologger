package parser

import (
	"Gologger/pkg/types"
	"context"
	"regexp"
	"strings"
	"sync"
)

var (
	errorRegex = regexp.MustCompile(`(?i)\bERROR\b`)
	warnRegex  = regexp.MustCompile(`(?i)\bWARN(?:ING)?\b`)
)

func StartParser(ctx context.Context, wg *sync.WaitGroup, in <-chan types.LogLine, out chan<- types.ParsedLog) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case logLine := <-in:
				level := ""
				switch {
				case errorRegex.MatchString(logLine.Line):
					level = "ERROR"
				case warnRegex.MatchString(logLine.Line):
					level = "WARN"
				default:
					continue // skip non-matching lines
				}

				out <- types.ParsedLog{
					Filename:  logLine.Filename,
					Level:     level,
					Message:   strings.TrimSpace(logLine.Line),
					Timestamp: "",
				}
			}
		}
	}()
}

