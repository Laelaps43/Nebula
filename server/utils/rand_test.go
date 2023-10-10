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

func TestRandInt(t *testing.T) {
	fmt.Println(RandInt(3))
	fmt.Println(RandInt(3))
	fmt.Println(RandInt(3))
	fmt.Println(RandInt(3))
}
