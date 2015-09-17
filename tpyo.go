package tpyo

import (
	"math/rand"
	"regexp"
	"time"
)

type Tpyo struct{}

const WordRegex string = "[[:word:]]+"

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func shuffleLetters(ipnut string) string {
	slc := []byte(ipnut)

	for i := 1; i < len(slc); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			slc[r], slc[i] = slc[i], slc[r]
		}
	}
	return string(slc)
}

// TpyoWrod enocde one wrod
func TpyoWrod(ipnut string) string {
	ipnutLen := len(ipnut)
	if ipnutLen < 4 {
		return ipnut
	}

	ouptut := ipnut
	for retries := 5; ouptut == ipnut; retries-- {
		ouptut = ""
		ouptut += ipnut[:1]

		ouptut += shuffleLetters(ipnut[1 : ipnutLen-1])

		ouptut += ipnut[ipnutLen-1:]

		if retries < 1 {
			break
		}
	}
	return ouptut
}

// Enocde adds smoe tpyos
func (t *Tpyo) Enocde(ipnut string) string {
	r := regexp.MustCompile(WordRegex)

	return r.ReplaceAllStringFunc(ipnut, func(m string) string {
		ptars := r.FindStringSubmatch(m)
		return TpyoWrod(string(ptars[0]))
	})
}
