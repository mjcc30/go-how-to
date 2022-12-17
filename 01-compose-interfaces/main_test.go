package main

import (
	"testing"
)

// test NewHashReader function
func TestNewHashReader(t *testing.T) {
	var test string = "hello high value software enginer"
	payload := []byte(test)
	h := NewHashReader(payload)
	if h == nil {
		t.Error("NewHashReader failed")
	}
}

// test hash function
func TestHash(t *testing.T) {
	var test string = "hello high value software enginer"
	payload := []byte(test)
	h := NewHashReader(payload)

	got := h.Hash()
	want := "68656c6c6f20686967682076616c756520736f66747761726520656e67696e6572"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// test broacast function
func TestBroadcast(t *testing.T) {
	var test string = "hello high value software enginer"
	payload := []byte(test)
	h := NewHashReader(payload)
	got, _ := Broadcast(h)
	want := test
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// test hashAndBroadcast function
func TestHashAndBroadcast(t *testing.T) {
	var test string = "hello high value software enginer"
	payload := []byte(test)
	h := NewHashReader(payload)
	got, _ := HashAndBroadcast(h)
	want := test
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}