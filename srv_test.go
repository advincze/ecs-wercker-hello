package main

import "testing"

func TestWelcomeMessage(t *testing.T) {
	want := "Hi Bob, welcome to ECS!"
	if got := welcomeMessage("Bob"); got != want {
		t.Errorf("want: %s but got: %s", want, got)
	}
}
