package utils_test

import (
	"fmt"
	"nebula.xyz/utils"
	"testing"
)

func TestBcryptHash(t *testing.T) {
	fmt.Println(utils.BcryptHash("nebula"))
}
