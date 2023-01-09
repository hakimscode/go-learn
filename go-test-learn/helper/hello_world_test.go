package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Heri")

	if result != "Hello Heri" {
		t.Error("Result should be 'Hello Heri")
	}
}

func TestHelloWorldHakim(t *testing.T) {
	result := HelloWorld("Hakim")

	if result != "Hello Hakim" {
		t.Fatal("Result should be 'Hello Hakim'")
	}
}
