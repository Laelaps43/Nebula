package utils

import (
	"strconv"
	"strings"
	"time"
)

// 将传递过来的时间转换成Duration
func ParseExpireTime(d string) (time.Duration, error) {
	d = strings.TrimSpace(d)
	dr, err := time.ParseDuration(d)
	if err == nil {
		return dr, nil
	}
	if strings.Contains(d, "d") {
		index := strings.Index(d, "d")

		num, _ := strconv.Atoi(d[:index])
		dr = time.Hour * 24 * time.Duration(num)
		ndr, err := time.ParseDuration(d[index+1:])
		if err != nil {
			return dr, nil
		}
		return dr + ndr, nil
	}
	dv, err := strconv.ParseInt(d, 10, 64)
	return time.Duration(dv), err
}
