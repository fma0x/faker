package fma_test

import (
	"reflect"
	"testing"

	"github.com/fma0x/faker/lib/anime/fma"
)

func TestFMARandomData(t *testing.T) {
	t.Run("GenerateRandomCharacters", func(t *testing.T) {
		got := reflect.TypeOf(fma.RandomCharacter())
		expected := reflect.String

		if got.Kind() != expected {
			t.Errorf("RandomCharacter test should return %v, but it returned %v.", expected, got)
		}
	})

	t.Run("GenerateRandomCity", func(t *testing.T) {
		got := reflect.TypeOf(fma.RandomCity())
		expected := reflect.String

		if got.Kind() != expected {
			t.Errorf("RandomCity test should return %v, but it returned %v.", expected, got)
		}
	})

	t.Run("GenerateRandomCountry", func(t *testing.T) {
		got := reflect.TypeOf(fma.RandomCountry())
		expected := reflect.String

		if got.Kind() != expected {
			t.Errorf("RandomCountry test should return %v, but it returned %v.", expected, got)
		}
	})
}
