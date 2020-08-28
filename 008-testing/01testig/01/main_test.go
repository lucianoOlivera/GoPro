package main

import (
	"testing"
)

func Testsuma(t *testing.T) {
	v := Mysuma(2, 8)
	if v != 10 {
		t.Error("expected", 10, "got", Mysuma(2, 8))
	}

}
