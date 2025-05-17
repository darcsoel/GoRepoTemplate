package main

import (
	"testing"
)

func TestStrConcat(t *testing.T) {
	first := "hello"
	second := "world"

	result := strConcat(first, second)

	if result != "helloworld" {
		t.Error("Test failed")
	}
}
