package logging

import (
	"designer-api/pkg/setting"
	"fmt"
	"time"
)

// getLogFilePath get the log file save path
func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.LogSetting.RootPath, setting.LogSetting.SavePath)
}

// getLogFileName get the save name of the log file
func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		setting.LogSetting.SaveName,
		time.Now().Format(setting.LogSetting.TimeFormat),
		setting.LogSetting.FileExt,
	)
}
