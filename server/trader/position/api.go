package position

import (
	"trader/common"
)

//策略接口
type API interface {
	Name() string
	Num() float64

	IsEmpty() bool

	Direction() common.OperType
	KaiCang_money() float64

	Init()
	PullData()
	Clear()
	AvgPx() float64
	GetOperDetails(op common.OperType) *common.OperDetails

	FristPosition(p, n float64)
}


