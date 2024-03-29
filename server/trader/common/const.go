package common

import "goex"

type OperType int32

const(
	OperType_None OperType = iota
	OperType_Buy    				//做多
	OperType_Sell					 //做空
	OperType_Add					 //加仓
	OperType_Cover  				 //平仓

)

type OperDetails struct{
	Op OperType
	Money float64 //钱
	Num int //购买的张数
	Price float64 //价格
}




var Corn  map[int32]string = map[int32]string{
	goex.KLINE_PERIOD_1MIN: "0 */1 * * * ?",
	goex.KLINE_PERIOD_3MIN: "0 */3 * * * ?",
	goex.KLINE_PERIOD_5MIN: "0 */5 * * * ?",
	goex.KLINE_PERIOD_15MIN: "0 */15 * * * ?",
	goex.KLINE_PERIOD_30MIN: "0 */30 * * * ?",
	goex.KLINE_PERIOD_60MIN: "0 */60 * * * ?",
	goex.KLINE_PERIOD_1H: "0 0 */1 * * ?",
	goex.KLINE_PERIOD_2H: "0 0 */2 * * ?",
	goex.KLINE_PERIOD_3H: "0 0 */3 * * ?",
	goex.KLINE_PERIOD_4H: "0 0 */4 * * ?",
	goex.KLINE_PERIOD_6H: "0 0 */6 * * ?",
	goex.KLINE_PERIOD_8H: "0 0 */8 * * ?",
	goex.KLINE_PERIOD_12H: "0 0 */12 * * ?",
	goex.KLINE_PERIOD_1DAY: "0 0 0 * * *",
}



