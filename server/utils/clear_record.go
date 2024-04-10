package utils

import (
	"github.com/shirou/gopsutil/v3/disk"
	"go.uber.org/zap"
	"nebula.xyz/global"
	"nebula.xyz/helper"
	"nebula.xyz/model/system"
	"os"
	"strconv"
	"sync"
)

var mutex sync.Mutex

func ClearRecordVideo() {
	recordPath := global.CONFIG.Media.RecordPath

	mutex.Lock()
	defer mutex.Unlock()

	usage, err := disk.Usage(recordPath)
	global.Logger.Debug("磁盘使用", zap.Float64("", usage.UsedPercent))
	if err != nil {
		global.Logger.Warn("文件路径不存在", zap.String("path", recordPath))
		return
	}
	errDelete := 0
	for usage.UsedPercent > global.CONFIG.Media.StorageThreshold {
		if errDelete > helper.MaxDeleteCount {
			return
		}
		var record system.Record
		err := global.DB.Model(&system.Record{}).Order("start_time asc").First(&record).Error
		global.Logger.Debug("文件路径", zap.String("path", record.FilePath))
		if err != nil {
			global.Logger.Error("获取录像记录失败", zap.Error(err))
			return
		}
		err = os.Remove(record.FilePath)
		if err != nil {
			global.Logger.Error("删除文件失败", zap.String("path", record.FilePath+"/ "+err.Error()))
			return
		}
		err = global.DB.Model(&system.Record{}).Where("id = ?", record.ID).Delete(nil).Error
		if err != nil {
			global.Logger.Error("删除文件错误", zap.String("文件信息", strconv.Itoa(int(record.ID))+"/ "+err.Error()))
			errDelete += 1
		}
		usage, err = disk.Usage(recordPath)
		if err != nil {
			global.Logger.Warn("文件路径不存在", zap.String("path", recordPath))
			return
		}
		global.Logger.Debug("磁盘使用", zap.Float64("", usage.UsedPercent))
	}

}
