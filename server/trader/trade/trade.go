package trade

import (
	"fmt"
	"github.com/spf13/viper"
	"goex/mail"
	"strconv"
	"strings"
	"trader/common"
	"trader/position"
)
import "trader/strategy"

type Trade struct{
	positon_interface position.API
	strategy_interface strategy.API
	account *Account
}

func NewTrade(p position.API, s strategy.API, a *Account ) *Trade{
	t := &Trade{
		positon_interface:p,
		strategy_interface:s,
		account: a,
	}
	return t
}

func (this *Trade) Init(){
	this.positon_interface.Init()
	this.strategy_interface.Init()
}
func (this *Trade) Tick(){

	if v, ok :=  common.PairMap[this.positon_interface.Name()]; ok{

		lines, err := common.GoexApi.GetKlineRecords(v, this.strategy_interface.Period(), this.strategy_interface.Avg(), 0, 0)
		if err != nil{
			common.Logger.Error("获取产品历史K线数据失败 %s", this.positon_interface.Name())
			return
		}

		oper := this.strategy_interface.Tick(lines)
		details := this.positon_interface.GetOperDetails(oper)
		if details !=nil{
			this.DoTrade(details)
		}
	}else{
		common.Logger.Error("产品类型错误 %s", this.positon_interface.Name())
	}

}

//执行操作
func (this *Trade) DoTrade( details *common.OperDetails){

	if v, ok := common.PairMap[this.positon_interface.Name()]; ok{

		switch details.Op{
		case common.OperType_Buy:

			isEmpty := this.positon_interface.IsEmpty()

			_, err := common.GoexApi.MarketBuy( strconv.Itoa(details.Num), "buy", v)

			if err != nil{

				desc := fmt.Sprintf("做多产品 %s 数量%v, price = %f 失败 %s", this.positon_interface.Name(), details.Num, details.Price, err.Error())
				common.Logger.Error(desc)

				mails := strings.Split(viper.GetString("server.mail"), ",")
				for _, recv :=range mails{
					mail.SendMail(recv, "交易失败", desc)
				}


			}else{


				if isEmpty{
					this.positon_interface.FristPosition(details.Price, float64(details.Num))
				}

				desc := fmt.Sprintf("做多产品 %s 数量%v, price = %f  money = %v 成功 ", this.positon_interface.Name(), details.Num, details.Price, details.Money)
				common.Logger.Info(desc)

				mails := strings.Split(viper.GetString("server.mail"), ",")
				for _, recv :=range mails{
					mail.SendMail(recv, "做多交易成功", desc)
				}

			}


		case common.OperType_Sell:

			isEmpty := this.positon_interface.IsEmpty()
			_, err := common.GoexApi.MarketBuy( strconv.Itoa(details.Num), "sell", v)
			if err != nil{

				desc := fmt.Sprintf("做空产品 %s 数量%v, price = %v 失败 %v", this.positon_interface.Name(), details.Num, details.Price, err.Error())
				common.Logger.Error(desc)


				mails := strings.Split(viper.GetString("server.mail"), ",")
				for _, recv :=range mails{
					mail.SendMail(recv, "交易失败", desc)
				}


			}else{

				if isEmpty{
					this.positon_interface.FristPosition(details.Price, float64(details.Num))
				}

				desc := fmt.Sprintf("做空产品 %s 数量%v, price = %v money = %v 成功 ", this.positon_interface.Name(), details.Num,  details.Price,details.Money)
				common.Logger.Info(desc)


				mails := strings.Split(viper.GetString("server.mail"), ",")
				for _, recv :=range mails{
					mail.SendMail(recv, "做空交易成功", desc)
				}

			}
		case common.OperType_Cover:

			ret, err := common.GoexApi.ClosePosition(this.positon_interface.Name(), "cross", "USDT")
			if err != nil{
				desc := fmt.Sprintf("平仓产品 %v 数量 %v,失败 %v", this.positon_interface.Name(), details.Num,  err.Error())
				common.Logger.Info(desc)

				mails := strings.Split(viper.GetString("server.mail"), ",")
				for _, recv :=range mails{
					mail.SendMail(recv, "平仓失败", desc)
				}


			}
			if ret == true{

				this.positon_interface.Clear()

				desc := fmt.Sprintf("平仓产品 %v 数量%v,  成功 ", this.positon_interface.Name(), details.Num)
				common.Logger.Info(desc)


				mails := strings.Split(viper.GetString("server.mail"), ",")
				for _, recv :=range mails{
					mail.SendMail(recv, "平仓成功", desc)
				}

			}
		default:
		}
	}
}
