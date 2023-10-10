package global

type Media interface {
	GetAddress() string
	SetAddress(a string)

	GetRestful() string
	SetRestful(p string)

	GetRTSPPort() string
	SetRTSPPort(p string)

	GetRTP() string
	SetRTP(p string)

	GetRTMPPort() string
	SetRTMPPort(p string)

	GetSecret() string
	SetSecret(s string)
}
