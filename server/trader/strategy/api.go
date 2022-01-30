package strategy

import (
	"goex"
	"trader/common"
)

//策略接口
type API interface {
	Tick(data []goex.Kline) (op common.OperType)
	Period() goex.KlinePeriod
	Init()
	Avg() int
}

