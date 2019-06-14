package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	output := helloWorld()
	if output != "Hello World" {

	}

	if output != "Hello World" {
		t.Errorf("main() = %s; want \"Hello World\"", output)
	}
}
