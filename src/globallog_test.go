package globallog

import (
	"os"
	"testing"
)

func TestGetLogger(t *testing.T) {
	TAG := "TestLog"

	log := GetLogger()
	defer CloseLogger()
	if log == nil {
		t.Log("获取日志记录器失败。")
		t.Fail()
	}

	log.Info("测试临时路径：", getExeFilePath())
	log.Finest(TAG, "Finest")
	log.Fine(TAG, "Fine")
	log.Info(TAG, "Info")
	log.Warn(TAG, "Warn")
	log.Trace(TAG, "Trace")
	log.Error(TAG, "Error")
	log.Critical(TAG, "Critical")

	_, err := os.Open(getLogFilePath())
	if err != nil && os.IsNotExist(err) {
		log.Info(TAG, "日志文件不存在。")
		t.Fail()
	} else {
		log.Info(TAG, "检测到了日志文件：", getLogFilePath())
	}
}
