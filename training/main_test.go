package main

import (
  "fmt"
  "testing"
)

func TestSquare(t *testing.T) {
  for input, want := range map[int]int {
    2: 4,
    3: 9,
    5: 25,
  }{
    name := fmt.Sprintf("square(%d)", input)
    t.Run(name, func(t *testing.T) {
      if have := square(input); want != have {
        t.Errorf("want %d, have %d", want, have)
      }
    })
  }
}

func BenchmarkSquare9(b *testing.B) {
  for i := 0; i < b.N; i++ {
    square(9)
  }
}
