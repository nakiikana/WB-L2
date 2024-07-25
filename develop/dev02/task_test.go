package main

import "testing"

func TestSimple(t *testing.T) {
	got, err := unpackString("a4bc2d5e3")
	want := "aaaabccdddddeee"
	if got != want && err != nil {
		t.Errorf("got %s, wanted %s", got, want)
	} 
}

func TestError(t *testing.T) {
	_, err := unpackString("45aaa")
	if err == nil {
		t.Errorf("Expect error but it isn't")
	}
}