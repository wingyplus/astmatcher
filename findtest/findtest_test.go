package main

import "testing"

func TestFindTests(t *testing.T) {
	src := `package main_test

func TestDoSomething(t *testing.T) {}`

	funcs := findTests(src)

	if len(funcs) != 1 {
		t.Error("Expect 1 function")
	}

	if name := funcs[0]; name != "TestDoSomething" {
		t.Errorf("Expect 'TestDoSomething' but got %s", name)
	}
}
