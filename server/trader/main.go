package main

import (
	"fmt"
	"goex"
	"trader/common"

	"github.com/robfig/cron"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"strconv"
	"time"
	"trader/position"
	"trader/strategy"
	"trader/trade"
)


var account *trade.Account



func initConfig() bool {
	viper.SetConfigName("config.toml")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() //根据上面配置加载文件
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}


func main(){

	if initConfig() != true{
		return
	}


	account = trade.NewAccount("c2eeb999-72f2-4266-9ac0-92c92b3a0404",
		"668BD1201C828FD516FE6D2538FC08A7",
		"123456789qq")


	tick := time.Tick(time.Minute)

	c := cron.New()


	product_size := viper.GetInt("product.num")


	for i := 0; i < product_size; i++{

		cfg_key := "product" + strconv.Itoa(i)
		cfg_name := viper.GetString(cfg_key + ".name")
		cfg_money := viper.GetFloat64(cfg_key + ".money")
		cfg_strategy := viper.GetString(cfg_key + ".strategy")
		cfg_cron	:= viper.GetInt32(cfg_key + ".cron")
		cfg_avg :=viper.GetInt(cfg_key + ".avg")
		cfg_breakeven := viper.GetBool(cfg_key + ".breakeven")


		//如果策略是均线策略
		if cfg_strategy == "ma"{

			position := position.NewPositionV(cfg_name, cfg_money, cfg_breakeven)

			trade := trade.NewTrade(
					position,
					strategy.NewMa( goex.KlinePeriod(cfg_cron), cfg_avg, position),
					account)

			if v, ok := common.Corn[cfg_cron]; ok{
				c.AddFunc(v, func(){
					trade.Tick()
				})
			}

			ret := account.Add(trade)
			if ret == false{
				return
			}
		}
	}

	c.Run()

	//tick
	go func(){
		for {
			<-tick
			account.Tick()
		}
	}()
	//开启web服务
	http.HandleFunc("/acc", AccountInfo)
	http.HandleFunc("/trade", TradeInfo)
	err := http.ListenAndServe(viper.GetString("server.port"), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}

}


func AccountInfo(w http.ResponseWriter, r *http.Request){

	w.Write([]byte("AccountInfo"))

	common.Logger.Info("AccountInfo")

}

func TradeInfo(w http.ResponseWriter, r *http.Request){

	w.Write([]byte("TradeInfo"))

	common.Logger.Info("AccountInfo")
}