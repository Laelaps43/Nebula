package sdp

import (
	"strconv"
	"strings"
)

// RFC 4566 SDP协议，在这里许多没有根据RFC 4566 来实现
// 目前满足能就行

const newLine = '\n'

// TODO 完善RFC 4566规范
type SdpSession struct {
	Version            int            // v= 协议版本，RFC4566规定为0
	Origin             Origin         // o= Origin 发起人
	SessionName        string         // s= 每个对话只能拥有一个，对对话的描述
	SessionInformation string         // i= 会话描述
	URI                string         // 资源定位符
	EmailAddress       string         //
	PhoneNumber        string         //
	ConnectionData     ConnectionData // c=
	BandWidth          any            // b= 未实现
	Timing             []Timing       // t=
	RepeatTimes        any            // r= 未实现
	TimeZones          any            // z= 未实现
	Encryption         any            // k= 未实现
	Attributes         []Attribute    // a=
	MediaName          MediaName      // m =
	SSRC               string         // y=
}

func (s SdpSession) String() (result string) {
	sb := strings.Builder{}
	sb.WriteString("v=" + strconv.Itoa(s.Version))
	sb.WriteRune(newLine)
	sb.WriteString(s.Origin.String())
	sb.WriteRune(newLine)
	sb.WriteString("s=" + s.SessionName)
	sb.WriteRune(newLine)

	if len(s.SessionInformation) > 0 {
		sb.WriteString(s.SessionInformation)
		sb.WriteRune(newLine)
	}

	if len(s.URI) > 0 {
		sb.WriteString(s.URI)
		sb.WriteRune(newLine)
	}
	if len(s.EmailAddress) > 0 {
		sb.WriteString(s.EmailAddress)
		sb.WriteRune(newLine)
	}
	if len(s.PhoneNumber) > 0 {
		sb.WriteString(s.PhoneNumber)
		sb.WriteRune(newLine)
	}
	if !s.ConnectionData.IsNil() {
		sb.WriteString(s.ConnectionData.String())
		sb.WriteRune(newLine)
	}
	if s.BandWidth != nil {
		//sb.WriteString(s.BandWidth)
	}
	if len(s.Timing) > 0 {
		for _, time := range s.Timing {
			sb.WriteString(time.String())
			sb.WriteRune(newLine)
		}

	}
	if s.RepeatTimes != nil {

	}
	if s.TimeZones != nil {

	}
	if s.Encryption != nil {

	}
	if !s.MediaName.IsNil() {
		sb.WriteString(s.MediaName.String())
		sb.WriteRune(newLine)
	}
	if len(s.Attributes) > 0 {
		for _, attribute := range s.Attributes {
			sb.WriteString(attribute.String())
			sb.WriteRune(newLine)
		}
	}

	if len(s.SSRC) > 0 {
		sb.WriteString("y=" + s.SSRC)
	}
	return sb.String()
}