var PairMap map[string]goex.CurrencyPair = map[string]goex.CurrencyPair{

	"BTC-USDT-SWAP" : goex.BTC_USDT_SWAP,
	"SOS-USDT-SWAP": goex.SOS_USDT_SWAP,
	"BABYDOGE-USDT-SWAP": goex.BABYDOGE_USDT_SWAP,
	"ETH-USDT-SWAP": goex.ETC_USDT_SWAP,
	"LTC-USDT-SWAP": goex.LTC_USDT_SWAP,
	"DOT-USDT-SWAP": goex.DOT_USDT_SWAP,
	"DOGE-USDT-SWAP": goex.DOGE_USDT_SWAP,
	"ADA-USDT-SWAP": goex.ADA_USDT_SWAP,
	"SOL-USDT-SWAP": goex.SOL_USDT_SWAP,
	"ALGO-USDT-SWAP": goex.ALGO_USDT_SWAP,
	"ATOM-USDT-SWAP": goex.ATOM_USDT_SWAP,
	"AVAX-USDT-SWAP": goex.AVAX_USDT_SWAP,
	"ANC-USDT-SWAP": goex.ANC_USDT_SWAP,
	"ANT-USDT-SWAP": goex.ANT_USDT_SWAP,
	"AXS-USDT-SWAP": goex.AXS_USDT_SWAP,
	"BADGER-USDT-SWAP": goex.BADGER_USDT_SWAP,
	"BAL-USDT-SWAP": goex.BAL_USDT_SWAP,
	"BAND-USDT-SWAP": goex.BAND_USDT_SWAP,
	"BAT-USDT-SWAP": goex.BAT_USDT_SWAP,
	"BCH-USDT-SWAP": goex.BCH_USDT_SWAP,
	"BICO-USDT-SWAP": goex.BICO_USDT_SWAP,
	"BNT-USDT-SWAP": goex.BNT_USDT_SWAP,
	"BSV-USDT-SWAP": goex.BSV_USDT_SWAP,
	"BTM-USDT-SWAP": goex.BTC_USDT_SWAP,
	"BZZ-USDT-SWAP": goex.BZZ_USDT_SWAP,
	"CEL-USDT-SWAP": goex.CEL_USDT_SWAP,
	"CFX-USDT-SWAP": goex.CFX_USDT_SWAP,
	"CRV-USDT-SWAP": goex.CRV_USDT_SWAP,
	"CHZ-USDT-SWAP": goex.CHZ_USDT_SWAP,
	"COMP-USDT-SWAP": goex.COMP_USDT_SWAP,
	"CONV-USDT-SWAP": goex.CONV_USDT_SWAP,
	"CQT-USDT-SWAP": goex.CQT_USDT_SWAP,
	"CRO-USDT-SWAP": goex.CRO_USDT_SWAP,
	"CSPR-USDT-SWAP": goex.CSPR_USDT_SWAP,
	"CVC-USDT-SWAP": goex.CVC_USDT_SWAP,
	"DASH-USDT-SWAP": goex.DASH_USDT_SWAP,
	"DORA-USDT-SWAP": goex.DORA_USDT_SWAP,
	"DYDX-USDT-SWAP": goex.DYDX_USDT_SWAP,
	"EFI-USDT-SWAP": goex.EFI_USDT_SWAP,
	"EGLD-USDT-SWAP": goex.EGLD_USDT_SWAP,
	"ELON-USDT-SWAP": goex.ELON_USDT_SWAP,
	"ENJ-USDT-SWAP": goex.ENJ_USDT_SWAP,
	"ENS-USDT-SWAP": goex.ENS_USDT_SWAP,
	"EOS-USDT-SWAP": goex.EOS_USDT_SWAP,
	"ETC-USDT-SWAP": goex.ETC_USDT_SWAP,
	"FIL-USDT-SWAP": goex.FIL_USDT_SWAP,
	"FLM-USDT-SWAP": goex.FLM_USDT_SWAP,
	"GALA-USDT-SWAP": goex.GALA_USDT_SWAP,
	"GOD-USDT-SWAP": goex.GOD_USDT_SWAP,
	"GRT-USDT-SWAP": goex.GRT_USDT_SWAP,
	"ICP-USDT-SWAP": goex.ICP_USDT_SWAP,
	"IMX-USDT-SWAP": goex.IMX_USDT_SWAP,
	"IOST-USDT-SWAP": goex.IOST_USDT_SWAP,
	"IOTA-USDT-SWAP": goex.IOTA_USDT_SWAP,
	"JST-USDT-SWAP": goex.JST_USDT_SWAP,
	"KISHU-USDT-SWAP": goex.KISHU_USDT_SWAP,
	"KNC-USDT-SWAP": goex.KNC_USDT_SWAP,
	"KSM-USDT-SWAP": goex.KSM_USDT_SWAP,
	"LAT-USDT-SWAP": goex.LAT_USDT_SWAP,
	"LINK-USDT-SWAP": goex.LINK_USDT_SWAP,
	"LON-USDT-SWAP": goex.LON_USDT_SWAP,
	"LPT-USDT-SWAP": goex.LPT_USDT_SWAP,
	"LRC-USDT-SWAP": goex.LRC_USDT_SWAP,
	"MANA-USDT-SWAP": goex.MANA_USDT_SWAP,
	"MASK-USDT-SWAP": goex.MASK_USDT_SWAP,
	"MATIC-USDT-SWAP": goex.MATIC_USDT_SWAP,
	"MINA-USDT-SWAP": goex.MINA_USDT_SWAP,
	"MIR-USDT-SWAP": goex.MIR_USDT_SWAP,
	"MKR-USDT-SWAP": goex.MKR_USDT_SWAP,
	"NEAR-USDT-SWAP": goex.NEAR_USDT_SWAP,
	"NEO-USDT-SWAP":goex.NEO_USDT_SWAP,
	"NFT-USDT-SWAP":goex.NFT_USDT_SWAP,
	"OMG-USDT-SWAP":goex.OMG_USDT_SWAP,
	"ONT-USDT-SWAP":goex.ONT_USDT_SWAP,
	"PEOPLE-USDT-SWAP":goex.PEOPLE_USDT_SWAP,
	"PERP-USDT-SWAP":goex.PERP_USDT_SWAP,
	"QTUM-USDT-SWAP": goex.QTUM_USDT_SWAP,
	"REN-USDT-SWAP": goex.REN_USDT_SWAP,
	"RSR-USDT-SWAP": goex.RSR_USDT_SWAP,
	"RVN-USDT-SWAP": goex.RVN_USDT_SWAP,
	"SAND-USDT-SWAP":goex.SAND_USDT_SWAP,
	"SC-USDT-SWAP":goex.SC_USDT_SWAP,
	"SLP-USDT-SWAP":goex.SLP_USDT_SWAP,
	"SNX-USDT-SWAP":goex.SNX_USDT_SWAP,
	"SRM-USDT-SWAP":goex.SRM_USDT_SWAP,
	"STARL-USDT-SWAP":goex.STARL_USDT_SWAP,
	"STORJ-USDT-SWAP":goex.STORJ_USDT_SWAP,
	"SUSHI-USDT-SWAP":goex.SUSHI_USDT_SWAP,
	"SWRV-USDT-SWAP":goex.SWRV_USDT_SWAP,
	"THETA-USDT-SWAP":goex.THETA_USDT_SWAP,
	"TONCOIN-USDT-SWAP":goex.TONCOIN_USDT_SWAP,
	"TORN-USDT-SWAP":goex.TORN_USDT_SWAP,
	"TRB-USDT-SWAP":goex.TRB_USDT_SWAP,
	"TRX-USDT-SWAP":goex.TRX_USDT_SWAP,
	"UMA-USDT-SWAP":goex.UMA_USDT_SWAP,
	"UNI-USDT-SWAP":goex.UNI_USDT_SWAP,
	"WAVES-USDT-SWAP":goex.WAVES_USDT_SWAP,
	"WNCG-USDT-SWAP":goex.WNCG_USDT_SWAP,
	"WNXM-USDT-SWAP":goex.WNXM_USDT_SWAP,
	"XCH-USDT-SWAP":goex.XCH_USDT_SWAP,
	"XEM-USDT-SWAP":goex.XEM_USDT_SWAP,
	"XLM-USDT-SWAP":goex.XLM_USDT_SWAP,
	"XMR-USDT-SWAP":goex.XMR_USDT_SWAP,
	"XRP-USDT-SWAP":goex.XRP_USDT_SWAP,
	"XTZ-USDT-SWAP":goex.XTZ_USDT_SWAP,
	"YFI-USDT-SWAP":goex.YFI_USDT_SWAP,
	"YFII-USDT-SWAP":goex.YFII_USDT_SWAP,
	"YGG-USDT-SWAP":goex.YGG_USDT_SWAP,
	"ZEC-USDT-SWAP":goex.ZEC_USDT_SWAP,
	"ZEN-USDT-SWAP":goex.ZEN_USDT_SWAP,
	"ZIL-USDT-SWAP":goex.ZIL_USDT_SWAP,
	"ZRX-USDT-SWAP":goex.ZRX_USDT_SWAP,

}


var GoexApi goex.API