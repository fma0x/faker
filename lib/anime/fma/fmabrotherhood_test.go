package fma_test

import (
	"testing"

	"github.com/sjerusallem/faker/lib/anime/fma"
)

var got interface{}

//func TesndomCharacter(t *testing.T) {
//got := reflect.TypeOf(fma.RandomCharacter())
//expected := reflect.String

//if got.Kind() != expected {
//t.Errorf("RandomCharacter test should return %s, but it returned %s.", expected, got)
//}
//}

func TestRandomCity(t *testing.T) {
	got = fma.RandomCity()
	expected := true

	_, ok := got.(string)

	if ok != expected {
		t.Errorf("RandomCity test should return %t, but it returned %t.", expected, ok)
	}
}

func TestRandomCountry(t *testing.T) {
	got = fma.RandomCountry()
	expected := true

	_, ok := got.(string)

	if ok != expected {
		t.Errorf("RandomCountry test should return %t, but it returned %t.", expected, ok)
	}
}
