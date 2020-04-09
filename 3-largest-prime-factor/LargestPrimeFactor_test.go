package main

import (
	"testing"
	"strconv"
)

func Test190(t *testing.T) {
	actual := LargestPrimeFactor(190)
	expected := 19
	if actual != expected {
		t.Error("Returned " + strconv.Itoa(actual) + " instead of " + strconv.Itoa(expected))
	}
}

func Test13195(t *testing.T) {
	actual := LargestPrimeFactor(13195)
	expected := 29
	if actual != expected {
		t.Error("Returned " + strconv.Itoa(actual) + " instead of " + strconv.Itoa(expected))
	}
}
