package main

import (
	"fmt"
	"time"
)

func main() {
	alphabets := []string{
		"0123456789",
		"abcdefghijklmnopqrstuvwxyz",
		"aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ",
		"aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ0123456789",
		"aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ0123456789!@#$%%^&*",
	}

	// Усі паролі згенеровані через BitWarden
	passwords := []string{
		"87355",
		"zqcyb",
		"GpaTJ",
		"RE5Pd",
		"W9#j!",
	}

	for i, password := range passwords {
		breakPassword(password, alphabets[i])
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
