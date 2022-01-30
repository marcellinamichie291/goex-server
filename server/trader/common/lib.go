package common


import (
	"fmt"
	"github.com/markcheno/go-talib"
"goex"
	"math"
)

type PriceType int

const (
	InClose PriceType = iota + 1
	InHigh
	InLow
	InOpen
)

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

func Macd(data []goex.Kline, inFastPeriod int, inSlowPeriod int, inSignalPeriod int, priceTy PriceType) (DIF, DEA, MACD []float64) {
	var macd []float64
	dif, dea, hist := talib.Macd(realData(data, priceTy), inFastPeriod, inSlowPeriod, inSignalPeriod)
	for _, item := range hist {
		macd = append(macd, item*2)
	}
	return dif, dea, macd
}

func Boll(data []goex.Kline, inTimePeriod int, deviation float64, priceTy PriceType) (up, middle, low []float64) {
	return talib.BBands(realData(data, priceTy), inTimePeriod, deviation, deviation, 0)
}

func Rsi(data []goex.Kline, inTimePeriod int, priceTy PriceType) []float64 {
	return talib.Rsi(realData(data, priceTy), inTimePeriod)
}

func realData(data []goex.Kline, priceTy PriceType) []float64 {
	var inReal []float64
	for i := len(data) - 1; i >= 0; i-- {
		k := data[i]
		switch priceTy {
		case InClose:
			inReal = append(inReal, k.Close)
		case InHigh:
			inReal = append(inReal, k.High)
		case InLow:
			inReal = append(inReal, k.Low)
		case InOpen:
			inReal = append(inReal, k.Open)
		default:
			panic("please set ema type")
		}
	}
	return inReal
}