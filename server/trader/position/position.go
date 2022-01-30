package position

import (
	"goex"
	"strconv"
	"strings"
	"trader/common"
)

//////////////////////////////////////////////
//持仓 &仓位管理
type Position struct{
	name string   //产品名字
	//num float64   //当前数量 仓位
	cost float64  //成本价格
	direction common.OperType  //当前方向
	kaicang_money float64//第一次开多少钱的仓
	positions *goex.Positions
	coin_info *goex.Coin

}


func (this *Position) IsEmpty() bool{

	if this.positions == nil{
		return true
	}
	return false

}

func (this *Position) Init(){
	this.PullData()
}

//初始化 拉取数据
func (this *Position) PullData()(){

	p , err := common.GoexApi.GetPositions(this.name)
	if err != nil{
		common.Logger.Error("获取产品持仓信息失败 %s", this.name)
		return
	}else{
		this.positions = p
	}

	if v ,ok := common.PairMap[this.name]; ok{

		contain := strings.ContainsAny(this.name, "SWAP")
		if contain{
			ci, err := common.GoexApi.GetCoinInfo(v, "SWAP")  //永续合约
			if err != nil{
				common.Logger.Error("获取永续合约 SWAP corninfo失败 %s", this.name)
			}else{
				this.coin_info = ci
			}
		}else{

			ci, err := common.GoexApi.GetCoinInfo(v, "SPOT")  //币币
			if err != nil{
				common.Logger.Error("获取币币交易 SPOT corninfo失败 %s", this.name)
			}else{
				this.coin_info = ci
			}
		}
	}

	//设置杠杆倍数
	common.GoexApi.SetLeverage(this.name, "USDT", strconv.Itoa(int(this.coin_info.Lever)), "cross")



}
