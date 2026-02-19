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

func Count(text []byte) (Counter, error) {
	var counter Counter
	tokens := strings.Fields(string(text))
	tokenMap := make(map[string]uint)
	for _, t := range tokens {
		tokenMap[t]++
	}

	for token, count := range tokenMap {
		counter = append(counter, Token{Value: token, Num: count})
	}
	return counter, nil
}

func Probability(data, token []byte) (float32, error) {
	counter, _ := Count(data)

	println(len(counter))

	i := slices.IndexFunc(counter, func(t Token) bool {
		return t.Value == string(token)
	})
	if i == -1 {
		return 0, nil
	}

	var total uint
	for _, v := range counter {
		total += v.Num
	}

	println(total)
	println(len(counter))

	return float32(counter[i].Num) / float32(total), nil
}
