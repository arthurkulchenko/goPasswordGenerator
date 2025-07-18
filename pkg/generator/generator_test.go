package generator

import(
	"testing"
	"pass_gen/pkg/password"
	"pass_gen/pkg/generator"
)

func Test_Generate(t *testing.T) {
	length := 4
	valid_pass_conf := password.PasswordConfig { length, true, false, false, true }
	output := generator.Generate(valid_pass_conf)
	if len(output) != length {
		t.Errorf("expected 4 digits password")
	}
}
func Test_GenNextSymbol(t *testing.T) {
	conf := generator.RandGenParam{generator.Digit, generator.Digits}


	output := generator.GenNextSymbol(&conf)

	if len(output) > 1 {
		t.Errorf("expected 1 symbol")
	}
}
