package goex

import "strings"

type Currency struct {
	Symbol string
	Desc   string
}

func (c Currency) String() string {
	return c.Symbol
}

func (c Currency) Eq(c2 Currency) bool {
	return c.Symbol == c2.Symbol
}

// A->B(A兑换为B)
type CurrencyPair struct {
	CurrencyA      Currency
	CurrencyB      Currency
	AmountTickSize int    // 下单量精度
	PriceTickSize  int    // 交易对价格精度
	MinSize        string //最小开单数量
}

var (
	UNKNOWN   = Currency{"UNKNOWN", ""}
	CNY       = Currency{"CNY", ""}
	USD       = Currency{"USD", ""}
	USDT      = Currency{"USDT", ""}
	USDT_SWAP = Currency{"USDT-SWAP", ""}
	PAX       = Currency{"PAX", "https://www.paxos.com/"}
	USDC      = Currency{"USDC", "https://www.centre.io/"}
	EUR       = Currency{"EUR", ""}
	KRW       = Currency{"KRW", ""}
	JPY       = Currency{"JPY", ""}
	BTC       = Currency{"BTC", "https://bitcoin.org/"}
	XBT       = Currency{"XBT", ""}
	BCC       = Currency{"BCC", ""}
	BCH       = Currency{"BCH", ""}
	BCX       = Currency{"BCX", ""}
	LTC       = Currency{"LTC", ""}
	ETH       = Currency{"ETH", ""}
	ETC       = Currency{"ETC", ""}
	EOS       = Currency{"EOS", ""}
	BTS       = Currency{"BTS", ""}
	QTUM      = Currency{"QTUM", ""}
	SC        = Currency{"SC", ""}
	ANS       = Currency{"ANS", ""}
	ZEC       = Currency{"ZEC", ""}
	DCR       = Currency{"DCR", ""}
	XRP       = Currency{"XRP", ""}
	BTG       = Currency{"BTG", ""}
	BCD       = Currency{"BCD", ""}
	NEO       = Currency{"NEO", ""}
	HSR       = Currency{"HSR", ""}
	BSV       = Currency{"BSV", ""}
	OKB       = Currency{"OKB", "OKB is a global utility token issued by OK Blockchain Foundation"}
	HT        = Currency{"HT", "HuoBi Token"}
	BNB       = Currency{"BNB", "BNB, or Binance Coin, is a cryptocurrency created by Binance."}
	TRX       = Currency{"TRX", ""}
	GBP       = Currency{"GBP", ""}
	XLM       = Currency{"XLM", ""}
	DOT       = Currency{"DOT", ""}
	DASH      = Currency{"DASH", ""}
	CRV       = Currency{"CRV", ""}
	ALGO      = Currency{"ALGO", ""}
	SOS       = Currency{"SOS", ""}
	BABYDOGE  = Currency{"BABYDOGE", ""}
	CFX       = Currency{"CFX", ""}

	DOGE    = Currency{"DOGE", ""}
	ADA     = Currency{"ADA", ""}
	SOL     = Currency{"SOL", ""}
	ATOM    = Currency{"ATOM", ""}
	AVAX    = Currency{"AVAX", ""}
	ANC     = Currency{"ANC", ""}
	ANT     = Currency{"ANT", ""}
	AXS     = Currency{"AXS", ""}
	BADGER  = Currency{"BADGER", ""}
	BAL     = Currency{"BAL", ""}
	BAND    = Currency{"BAND", ""}
	BAT     = Currency{"BAT", ""}
	BICO    = Currency{"BICO", ""}
	BNT     = Currency{"BNT", ""}
	BTM     = Currency{"BTM", ""}
	BZZ     = Currency{"BZZ", ""}
	CEL     = Currency{"CEL", ""}
	CHZ     = Currency{"CHZ", ""}
	COMP    = Currency{"COMP", ""}
	CONV    = Currency{"CONV", ""}
	CQT     = Currency{"CQT", ""}
	CRO     = Currency{"CRO", ""}
	CSPR    = Currency{"CSPR", ""}
	CVC     = Currency{"CVC", ""}
	DORA    = Currency{"DORA", ""}
	DYDX    = Currency{"DYDX", ""}
	EFI     = Currency{"EFI", ""}
	EGLD    = Currency{"EGLD", ""}
	ELON    = Currency{"ELON", ""}
	ENJ     = Currency{"ENJ", ""}
	ENS     = Currency{"ENS", ""}
	FIL     = Currency{"FIL", ""}
	FLM     = Currency{"FLM", ""}
	GALA    = Currency{"GALA", ""}
	GOD     = Currency{"GOD", ""}
	GRT     = Currency{"GRT", ""}
	ICP     = Currency{"ICP", ""}
	IMX     = Currency{"IMX", ""}
	IOST    = Currency{"IOST", ""}
	IOTA    = Currency{"IOTA", ""}
	JST     = Currency{"JST", ""}
	KISHU   = Currency{"KISHU", ""}
	KNC     = Currency{"KNC", ""}
	KSM     = Currency{"KSM", ""}
	LAT     = Currency{"LAT", ""}
	LINK    = Currency{"LINK", ""}
	LON     = Currency{"LON", ""}
	LPT     = Currency{"LPT", ""}
	LRC     = Currency{"LRC", ""}
	MANA    = Currency{"MANA", ""}
	MASK    = Currency{"MASK", ""}
	MATIC   = Currency{"MATIC", ""}
	MINA    = Currency{"MINA", ""}
	MIR     = Currency{"MIR", ""}
	MKR     = Currency{"MKR", ""}
	NEAR    = Currency{"NEAR", ""}
	NFT     = Currency{"NFT", ""}
	OMG     = Currency{"OMG", ""}
	ONT     = Currency{"ONT", ""}
	PEOPLE  = Currency{"PEOPLE", ""}
	PERP    = Currency{"PERP", ""}
	REN     = Currency{"REN", ""}
	RSR     = Currency{"RSR", ""}
	RVN     = Currency{"RVN", ""}
	SAND    = Currency{"SAND", ""}
	SLP     = Currency{"SLP", ""}
	SNX     = Currency{"SNX", ""}
	SRM     = Currency{"SRM", ""}
	STARL   = Currency{"STARL", ""}
	STORJ   = Currency{"STORJ", ""}
	SUSHI   = Currency{"SUSHI", ""}
	SWRV    = Currency{"SWRV", ""}
	THETA   = Currency{"THETA", ""}
	TONCOIN = Currency{"TONCOIN", ""}
	TORN    = Currency{"TORN", ""}
	TRB     = Currency{"TRB", ""}
	UMA     = Currency{"UMA", ""}
	UNI     = Currency{"UNI", ""}
	WAVES   = Currency{"WAVES", ""}
	WNCG    = Currency{"WNCG", ""}
	WNXM    = Currency{"WNXM", ""}
	XCH     = Currency{"XCH", ""}
	XEM     = Currency{"XEM", ""}
	XMR     = Currency{"XMR", ""}
	XTZ     = Currency{"XTZ", ""}
	YFI     = Currency{"YFI", ""}
	YFII    = Currency{"YFII", ""}
	YGG     = Currency{"YGG", ""}
	ZEN     = Currency{"ZEN", ""}
	ZIL     = Currency{"ZIL", ""}
	ZRX     = Currency{"ZRX", ""}

	//currency pair
	BTC_KRW = CurrencyPair{CurrencyA: BTC, CurrencyB: KRW, AmountTickSize: 2, PriceTickSize: 1}
	ETH_KRW = CurrencyPair{CurrencyA: ETH, CurrencyB: KRW, AmountTickSize: 2, PriceTickSize: 2}
	ETC_KRW = CurrencyPair{CurrencyA: ETC, CurrencyB: KRW, AmountTickSize: 2, PriceTickSize: 2}
	LTC_KRW = CurrencyPair{CurrencyA: LTC, CurrencyB: KRW, AmountTickSize: 2, PriceTickSize: 2}
	BCH_KRW = CurrencyPair{CurrencyA: BCH, CurrencyB: KRW, AmountTickSize: 2, PriceTickSize: 2}

	// Swap pair
	BTC_USDT_SWAP      = CurrencyPair{CurrencyA: BTC, CurrencyB: USDT_SWAP}
	SOS_USDT_SWAP      = CurrencyPair{CurrencyA: SOS, CurrencyB: USDT_SWAP}
	BABYDOGE_USDT_SWAP = CurrencyPair{CurrencyA: BABYDOGE, CurrencyB: USDT_SWAP}
	ETH_USDT_SWAP      = CurrencyPair{CurrencyA: ETH, CurrencyB: USDT_SWAP}
	LTC_USDT_SWAP      = CurrencyPair{CurrencyA: LTC, CurrencyB: USDT_SWAP}
	DOT_USDT_SWAP      = CurrencyPair{CurrencyA: DOT, CurrencyB: USDT_SWAP}
	DOGE_USDT_SWAP     = CurrencyPair{CurrencyA: DOGE, CurrencyB: USDT_SWAP}
	ADA_USDT_SWAP      = CurrencyPair{CurrencyA: ADA, CurrencyB: USDT_SWAP}
	SOL_USDT_SWAP      = CurrencyPair{CurrencyA: SOL, CurrencyB: USDT_SWAP}
	ALGO_USDT_SWAP     = CurrencyPair{CurrencyA: ALGO, CurrencyB: USDT_SWAP}
	ATOM_USDT_SWAP     = CurrencyPair{CurrencyA: ATOM, CurrencyB: USDT_SWAP}
	AVAX_USDT_SWAP     = CurrencyPair{CurrencyA: AVAX, CurrencyB: USDT_SWAP}
	ANC_USDT_SWAP      = CurrencyPair{CurrencyA: ANC, CurrencyB: USDT_SWAP}
	ANT_USDT_SWAP      = CurrencyPair{CurrencyA: ANT, CurrencyB: USDT_SWAP}
	AXS_USDT_SWAP      = CurrencyPair{CurrencyA: AXS, CurrencyB: USDT_SWAP}
	BADGER_USDT_SWAP   = CurrencyPair{CurrencyA: BADGER, CurrencyB: USDT_SWAP}
	BAL_USDT_SWAP      = CurrencyPair{CurrencyA: BAL, CurrencyB: USDT_SWAP}
	BAND_USDT_SWAP     = CurrencyPair{CurrencyA: BAND, CurrencyB: USDT_SWAP}
	BAT_USDT_SWAP      = CurrencyPair{CurrencyA: BAT, CurrencyB: USDT_SWAP}
	BCH_USDT_SWAP      = CurrencyPair{CurrencyA: BCH, CurrencyB: USDT_SWAP}
	BICO_USDT_SWAP     = CurrencyPair{CurrencyA: BICO, CurrencyB: USDT_SWAP}
	BNT_USDT_SWAP      = CurrencyPair{CurrencyA: BNT, CurrencyB: USDT_SWAP}
	BSV_USDT_SWAP      = CurrencyPair{CurrencyA: BSV, CurrencyB: USDT_SWAP}
	BTM_USDT_SWAP      = CurrencyPair{CurrencyA: BTM, CurrencyB: USDT_SWAP}
	BZZ_USDT_SWAP      = CurrencyPair{CurrencyA: BZZ, CurrencyB: USDT_SWAP}
	CEL_USDT_SWAP      = CurrencyPair{CurrencyA: CEL, CurrencyB: USDT_SWAP}
	CFX_USDT_SWAP      = CurrencyPair{CurrencyA: CFX, CurrencyB: USDT_SWAP}
	CRV_USDT_SWAP      = CurrencyPair{CurrencyA: CRV, CurrencyB: USDT_SWAP}
	CHZ_USDT_SWAP      = CurrencyPair{CurrencyA: CHZ, CurrencyB: USDT_SWAP}
	COMP_USDT_SWAP     = CurrencyPair{CurrencyA: COMP, CurrencyB: USDT_SWAP}
	CONV_USDT_SWAP     = CurrencyPair{CurrencyA: CONV, CurrencyB: USDT_SWAP}
	CQT_USDT_SWAP      = CurrencyPair{CurrencyA: CQT, CurrencyB: USDT_SWAP}
	CRO_USDT_SWAP      = CurrencyPair{CurrencyA: CRO, CurrencyB: USDT_SWAP}
	CSPR_USDT_SWAP     = CurrencyPair{CurrencyA: CSPR, CurrencyB: USDT_SWAP}
	CVC_USDT_SWAP      = CurrencyPair{CurrencyA: CVC, CurrencyB: USDT_SWAP}
	DASH_USDT_SWAP     = CurrencyPair{CurrencyA: DASH, CurrencyB: USDT_SWAP}
	DORA_USDT_SWAP     = CurrencyPair{CurrencyA: DORA, CurrencyB: USDT_SWAP}
	DYDX_USDT_SWAP     = CurrencyPair{CurrencyA: DYDX, CurrencyB: USDT_SWAP}
	EFI_USDT_SWAP      = CurrencyPair{CurrencyA: EFI, CurrencyB: USDT_SWAP}
	EGLD_USDT_SWAP     = CurrencyPair{CurrencyA: EGLD, CurrencyB: USDT_SWAP}
	ELON_USDT_SWAP     = CurrencyPair{CurrencyA: ELON, CurrencyB: USDT_SWAP}
	ENJ_USDT_SWAP      = CurrencyPair{CurrencyA: ENJ, CurrencyB: USDT_SWAP}
	ENS_USDT_SWAP      = CurrencyPair{CurrencyA: ENS, CurrencyB: USDT_SWAP}
	EOS_USDT_SWAP      = CurrencyPair{CurrencyA: EOS, CurrencyB: USDT_SWAP}
	ETC_USDT_SWAP      = CurrencyPair{CurrencyA: ETC, CurrencyB: USDT_SWAP}
	FIL_USDT_SWAP      = CurrencyPair{CurrencyA: FIL, CurrencyB: USDT_SWAP}
	FLM_USDT_SWAP      = CurrencyPair{CurrencyA: FLM, CurrencyB: USDT_SWAP}
	GALA_USDT_SWAP     = CurrencyPair{CurrencyA: GALA, CurrencyB: USDT_SWAP}
	GOD_USDT_SWAP      = CurrencyPair{CurrencyA: GOD, CurrencyB: USDT_SWAP}
	GRT_USDT_SWAP      = CurrencyPair{CurrencyA: GRT, CurrencyB: USDT_SWAP}
	ICP_USDT_SWAP      = CurrencyPair{CurrencyA: ICP, CurrencyB: USDT_SWAP}
	IMX_USDT_SWAP      = CurrencyPair{CurrencyA: IMX, CurrencyB: USDT_SWAP}
	IOST_USDT_SWAP     = CurrencyPair{CurrencyA: IOST, CurrencyB: USDT_SWAP}
	IOTA_USDT_SWAP     = CurrencyPair{CurrencyA: IOTA, CurrencyB: USDT_SWAP}
	JST_USDT_SWAP      = CurrencyPair{CurrencyA: JST, CurrencyB: USDT_SWAP}
	KISHU_USDT_SWAP    = CurrencyPair{CurrencyA: KISHU, CurrencyB: USDT_SWAP}
	KNC_USDT_SWAP      = CurrencyPair{CurrencyA: KNC, CurrencyB: USDT_SWAP}
	KSM_USDT_SWAP      = CurrencyPair{CurrencyA: KSM, CurrencyB: USDT_SWAP}
	LAT_USDT_SWAP      = CurrencyPair{CurrencyA: LAT, CurrencyB: USDT_SWAP}
	LINK_USDT_SWAP     = CurrencyPair{CurrencyA: LINK, CurrencyB: USDT_SWAP}
	LON_USDT_SWAP      = CurrencyPair{CurrencyA: LON, CurrencyB: USDT_SWAP}
	LPT_USDT_SWAP      = CurrencyPair{CurrencyA: LPT, CurrencyB: USDT_SWAP}
	LRC_USDT_SWAP      = CurrencyPair{CurrencyA: LRC, CurrencyB: USDT_SWAP}
	MANA_USDT_SWAP     = CurrencyPair{CurrencyA: MANA, CurrencyB: USDT_SWAP}
	MASK_USDT_SWAP     = CurrencyPair{CurrencyA: MASK, CurrencyB: USDT_SWAP}
	MATIC_USDT_SWAP    = CurrencyPair{CurrencyA: MATIC, CurrencyB: USDT_SWAP}
	MINA_USDT_SWAP     = CurrencyPair{CurrencyA: MINA, CurrencyB: USDT_SWAP}
	MIR_USDT_SWAP      = CurrencyPair{CurrencyA: MIR, CurrencyB: USDT_SWAP}
	MKR_USDT_SWAP      = CurrencyPair{CurrencyA: MKR, CurrencyB: USDT_SWAP}
	NEAR_USDT_SWAP     = CurrencyPair{CurrencyA: NEAR, CurrencyB: USDT_SWAP}
	NEO_USDT_SWAP      = CurrencyPair{CurrencyA: NEO, CurrencyB: USDT_SWAP}
	NFT_USDT_SWAP      = CurrencyPair{CurrencyA: NFT, CurrencyB: USDT_SWAP}
	OMG_USDT_SWAP      = CurrencyPair{CurrencyA: OMG, CurrencyB: USDT_SWAP}
	ONT_USDT_SWAP      = CurrencyPair{CurrencyA: ONT, CurrencyB: USDT_SWAP}
	PEOPLE_USDT_SWAP   = CurrencyPair{CurrencyA: PEOPLE, CurrencyB: USDT_SWAP}
	PERP_USDT_SWAP     = CurrencyPair{CurrencyA: PERP, CurrencyB: USDT_SWAP}
	QTUM_USDT_SWAP     = CurrencyPair{CurrencyA: QTUM, CurrencyB: USDT_SWAP}
	REN_USDT_SWAP      = CurrencyPair{CurrencyA: REN, CurrencyB: USDT_SWAP}
	RSR_USDT_SWAP      = CurrencyPair{CurrencyA: RSR, CurrencyB: USDT_SWAP}
	RVN_USDT_SWAP      = CurrencyPair{CurrencyA: RVN, CurrencyB: USDT_SWAP}
	SAND_USDT_SWAP     = CurrencyPair{CurrencyA: SAND, CurrencyB: USDT_SWAP}
	SC_USDT_SWAP       = CurrencyPair{CurrencyA: SC, CurrencyB: USDT_SWAP}
	SLP_USDT_SWAP      = CurrencyPair{CurrencyA: SLP, CurrencyB: USDT_SWAP}
	SNX_USDT_SWAP      = CurrencyPair{CurrencyA: SNX, CurrencyB: USDT_SWAP}
	SRM_USDT_SWAP      = CurrencyPair{CurrencyA: SRM, CurrencyB: USDT_SWAP}
	STARL_USDT_SWAP    = CurrencyPair{CurrencyA: STARL, CurrencyB: USDT_SWAP}
	STORJ_USDT_SWAP    = CurrencyPair{CurrencyA: STORJ, CurrencyB: USDT_SWAP}
	SUSHI_USDT_SWAP    = CurrencyPair{CurrencyA: SUSHI, CurrencyB: USDT_SWAP}
	SWRV_USDT_SWAP     = CurrencyPair{CurrencyA: SWRV, CurrencyB: USDT_SWAP}
	THETA_USDT_SWAP    = CurrencyPair{CurrencyA: THETA, CurrencyB: USDT_SWAP}
	TONCOIN_USDT_SWAP  = CurrencyPair{CurrencyA: TONCOIN, CurrencyB: USDT_SWAP}
	TORN_USDT_SWAP     = CurrencyPair{CurrencyA: TORN, CurrencyB: USDT_SWAP}
	TRB_USDT_SWAP      = CurrencyPair{CurrencyA: TRB, CurrencyB: USDT_SWAP}
	TRX_USDT_SWAP      = CurrencyPair{CurrencyA: TRX, CurrencyB: USDT_SWAP}
	UMA_USDT_SWAP      = CurrencyPair{CurrencyA: UMA, CurrencyB: USDT_SWAP}
	UNI_USDT_SWAP      = CurrencyPair{CurrencyA: UNI, CurrencyB: USDT_SWAP}
	WAVES_USDT_SWAP    = CurrencyPair{CurrencyA: WAVES, CurrencyB: USDT_SWAP}
	WNCG_USDT_SWAP     = CurrencyPair{CurrencyA: WNCG, CurrencyB: USDT_SWAP}
	WNXM_USDT_SWAP     = CurrencyPair{CurrencyA: WNXM, CurrencyB: USDT_SWAP}
	XCH_USDT_SWAP      = CurrencyPair{CurrencyA: XCH, CurrencyB: USDT_SWAP}
	XEM_USDT_SWAP      = CurrencyPair{CurrencyA: XEM, CurrencyB: USDT_SWAP}
	XLM_USDT_SWAP      = CurrencyPair{CurrencyA: XLM, CurrencyB: USDT_SWAP}
	XMR_USDT_SWAP      = CurrencyPair{CurrencyA: XMR, CurrencyB: USDT_SWAP}
	XRP_USDT_SWAP      = CurrencyPair{CurrencyA: XRP, CurrencyB: USDT_SWAP}
	XTZ_USDT_SWAP      = CurrencyPair{CurrencyA: XTZ, CurrencyB: USDT_SWAP}
	YFI_USDT_SWAP      = CurrencyPair{CurrencyA: YFI, CurrencyB: USDT_SWAP}
	YFII_USDT_SWAP     = CurrencyPair{CurrencyA: YFII, CurrencyB: USDT_SWAP}
	YGG_USDT_SWAP      = CurrencyPair{CurrencyA: YGG, CurrencyB: USDT_SWAP}
	ZEC_USDT_SWAP      = CurrencyPair{CurrencyA: ZEC, CurrencyB: USDT_SWAP}
	ZEN_USDT_SWAP      = CurrencyPair{CurrencyA: ZEN, CurrencyB: USDT_SWAP}
	ZIL_USDT_SWAP      = CurrencyPair{CurrencyA: ZIL, CurrencyB: USDT_SWAP}
	ZRX_USDT_SWAP      = CurrencyPair{CurrencyA: ZRX, CurrencyB: USDT_SWAP}

	BTC_USD   = CurrencyPair{CurrencyA: BTC, CurrencyB: USD, AmountTickSize: 2, PriceTickSize: 1}
	LTC_USD   = CurrencyPair{CurrencyA: LTC, CurrencyB: USD, AmountTickSize: 2, PriceTickSize: 2}
	ETH_USD   = CurrencyPair{CurrencyA: ETH, CurrencyB: USD, AmountTickSize: 2, PriceTickSize: 2}
	ETC_USD   = CurrencyPair{CurrencyA: ETC, CurrencyB: USD, AmountTickSize: 2, PriceTickSize: 2}
	BCH_USD   = CurrencyPair{CurrencyA: BCH, CurrencyB: USD, AmountTickSize: 2, PriceTickSize: 2}
	XRP_USD   = CurrencyPair{CurrencyA: XRP, CurrencyB: USD, AmountTickSize: 3, PriceTickSize: 3}
	BCD_USD   = CurrencyPair{CurrencyA: BCD, CurrencyB: USD, AmountTickSize: 2, PriceTickSize: 3}
	EOS_USD   = CurrencyPair{CurrencyA: EOS, CurrencyB: USD, AmountTickSize: 2, PriceTickSize: 2}
	BTG_USD   = CurrencyPair{CurrencyA: BTG, CurrencyB: USD, AmountTickSize: 2, PriceTickSize: 2}
	BSV_USD   = CurrencyPair{CurrencyA: BSV, CurrencyB: USD, AmountTickSize: 2, PriceTickSize: 2}
	DOT_USD   = CurrencyPair{CurrencyA: DOT, CurrencyB: USD, AmountTickSize: 3, PriceTickSize: 2}
	DASH_USD  = CurrencyPair{CurrencyA: DASH, CurrencyB: USD, AmountTickSize: 2, PriceTickSize: 2}
	CRV_USD   = CurrencyPair{CurrencyA: CRV, CurrencyB: USD, AmountTickSize: 4, PriceTickSize: 3}
	ALGO_USD  = CurrencyPair{CurrencyA: ALGO, CurrencyB: USD, AmountTickSize: 4, PriceTickSize: 4}
	BTC_USDT  = CurrencyPair{CurrencyA: BTC, CurrencyB: USDT, AmountTickSize: 2, PriceTickSize: 1}
	LTC_USDT  = CurrencyPair{CurrencyA: LTC, CurrencyB: USDT, AmountTickSize: 2, PriceTickSize: 2}
	BCH_USDT  = CurrencyPair{CurrencyA: BCH, CurrencyB: USDT, AmountTickSize: 2, PriceTickSize: 2}
	ETC_USDT  = CurrencyPair{CurrencyA: ETC, CurrencyB: USDT, AmountTickSize: 2, PriceTickSize: 3}
	ETH_USDT  = CurrencyPair{CurrencyA: ETH, CurrencyB: USDT, AmountTickSize: 2, PriceTickSize: 2}
	BCD_USDT  = CurrencyPair{CurrencyA: BCD, CurrencyB: USDT, AmountTickSize: 2, PriceTickSize: 2}
	NEO_USDT  = CurrencyPair{CurrencyA: NEO, CurrencyB: USDT, AmountTickSize: 2, PriceTickSize: 2}
	EOS_USDT  = CurrencyPair{CurrencyA: EOS, CurrencyB: USDT, AmountTickSize: 2, PriceTickSize: 2}
	XRP_USDT  = CurrencyPair{CurrencyA: XRP, CurrencyB: USDT, AmountTickSize: 2, PriceTickSize: 2}
	HSR_USDT  = CurrencyPair{CurrencyA: HSR, CurrencyB: USDT, AmountTickSize: 2, PriceTickSize: 2}
	BSV_USDT  = CurrencyPair{CurrencyA: BSV, CurrencyB: USDT, AmountTickSize: 2, PriceTickSize: 2}
	OKB_USDT  = CurrencyPair{CurrencyA: OKB, CurrencyB: USDT, AmountTickSize: 2, PriceTickSize: 2}
	HT_USDT   = CurrencyPair{CurrencyA: HT, CurrencyB: USDT, AmountTickSize: 2, PriceTickSize: 4}
	BNB_USDT  = CurrencyPair{CurrencyA: BNB, CurrencyB: USDT, AmountTickSize: 2, PriceTickSize: 2}
	PAX_USDT  = CurrencyPair{CurrencyA: PAX, CurrencyB: USDT, AmountTickSize: 2, PriceTickSize: 3}
	TRX_USDT  = CurrencyPair{CurrencyA: TRX, CurrencyB: USDT, AmountTickSize: 2, PriceTickSize: 3}
	DOT_USDT  = CurrencyPair{CurrencyA: DOT, CurrencyB: USDT, AmountTickSize: 3, PriceTickSize: 2}
	DASH_USDT = CurrencyPair{CurrencyA: DASH, CurrencyB: USDT, AmountTickSize: 2, PriceTickSize: 2}
	CRV_USDT  = CurrencyPair{CurrencyA: CRV, CurrencyB: USDT, AmountTickSize: 3, PriceTickSize: 3}
	ALGO_USDT = CurrencyPair{CurrencyA: ALGO, CurrencyB: USDT, AmountTickSize: 3, PriceTickSize: 4}

	XRP_EUR = CurrencyPair{CurrencyA: XRP, CurrencyB: EUR, AmountTickSize: 2, PriceTickSize: 4}

	BTC_JPY = CurrencyPair{CurrencyA: BTC, CurrencyB: JPY, AmountTickSize: 2, PriceTickSize: 0}
	LTC_JPY = CurrencyPair{CurrencyA: LTC, CurrencyB: JPY, AmountTickSize: 2, PriceTickSize: 0}
	ETH_JPY = CurrencyPair{CurrencyA: ETH, CurrencyB: JPY, AmountTickSize: 2, PriceTickSize: 0}
	ETC_JPY = CurrencyPair{CurrencyA: ETC, CurrencyB: JPY, AmountTickSize: 2, PriceTickSize: 0}
	BCH_JPY = CurrencyPair{CurrencyA: BCH, CurrencyB: JPY, AmountTickSize: 2, PriceTickSize: 0}

	LTC_BTC = CurrencyPair{CurrencyA: LTC, CurrencyB: BTC, AmountTickSize: 2, PriceTickSize: 4}
	ETH_BTC = CurrencyPair{CurrencyA: ETH, CurrencyB: BTC, AmountTickSize: 2, PriceTickSize: 4}
	ETC_BTC = CurrencyPair{CurrencyA: ETC, CurrencyB: BTC, AmountTickSize: 2, PriceTickSize: 4}
	BCC_BTC = CurrencyPair{CurrencyA: BCC, CurrencyB: BTC, AmountTickSize: 2, PriceTickSize: 4}
	BCH_BTC = CurrencyPair{CurrencyA: BCH, CurrencyB: BTC, AmountTickSize: 2, PriceTickSize: 4}
	DCR_BTC = CurrencyPair{CurrencyA: DCR, CurrencyB: BTC, AmountTickSize: 2, PriceTickSize: 4}
	XRP_BTC = CurrencyPair{CurrencyA: XRP, CurrencyB: BTC, AmountTickSize: 2, PriceTickSize: 6}
	BTG_BTC = CurrencyPair{CurrencyA: BTG, CurrencyB: BTC, AmountTickSize: 2, PriceTickSize: 4}
	BCD_BTC = CurrencyPair{CurrencyA: BCD, CurrencyB: BTC, AmountTickSize: 2, PriceTickSize: 4}
	NEO_BTC = CurrencyPair{CurrencyA: NEO, CurrencyB: BTC, AmountTickSize: 2, PriceTickSize: 4}
	EOS_BTC = CurrencyPair{CurrencyA: EOS, CurrencyB: BTC, AmountTickSize: 2, PriceTickSize: 5}
	HSR_BTC = CurrencyPair{CurrencyA: HSR, CurrencyB: BTC, AmountTickSize: 2, PriceTickSize: 4}
	BSV_BTC = CurrencyPair{CurrencyA: BSV, CurrencyB: BTC, AmountTickSize: 2, PriceTickSize: 4}
	OKB_BTC = CurrencyPair{CurrencyA: OKB, CurrencyB: BTC, AmountTickSize: 2, PriceTickSize: 6}
	HT_BTC  = CurrencyPair{CurrencyA: HT, CurrencyB: BTC, AmountTickSize: 2, PriceTickSize: 7}
	BNB_BTC = CurrencyPair{CurrencyA: BNB, CurrencyB: BTC, AmountTickSize: 2, PriceTickSize: 6}
	TRX_BTC = CurrencyPair{CurrencyA: TRX, CurrencyB: BTC, AmountTickSize: 2, PriceTickSize: 7}
	DOT_BTC = CurrencyPair{CurrencyA: DOT, CurrencyB: BTC, AmountTickSize: 3, PriceTickSize: 6}

	ETC_ETH = CurrencyPair{CurrencyA: ETC, CurrencyB: ETH, AmountTickSize: 2, PriceTickSize: 4}
	EOS_ETH = CurrencyPair{CurrencyA: EOS, CurrencyB: ETH, AmountTickSize: 2, PriceTickSize: 4}
	ZEC_ETH = CurrencyPair{CurrencyA: ZEC, CurrencyB: ETH, AmountTickSize: 2, PriceTickSize: 4}
	NEO_ETH = CurrencyPair{CurrencyA: NEO, CurrencyB: ETH, AmountTickSize: 2, PriceTickSize: 4}
	HSR_ETH = CurrencyPair{CurrencyA: HSR, CurrencyB: ETH, AmountTickSize: 2, PriceTickSize: 4}
	LTC_ETH = CurrencyPair{CurrencyA: LTC, CurrencyB: ETH, AmountTickSize: 2, PriceTickSize: 4}

	UNKNOWN_PAIR = CurrencyPair{CurrencyA: UNKNOWN, CurrencyB: UNKNOWN}
)

