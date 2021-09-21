package util

import "testing"

func TestString(t *testing.T) {
	
	s := ReverseRunes("ollaH")
	
	if s != "Hallo" {
		t.Errorf("Hallo falsch")
	}
}
