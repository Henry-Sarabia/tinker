package tinker

import (
	"math/rand"
)

func randomString(s []string) string {
	i := rand.Intn(len(s))
	return s[i]
}
