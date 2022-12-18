package main

import (
	"testing"
)

// test fetchUser function
func TestFetchUser(t *testing.T) {
	got := FetchUser()
	want := "BOB"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// TODO: test fetchUserLikes function
// TODO: test fetchUserMatch function
