package one_piece

import (
	"math/rand"
)

// Random One Piece Akuma No Mi.
func RandomAkumaNoMi() string {
	akumaMoMi := akuma_No_Mi[rand.Intn(len(akuma_No_Mi))]
	return akumaMoMi
}

// Random One Piece character names.
func RandomCharacter() string {
	character := characters[rand.Intn(len(characters))]
	return character
}

// Random One Piece island names.
func RandomIsland() string {
	island := islands[rand.Intn(len(islands))]
	return island
}

// Random One Piece location names.
func RandomLocation() string {
	location := locations[rand.Intn(len(locations))]
	return location
}

// Random One Piece quotes.
func RandomQuote() string {
	quote := quotes[rand.Intn(len(quotes))]
	return quote
}

// Random One Piece sea names.
func RandomSea() string {
	sea := seas[rand.Intn(len(seas))]
	return sea
}
