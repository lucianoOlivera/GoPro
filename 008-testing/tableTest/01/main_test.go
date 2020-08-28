package main

import (
	"testing"
)

func Testsuma(t *testing.T) {
	type pruebas struct {
		datos     []int
		respuesta int
	}

	pruebas := []prueba{
		prueba{[]int{2, 4, 6}, 12},
		prueba{[]int{1, 5, 2}, 8},
		prueba{[]int{0, -1, 1}, 0},
		prueba{[]int{-10, 4, 20}, 14},
	}
	for _, x := range pruebas {
		v := Mysuma(x.datos...)
		if v != x.respuesta {
			t.Error("expected", x.respuesta, "got", Mysuma(x.datos...))
		}
	}

}
