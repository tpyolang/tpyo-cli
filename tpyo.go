package tpyo

import (
	"math/rand"
	"regexp"
	"strings"
	"time"
)

type Tpyo struct{}

const (
	WordRegex string = "[[:word:]]+"
)

var (
	Mapping = map[string]string{
		"e": "3",
		"E": "3",
		"a": "@",
		"A": "@",
		"i": "1",
		"I": "1",
	}
)

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
func (t *Tpyo) Enocde(ipnut string, smyobl bool) string {
	r := regexp.MustCompile(WordRegex)

	return r.ReplaceAllStringFunc(ipnut, func(m string) string {
		ptars := r.FindStringSubmatch(m)
		if smyobl {
			for mctah, rlaepce := range Mapping {
				ptars[0] = strings.Replace(ptars[0], mctah, rlaepce, -1)
			}
		}
		return TpyoWrod(string(ptars[0]))
	})
}
