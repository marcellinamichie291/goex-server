package common

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var Logger  = logrus.New()

func init(){

	src, err := os.OpenFile( "./sys.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("err", err)
	}
	fileAndStdoutWriter := io.MultiWriter(os.Stdout, src)
	//设置输出
	Logger.Out = fileAndStdoutWriter

	//设置日志级别
	Logger.SetLevel(logrus.DebugLevel)

	//设置日志格式
	Logger.SetFormatter(&logrus.TextFormatter{})

}

