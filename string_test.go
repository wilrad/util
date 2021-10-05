package util

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {

	fmt.Printf("in TestString\n")
	s := ReverseRunes("ollaH")

	if s != "Hallo" {
		t.Errorf("Hallo falsch")
	}
}
