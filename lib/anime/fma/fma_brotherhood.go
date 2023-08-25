package fma

import (
	"math/rand"
)

// Random FullMetal Alchemist Brotherhood character names.
func RandomCharacter() string {
	character := characters[rand.Intn(len(characters))]
	return character
}

// Random FullMetal Alchemist Brotherhood cities names.
func RandomCity() string {
	city := cities[rand.Intn(len(cities))]
	return city
}

// Random FullMetal Alchemist Brotherhood countries names.
func RandomCountry() string {
	country := countries[rand.Intn(len(countries))]
	return country
}
