package main

import (
	"bytes"
	"slices"
	"testing"
)

func TestCountDown(t *testing.T) {
	t.Run("Checking print content", func(t *testing.T) {
		buffer := bytes.Buffer{}
		spySleeperCumPrinter := &SpySleeperCumPrinter{}
		CountDown(&buffer, spySleeperCumPrinter)

		expected := "1\n2\n3\nGo!\n"
		actual := buffer.String()

		if expected != actual {
			t.Errorf("Expected: %q but Actual: %q", expected, actual)
		}
	})

	t.Run("Checking order of prints and sleeps", func(t *testing.T) {
		spySleeperCumPrinter := &SpySleeperCumPrinter{}
		CountDown(spySleeperCumPrinter, spySleeperCumPrinter)

		expOps := []string{write, sleep, write, sleep, write, sleep, write}
		actOps := spySleeperCumPrinter.operations

		if !slices.Equal(expOps, actOps) {
			t.Errorf("Expected calls: %s but Actual calls: %s", expOps, actOps)
		}
	})

}
