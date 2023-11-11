package sip

import (
	"crypto/md5"
	"fmt"
	"github.com/ghettovoice/gosip/sip"
	"go.uber.org/zap"
	"nebula.xyz/global"
	"nebula.xyz/helper"
	m "nebula.xyz/model/sip"
	"nebula.xyz/model/system"
	"nebula.xyz/utils"
	"net/http"
	"strings"
	"sync"
	"time"
)

const (
	ExpiresConstant       = "Expires"
	AuthorizationConstant = "Authorization"
	MD5                   = "MD5"
	WWWHeader             = "WWW-Authenticate"
	ExpiresTime           = 3600
)

var Wait = sync.WaitGroup{}

func Register(req sip.Request, tx sip.ServerTransaction) {
	global.Logger.Info("收到SIP Register请求，正在处理")

	// 获取请求的消息头
	authorization := req.GetHeaders(AuthorizationConstant)
	if len(authorization) == 0 {
		// 没有认证信息，返回401
		global.Logger.Info("请求未认证，返回认证请求401。")
		response := sip.NewResponseFromRequest("", req, http.StatusUnauthorized, "Unauthorized", "")
		auth := fmt.Sprintf(
			`Digest realm="%s",qop="%s",algorithm="%s",nonce="%s"`,
			sipServer.Realm,
			"auth",
			MD5,
			utils.RandString(32),
		)
		response.AppendHeader(&sip.GenericHeader{
			HeaderName: WWWHeader,
			Contents:   auth,
		})
		_ = tx.Respond(response)
		return
	}

	// 第二次请求
	fromInfo, err := DeviceFormRequest(req)
	if err != nil {
		global.Logger.Error("解析设备信息错误。", zap.Error(err))
		return
	}
	if len(fromInfo.DeviceId) != 20 {
		global.Logger.Error("错误的设备ID", zap.String("ID", fromInfo.DeviceId))
		_ = tx.Respond(sip.NewResponseFromRequest("", req, http.StatusForbidden, "Forbidden", ""))
		return
	}

	global.Logger.Info("开始注册认证...")
	authorizationHeader := authorization[0].(*sip.GenericHeader)
	authorizationInfo := m.Auth(authorizationHeader.Contents)
	if authorizationInfo.Password == "" {
		authorizationInfo.Password = sipServer.Password
	}
	// TODO 完成设备认证出错
	// 没有通过认证，非法请求
	//if !verify(authorizationInfo) {
	if false {
		global.Logger.Info("设备认证失败。")
		_ = tx.Respond(sip.NewResponseFromRequest("", req, http.StatusForbidden, "Forbidden", ""))
		return
	}
	expires := req.GetHeaders(ExpiresConstant)
	if len(expires) != 1 {
		global.Logger.Error("获取Expires失败")
		return
	}
	idx := strings.Index(req.Source(), ":")
	fromIp := req.Source()[:idx]
	fromPort := req.Source()[idx+1:]

	if expires[0].(*sip.Expires).Equals(new(sip.Expires)) {
		// 请求为注销请求
		if err := fromInfo.DeviceById(); err != nil {
			global.Logger.Error("获取设备失败")
			return
		}

		device := fromInfo
		// 判断设备是否存储数据库中
		if device.DeviceId == "" {
			global.Logger.Info("设备注销失败，设备不存在系统中。", zap.String("device Id", fromInfo.DeviceId))
			return
		}
		device.Status = helper.DeviceOffline
		global.Logger.Info("设备注销", zap.String("device Id", device.DeviceId))
		_ = device.DeviceUpdate()
		response := sip.NewResponseFromRequest("", req, http.StatusOK, http.StatusText(http.StatusOK), "")
		_ = tx.Respond(response)
		return
	}

	// 在数据库中是否存在对应的设备
	if err := fromInfo.DeviceById(); err != nil {
		global.Logger.Error("获取设备失败")
		return
	}

	device := fromInfo
	nowTime := time.Now()
	if device.DeviceId == "" {
		device = fromInfo
		device.RegisterAt = &nowTime
		device.KeepLiveAt = &nowTime
		device.Expires = expires[0].Value()
		device.Status = helper.DeviceOnline
		device.IP = fromIp
		device.Port = fromPort
		_ = device.DeviceAdd()
	} else {
		var deviceSql system.Device
		deviceSql.RegisterAt = &nowTime
		deviceSql.KeepLiveAt = &nowTime
		deviceSql.Status = helper.DeviceOnline
		deviceSql.IP = fromIp
		deviceSql.Port = fromPort
		deviceSql.DeviceId = device.DeviceId
		_ = deviceSql.DeviceUpdate()
	}
	response := sip.NewResponseFromRequest("", req, http.StatusOK, "OK", "")
	to, _ := response.To()
	response.ReplaceHeaders("To",
		[]sip.Header{&sip.ToHeader{Address: to.Address,
			Params: sip.NewParams().Add("tag", sip.String{Str: utils.RandString(9)})}})
	response.RemoveHeader("Allow")
	es := sip.Expires(ExpiresTime)
	response.AppendHeader(&es)
	response.AppendHeader(&sip.GenericHeader{
		HeaderName: "Date",
		Contents:   time.Now().String(),
	})
	global.Logger.Info("设备添加成功！")
	_ = tx.Respond(response)
	Wait.Add(2)
	go QueryDeviceInfo(device)
	go QueryDeviceCatalog(device)
	Wait.Wait()
}

