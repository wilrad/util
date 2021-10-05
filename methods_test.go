package util

import (
	"fmt"
	//	"os"
	"testing"
	"time"
)

type xx struct{}

func (x *xx) Print() int { return 1 }

var x xx

func TestMethods(t *testing.T) {

	fmt.Println("in TestMethods")

	Methods(x)
	Methods(time.Hour)
}
