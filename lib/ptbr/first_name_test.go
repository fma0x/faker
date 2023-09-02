package ptbr_test

import (
	"reflect"
	"testing"

	"github.com/fma0x/faker/lib/ptbr"
)

func TestPtBrRandomName(t *testing.T) {
	t.Run("GenerateRandomName", func(t *testing.T) {
		got := reflect.TypeOf(ptbr.RandomFirstName())
		expected := reflect.String

		if got.Kind() != expected {
			t.Errorf("RandomName test should return %v, but it returned %v.", expected, got)
		}
	})
}