func (pair CurrencyPair) String() string {
	return pair.ToSymbol("_")
}

func (pair CurrencyPair) Eq(c2 CurrencyPair) bool {
	return pair.String() == c2.String()
}

func (c Currency) AdaptBchToBcc() Currency {
	if c.Symbol == "BCH" || c.Symbol == "bch" {
		return BCC
	}
	return c
}

func (c Currency) AdaptBccToBch() Currency {
	if c.Symbol == "BCC" || c.Symbol == "bcc" {
		return BCH
	}
	return c
}

func NewCurrency(symbol, desc string) Currency {
	switch symbol {
	case "cny", "CNY":
		return CNY
	case "usdt", "USDT":
		return USDT
	case "usd", "USD":
		return USD
	case "usdc", "USDC":
		return USDC
	case "pax", "PAX":
		return PAX
	case "jpy", "JPY":
		return JPY
	case "krw", "KRW":
		return KRW
	case "eur", "EUR":
		return EUR
	case "btc", "BTC":
		return BTC
	case "xbt", "XBT":
		return XBT
	case "bch", "BCH":
		return BCH
	case "bcc", "BCC":
		return BCC
	case "ltc", "LTC":
		return LTC
	case "sc", "SC":
		return SC
	case "ans", "ANS":
		return ANS
	case "neo", "NEO":
		return NEO
	case "okb", "OKB":
		return OKB
	case "ht", "HT":
		return HT
	case "bnb", "BNB":
		return BNB
	case "trx", "TRX":
		return TRX
	case "dot", "DOT":
		return DOT
	default:
		return Currency{strings.ToUpper(symbol), desc}
	}
}

