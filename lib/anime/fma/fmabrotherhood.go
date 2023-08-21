package fma

import (
	"math/rand"

	"github.com/sjerusallem/faker/lib/locale/en/fma"
)

// Random FullMetal Alchemist Brotherhood character names.
func RandomCharacter() string {
	character := fma.Characters[rand.Intn(len(fma.Characters))]
	return character
}

// Random FullMetal Alchemist Brotherhood cities names.
func RandomCity() string {
	city := fma.Cities[rand.Intn(len(fma.Cities))]
	return city
}

// Random FullMetal Alchemist Brotherhood countries names.
func RandomCountry() string {
	country := fma.Countries[rand.Intn(len(fma.Countries))]
	return country
}
