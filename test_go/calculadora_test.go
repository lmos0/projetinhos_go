package main

import "testing"

func TestShouldSumCorrect(t *testing.T) {
	teste := soma(3, 2, 1)
	resultado := 6
	if teste != resultado {
		t.Error("Valor esperado", resultado, "Valor retornado", teste)
	}
}

func TestShouldSumIncorrect(t *testing.T) {
	teste := soma(3, 2, 1)
	resultado := 7
	if teste != resultado {
		t.Error("Valor esperado", resultado, "Valor retornado", teste)
	}
}

func TestShouldMinusCorrect(t *testing.T) {
	teste := subtrai(12, 2, 2)
	resultado := -4
	if teste != resultado {
		t.Error("Valor esperado", resultado, "Valor retornado", teste)
	}
}

func TestShouldMinusIncorrect(t *testing.T) {
	teste := subtrai(12, 2, 2)
	resultado := 28
	if teste != resultado {
		t.Error("Valor esperado", resultado, "Valor retornado", teste)
	}
}

func TestShouldMultCorrect(t *testing.T) {
	teste := multiplca(5, 3)
	resultado := 15
	if teste != resultado {
		t.Error("Valor esperado", resultado, "Valor retornado", teste)
	}
}

func TestShouldMultIncorrect(t *testing.T) {
	teste := multiplca(5, 3)
	resultado := 45
	if teste != resultado {
		t.Error("Valor esperado", resultado, "Valor retornado", teste)
	}
}

func TestShouldDivideCorrect(t *testing.T) {
	teste := dividi(8, 2)
	resultado := 4
	if teste != resultado {
		t.Error("Valor esperado", resultado, "Valor retornado", teste)
	}
}

func TestShouldDivideIncorrect(t *testing.T) {
	teste := dividi(8, 2)
	resultado := 25
	if teste != resultado {
		t.Error("Valor esperado", resultado, "Valor retornado", teste)
	}
}
