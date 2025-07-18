package password

import(
	"testing"
	"os"
	"pass_gen/pkg/password"
)

func Test_SetLength_case_12(t *testing.T) {
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	r, w, _ := os.Pipe()
	os.Stdin = r
	w.Write([]byte("12\n"))
	w.Close()

	var config password.PasswordConfig
	config.SetLength()

	if config.Length != 12 {
		t.Errorf("expected 12, got %d", config.Length)
	}
}

func Test_WillUseDigits(t *testing.T) {
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	r, w, _ := os.Pipe()
	os.Stdin = r
	w.Write([]byte("y\n"))
	w.Close()

	var config password.PasswordConfig
	config.WillUseDigits()

	if config.UseDigits == false {
		t.Errorf("expected to get true")
	}
}

func Test_WillUseUppercase(t *testing.T) {
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	r, w, _ := os.Pipe()
	os.Stdin = r
	w.Write([]byte("y\n"))
	w.Close()

	var config password.PasswordConfig
	config.WillUseUppercase()

	if config.UseUppercase == false {
		t.Errorf("expected to get true")
	}
}
func Test_WillUseLowercase(t *testing.T) {
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	r, w, _ := os.Pipe()
	os.Stdin = r
	w.Write([]byte("y\n"))
	w.Close()

	var config password.PasswordConfig
	config.WillUseLowercase()

	if config.UseLowercase == false {
		t.Errorf("expected to get true")
	}
}
func Test_Validate(t *testing.T) {
	config := password.PasswordConfig {0, false, false, false, false }
	config.WillUseLowercase()
	if config.Valid == true {
		t.Errorf("expected to get true")
	}
}
func Test_BoolSum(t *testing.T) {
	config := password.PasswordConfig {0, false, false, false, false }
	sum := config.BoolSum()
	if sum != 0 {
		t.Errorf("expected to get true")
	}
}
