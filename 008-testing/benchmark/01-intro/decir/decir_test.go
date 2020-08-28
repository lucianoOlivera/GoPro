package decir

import "testing"

func TestSaludar(t *testing.T) {
	s := Saludar("luciano")
	if s != "binvenido" {
		t.Error("expected", "binvenido", "got", s)
	}

}
