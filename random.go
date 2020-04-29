package tinker

import (
	"math/rand"
)

func randString(s []string) string {
	i := rand.Intn(len(s))
	return s[i]
}