// verify 验证请求
func verify(a *m.Authorization) bool {
	// 将 username, realm, password 加密后获取密文
	s1 := fmt.Sprintf("%s:%s:%s", a.Username, a.Realm, a.Password)
	h1 := getDigest(a.Algorithm, s1)

	// 对请求的方法和请求的Uri进行加密获取密文
	s2 := fmt.Sprintf("REGISTER:%s", a.Uri)
	h2 := getDigest(a.Algorithm, s2)

	if h2 == "" || h1 == "" {
		global.Logger.Error("请求验证错误")
		return false
	}
	// 对密文1，nonce，密文2 进行加密，然后和Response比较是否相等
	s3 := fmt.Sprintf("%s:%s", h1, a.Nonce)
	if a.Qop != "" && strings.ToLower(a.Qop) == "auth" {
		if a.Nc != "-1" {
			s3 += ":" + a.Nc
		}
		if a.CNonce != "" {
			s3 += ":" + a.CNonce
		}
		s3 += ":" + a.Qop
	}
	s3 += ":" + h2
	//s3 := fmt.Sprintf("%s:%s:%s", h1, a.Nonce, h2)
	h3 := getDigest(a.Algorithm, s3)

	isEqual := h3 == a.Response
	global.Logger.Error("aas", zap.Any("abc", a))
	global.Logger.Info("验证：", zap.String("处理后", h3), zap.String("处理前", a.Response))
	return isEqual
}

// 使用对应的算法加密raw
func getDigest(algorithm, raw string) string {
	switch algorithm {
	case MD5:
		return fmt.Sprintf("%x", md5.Sum([]byte(raw)))
	default:
		return fmt.Sprintf("%x", md5.Sum([]byte(raw)))
	}
}

// DeviceFormRequest 从请求获取设备信息
func DeviceFormRequest(request sip.Request) (system.Device, error) {
	d := system.Device{}
	from, ok := request.From()
	if !ok {
		global.Logger.Error("无法从请求中解析出from信息", zap.String("请求", request.String()))
		return d, fmt.Errorf("无法解析出from信息，%s", request.String())
	}
	if from.Address == nil {
		global.Logger.Error("无法从请求中解析出address信息", zap.String("请求", request.String()))
		return d, fmt.Errorf("无法解析出address信息，%s", request.String())
	}
	if from.Address.User() == nil {
		global.Logger.Error("无法从请求中解析出user信息", zap.String("请求", request.String()))
		return d, fmt.Errorf("无法解析出user信息，%s", request.String())
	}
	d.DeviceId = from.Address.User().String()
	d.Realm = from.Address.Host()
	via, ok := request.ViaHop()
	if !ok {
		global.Logger.Error("无法从请求中解析出via头部信息", zap.String("请求", via.String()))
		return d, fmt.Errorf("无法解析出via头部信息，%s", via.String())
	}
	d.Transport = via.Transport
	global.Logger.Debug("解析出来的数据。", zap.Any("设备信息", d))
	return d, nil
}
