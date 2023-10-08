package sip

import (
	"errors"
	"fmt"
	"github.com/ghettovoice/gosip/sip"
	"math/big"
	"strings"

	//"nebula.xyz/sip/sdp"
	"net"
	"strconv"
	"time"

	sdp "github.com/panjjo/gosdp"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"nebula.xyz/global"
	"nebula.xyz/helper"
	"nebula.xyz/model/system"

	//"nebula.xyz/sip/sdp"

	"nebula.xyz/utils"
	"sync"
)

var ssrcLock *sync.Mutex = &sync.Mutex{}

// Play 点播
func Play(stream *system.Stream) (*system.Stream, error) {
	global.Logger.Debug("开始点播.....")
	channelTmp := &system.DeviceChannel{ChannelId: stream.ChannelId}
	var channel system.DeviceChannel
	var err error
	// 数据库中不存在对应通道
	if channel, err = channelTmp.DeviceChannelById(); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("通道不存在")
		}
		return nil, err
	}
	stream.DeviceId = channel.DeviceId
	deviceTmp := &system.Device{DeviceId: channel.DeviceId}

	// TODO 拉流
	// 推流处理
	if channel.Status != helper.ChannelStatusON {
		return nil, errors.New("通道已离线")
	}
	var device system.Device
	// 判断设备是否存在
	if device, err = deviceTmp.DeviceById(); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("设备不能存在")
		}
		return nil, err
	}
	if device.Status != helper.DeviceOnline {
		return nil, errors.New("设备离线")
	}

	//
	if stream.StreamId == "" {
		global.Logger.Debug("stream 不存在StreamId")
		ssrcLock.Lock()
		stream.StreamId = GetSSRC(&channel)
		if err := stream.Save(); err != nil {
			ssrcLock.Unlock()
			return nil, err
		}
		ssrcLock.Unlock()
	}
	stream, err = sipPlayPush(stream, &channel, &device)
	if err != nil {
		global.Logger.Info("点播失败", zap.Error(err))
		return nil, err
	}
	global.Logger.Info("点播成功......")

	media := global.CONFIG.Media
	zlmId := ToHex(stream.StreamId)
	stream.HTTP = fmt.Sprintf("http://%s:%s/rtp/%s/hls.m3u8", media.Address, media.Restful, zlmId)
	stream.RTMP = fmt.Sprintf("rtmp://%s:%s/rtp/%s", media.Address, media.RTMPPort, zlmId)
	stream.RTSP = fmt.Sprintf("rtsp://%s:%s/rtp/%s", media.Address, media.RTSPPort, zlmId)
	stream.WSFLV = fmt.Sprintf("ws://%s:%s/rtp/%s.live.flv", media.Address, media.Restful, zlmId)
	global.Logger.Info("ab", zap.Strings("直播流", []string{stream.HTTP, stream.RTMP, stream.RTSP, stream.WSFLV}))
	err = stream.Save()
	if err != nil {
		return nil, err
	}
	return stream, nil
}

