package main

import "testing"

func TestAdd(t *testing.T) {
	if Add(2, 3) != 5 {
		t.Errorf("Add(2, 3) = %d; want 5", Add(2, 3))
	}
	if Add(-1, 1) != 0 {
		t.Errorf("Add(-1, 1) = %d; want 0", Add(-1, 1))
	}
	if Add(0, 0) != 0 {
		t.Errorf("Add(0, 0) = %d; want 0", Add(0, 0))
	}
}

func TestSub(t *testing.T) {
	if Sub(5, 3) != 2 {
		t.Errorf("Sub(5, 3) = %d; want 2", Sub(5, 3))
	}
	if Sub(0, 1) != -1 {
		t.Errorf("Sub(0, 1) = %d; want -1", Sub(0, 1))
	}
	if Sub(-1, -1) != 0 {
		t.Errorf("Sub(-1, -1) = %d; want 0", Sub(-1, -1))
	}
}

func TestMul(t *testing.T) {
	if Mul(2, 3) != 6 {
		t.Errorf("Mul(2, 3) = %d; want 6", Mul(2, 3))
	}
	if Mul(-1, 1) != -1 {
		t.Errorf("Mul(-1, 1) = %d; want -1", Mul(-1, 1))
	}
	if Mul(0, 5) != 0 {
		t.Errorf("Mul(0, 5) = %d; want 0", Mul(0, 5))
	}
}

func TestDiv(t *testing.T) {
	result, err := Div(6, 3)
	if err != nil || result != 2 {
		t.Errorf("Div(6, 3) = %f, %v; want 2, nil", result, err)
	}

	result, err = Div(5, 0)
	if err == nil {
		t.Errorf("Div(5, 0) = %f, %v; want error", result, err)
	}

	result, err = Div(-6, -3)
	if err != nil || result != 2 {
		t.Errorf("Div(-6, -3) = %f, %v; want 2, nil", result, err)
	}
}
