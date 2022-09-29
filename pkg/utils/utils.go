package utils

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
)

const CommentChar = "#"

type Tokens []string

func IsEqual(first, second []string) bool {
	if len(first) != len(second) {
		return false
	}
	for i := 0; i < len(first); i++ {
		if first[i] != second[i] {
			return false
		}
	}
	return true
}

func ReadTokensFromFile(filepath string) (Tokens, error) {
	tokens := []string{}
	file, err := os.Open(filepath)
	if err != nil {
		return tokens, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		token := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(token, CommentChar) {
			continue
		}
		tokens = append(tokens, strings.TrimSpace(token))
	}

	return tokens, scanner.Err()
}

func (t Tokens) GetRandom() string {
	return t[rand.Intn(len(t))]
}
