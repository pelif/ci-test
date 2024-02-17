package main 

import "testing"
import "math"

func TestSoma(t *testing.T) {
	total := soma(15, 15)

	if total != 30 {
		t.Errorf("Result of sum is invalid: Result %d. Wait Value : %d", total, 30)
	}
}