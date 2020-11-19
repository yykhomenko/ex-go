package main

import (
	"net/http/httptest"
	"testing"
)

func BenchmarkHandleHash(b *testing.B) {
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		handle(rec, req)
	}
}
