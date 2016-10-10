package bcGolang

/*
beecloud entity
~~~~~~~~~
This module contains data entity.
:created by sunqi on 2016/08/15.
:copyright (c) 2016 BeeCloud.
:license struct { MIT, see LICENSE for more details.
*/

import (
	"time"
	"go/types"
	"net/url"
)

type BCApp struct {
	//correspond to console app
	AppId        string
	AppSecret    string
	TestSecret   string
	MasterSecret string
	IsTestMode  bool
	Timeout       int
}

type BCChannelType int
const (
	// BCChannelType enum

	WX BCChannelType = iota
	WX_APP
	WX_NATIVE
	WX_JSAPI
	WX_REDPACK
	WX_TRANSFER

	ALI
	ALI_APP
	ALI_WEB
	ALI_WAP
	ALI_QRCODE
	ALI_OFFLINE_QRCODE
	ALI_TRANSFER

	UN
	UN_APP
	UN_WEB
	UN_WAP

	JD
	JD_WAP
	JD_WEB

	YEE
	YEE_WAP
	YEE_WEB
	YEE_NOBANKCARD

	KUAIQIAN
	KUAIQIAN_WEB
	KUAIQIAN_WAP

	PAYPAL
	// 生产环境支付，用于手机APP
	PAYPAL_LIVE
	// 沙箱环境支付，用于手机APP
	PAYPAL_SANDBOX
	// 以下用于PC
	// 跳转到paypal使用paypal内支付
	PAYPAL_PAYPAL
	// 直接使用信用卡支付（paypal渠道）
	PAYPAL_CREDITCARD
	// 使用存储的行用卡id支付（信用卡信息存储在PAYPAL）
	PAYPAL_SAVED_CREDITCARD

	BD
	BD_APP
	BD_WAP
	BD_WEB

	BC
	BC_GATEWAY
	BC_APP
	BC_EXPRESS
	BC_TRANSFER
)

type BCReqestType int
const (
	PAY BCReqestType = iota
	QUERY
	REFUND
	TRANSFER
)

type CardType int
const (
	DE CardType = iota
	CR
)

type CurrencyType int
const (
// 三位货币种类代码
)

type stringSlice struct {
	s []string
}
func (ss *stringSlice) Append(x string) {
	ss.s = append(ss.s, x)
}

type BCReqParams struct {
	AppId 		string
	AppSign		string
	Timestamp	int64
}

type BCPayReqParams struct {
	// app_id, app_sign, timestamp
	BCReqParams
	// 渠道类型
	Channel     BCChannelType
	// 订单总金额
	TotalFee    int
	// 商户订单号
	BillNo      string
	// 订单标题
	Title        string
	// 分析数据
	Analysis     MapObject
	// 同步返回页面
	ReturnUrl   string
	// 订单失效时间
	BillTimeout int64
	// 附加数据
	Optional     MapObject // This is a nil map. Need to initialize when using.
}

type BCPayParams struct {
	BCPayReqParams
	// WX_JSAPI, req
	Openid		string
	// ALI_WEB, opt
	ShowUrl		string
	// ALI_WAP, opt
	UseApp		bool
	// ALI_QRCODE, req
	QrPayMode 	string
	// YEE_WAP, req
	IdentifyId	string
	// YEE_NOBANKCARD, req
	CardNo		string
	CardPwd		string
	FrqId		string
}

type BCRefundReqParams struct {
	// app_id, app_sign, timestamp
	BCReqParams
	// 渠道类型
	Channel       BCChannelType
	// 商户退款单号
	RefundNo     string
	// 商户订单号
	BillNo       string
	// 退款金额
	RefundFee    int
	// 是否为预退款
	NeedApproval bool
	// 附加数据
	Optional     MapObject // This is a nil map. Need to initialize when using.
}

type BCPreRefundParams struct {
	// app_id, app_sign, timestamp
	BCReqParams
	// 渠道类型
	Channel     BCChannelType
	// 预退款记录id列表
	Ids         stringSlice
	// 同意或者驳回
	Agree       bool
	// 驳回理由; optional
	DenyReason string
}


type BCQueryReqParams struct {
	// app_id, app_sign, timestamp
	BCReqParams
	// 渠道类型
	Channel       BCChannelType
	// 商户订单号
	BillNo       string
	// 商户退款单号 仅对查询退款有效
	RefundNo     string
	// 订单是否成功,仅对支付查询有效
	spayResult   bool
	// 标识退款记录是否为预退款 仅对查询退款有效
	NeedApproval bool
	// 是否需要返回渠道详细信息
	NeedDetail   bool
	// 起始时间
	startTime    time.Time
	// 结束时间
	endTime      time.Time
	// 查询起始位置
	Skip          int
	// 查询的条数
	Limit         int
}

type BCResult struct {
	ResultCode int		`json:"result_code"`
	ResultMsg  string	`json:"result_msg"`
	ErrDetail  string	`json:"err_detail"`

}

type BCPayResult struct {
	BCResult
	Id		   string	`json:"id"`
	// WX_APP, WX_JSAPI
	AppId		string	`json:"app_id"`
		// WX_APP
	PartnerId	string 	`json:"partner_id"`
	Package		string	`json:"package"`
	NonceStr	string	`json:"nonce_str"`
	Timestamp	string	`json:"timestamp"`
	PaySign		string	`json:"pay_sign"`
		// WX_APP
	PrepayId	string 	`json:"prepay_id"`
		// WX_JSAPI
	SignType	string	`json:"sign_type"`
	// WX_NATIVE
	CodeUrl		string  `json:"code_url"`
	// ALI_APP
	OrderString	string	`json:"order_string"`
	// ALI_WEB，ALI_WAP, ALI_QRCODE |
	// Html only -> UN_WEB、UN_WAP、JD_WAP、JD_WEB、KUAIQIAN_WAP、KUAIQIAN_WEB
	// Url only -> YEE_WAP、YEE_WEB、BD_WAP、BD_WEB
	Html		string	`json:"html"`
	Url			string	`json:"url"`
	// ALI_OFFLINE_QRCODE
	QrCode		string	`json:"qr_code"`
	// UN_APP
	Tn			string	`json:"tn"`
	// BD_APP
	OrderInfo	string	`json:"orderInfo"`
}

