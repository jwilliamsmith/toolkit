package toolkit

import "crypto/rand"

const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ*_+%&!$#"

// Tools is the type used to instantiate this module.  Any var of this type will have access to all methods with receiver *Tools
type Tools struct{}

func (t *Tools) RandomString(n int) string {
	var s, r = make([]rune, n), []rune(randomStringSource)
	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x, y := p.Uint64(), uint64(len(r))
		s[i] = r[x%y]
	}
	return string(s)
}
