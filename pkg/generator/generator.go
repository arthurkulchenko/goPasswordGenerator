package pass_gen

import(
	"pass_gen/pkg/password"
	"time"
	"math/rand"
	"strings"
)

var Digits = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
var AlphabetLower = []string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
}
var AlphabetUpper = []string{
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
}

type SymbolType int

const (
	Digit SymbolType = iota
	LowerCase
	UpperCase
)

type RandGenParam struct {
	entity SymbolType
	dictionary []string
}

func Generate(config pass_gen.PasswordConfig) string {
	password := make([]string, config.Length)
	rand.Seed(time.Now().UnixNano())

	genParams := make([]*RandGenParam, 0)
	if config.UseDigits { genParams = append(genParams, &RandGenParam{Digit, Digits}) }
	if config.UseLowercase { genParams = append(genParams, &RandGenParam{LowerCase, AlphabetLower}) }
	if config.UseUppercase { genParams = append(genParams, &RandGenParam{UpperCase, AlphabetUpper}) }

	for index, _ := range password {
		if len(genParams) == 0 { break }

		paramsIndex := rand.Intn(len(genParams))
		params := genParams[paramsIndex]
		if len(params.dictionary) == 0 {
			genParams = append(genParams[:paramsIndex], genParams[paramsIndex+1:]...)
			continue
		}

		nextSymbol := genNextSymbol(params)
		password[index] = nextSymbol
	}
	return strings.Join(password, "")
}

func genNextSymbol(params *RandGenParam) string {
	dic := params.dictionary
	var nextSymbol string
	randomizedNumber := rand.Intn(len(dic))
	switch params.entity {
		case Digit:
			nextSymbol = dic[randomizedNumber]
		case LowerCase:
			nextSymbol = string('a' + randomizedNumber)
		case UpperCase:
			nextSymbol = strings.ToUpper(string('a' + randomizedNumber))
	}
	params.dictionary = append(dic[:randomizedNumber], dic[randomizedNumber+1:]...)
	return nextSymbol
}
