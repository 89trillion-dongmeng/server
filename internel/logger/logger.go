package logger

import "github.com/beego/beego/v2/core/logs"

var Logger= logs.NewLogger()
func init() {
	Logger.SetLogger(logs.AdapterConsole)
	Logger.Debug("this is a debug message")
}
