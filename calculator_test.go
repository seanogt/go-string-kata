package main

import (
	"testing"
)

func TestAdd12(t *testing.T) {
	got, err := Add("1,2")
	if got != 3 {
		t.Errorf("Add(\"1,2\") = %d; want 3", got)
	}

	got, _ = Add("3,5")
	if got != 8 {
		t.Errorf("Add(\"3,5\") = %d; want 8", got)
	}

	got, err = Add("3,FIVE")
	if err == nil {
		t.Errorf("Add(\"3,FIVE\") = %d; wanted err", got)
	}

	got, _ = Add("3, 6")
	if got != 9 {
		t.Errorf("Add(\"3, 6\") = %d; want 9", got)
	}

	got, _ = Add("3,5,6,4,2")
	if got != 20 {
		t.Errorf("Add(\"3,5,6,4,2\") = %d; want 20", got)
	}

	got, _ = Add("3,\n5,\n6,4,2")
	if got != 20 {
		t.Errorf("Add(\"3,5,6,4,2\") = %d; want 20", got)
	}

	got, _ = Add("3,")
	if err == nil {
		t.Errorf("Add(\"3,\") = %d; wanted err", got)
	}
}

func TestAddChangeDelimiter(t *testing.T) {
	got, _ := Add("//;\n1;2")
	if got != 3 {
		t.Errorf("Add(\"//;\n1;2\") = %d; want 3", got)
	}

	got, _ = Add("//[***]\n1***2***3")
	if got != 6 {
		t.Errorf("Add(\"//[***]\n1***2***3\") = %d; want 6", got)
	}

	got, _ = Add("//[*][%]\n1*2%3")
	if got != 6 {
		t.Errorf("Add(\"//[*][percent]\n1*2percent3\") = %d; want 6", got)
	}

	got, _ = Add("//[***][%]\n1***2%3")
	if got != 6 {
		t.Errorf("Add(\"//[***][percent]\n1*2percent3\") = %d; want 6", got)
	}

}

func TestAddZero(t *testing.T) {
	got, _ := Add("")
	if got != 0 {
		t.Errorf("Add(\"\") = %d; want 0", got)
	}
}

func TestAddNegatives(t *testing.T) {
	got, err := Add("3,-4,5,-5")
	if err == nil {
		t.Errorf("Add(\"3,-4,5,-5\") = %d; wanted err", got)
	}
}

func TestAddIgnoreGreaterThan100(t *testing.T) {
	got, _ := Add("2,1001")
	if got != 2 {
		t.Errorf("Add(\"2,1001\") = %d; want 2", got)
	}
}
