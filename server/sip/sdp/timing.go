package sdp

// t=<start-time> <stop-time>

import (
	"fmt"
	"time"
)

type Timing struct {
	Start time.Time // 会话开始时间
	End   time.Time // 会话结束时间
}

func (t *Timing) String() string {

	return fmt.Sprintf("t=%d %d", toNTP(t.Start), toNTP(t.End))
}

func toNTP(c time.Time) uint64 {
	if c.IsZero() {
		return 0
	}
	seconds := c.Unix()

	// UnixNano 时间戳（纳秒）部分，将其除以 1e9 转换为秒
	nanoseconds := c.UnixNano() / int64(time.Second)

	// 计算 NTP 时间戳
	ntpTimestamp := uint64(seconds)<<32 | uint64(nanoseconds)
	return ntpTimestamp
}
