module trader

go 1.16

require goex v0.0.0

replace goex => ./../goex

require (
	github.com/markcheno/go-talib v0.0.0-20190307022042-cd53a9264d70
	github.com/robfig/cron v1.2.0
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/viper v1.10.1
)
