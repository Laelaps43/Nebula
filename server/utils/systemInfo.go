package utils

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"nebula.xyz/global"
	"time"
)

func GetSystemInfo() {
	for {
		currentTime := time.Now()
		global.Info.TimeList = append(global.Info.TimeList[1:60], currentTime.Format("15:04"))
		//fmt.Println(global.Info.TimeList)

		// 获取CPU百分比
		//fmt.Println("-----------CPU-----------------")
		percent, _ := cpu.Percent(0, false)
		global.Info.CPUList = append(global.Info.CPUList[1:60], fmt.Sprintf("%.2f", percent[0]))
		//fmt.Println(global.CPUList)

		// 内存百分比
		//fmt.Println("-----------MEM-----------------")
		memory, _ := mem.SwapMemory()
		global.Info.MemList = append(global.Info.MemList[1:60], fmt.Sprintf("%.2f", memory.UsedPercent))
		//fmt.Println(global.MemList)

		// 硬盘占比
		//fmt.Println("-----------Disk-----------------")
		//usage, _ := disk.Usage("/")
		partitions, _ := disk.Partitions(true)
		for _, partition := range partitions {
			if partition.Device == "/dev/disk3s1" {
				usage, _ := disk.Usage(partition.Mountpoint)
				global.Info.DiskList = append(global.Info.DiskList[1:60], fmt.Sprintf("%.2f",
					usage.UsedPercent))
			}
		}
		//global.DiskList = append(global.DiskList[1:60], fmt.Sprintf("%.2f", usage.UsedPercent))
		//fmt.Println(global.DiskList)

		//fmt.Println("-----------Net-----------------")
		netStatus, _ := net.IOCounters(true)
		var oldStats net.IOCountersStat
		for _, stat := range netStatus {
			if stat.Name == "en0" {
				oldStats = stat
			}
		}
		time.Sleep(5 * time.Second)

		netStatus, _ = net.IOCounters(true)
		var newStats net.IOCountersStat
		for _, stat := range netStatus {
			if stat.Name == "en0" {
				newStats = stat
			}
		}
		calculateSpeed(oldStats, newStats)
		//fmt.Println(global.DownList)
		//fmt.Println(global.UpList)
		time.Sleep(time.Minute)
	}
}

// 计算速度（Mbps）
func calculateSpeed(oldStats, newStats net.IOCountersStat) {
	// 计算字节数差值
	bytesSentDiff := newStats.BytesSent - oldStats.BytesSent
	bytesRecvDiff := newStats.BytesRecv - oldStats.BytesRecv

	// 计算速度（Mbps）
	speedMbpsSent := float64(bytesSentDiff) / 1e6 / 5.0 * 8
	speedMbpsRecv := float64(bytesRecvDiff) / 1e6 / 5.0 * 8

	global.Info.UpList = append(global.Info.UpList[1:60], fmt.Sprintf("%.2f", speedMbpsSent))
	global.Info.DownList = append(global.Info.DownList[1:60], fmt.Sprintf("%.2f", speedMbpsRecv))
}
