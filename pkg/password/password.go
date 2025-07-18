package pass_gen

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
	ExcludeAmbigiousDigits bool
	ExcludeAmbigiousUppercase bool
	ExcludeAmbigiousLowercase bool
	Valid bool
}

func (p *PasswordConfig) BoolSum() int {
	boolMap := map[bool]int{true: 1, false: 0}
	return boolMap[p.UseDigits] + boolMap[p.UseUppercase] + boolMap[p.UseLowercase]
}

func (p *PasswordConfig) Validate() {
	sum := p.BoolSum()
	//           && p.length > 0
	if sum > 0 {
		p.Valid = true
	}
}

func (p *PasswordConfig) SetLength() {
	var userInput string
	var err error
	for {
		fmt.Println("Type number of symbols:")
		fmt.Print("-> ")
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
		fmt.Println("Type 'y' or 'yes' if you want use digits:")
		fmt.Print("-> ")
		_, _ = fmt.Scanln(&userInput)
		decision := strings.ToLower(userInput)
		if decision == "y" || decision == "yes" {
			p.UseDigits = true
		}
		break
	}
}

func (p *PasswordConfig) WillUseUppercase() {
	var userInput string
	for {
		fmt.Println("Type 'y' or 'yes' if you want use uppercase:")
		fmt.Print("-> ")
		_, _ = fmt.Scanln(&userInput)
		decision := strings.ToLower(userInput)
		if decision == "y" || decision == "yes" {
			p.UseUppercase = true
		}
		break
	}
}

func (p *PasswordConfig) WillUseLowercase() {
	var userInput string
	for {
		fmt.Println("Type 'y' or 'yes' if you want use lowercase:")
		fmt.Print("-> ")
		_, _ = fmt.Scanln(&userInput)
		decision := strings.ToLower(userInput)
		if decision == "y" || decision == "yes" {
			p.UseLowercase = true
		}
		break
	}
}
