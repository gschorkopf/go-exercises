package main

import (
	"fmt"
	"testing"
)

func TestRunnerGeoff(t *testing.T) {
	geoff := geoff{}
	if v := geoff.name(); v != "Geoff" {
		t.Errorf("runner.name(): want Geoff, have %s", v)
	}
	if v := geoff.run(10); v != 20 {
		t.Errorf("runner.run(10): want 20, have %d", v)
	}
}

func TestMultipleRunners(t *testing.T) {
	for runner, want := range map[runner]int{
		geoff{}:  20,
		blithe{}: 10,
		thom{}:   5,
	} {
		name := fmt.Sprintf("%s{}.run(10)", runner.name())
		t.Run(name, func(t *testing.T) {
			if have := runner.run(10); want != have {
				t.Errorf("want %d, have %d", want, have)
			}
		})
	}
}

func TestRace(t *testing.T) {
	winner, fastest := race(100, geoff{}, blithe{}, thom{})
	if fastest != 50 {
		t.Errorf("fastest for run(100): want 50, have %d", fastest)
	}
	if winner != "Thom" {
		t.Errorf("winner for run(100): want Thom, have %s", winner)
	}
}
