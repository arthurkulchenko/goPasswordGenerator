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

type RandGenerator struct {
	generator *rand.Rand
	generationRange int
	entity SymbolType
}

func Generate(config pass_gen.PasswordConfig) string {
	password := make([]string, config.Length)
	source1 := rand.NewSource(time.Now().UnixNano())
	randUppercaseGenerator := rand.New(source1)

	source2 := rand.NewSource(time.Now().UnixNano())
	randLowercaseGenerator := rand.New(source2)

	source3 := rand.NewSource(time.Now().UnixNano())
	randDigitGenerator := rand.New(source3)

	source4 := rand.NewSource(time.Now().UnixNano())
	symbolTurnGenerator := rand.New(source4)

	generators := make([]RandGenerator, 0)
	if config.UseDigits {
		generators = append(generators, RandGenerator{randDigitGenerator, 10, Digit})
	}
	if config.UseUppercase {
		generators = append(generators, RandGenerator{randUppercaseGenerator, 26, LowerCase})
	}
	if config.UseLowercase {
		generators = append(generators, RandGenerator{randLowercaseGenerator, 26, UpperCase})
	}

	var nextSymbol string
	for index, _ := range password {
		nextGenerator := symbolTurnGenerator.Intn(len(generators))
		currentRandGenerator := generators[nextGenerator]

		randomizedNumber := currentRandGenerator.generator.Intn(currentRandGenerator.generationRange)
		switch currentRandGenerator.entity {
			case Digit:
				nextSymbol = strconv.Itoa(randomizedNumber)
			case LowerCase:
				nextSymbol = string('a' + randomizedNumber)
			case UpperCase:
				nextSymbol = strings.ToUpper(string('a' + randomizedNumber))
		}

		password[index] = nextSymbol

	}
	
	return strings.Join(password, "")
}
