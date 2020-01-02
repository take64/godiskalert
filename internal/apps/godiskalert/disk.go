package godiskalert

import (
	"fmt"
	"syscall"
)

// ディスク情報
type DiskInfo struct {
	Total uint64
	Free uint64
	Used uint64
	FreePercent float64
	UsedPercent float64
}

// フォーマット化されたディスク情報
type FormattedDiskInfo struct {
	Total string `json:"total"`
	Free string `json:"free"`
	Used string `json:"used"`
	FreePercent string `json:"free_percent"`
	UsedPercent string `json:"used_percent"`
}

// サイズ定数
const (
	B = 1
	KB = B * 1024
	MB = KB * 1024
	GB = MB * 1024
	TB = GB * 1024
	)

// ディスクの空きスペース等の情報を取得
func Info(path string) (result DiskInfo) {
	// ディスク情報の取得
	stat := syscall.Statfs_t{}
	err := syscall.Statfs(path, &stat)
	if err != nil {
		fmt.Println("ディスク情報の取得に失敗しました")
		return
	}

	// disk情報の設定
	result.Total = stat.Blocks * uint64(stat.Bsize)
	result.Free = stat.Bavail * uint64(stat.Bsize)
	result.Used = result.Total - result.Free
	result.FreePercent = float64(result.Free) / float64(result.Total) * 100
	result.UsedPercent = 100 - result.FreePercent

	return result
}

// フォーマット化されたディスク情報を取得
func (disk DiskInfo) Format() (result FormattedDiskInfo) {
	// フォーマットする
	result.Total = formatSize(disk.Total)
	result.Free = formatSize(disk.Free)
	result.Used = formatSize(disk.Used)
	result.FreePercent = formatPercent(disk.FreePercent)
	result.UsedPercent = formatPercent(disk.UsedPercent)
	return result
}

// フォーマットする(サイズ)
func formatSize(size uint64) string {
	if TB < size {
		return fmt.Sprintf("%d TB", size / TB)
	}
	if GB < size {
		return fmt.Sprintf("%d GB", size / GB)
	}
	if MB < size {
		return fmt.Sprintf("%d MB", size / MB)
	}
	if KB < size {
		return fmt.Sprintf("%d KB", size / KB)
	}
	return fmt.Sprintf("%d B", size)
}

// フォーマットする(パーセント)
func formatPercent(percent float64) string {
	return fmt.Sprintf("%.2f %%", percent)
}
