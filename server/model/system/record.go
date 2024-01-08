package system

import (
	"nebula.xyz/model"
)

type Record struct {
	model.NEBULA
	ID            uint    `gorm:"primaryKey comment:主键"`
	MediaServerId string  `gorm:"comment:媒体服务器Id" json:"mediaServerId"`
	App           string  `gorm:"comment:录制的流应用名" json:"app"`
	FileName      string  `gorm:"comment:文件名称" json:"file_name"`
	FilePath      string  `gorm:"comment:文件绝对路径" json:"file_path"`
	FileSize      int64   `gorm:"comment:文件大小，单位字节" json:"file_size"`
	Folder        string  `gorm:"comment:文件所在目录路径" json:"folder"`
	StartTime     int64   `gorm:"comment:开始录制时间戳" json:"start_time"`
	Stream        string  `gorm:"comment:流Id" json:"stream"`
	TimeLen       float64 `gorm:"comment:录制时长，单位秒" json:"time_len"`
	URL           string  `gorm:"comment:http/rtsp/rtmp 点播相对 url 路径" json:"url"`
	Vhost         string  `gorm:"comment:流虚拟主机" json:"vhost"`
	RecordDate    string  `gorm:"comment:记录日期; type:date" json:"record_date"`
}

func (r *Record) TableName() string {
	return "record"
}
