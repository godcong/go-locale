package locale

import (
	"testing"
)

func Test_detectViaRegistry(t *testing.T) {
	langs, err := detectViaRegistry()

	t.Logf("langs: %v", langs)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if len(langs) == 0 {
		t.Error("Expected non-empty langs, got empty slice")
	}
}

func Test_detectViaPowershell(t *testing.T) {
	langs, err := detectViaPowershell()

	t.Logf("langs: %v", langs)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if len(langs) == 0 {
		t.Error("Expected non-empty langs, got empty slice")
	}
}

func Test_detectViaSyscall(t *testing.T) {
	langs, err := detectViaSyscall()

	t.Logf("langs: %v", langs)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if len(langs) == 0 {
		t.Error("Expected non-empty langs, got empty slice")
	}
}
