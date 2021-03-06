package main

import "testing"

func TestIsCorrect(t *testing.T) {
	var v bool
	v = IsCorrect("10000000+555555555")

	if v!=true {
		t.Error("Expected true, got", v)
	}
}

func TestGetNumbers(t *testing.T) {
	var v1, v2, v3 string

	v1, v2, v3 = GetNumbers("200+3000")

	if v1!="0200" || v2 != "3000" || v3 != "+" {
		t.Error("Expected 02003000+, got", v1 + v2 + v3)
	}
}

func TestConverToString(t *testing.T) {
	var v string

	v = ConverToString([]int{1, 2, 9, 5, 4, 6})

	if v!="129546" {
		t.Error("Expected 129546, got", v)
	}
}

func TestAdd(t *testing.T) {
	var v string

	v = Add([]int{1, 2, 9}, []int{6, 3, 4})

	if v!="763" {
		t.Error("Expected 763, got", v)
	}
}

func TestIsNegative(t *testing.T) {
	var v bool

	v = IsNegative([]int{1, 2, 9}, []int{6, 3, 4})

	if v!=true {
		t.Error("Expected true, got", v)
	}
}

func TestSub(t *testing.T) {
	var v string

	v = Sub([]int{1, 2, 9, 2, 2}, []int{6, 3, 4, 2, 2})

	if v!="-50500" {
		t.Error("Expected -50500, got", v)
	}
}
