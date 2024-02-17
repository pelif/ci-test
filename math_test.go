package main

import "testing"

func TestSoma(t *testing.T) {
	total := soma(15, 15)

	if total != 30 {
		t.Errorf("Resultado de Soma foi inválido: resultado %d. Valor esperado : %d", total, 30)
	}
}
