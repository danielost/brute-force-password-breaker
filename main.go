package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	passwordbreaker "github.com/danielost/brute-force-password-breaker/src"
)

func main() {
	file, err := os.Open("./passwords.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		password := line[0]
		alphabet := line[1]
		passwordbreaker.BreakPassword(password, alphabet)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
