package main

import (
	"fmt"
	"bytes"
	"testing"
	."./tpl"
)

func BenchmarkMain(*testing.B) {
	// Parallel benchmark for text/template.Template.Execute on a single object.
	testing.Benchmark(func(b *testing.B) {

		page := &Page{
			Title: "Bob",
			FavoriteColors: []string{"blue", "green", "mauve"},
		}

		// RunParallel will create GOMAXPROCS goroutines
		// and distribute work among them.
		b.RunParallel(func(pb *testing.PB) {
			// Each goroutine has its own bytes.Buffer.
			var buf bytes.Buffer

			for pb.Next() {
				// The loop body is executed b.N times total across all goroutines.
				buf.Reset()

				_ = Home(&buf, page)

				fmt.Println(&buf)
			}
		})

	})
}
