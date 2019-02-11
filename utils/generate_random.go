package utils

import "math/rand"

func GenerateRandomString(length int) string {
	var s string
	for i := 0; i < length; i++ {
		s += string('A' - 1 + rand.Intn(26))
	}
	return s
}

func GenerateRandomBytes(length int) []byte {
	p := make([]byte, length)
	rand.Read(p)
	return p
}
