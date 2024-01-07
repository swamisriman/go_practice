package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	write = "write"
	sleep = "sleep"
)

func main() {
	CountDown(os.Stdout, &DefaultSleeper{})
}

func CountDown(writer io.Writer, sleeper Sleeper) {
	for i := 1; i <= 3; i++ {
		fmt.Fprintln(writer, i)
		sleeper.Sleep(1 * time.Second)
	}
	fmt.Fprintln(writer, "Go!")
}

type Sleeper interface {
	Sleep(time.Duration)
}

type SpySleeperCumPrinter struct {
	operations []string
}

func (s *SpySleeperCumPrinter) Sleep(dur time.Duration) {
	s.operations = append(s.operations, sleep)
}

func (s *SpySleeperCumPrinter) Write(str []byte) (int, error) {
	s.operations = append(s.operations, write)
	return 0, nil
}

type DefaultSleeper struct {
}

func (d *DefaultSleeper) Sleep(dur time.Duration) {
	time.Sleep(dur)
}
