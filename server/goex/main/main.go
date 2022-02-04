package main

import (
	"fmt"
	"goex"
	"goex/builder"
	"log"
	"math"
)

var (
	apiKey     = "c2eeb999-72f2-4266-9ac0-92c92b3a0404"
	secretKey  = "668BD1201C828FD516FE6D2538FC08A7"
	passPhrase = "123456789qq"
	leverage   = "100"
)

func main() {
	//创建现货api实例
	api := builder.DefaultAPIBuilder.APIKey(apiKey).APISecretkey(secretKey).ApiPassphrase(passPhrase).Build(goex.OKEX)
	/********* 获取深度，（深度数量，交易对）bidlist卖价，asklist买价 **********/
	//depth, _ := api.GetDepth(20, goex.BTC_USDT)

	/*********  **********/
	/********* 获取账号币种权益信息，account.SubAccounts[goex.USDT] **********/
	//account, _ := api.GetAccount()

	/********* Amount可用资金，ForzenAmount被冻结资金 **********/
	//balanceAvailable := account.SubAccounts[goex.USDT].Amount

	/********* 获取获取单个产品行情信息，ticker.last = 市价 **********/
	//ticker, _ := api.GetTicker(goex.BTC_USDT)

	/********* 设置杠杆倍数 **********/
	//api.SetLeverage("ETH-USDT-SWAP", "USDT", "125", "cross")
	//api.SetLeverage("BTC-USDT", "USDT", "10", "cross")
	//api.SetLeverage("BABYDOGE-USDT-SWAP", "USDT", "75", "cross")

	/*********
			   amount = 购买的金额 （如：买100美金10倍杠杆， 则保证金为10美金，amount填写100）
	           市价购买 MarketBuy(amount, price string, currency CurrencyPair)
			   BTC_USDT_SWAP 为永续合约交易， BTC_USDT 只是杠杆交易， 如需要其他币种还得去currencyPair添加
	**********/
	//tradeAmount := balanceAvailable * 0.01
	//tradeAmountStr := strconv.FormatFloat(tradeAmount, 'f', 0, 32)
	//order, buyErr := api.MarketBuy(tradeAmountStr, "buy", goex.BTC_USDT_SWAP)
	//api.MarketBuy("1", "sell", goex.BTC_USDT_SWAP)
	/********* BTC_USDT_SWAP 永续合约交易amount参数为btc 合约张数，最低1 **********/
	//tradeAmount := balanceAvailable * 0.01
	//tradeAmountStr := strconv.FormatFloat(tradeAmount, 'f', 0, 32)
	order, buyErr := api.MarketBuy("1", "buy", goex.LRC_USDT_SWAP)
	log.Println(order)
	log.Println(buyErr)
	/********* 市价全平 **********/
	//api.ClosePosition("BTC-USDT-SWAP", "cross", "USDT")

	/********* 查询当前持仓 **********/
	//pos, err := api.GetPositions("BTC-USDT-SWAP")
	//if err == nil {
	//	log.Println(pos.AvgPx)
	//} else {
	//	log.Println(err)
	//}

	/********* 查询历史k线数据 **********/
	//lines, _ := api.GetKlineRecords(goex.BTC_USDT_SWAP, goex.KLINE_PERIOD_1H, 120, 0, 0)
	//log.Println(lines)

	/********* 获取币种信息 **********/
	//info, err := api.GetCoinInfo(goex.BABYDOGE_USDT_SWAP, "SWAP")
	//log.Println(err)
	//log.Println(info.LotSz)

	/********* 获取MA **********/
	//lines, _ := api.GetKlineRecords(goex.BTC_USDT, goex.KLINE_PERIOD_1DAY, 120, 0, 0)
	//ma, _ := MA(lines, 120, "sma", "close")
	//log.Println(ma)

	/********* 获取ATR **********/
	//lines, _ := api.GetKlineRecords(goex.BTC_USDT, goex.KLINE_PERIOD_1DAY, 120, 0, 0)
	//atr, _ := ATR(lines, 14)
	//log.Println(atr)

	/********* 发送邮件 **********/
	//mail.SendMail("842588355@qq.com", "hello", "test")
}

func ATR(lines []goex.Kline, inTimePeriod int) (float64, error) {
	if len(lines) < inTimePeriod {
		return 0, fmt.Errorf("klines not enough error")
	}
	var sum float64
	sum = 0

	for i := 0; i < inTimePeriod; i++ {
		tmp := math.Max(lines[i].High, lines[i].Close-1)
		tmp -= math.Min(lines[i].Low, lines[i].Close-1)
		sum += tmp
	}
	return sum / float64(inTimePeriod), nil
}

func MA(lines []goex.Kline, size int, mode, lineType string) (float64, error) {
	if len(lines) < size {
		return 0, fmt.Errorf("klines not enough error")
	}

	if lineType == "close" {
		if mode == "sma" {
			var closeSum float64
			closeSum = 0
			for i := 0; i < size; i++ {
				closeSum += lines[i].Close
			}
			return closeSum / float64(size), nil
		}
	} else if lineType == "open" {
		if mode == "sma" {
			var closeSum float64
			closeSum = 0
			for i := 0; i < size; i++ {
				closeSum += lines[i].Open
			}
			return closeSum / float64(size), nil
		}
	}

	return 0, fmt.Errorf("parameter error")
}
