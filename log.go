package archive

import (
	"io"
	"log"
	"os"
)

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	errFile, err := os.OpenFile("log/errors.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		Error.Fatalln("打开日志文件失败：", err.Error())
	}
	logFile, err := os.OpenFile("log/logs.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		Error.Fatalln("打开日志文件失败：", err.Error())
	}

	Info = log.New(os.Stderr, "【Info】", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(io.MultiWriter(os.Stderr, logFile), "【Warning】", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(os.Stderr, errFile), "【Error】", log.Ldate|log.Ltime|log.Lshortfile)

}
