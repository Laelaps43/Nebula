package utils

import (
	"fmt"
	"math/big"
	"strings"
)

// StreamToHex 将10禁止转换为16进制
func StreamToHex(s string) string {
	d := new(big.Int)
	d.SetString(s, 10)
	return fmt.Sprintf("%08s", strings.ToUpper(d.Text(16)))
}

// HexToStream 将16禁止转化为十进制
func HexToStream(hexStr string) string {
	// 将十六进制字符串转换为大写
	hexStr = strings.ToUpper(hexStr)

	// 移除前导零以确保正确的解析
	hexStr = strings.TrimLeft(hexStr, "0")

	// 将十六进制字符串解析为大整数
	d := new(big.Int)
	_, _ = d.SetString(hexStr, 16)
	// 将大整数转换为十进制字符串
	decimalStr := d.String()
	return fmt.Sprintf("%010s", decimalStr)
}
