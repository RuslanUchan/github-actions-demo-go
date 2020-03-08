package hello

import "testing"

func TestGreetsGitHub(t *testing.T) {
	result := Greet()
	if result != "Hello GitHub Actions from Uchan" {
		t.Errorf("Greet() = %s; want Hello GitHub Actions from Uchan", result)
	}
}
