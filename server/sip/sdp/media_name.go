package sdp

import "strings"

// MediaName describes the "m=" field storage structure.
type MediaName struct {
	Media   string
	Port    string
	Protos  []string
	Formats []string
}

func (c *MediaName) String() string {
	return "m=" + strings.Join(
		[]string{
			c.Media,
			c.Port,
			strings.Join(c.Protos, "/"),
			strings.Join(c.Formats, " "),
		}, " ")
}

func (c *MediaName) IsNil() bool {
	if len(c.Media) == 0 && len(c.Port) == 0 && len(c.Protos) == 0 && len(c.Formats) > 0 {
		return true
	}
	return false
}
