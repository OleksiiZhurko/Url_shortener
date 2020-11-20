package generators

import (
	"math/rand"
	"time"
)

const (
	values     = "abcdefghijklmnopqrstuvwxyz1234567890" // for generating
	valuesSize = len(values)
	stringSize = 6
)

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano())) // for non-pseudo-random

// Generates a random string 6 in length
func GenerateRandomString() string {
	result := make([]byte, stringSize)

	for one := 0; one < stringSize; one++ {
		result[one] = values[seededRand.Intn(valuesSize)]
	}

	return string(result)
}
