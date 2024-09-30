package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 15)

	if total != 30 {
		t.Error("O resultado da soma foi inválido: ", total)
	}
}
