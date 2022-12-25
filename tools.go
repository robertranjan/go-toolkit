package toolkit

import "crypto/rand"

const (
	randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890_+"
)

// Tools is the type instantiate this module and have access to all functions in this module
type Tools struct{}

// RandomString returns a random string of length n from const: randomStringSource
func (t *Tools) RandomString(n int) string {

	s, r := make([]rune, n), []rune(randomStringSource)
	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x, y := p.Uint64(), uint64(len(r))
		s[i] = r[x%y]
	}
	return string(s)
}
