package main

import "testing"

func TestShouldSomaCorrect(t *testing.T) {
	result := soma(3, 2, 1)
	expectedResult := 6
	if result != expectedResult {
		t.Error("Valor esperado: ", expectedResult, "Valor retornado: ", result)
	}
}

func TestShouldSomaIncorrect(t *testing.T) {
	result := soma(3, 2, 1)
	expectedResult := 7
	if result != expectedResult {
		t.Error("Valor esperado: ", expectedResult, "Valor retornado: ", result)
	}
}

func TestShouldSubtraiCorrect(t *testing.T) {
	result := subtrai(10, 5)
	expectedResult := -5
	if result != expectedResult {
		t.Error("Valor esperado: ", expectedResult, "Valor retornado: ", result)
	}
}

func TestShouldSubtraiIncorrect(t *testing.T) {
	result := subtrai(10, 10)
	expectedResult := 2
	if result != expectedResult {
		t.Error("Valor esperado: ", expectedResult, "Valor retornado: ", result)
	}
}

func TestShouldMultiplicaCorrect(t *testing.T) {
	result := multiplica(10, 5)
	expectedResult := 50
	if result != expectedResult {
		t.Error("Valor esperado: ", expectedResult, "Valor retornado: ", result)
	}
}

func TestShouldMultiplicaIncorrect(t *testing.T) {
	result := multiplica(10, 5)
	expectedResult := 52
	if result != expectedResult {
		t.Error("Valor esperado: ", expectedResult, "Valor retornado: ", result)
	}
}

func TestShouldDivideCorrect(t *testing.T) {
	result := divide(10, 5)
	expectedResult := 2
	if result != expectedResult {
		t.Error("Valor esperado: ", expectedResult, "Valor retornado: ", result)
	}
}

func TestShouldDivideIncorrect(t *testing.T) {
	result := divide(10, 5)
	expectedResult := 3
	if result != expectedResult {
		t.Error("Valor esperado: ", expectedResult, "Valor retornado: ", result)
	}
}
