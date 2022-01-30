package trade

import (
	"fmt"
	"github.com/spf13/viper"
	"goex"
	"goex/builder"
	"goex/mail"
	"strings"
	"trader/common"
)

//账户信息
type Account struct {

	apiKey string
	secretKey string
	passPhrase string
	invest_money float64 //投入的钱

	trade map[string]*Trade //当前交易的品种

	account *goex.Account


}


func NewAccount(apiKey ,secretKey, passPhrase string)*Account{
	a := &Account{
		apiKey: apiKey,
		secretKey:secretKey,
		passPhrase:passPhrase,
		trade: make(map[string]*Trade, 0),
	}

	a.build()

	return a
}



//拉取数据
func (this *Account) build(){

	common.GoexApi  = builder.DefaultAPIBuilder.APIKey(this.apiKey).APISecretkey(this.secretKey).ApiPassphrase(this.passPhrase).Build(goex.OKEX)

	account , err := common.GoexApi.GetAccount()
	if err != nil{
		common.Logger.Info("获取账户信息失败")
	}else{
		this.account = account
	}

}

//加载服务器历史数据
func (this *Account)loadData(){

}

func (this *Account) Tick(){

	//打印账户信息

	acc, err := common.GoexApi.GetAccount()
	if err != nil{
		common.Logger.Error("获取账户信息失败")
	}else{

		desc := fmt.Sprintf("只打印USDT资产  %v USDT", acc.SubAccounts[goex.USDT].Amount + acc.SubAccounts[goex.USDT].ForzenAmount)
		common.Logger.Info(desc)

		mails := strings.Split(viper.GetString("server.mail"), ",")
		for _, recv :=range mails{
			mail.SendMail(recv, "账户信息", desc)
		}


	}


}

func (this *Account) Add(trade *Trade) bool{

	if _, ok := this.trade[trade.positon_interface.Name()]; ok{
		fmt.Println("加入产品失败 之前存在这个产品")
		return false
	}
	this.trade[trade.positon_interface.Name()] = trade

	return true

}

