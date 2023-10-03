package utils

import (
	"fmt"
	"testing"
)

func TestRandString(t *testing.T) {
	fmt.Println(RandString(3))
	fmt.Println(RandString(6))
	fmt.Println(RandString(12))
	fmt.Println(RandString(2))
}
