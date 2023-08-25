package ptbr

import (
	"math/rand"
)

// returns random name of Brazilian person.
func RandomFirstName() string {
	name := firstName[rand.Intn(len(firstName))]
	return name
}
