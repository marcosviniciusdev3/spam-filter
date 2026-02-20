package text

import (
	"slices"
	"strings"
)

type Counter []Token

type Token struct {
	Value string
	Num   uint
}

/* TODO: Make tokenize accept parameter
   To process insensitive tokens */

func Tokenize(data []byte) []string {
	// sentitive
	return strings.Fields(string(data))
}

func Count(text *[]byte) (Counter, error) {
	var counter Counter
	tokens := Tokenize(*text)
	tokenMap := make(map[string]uint)
	for _, t := range tokens {
		tokenMap[t]++
	}

	for token, count := range tokenMap {
		counter = append(counter, Token{Value: token, Num: count})
	}
	return counter, nil
}

// Work with small chucks of data
/* TODO: Make data parameter be a hashmap of type map[string]uint */

func Probability(data, token []byte) (float32, error) {
	counter, _ := Count(&data)

	i := slices.IndexFunc(counter, func(t Token) bool {
		return t.Value == string(token)
	})
	if i == -1 {
		// panic("An error has occurred!")
		return 0.0, nil
	}

	var total uint
	for _, v := range counter {
		total += v.Num
	}

	return float32(counter[i].Num) / float32(total), nil
}