func NewCurrencyPair(currencyA Currency, currencyB Currency) CurrencyPair {
	return CurrencyPair{CurrencyA: currencyA, CurrencyB: currencyB}
}

func NewCurrencyPair2(currencyPairSymbol string) CurrencyPair {
	return NewCurrencyPair3(currencyPairSymbol, "_")
}

func NewCurrencyPair3(currencyPairSymbol string, sep string) CurrencyPair {
	currencys := strings.Split(currencyPairSymbol, sep)
	if len(currencys) >= 2 {
		return CurrencyPair{CurrencyA: NewCurrency(currencys[0], ""),
			CurrencyB: NewCurrency(currencys[1], "")}
	}
	return UNKNOWN_PAIR
}

func (pair *CurrencyPair) SetAmountTickSize(tickSize int) CurrencyPair {
	pair.AmountTickSize = tickSize
	return *pair
}

func (pair *CurrencyPair) SetPriceTickSize(tickSize int) CurrencyPair {
	pair.PriceTickSize = tickSize
	return *pair
}

func (pair CurrencyPair) ToSymbol(joinChar string) string {
	return strings.Join([]string{pair.CurrencyA.Symbol, pair.CurrencyB.Symbol}, joinChar)
}

func (pair CurrencyPair) ToSymbol2(joinChar string) string {
	return strings.Join([]string{pair.CurrencyB.Symbol, pair.CurrencyA.Symbol}, joinChar)
}

