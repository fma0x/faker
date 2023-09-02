package one_piece_test

import (
	"reflect"
	"testing"

	"github.com/fma0x/faker/lib/anime/one_piece"
)

func TestOnePieceRandomData(t *testing.T) {
	t.Run("GenerateRandomCharacters", func(t *testing.T) {
		got := reflect.TypeOf(one_piece.RandomCharacter())
		expected := reflect.String

		if got.Kind() != expected {
			t.Errorf("RandomCharacter test should return %v, but it returned %v.", expected, got)
		}
	})

	t.Run("GenerateRandomAkumaNoMi", func(t *testing.T) {
		got := reflect.TypeOf(one_piece.RandomAkumaNoMi())
		expected := reflect.String

		if got.Kind() != expected {
			t.Errorf("RandomAkumaNoMi test should return %v, but it returned %v.", expected, got)
		}
	})

	t.Run("GenerateRandomIsland", func(t *testing.T) {
		got := reflect.TypeOf(one_piece.RandomIsland())
		expected := reflect.String

		if got.Kind() != expected {
			t.Errorf("RandomIsland test should return %v, but it returned %v.", expected, got)
		}
	})

	t.Run("GenerateRandomLocation", func(t *testing.T) {
		got := reflect.TypeOf(one_piece.RandomLocation())
		expected := reflect.String

		if got.Kind() != expected {
			t.Errorf("RandomLocation test should return %v, but it returned %v.", expected, got)
		}
	})

	t.Run("GenerateRandomQoute", func(t *testing.T) {
		got := reflect.TypeOf(one_piece.RandomQuote())
		expected := reflect.String

		if got.Kind() != expected {
			t.Errorf("RandomQuote test should return %v, but it returned %v.", expected, got)
		}
	})

	t.Run("GenerateRandomSea", func(t *testing.T) {
		got := reflect.TypeOf(one_piece.RandomSea())
		expected := reflect.String

		if got.Kind() != expected {
			t.Errorf("RandomSea test should return %v, but it returned %v.", expected, got)
		}
	})
}
