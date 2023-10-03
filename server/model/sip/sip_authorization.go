package sip

import (
	"go.uber.org/zap"
	"nebula.xyz/global"
	"regexp"
	"strings"
)

type Authorization struct {
	Realm     string
	Nonce     string
	Algorithm string
	Username  string
	Password  string
	Uri       string
	Response  string
	Method    string
	Qop       string
	Nc        string
	CNonce    string
	Other     map[string]string
}

func Auth(value string) *Authorization {
	auth := &Authorization{
		Algorithm: "MD5",
		Other:     make(map[string]string),
	}
	global.Logger.Info("aaa", zap.Any("aaa", value))
	re := regexp.MustCompile(`([\w]+)="([^"]+)"`)
	matches := re.FindAllStringSubmatch(value, -1)
	for _, match := range matches {
		switch match[1] {
		case "realm":
			auth.Realm = match[2]
		case "algorithm":
			auth.Algorithm = match[2]
		case "nonce":
			auth.Nonce = match[2]
		case "uri":
			auth.Uri = match[2]
		case "response":
			auth.Response = match[2]
		case "qop":
			for _, v := range strings.Split(match[2], ",") {
				v = strings.Trim(v, " ")
				if v == "auth" || v == "auth-int" {
					auth.Qop = "auth"
					break
				}
			}
		case "username":
			auth.Username = match[2]
		case "nc":
			auth.Nc = match[2]
		case "cnonce":
			auth.CNonce = match[2]
		default:
			auth.Other[match[1]] = match[2]
		}
	}
	return auth
}
