package fma

import (
	"math/rand"

	"github.com/sjerusallem/faker/lib/locale/en"
)

// Random FullMetal Alchemist Brotherhood character names.
func RandomCharacter() string {
	character := en.Characters[rand.Intn(len(en.Characters))]
	return character
}

// Random FullMetal Alchemist Brotherhood cities names.
func RandomCity() string {
	city := en.Cities[rand.Intn(len(en.Cities))]
	return city
}