type BCRefundResult struct {
	BCResult
	Id		string		`json:"id"`
	// ALI
	Url 	string		`json:"url"`
}

type BCPreRefundResult struct {
	BCResult
	// when agree == true && result_code == 0
	ResultMap	MapObject	`json:"result_map"`
	// when agree == true && result_code == 0 && channel == ALI
	Url 		string		`json:"url"`
}

type BCBill struct {
	// 订单记录的唯一标识
	id             string
	// 订单号
	BillNo        string
	Channel        BCChannelType
	// 子渠道
	SubChannel    BCChannelType
	// 渠道交易号，当支付成功时有值
	TradeNo       string
	// 订单创建时间，毫秒时间戳，13位
	createTime    time.Time
	// 订单是否成功
	spayResult    bool
	// 商品标题
	Title          string
	// 订单金额，单位为分
	TotalFee      int
	// 渠道详细信息
	MessageDetail string
	// 订单是否已经撤销
	revertResult  bool
	// 订单是否已经退款
	refundResult  bool
	// 附加数据
	Optional MapObject
}

type BCRefund struct {
	// 退款记录的唯一标识
	id             string
	// 支付订单号
	BillNo        string
	Channel        BCChannelType
	// 子渠道
	SubChannel    BCChannelType
	// 退款是否完成
	Finish         bool
	// 退款创建时间
	createTime    time.Time
	// 退款是否成功
	result         bool
	// 商品标题
	Title          string
	// 订单金额，单位为分
	TotalFee      int
	// 退款金额，单位为分
	RefundFee     int
	// 退款单号
	RefundNo      string
	// 渠道详细信息
	MessageDetail string
	// 附加数据
	Optional     MapObject // This is a nil map. Need to initialize when using.
}

type BCTransferReqParams struct {
	// app_id, app_sign, timestamp
	BCReqParams
	// 渠道类型 WXREDPACK, WXTRANSFER, ALITRANSFER
	Channel           BCChannelType
	// 打款单号
	TransferNo       string
	// 打款金额
	TotalFee         int
	// 打款说明
	Desc              string
	// 支付渠道方内收款人的标示，微信为openid，支付宝为支付宝账户
	ChannelUserId   string
	// 支付渠道内收款人账户名， 支付宝必填
	ChannelUserName string
	// 打款方账号名全称，支付宝必填
	AccountName      string
	// 微信红包的详细描述，Map类型，微信红包必填
	RedpackInfo      MapObject
}

type BCTransferRedPack struct {
	// 红包发送者名称
	SendName string
	// 红包祝福语
	Wishing   string
	// 红包活动名称
	ActName  string
}

// 用于bcTransfer
type BCCardTransferParams struct {
	// 下发订单总金额，正整数
	TotalFee     int
	// 商户订单号
	BillNo       string
	// 下发订单标题
	Title         string
	// 银行全名
	BankFullname string
	// 银行卡类型，DE代表借记卡，CR代表信用卡
	CardType     CardType
	// 收款帐户类型
	AccountType  string
	// 收款帐户号
	AccountNo    string
	// 收款帐户名称
	AccountName  string
	// 银行绑定的手机号
	Mobile        string
	// 附加数据，Map类型
	Optional     MapObject // This is a nil map. Need to initialize (make) when using.
	// 交易源 目前只能填写'OUT_PC'
	// note: 需在处理此类时设置为OUT_PC
	TradeSource  string
}


type BCBatchTransferParams struct {
	// 渠道类型 目前只支持ALI
	// note: 需要在处理此类变量时,设置Channel = ALI
	Channel       BCChannelType
	// 打款单号
	BatchNo      string
	// 付款账号账户全称
	AccountName  string
	// 包含每一笔的具体信息，List类型
	TransferData stringSlice
}

type BCBatchTransferItem struct {
	// 打款流水号
	transferId      string
	// 收款方账户
	ReceiverAccount string
	// 收款方账号姓名
	ReceiverName    string
	// 打款金额，单位为分
	TransferFee     int
	// 打款备注
	TransferNote    string
}

type BCInternationalPayParams struct {
	// 渠道类型
	Channel          BCChannelType
	// 订单总金额，单位分
	TotalFee        int
	// 三位货币种类代码
	Currency         CurrencyType
	// 商户订单号
	BillNo          string
	// 订单标题
	Title            string
	// 信用卡信息 BCPayPalCreditCard
	CreditCardInfo BCPayPalCreditCard
	// 信用卡id
	CreditCardId   string
	// 同步返回页面
	ReturnUrl       string
}

type BCPayPalCreditCard struct {
	// 卡号
	CardNumber  string
	// 过期时间中的月
	ExpireMonth int
	// 过期时间中的年
	ExpireYear  int
	// 信用卡的三位cvv码
	Cvv          string
	// 用户名字
	FirstName   string
	// 用户的姓
	LastName    string
	// 卡类别
	// 什么类别???VISA MASTER?
	CardType    string
}

type MapObject map[string]interface{}

