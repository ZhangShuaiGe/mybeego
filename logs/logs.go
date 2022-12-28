package logs

import "github.com/beego/beego/v2/adapter/logs"

func init() {

	// 配置输出的位置
	// logs.AdapterConsole 只输出到控制台
	// logs.SetLogger(logs.AdapterConsole)

	// 输出代码行号
	logs.EnableFuncCallDepth(true)
	// 异步输出日志 提升性能
	logs.Async()
	//an official log.Logger
	l := logs.GetLogger()
	l.Println("this is a message of http")
	//an official log.Logger with prefix ORM
	logs.GetLogger("ORM").Println("this is a message of orm")

	logs.Debug("my book is bought in the year of ", 2016)
	logs.Info("this %s cat is %v years old", "yellow", 3)
	logs.Warn("json is a type of kv like", map[string]int{"key": 2016})
	logs.Error(1024, "is a very", "good game")
	logs.Critical("oh,crash")
}