// SIP
func sipPlayPush(
	stream *system.Stream,
	channel *system.DeviceChannel,
	device *system.Device) (*system.Stream, error) {

	//name := "Play"
	//protos := []string{"TCP/RTP/AVP"}
	//mediaName := sdp.MediaName{
	//	Media:   "video",
	//	Port:    global.CONFIG.Media.RTP,
	//	Protos:  protos,
	//	Formats: []string{"96", "98", "97"},
	//}
	//attributes := []sdp.Attribute{
	//	{Key: "recvonly"},
	//	{Key: "setup", Value: "passive"},
	//	{Key: "connection", Value: "new"},
	//	{Key: "rtpmap", Value: "96 PS/90000"},
	//	{Key: "rtpmap", Value: "98 H264/90000"},
	//	{Key: "rtpmap", Value: "97 MPEG4/90000"},
	//}
	//origin := sdp.Origin{
	//	Username:       device.DeviceId,
	//	SessionVersion: 0,
	//	SessionId:      0,
	//	NetType:        "IN",
	//	AddrType:       "IP4",
	//	UnicastAddress: global.CONFIG.Media.Address,
	//}
	//
	//body := sdp.SdpSession{
	//	Version:     0,
	//	Origin:      origin,
	//	SessionName: name,
	//	ConnectionData: sdp.ConnectionData{
	//		NetType:           "IN",
	//		AddrType:          "IP4",
	//		ConnectionAddress: global.CONFIG.Media.Address,
	//	},
	//	Timing:     []sdp.Timing{sdp.Timing{Start: stream.Start, End: stream.End}},
	//	MediaName:  mediaName,
	//	Attributes: attributes,
	//	SSRC:       stream.StreamId,
	//}
	//global.Logger.Info(body.String())
	var (
		s sdp.Session
		b []byte
	)
	name := "Play"
	protocal := "TCP/RTP/AVP"
	//if data.T == 1 {
	//	name = "Playback"
	//	protocal = "RTP/RTCP"
	//}
	port, _ := strconv.Atoi(global.CONFIG.Media.RTP)
	video := sdp.Media{
		Description: sdp.MediaDescription{
			Type:     "video",
			Port:     port,
			Formats:  []string{"96", "98", "97"},
			Protocol: protocal,
		},
	}
	video.AddAttribute("recvonly")
	//if data.T == 0 {
	//	video.AddAttribute("setup", "passive")
	//	video.AddAttribute("connection", "new")
	//}
	video.AddAttribute("rtpmap", "96", "PS/90000")
	video.AddAttribute("rtpmap", "98", "H264/90000")
	video.AddAttribute("rtpmap", "97", "MPEG4/90000")

	// defining message
	msg := &sdp.Message{
		Origin: sdp.Origin{
			Username: device.DeviceId, // 媒体服务器id
			Address:  global.CONFIG.Media.Address,
		},
		Name: name,
		Connection: sdp.ConnectionData{
			IP:  net.ParseIP(global.CONFIG.Media.Address),
			TTL: 0,
		},
		Timing: []sdp.Timing{
			{
				Start: time.Time{},
				End:   time.Time{},
			},
		},
		Medias: []sdp.Media{video},
		SSRC:   stream.StreamId,
	}
	//if data.T == 1 {
	//	msg.URI = fmt.Sprintf("%s:0", channel.ChannelID)
	//}

	// appending message to session
	s = msg.Append(s)
	// appending session to byte buffer
	b = s.AppendTo(b)
	request, _ := createVideoRequest(channel, device, sip.INVITE, string(b))
	//request, _ := createVideoRequest(channel, device, sip.INVITE, body.String())
	global.Logger.Info("创建点播请求成功")
	tx, err := Server.Request(request)
	// TODO 需要重新考虑
	if err != nil {
		return nil, err
	}
	resp := <-tx.Responses()
	global.Logger.Info("接受到Invite回应")
	ack := sip.NewAckRequest("", request, resp, "", nil)
	ack.SetRecipient(request.Recipient())
	ack.AppendHeader(&sip.ContactHeader{
		Address: ack.Recipient(),
		Params:  nil,
	})
	global.Logger.Info("发送ACK请求")
	err = Server.Send(ack)
	if err != nil {
		global.Logger.Error("发送ACK请求失败", zap.Error(err))
	}
	return stream, nil
}

func GetSSRC(c *system.DeviceChannel) string {
	for {
		key := fmt.Sprintf("%s%s%s", c.ChannelId[17:20], sipServer.Realm[7:10], utils.Get4SSRC())
		fmt.Println(key)
		streamT := &system.Stream{StreamId: key}
		_, err := streamT.GetStreamById()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return key
		}
	}
}

func ToHex(s string) string {
	d := new(big.Int)
	d.SetString(s, 10)
	return fmt.Sprintf("%08s", strings.ToUpper(d.Text(16)))
}
