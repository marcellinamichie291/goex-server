package strategy

import (
	"goex"
	"sync"
	"trader/common"
	"trader/position"
)


//ma策略
type MA struct{
	period goex.KlinePeriod
	avg int  //均线刻度
	data *[]goex.Kline  //k线数据
	api_position position.API
	mutex sync.Mutex
}

func NewMa(c goex.KlinePeriod, avg int ,api position.API) *MA{
	ma := &MA{
		period: c,
		avg : avg,
		api_position :api,
	}
	return ma
}

func (this *MA)Avg() int{
	return this.avg
}

func (this *MA) Init(){

}

func (this *MA) Period() goex.KlinePeriod{
	return this.period
}

func (this *MA) Tick(data []goex.Kline) (op common.OperType){

	this.mutex.Lock()
	defer this.mutex.Unlock()

	if len(data) == 0{
		common.Logger.Debug("k线数据错误")
	}

	op = common.OperType_None

	ma, err := common.MA(data, int(this.avg), "sma", "close")
	if err != nil{
		common.Logger.Error("获取均线失败 %s", err.Error())
		return op
	}
	close := data[0].Close

	if close > ma{
		return common.OperType_Buy
	}else{
		return common.OperType_Sell
	}
	return op
}