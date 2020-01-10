package main

import (
  "testing"
)

var iterations int = 100000

func BenchmarkEchoJoin(b *testing.B) {
  for i := 0; i < iterations; i++ {
    EchoJoin()
  }
}

func BenchmarkEchoRange(b *testing.B) {
  for i := 0; i < iterations; i++ {
    EchoRange()
  }
}
