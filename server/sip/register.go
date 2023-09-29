package sip

import (
	"github.com/ghettovoice/gosip/sip"
	"nebula.xyz/global"
)

func Register(req sip.Request, tx sip.ServerTransaction) {
	global.Logger.Info("受到SIP Register请求，正在处理")
}
