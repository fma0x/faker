package ptbr

import (
	"math/rand"

	"github.com/sjerusallem/faker/lib/locale/pt_br"
)

// returns random name of Brazilian person.
func RandomFirstName() string {
	name := pt_br.FirstName[rand.Intn(len(pt_br.FirstName))]
	return name
}
