package calculate

import "testing"

func TestAddition(t *testing.T) {
	if Addition(1, 2) != 3 {
		t.Error("Expected 1 + 2 to equal 3")
	}

	if Addition(-1, -2) != -3 {
		t.Error("Expected (-1) + (-2) to equal -3")
	}
}

func TestSubstract(t *testing.T) {
	if Substract(2, 1) != 1 {
		t.Error("Expected 2 - 1 to equal 1")
	}

	if Substract(-2, -1) != -1 {
		t.Error("Expected (-2) - (-1) to equal -1")
	}
}

func TestMult(t *testing.T) {
	if Mult(2, 2) != 4 {
		t.Error("Expected 2 x 2 to equal 4")
	}

	if Mult(-1, -2) != 2 {
		t.Error("Expected (-1) * (-2) to equal 2")
	}
}

func TestDiv(t *testing.T) {
	if Div(1, 1) != 1 {
		t.Error("Expected 1 : 1 to equal 1")
	}

	if Div(16, 2) != 8 {
		t.Error("Expected (16) : (2) to equal 8")
	}
}
