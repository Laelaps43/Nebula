package utils_test

import (
	"fmt"
	"nebula.xyz/utils"
	"testing"
)

func TestHexToStream(t *testing.T) {
	fmt.Println(utils.HexToStream("00003039"))
	fmt.Println(utils.HexToStream("0001E240"))
	fmt.Println(utils.HexToStream("072E99F4"))
}

func TestStreamToHex(t *testing.T) {
	fmt.Println(utils.StreamToHex("012345"))
	fmt.Println(utils.StreamToHex("0123456"))
	fmt.Println(utils.StreamToHex("0120494580"))
}
