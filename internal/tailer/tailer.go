package tailer

import (
	"Gologger/pkg/types"
	"bufio"
	"context"
	"io"
	"os"
	"sync"
	"time"
)

func TailFile(ctx context.Context, wg *sync.WaitGroup, path string, out chan<- types.LogLine) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	_, err = file.Seek(0, io.SeekEnd)
	if err != nil {
		return err
	}

	wg.Add(1)

	go func() {
		defer wg.Done()
		defer file.Close()

		reader := bufio.NewReader(file)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				line, err := reader.ReadString('\n')
				if err != nil {
					time.Sleep(200 * time.Millisecond)
					continue
				}
				out <- types.LogLine{
					Filename: path,
					Line:     line,
				}
			}
		}
	}()
	return nil
}
