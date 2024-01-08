package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const numSet = "0123456789"

func RandString(n int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, n)
	for i := range result {
		result[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(result)
}

// RandInt 生成指定位数的Int
func RandInt(n int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, n)
	for i := range result {
		result[i] = numSet[seededRand.Intn(len(numSet))]
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

// GetRandStreamId 生成随机流ID
func GetRandStreamId() string {
	// 固定的前两位
	prefix := "18"

	// 获取当前时间
	currentTime := time.Now()

	// 使用时间信息生成后六位
	timeSuffix := fmt.Sprintf("%06X", currentTime.UnixNano()%1000000)

	// 组合生成的字符串
	return prefix + timeSuffix
}
