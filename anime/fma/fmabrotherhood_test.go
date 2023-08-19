package fma_test

import (
	"testing"

	"github.com/sjerusallem/faker/anime/fma"
)

var got interface{}

func TestRandomCharacter(t *testing.T) {
	got = fma.RandomCharacter()
	expected := true

	_, ok := got.(string)

	if ok != expected {
		t.Errorf("RandomCharacter test should return %t, but it returned %t.", expected, ok)
	}
}

func TestRandomCity(t *testing.T) {
	got = fma.RandomCity()
	expected := true

	_, ok := got.(string)

	if ok != expected {
		t.Errorf("RandomCity test should return %t, but it returned %t.", expected, ok)
	}
}
