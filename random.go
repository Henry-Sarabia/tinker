package tinker

import (
	"math/rand"
)

func randString(s []string) string {
	if len(s) <= 0 {
		return ""
	}

	i := rand.Intn(len(s))
	return s[i]
}
