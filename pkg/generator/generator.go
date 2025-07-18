package pass_gen

import(
	"pass_gen/pkg/password"
	"time"
	"math/rand"
	"strings"
	"strconv"
)

type SymbolType int

const (
	Digit SymbolType = iota
	LowerCase
	UpperCase
)

type RandGenParam struct {
	generationRange int
	entity SymbolType
}

func Generate(config pass_gen.PasswordConfig) string {
	password := make([]string, config.Length)
	uniqueKeys := make(map[string]int)
	rand.Seed(time.Now().UnixNano())
	genParams := make([]RandGenParam, 0)
	if config.UseDigits { genParams = append(genParams, RandGenParam{10, Digit}) }
	if config.UseUppercase { genParams = append(genParams, RandGenParam{26, UpperCase}) }
	if config.UseLowercase { genParams = append(genParams, RandGenParam{26, LowerCase}) }

	retry := 0
	// for index, _ := range password {
	for {
		params := genParams[rand.Intn(len(genParams))]
		nextSymbol := genNextSymbol(params)
		if checkUniqueness(nextSymbol, &uniqueKeys) {
			password[index] = nextSymbol
		} else {
			if retry > 59 { break }
			retry += 1
		}

	}
	return strings.Join(password, "")
}

func checkUniqueness(symbol string, uniqueKeys *map[string]int) bool {
	for k := range m {
		keys = append(keys, k)
	}
}

func genNextSymbol(params RandGenParam) string {
	var nextSymbol string
	randomizedNumber := rand.Intn(params.generationRange)
	switch params.entity {
		case Digit:
			nextSymbol = strconv.Itoa(randomizedNumber)
		case LowerCase:
			nextSymbol = string('a' + randomizedNumber)
		case UpperCase:
			nextSymbol = strings.ToUpper(string('a' + randomizedNumber))
	}
	return nextSymbol
}
