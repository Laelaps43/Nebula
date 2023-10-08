package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandString(n int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, n)
	for i := range result {
		result[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(result)
}

// GetSn 生成Sn
func GetSn() string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	return strconv.Itoa(seededRand.Intn(100) * 9876)
}

func Get4SSRC() string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%04d", seededRand.Intn(9999))
}