func (pair CurrencyPair) AdaptUsdtToUsd() CurrencyPair {
	CurrencyB := pair.CurrencyB
	if pair.CurrencyB.Eq(USDT) {
		CurrencyB = USD
	}
	pair.CurrencyB = CurrencyB
	return pair
}

func (pair CurrencyPair) AdaptUsdToUsdt() CurrencyPair {
	CurrencyB := pair.CurrencyB
	if pair.CurrencyB.Eq(USD) {
		CurrencyB = USDT
	}
	pair.CurrencyB = CurrencyB
	return pair
}

//for to symbol lower , Not practical '==' operation method
func (pair CurrencyPair) ToLower() CurrencyPair {
	return CurrencyPair{CurrencyA: Currency{Symbol: strings.ToLower(pair.CurrencyA.Symbol), Desc: pair.CurrencyA.Desc},
		CurrencyB: Currency{Symbol: strings.ToLower(pair.CurrencyB.Symbol), Desc: pair.CurrencyB.Desc}}
}

func (pair CurrencyPair) Reverse() CurrencyPair {
	return CurrencyPair{CurrencyA: pair.CurrencyB, CurrencyB: pair.CurrencyA,
		AmountTickSize: pair.AmountTickSize, PriceTickSize: pair.PriceTickSize}
}
