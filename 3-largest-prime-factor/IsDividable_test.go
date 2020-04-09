package main

import "testing"

// dividable
func Test1(t *testing.T) {
	if !IsDividable(9, []int{2, 3}) {
		t.Error("9 dividable by 3")
	} 
}

// not dividable
func Test2(t *testing.T) {
	if IsDividable(17, []int{2, 3}) {
		t.Error("17 not dividable by neither 2 nor 3")
	} 
}
