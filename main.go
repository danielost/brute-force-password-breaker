package main

import (
	"fmt"
	"time"
)

func main() {
	// Усі паролі згенеровані через BitWarden
	passwords := map[string]string{
		"87355": "0123456789",
		"zqcyb": "abcdefghijklmnopqrstuvwxyz",
		"GpaTJ": "aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ",
		"RE5Pd": "aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ0123456789",
		"W9#j!": "aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ0123456789!@#$%%^&*",
	}

	for password, alphabet := range passwords {
		breakPassword(password, alphabet)
	}
}

func breakPassword(password, alphabet string) {
	now := time.Now()
	defer func() {
		fmt.Printf("%s -> зламаний за %s\n", password, time.Since(now).String())
	}()

	length := 1
	for !permute(length, password, alphabet, make([]rune, 0)) {
		length++
	}
}

func permute(length int, passwordToBreak, alphabet string, currentPassword []rune) bool {
	if string(currentPassword) == passwordToBreak {
		return true
	}
	if length == 0 {
		return false
	}

	for _, r := range alphabet {
		currentPassword = append(currentPassword, r)
		if permute(length-1, passwordToBreak, alphabet, currentPassword) {
			return true
		}
		currentPassword = currentPassword[:len(currentPassword)-1]
	}
	return false
}
