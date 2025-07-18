package password

import(
	"fmt"
	"strconv"
	"strings"
	"github.com/fatih/color"
)

type PasswordConfig struct {
	Length int
	UseDigits bool
	UseUppercase bool
	UseLowercase bool
	// ExcludeAmbiguousDigits bool
	// ExcludeAmbiguousUppercase bool
	// ExcludeAmbiguousLowercase bool
	Valid bool
}

func (p *PasswordConfig) BoolSum() int {
	boolMap := map[bool]int{true: 1, false: 0}
	return boolMap[p.UseDigits] + boolMap[p.UseUppercase] + boolMap[p.UseLowercase]
}

func (p *PasswordConfig) Validate() {
	sum := p.BoolSum()
	if sum > 0 && p.Length > 0 {
		p.Valid = true
	}
}

func (p *PasswordConfig) SetLength() {
	var userInput string
	var err error
	for {
		color.New(color.FgCyan).Print("Type number of symbols: ")
		_, _ = fmt.Scanln(&userInput)
		p.Length, err = strconv.Atoi(userInput)
		if err != nil {
			color.Red("Wrong input, please provide a digit!")
		} else {
			if p.Length < 1 {
				color.Red("Sorry, password can't be zer0 length!")
			} else {
				break
			}
		}
	}
}

func (p *PasswordConfig) WillUseDigits() {
	var userInput string
	for {
		color.New(color.FgCyan).Print("Type 'y' to use digits:")
		_, _ = fmt.Scanln(&userInput)
		decision := strings.ToLower(userInput)
		if strings.Contains(strings.ToLower(decision), "y") {
			p.UseDigits = true
		} else {
			color.Yellow("Skipping digits")
		}
		break
	}
}

func (p *PasswordConfig) WillUseUppercase() {
	var userInput string
	for {
		color.New(color.FgCyan).Print("Type 'y' to use uppercase: ")
		_, _ = fmt.Scanln(&userInput)
		decision := strings.ToLower(userInput)
		if decision == "y" || decision == "yes" {
			p.UseUppercase = true
		} else {
			color.Yellow("Skipping uppercase")
		}
		break
	}
}

func (p *PasswordConfig) WillUseLowercase() {
	var userInput string
	for {
		color.New(color.FgCyan).Print("Type 'y' to use lowercase: ")
		_, _ = fmt.Scanln(&userInput)
		decision := strings.ToLower(userInput)
		if decision == "y" || decision == "yes" {
			p.UseLowercase = true
		} else {
			color.Yellow("Skipping lowercase")
		}
		break
	}
}
