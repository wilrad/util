package util

import (
	"fmt"
	"os"
	"testing"
)

func TestDisplay(t *testing.T) {

	fmt.Println("in TestDisplay")
	fmt.Println("vor Display")
	Display("os.Stderr", os.Stderr)

}
