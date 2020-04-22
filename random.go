package tinker

import (
	"math/rand"
)

func randomString(s []string) string {
	r := rand.Intn(len(s))
	return s[r]
}
