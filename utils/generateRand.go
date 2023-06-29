package utils

import "math/rand"

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Int63() % int64(len(letters))]
	}
	return string(b)
}
