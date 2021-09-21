package util

import "testing"

func TestInt(t *testing.T) {

	i := Add(15, 666)
	if i != 681 {
		t.Errorf("Add(15, 666) = %d", i)
	}
}
