package main

import(
	config "pass_gen/pkg/password"
	generator "pass_gen/pkg/generator"
	"github.com/fatih/color"
)

func main() {
	var pass_conf config.PasswordConfig
	pass_conf.SetLength()
	pass_conf.WillUseDigits()
	pass_conf.WillUseUppercase()
	pass_conf.WillUseLowercase()
	pass_conf.Validate()
	if !pass_conf.Valid {
		for {
			color.Red("You need to pick one or more collection of symbols!")
			pass_conf.WillUseDigits()
			pass_conf.WillUseUppercase()
			pass_conf.WillUseLowercase()
			pass_conf.Validate()
			if pass_conf.Valid { break }
		}
	}

	color.Cyan("Password generated!")
	color.Green("%v", generator.Generate(pass_conf))
}
