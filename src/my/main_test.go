package main

import "testing"

func TestHelloWorld(t *testing.T) {
	// t.Fatal("not implemented")

	result := Bar()
	if result != "bar" {
		t.Errorf("expecting bar, got %s", result)
	}
}

func Bar() string {
	return "bar"

}
