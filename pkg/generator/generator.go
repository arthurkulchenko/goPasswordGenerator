package generator

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

type RandGenParam struct {
	Dictionary []string
}

func Generate(config password.PasswordConfig) string {
	password := make([]string, config.Length)
	rand.Seed(time.Now().UnixNano())

	genParams := make([]*RandGenParam, 0)
	if config.UseDigits { genParams = append(genParams, &RandGenParam{Digits}) }
	if config.UseLowercase { genParams = append(genParams, &RandGenParam{AlphabetLower}) }
	if config.UseUppercase { genParams = append(genParams, &RandGenParam{AlphabetUpper}) }

	for index, _ := range password {
		if len(genParams) == 0 { break }

		paramsIndex := rand.Intn(len(genParams))
		params := genParams[paramsIndex]
		if len(params.Dictionary) == 0 {
			genParams = append(genParams[:paramsIndex], genParams[paramsIndex+1:]...)
			continue
		}

		nextSymbol := GenNextSymbol(params)
		password[index] = nextSymbol
	}
	return strings.Join(password, "")
}

func GenNextSymbol(params *RandGenParam) string {
	dic := params.Dictionary
	var nextSymbol string
	randomizedNumber := rand.Intn(len(dic))
	nextSymbol = dic[randomizedNumber]
	params.Dictionary = append(dic[:randomizedNumber], dic[randomizedNumber+1:]...)
	return nextSymbol
}
