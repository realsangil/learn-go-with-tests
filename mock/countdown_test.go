package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleeper := &SpySleeper{}

		Countdown(buffer, sleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}

		if sleeper.Calls != 4 {
			t.Errorf("not enough calls to sleeper, want 4 got %d", sleeper.Calls)
		}
	})

	t.Run("sleep after every print ", func(t *testing.T) {
		countOperationsSpy := &CountdownOperationsSpy{}
		Countdown(countOperationsSpy, countOperationsSpy)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(countOperationsSpy.Calls, want) {
			t.Errorf("wanted calls %v got %v", want, countOperationsSpy.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}