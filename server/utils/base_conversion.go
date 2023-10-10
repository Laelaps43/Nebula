package utils

import (
	"fmt"
	"math/big"
	"strings"
)

func StreamToHex(s string) string {
	d := new(big.Int)
	d.SetString(s, 10)
	return fmt.Sprintf("%08s", strings.ToUpper(d.Text(16)))
}
