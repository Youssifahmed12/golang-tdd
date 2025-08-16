package main

import "testing"

func TestHello(t *testing.T) {
	want := "Hello Youssif!"
	got := hello("Youssif")
	if want != got {
		t.Errorf("wanted %q but got %q", want, got)
	}
}
