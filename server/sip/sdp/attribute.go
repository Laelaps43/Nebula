package sdp

import "fmt"

// Attribute describes the "a=" field

type Attribute struct {
	Key   string
	Value string
}

func (a *Attribute) String() string {
	if a.Value == "" {
		return fmt.Sprintf("a=%s", a.Key)
	}
	return fmt.Sprintf("a=%s:%s", a.Key, a.Value)
}
