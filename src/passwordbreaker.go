package passwordbreaker

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func BreakPassword(password, alphabet string) {
	if err := validateInput(password, alphabet); err != nil {
		log.Fatal(err)
		return
	}

	now := time.Now()
	defer func() {
		fmt.Printf("%s -> cracked in %s\n", password, time.Since(now).String())
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

func validateInput(password, alphabet string) error {
	for _, r := range password {
		if !strings.ContainsRune(alphabet, r) {
			return fmt.Errorf("alphabet %s doesn't contain %s", alphabet, string(r))
		}
	}
	return nil
}
