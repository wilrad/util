package util

import (
	"testing"
//	"fmt"
)
func TestUtil(t *testing.T) {

	a := []string { "a", "b", "c"}
	
	ar := Array2string(a) 
	if ar != "a, b, c" {
		t.Errorf("Array %v", ar)
	} 

	q := QuestionMark(10)
//	fmt.Println(q)
	if q != "?, ?, ?, ?, ?, ?, ?, ?, ?, ?" {
		t.Errorf("Questionmark %v", q)
	}
}
