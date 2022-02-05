package position

import (
	"fmt"
	"goex"
	"math"
	"sync"
	"trader/common"
)

//倒金字塔
type PositionV struct {
	Position

	coin *goex.Coin

	First_price float64 //首次开仓价格
	First_num   float64 //首仓的数量

	Add_times   int32   //加仓看多少次
	Percent_add float64 //加仓涨跌幅比例  设置（比如涨跌1%开始加仓）
	mutex       sync.Mutex
	breakeven   bool //是否开启保本
}

func NewPositionV(name string, money float64, bk bool) *PositionV {
	p := &PositionV{
		Position: Position{
			name:          name,
			kaicang_money: money,
		},
		Percent_add: 0.3, //加仓比例 30%
		breakeven:   bk,
	}
	p.init()

	return p
}

func (this *PositionV) FristPosition(price, num float64) {
	this.First_price = price
	this.First_num = num

}

func (this *PositionV) Name() string {
	return this.name
}
func (this *PositionV) Num() float64 {

	if this.positions == nil {
		return 0.0
	}

	return math.Abs(this.positions.Pos)
}

func (this *PositionV) Direction() common.OperType {

	if this.positions == nil { //没有仓位
		return common.OperType_None
	}
	if this.positions.Pos > 0 {
		return common.OperType_Buy
	}

	return common.OperType_Sell

}
func (this *PositionV) KaiCang_money() float64 {

	return this.kaicang_money
}

func (this *PositionV) AvgPx() float64 {
	if this.positions == nil {
		return 0.0
	}
	return this.positions.AvgPx
}

//首次购买的价格
func (this *PositionV) SetFirstPrice(p float64) {
	this.First_price = p
}

func (this *PositionV) Clear() {
	this.First_price = 0
	this.First_num = 0
	this.Add_times = 0
}

func (this *PositionV) init() bool {

	if v, ok := common.PairMap[this.Name()]; ok {

		coin, err := common.GoexApi.GetCoinInfo(v, "SWAP")
		if err != nil {
			common.Logger.Info("产生数据GetCoinInfo 失败 %s", err.Error())
			return false
		}
		this.coin = coin
		return true
	}

	common.Logger.Info("产生数据GetCoinInfo 失败 无该产品 %s", this.Name())

	return false
}

//执行策略
func (this *PositionV) GetOperDetails(op common.OperType) (oper *common.OperDetails) {

	this.mutex.Lock()
	defer this.mutex.Unlock()

	this.PullData() //拉取数据

	var tick *goex.Ticker
	oper = &common.OperDetails{Op: common.OperType_None}

	if v, ok := common.PairMap[this.Name()]; ok {
		t, err := common.GoexApi.GetTicker(v)
		if err != nil {
			common.Logger.Error("获取产品 %s 价格失败 %s", err.Error())
			return oper
		} else {
			tick = t
		}
	}
	if tick == nil {
		common.Logger.Error("获取产品 %s ")
		return oper
	}

	if this.IsEmpty() { //空仓
		switch op {
		case common.OperType_Buy:
			num := int(this.kaicang_money / tick.Sell / this.coin_info.CtVal)
			if num < int(this.coin_info.MinSz) {
				common.Logger.Error("失败 产品 %s, 做多数量 %v 低于最低数量 %v, money = %v, 价格 %v",
					this.Name(), float64(num)*this.coin_info.CtVal, this.coin_info.MinSz, tick.Buy, this.kaicang_money)
				return oper
			} else {
				oper.Op = common.OperType_Buy
				oper.Num = num //合约张数量
				oper.Price = tick.Buy
				oper.Money = float64(num) * tick.Buy * this.coin_info.CtVal
			}

		case common.OperType_Sell:
			num := int(this.kaicang_money / tick.Buy / this.coin_info.CtVal)
			if num < int(this.coin_info.MinSz) {
				common.Logger.Error("失败 产品 %s, 做空数量 %v 低于最低数量 %v, money = %v, 价格 %v",
					this.Name(), float64(num)*this.coin_info.CtVal, this.coin_info.MinSz, tick.Buy, this.kaicang_money)
				return oper
			} else {
				oper.Op = common.OperType_Sell
				oper.Num = num //合约张数量
				oper.Price = tick.Buy
				oper.Money = float64(num) * tick.Buy * this.coin_info.CtVal
			}

		}
	} else { //如果有仓位

		switch op {
		case common.OperType_Buy:
			if this.Direction() == common.OperType_Buy { //方向和策略信号一直 加仓位

				danci_add := int(this.First_num * this.Percent_add) //单次加的数量
				if danci_add == 0 {
					desc := fmt.Sprintf("计算加仓的数量失败  %v, 首次数量 %v, 百分比 %v",
						danci_add, this.First_num, this.Percent_add)
					common.Logger.Error(desc)
					return
				}
				percent := (tick.Buy - this.First_price) / this.First_price * 100                         //涨幅
				need_add := float64(danci_add)*percent + this.First_num - (this.Num())/float64(danci_add) //需要加仓的次数

				if need_add <= 0 {
					desc := fmt.Sprintf("做多无需加仓 产品 %s, 低于最低数量 %v, 涨幅 = %v, 价格 %v",
						this.Name(), this.coin_info.CtVal, percent/100, tick.Buy)
					common.Logger.Error(desc)
				} else {
					if need_add > 10 { //涨幅过快 一次性加两次
						need_add = 2
					} else {
						need_add = 1
					}
					oper.Op = common.OperType_Buy
					oper.Num = int(need_add) //合约张数量

				}
			} else { //方向和策略信号 不一致  平仓或者保持不变

				oper.Op = common.OperType_Cover
				oper.Num = int(this.Num()) //合约张数量

			}
		case common.OperType_Sell:
			if this.Direction() == common.OperType_Sell { //方向和策略信号一直 加仓位

				danci_add := int(this.First_num * this.Percent_add)                                       //单次加的数量
				percent := (this.First_price - tick.Buy) / tick.Buy                                       //跌幅
				need_add := float64(danci_add)*percent + this.First_num - (this.Num())/float64(danci_add) //需要加仓的次数 //需要加仓的次数

				if need_add <= 0 {
					desc := fmt.Sprintf("做空无需加仓 产品 %s, 低于最低数量 %v, 涨幅 = %v, 价格 %v",
						this.Name(), this.coin_info.CtVal, percent/100, tick.Buy)
					common.Logger.Info(desc)
				} else {

					if need_add > 10 { //涨幅过快 一次性加两次
						need_add = 2
					} else {
						need_add = 1
					}
					oper.Op = common.OperType_Sell
					oper.Num = int(need_add) //合约张数量

				}
			} else { //方向和策略信号 不一致  平仓或者保持不变

				oper.Op = common.OperType_Cover
				oper.Num = int(this.Num()) //合约张数量

			}

		}

	}

	return oper
}
