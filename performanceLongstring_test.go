package main

import (
	"bytes"
	"strings"
	"testing"
)

func ThroughConcatenation(length int) string {
	var s string
	for i := 0; i < length; i++ {
		s += "text"
	}
	return s
}

func ThroughArrangment(length int) string {
	s := []string{}
	for i := 0; i < length; i++ {
		s = append(s, "text")
	}
	return strings.Join(s, "")
}

func ThroughBuffer(length int) string {
	var buffer bytes.Buffer
	for i := 0; i < length; i++ {
		buffer.WriteString("text")
	}
	return buffer.String()
}

func BenchmarkThroughConcatenation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ThroughConcatenation(100)
	}
}

func BenchmarkThroughArrangment(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ThroughArrangment(100)
	}
}

func BenchmarkThroughBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ThroughBuffer(100)
	}
}
