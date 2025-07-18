package main

import(
	"fmt"
	config "pass_gen/pkg/password"
	generator "pass_gen/pkg/generator"
	"github.com/fatih/color"
)



func main() {
	pass_conf := config.PasswordConfig { 0, false, false, false, false, false, false, false }
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

	fmt.Println(generator.Generate(pass_conf))
}
