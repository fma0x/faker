package fma_test

import (
	"testing"

	"github.com/sjerusallem/faker/anime/fma"
)

var got interface{}

func TestFMARandomCharacter(t *testing.T) {
	got = fma.RandomCharacter()
	expected := true

	_, ok := got.(string)

	if ok != expected {
		t.Errorf("FMARandomCharacter test should return %t, but it returned %t.", expected, ok)
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
