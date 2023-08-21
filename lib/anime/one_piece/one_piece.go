package one_piece

import (
	"math/rand"

	"github.com/sjerusallem/faker/lib/locale/en/one_piece"
)

// Random One Piece Akuma No Mi.
func RandomAkumaNoMi() string {
	akumaMoMi := one_piece.Akuma_No_Mi[rand.Intn(len(one_piece.Akuma_No_Mi))]
	return akumaMoMi
}

// Random One Piece character names.
func RandomCharacter() string {
	character := one_piece.Characters[rand.Intn(len(one_piece.Characters))]
	return character
}

// Random One Piece island names.
func RandomIsland() string {
	island := one_piece.Islands[rand.Intn(len(one_piece.Islands))]
	return island
}

// Random One Piece location names.
func RandomLocation() string {
	location := one_piece.Locations[rand.Intn(len(one_piece.Locations))]
	return location
}

// Random One Piece quotes.
func RandomQuote() string {
	quote := one_piece.Quotes[rand.Intn(len(one_piece.Quotes))]
	return quote
}

// Random One Piece sea names.
func RandomSea() string {
	sea := one_piece.Seas[rand.Intn(len(one_piece.Seas))]
	return sea
}
