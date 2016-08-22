package bcGolang

/*
beecloud entity
~~~~~~~~~
This module contains data entity.
:created by sunqi on 2016/08/15.
:copyright (c) 2016 BeeCloud.
:license struct { MIT, see LICENSE for more details.
*/



type BCApp struct {
	//correspond to console app
	app_id        string
	app_secret    string
	test_secret   string
	master_secret string
	is_test_mode  bool

	timeout       int
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

type BCReqType int

const (
	PAY BCReqType = iota
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

type StringSlice struct {
	s []string
}
func (ss *StringSlice) Append(x string) {
	ss.s = append(ss.s, x)
}

type BCPayReqParams struct {
	// 渠道类型
	channel      string

	// 订单总金额
	total_fee    int

	// 商户订单号
	bill_no      string

	// 订单标题
	title        string

	// 分析数据
	analysis     string

	// 同步返回页面
	return_url   string

	// 订单失效时间
	bill_timeout int64

	// 附加数据
	optional     map[string]string // This is a nil map. Need to initialize when using.
}


type BCRefundReqParams struct {
	// 渠道类型
	channel       BCChannelType

	// 商户退款单号
	refund_no     string

	// 商户订单号
	bill_no       string

	// 退款金额
	refund_fee    int

	// 是否为预退款
	need_approval bool

	// 附加数据
	optional     map[string]string // This is a nil map. Need to initialize when using.
}


type BCPreRefundAuditParams struct {
	// 渠道类型
	channel     BCChannelType

	// 预退款记录id列表
	ids         StringSlice

	// 同意或者驳回
	agree       bool

	// 驳回理由
	deny_reason string
}


type BCQueryReqParams struct {
	// 渠道类型
	channel       BCChannelType

	// 商户订单号
	bill_no       string

	// 商户退款单号
	// 仅对查询退款有效
	refund_no     string

	// 订单是否成功
	// 仅对支付查询有效
	spay_result   bool

	// 标识退款记录是否为预退款
	// 仅对查询退款有效
	need_approval bool

	// 是否需要返回渠道详细信息
	need_detail   bool

	// 起始时间
	start_time    int64

	// 结束时间
	end_time      int64

	// 查询起始位置
	skip          int

	// 查询的条数
	limit         int
}

type BCResult struct {
	result_code int
	result_msg  string
	err_detail  string
}


type BCBill struct {
	// 订单记录的唯一标识
	id             string

	// 订单号
	bill_no        string

	channel        BCChannelType

	// 子渠道
	sub_channel    BCChannelType

	// 渠道交易号，当支付成功时有值
	trade_no       string

	// 订单创建时间，毫秒时间戳，13位
	create_time    int64

	// 订单是否成功
	spay_result    bool

	// 商品标题
	title          string

	// 订单金额，单位为分
	total_fee      int

	// 渠道详细信息
	message_detail string

	// 订单是否已经撤销
	revert_result  bool

	// 订单是否已经退款
	refund_result  bool
	// 附加数据
}

type BCRefund struct {
	// 退款记录的唯一标识
	id             string
	// 支付订单号
	bill_no        string
	channel        BCChannelType

	// 子渠道
	sub_channel    BCChannelType

	// 退款是否完成
	finish         bool

	// 退款创建时间
	create_time    int64

	// 退款是否成功
	result         bool

	// 商品标题
	title          string

	// 订单金额，单位为分
	total_fee      int

	// 退款金额，单位为分
	refund_fee     int

	// 退款单号
	refund_no      string

	// 渠道详细信息
	message_detail string
	// 附加数据
	optional     map[string]string // This is a nil map. Need to initialize when using.
}

type BCTransferReqParams struct {
	// 渠道类型 WX_REDPACK, WX_TRANSFER, ALI_TRANSFER
	channel           BCChannelType

	// 打款单号
	transfer_no       string

	// 打款金额
	total_fee         int

	// 打款说明
	desc              string

	// 支付渠道方内收款人的标示，微信为openid，支付宝为支付宝账户
	channel_user_id   string

	// 支付渠道内收款人账户名， 支付宝必填
	channel_user_name string

	// 打款方账号名全称，支付宝必填
	account_name      string
	// 微信红包的详细描述，Map类型，微信红包必填
	redpack_info      map[string]string
}


type BCTransferRedPack struct {
	// 红包发送者名称
	send_name string

	// 红包祝福语
	wishing   string

	// 红包活动名称
	act_name  string
}


// 用于bc_transfer
type BCCardTransferParams struct {
	// 下发订单总金额，正整数
	total_fee     int

	// 商户订单号
	bill_no       string

	// 下发订单标题
	title         string

	// 银行全名
	bank_fullname string

	// 银行卡类型，DE代表借记卡，CR代表信用卡
	card_type     CardType

	// 收款帐户类型 ????????????
	account_type  string

	// 收款帐户号
	account_no    string

	// 收款帐户名称
	account_name  string

	// 银行绑定的手机号
	mobile        string

	// 附加数据，Map类型
	optional     map[string]string // This is a nil map. Need to initialize (make) when using.
	// 交易源????????????
	trade_source = 'OUT_PC'
}


type BCBatchTransferParams struct {
	// 渠道类型 目前只支持ALI
	// type正确么???
	channel       BCChannelType.ALI

	// 打款单号
	batch_no      string

	// 付款账号账户全称
	account_name  string

	// 包含每一笔的具体信息，List类型
	transfer_data StringSlice
}



type BCBatchTransferItem struct {
	// 打款流水号
	transfer_id      string

	// 收款方账户
	receiver_account string

	// 收款方账号姓名
	receiver_name    string

	// 打款金额，单位为分
	transfer_fee     int

	// 打款备注
	transfer_note    string
}


type BCInternationalPayParams struct {
	// 渠道类型
	channel          BCChannelType

	// 订单总金额，单位分
	total_fee        int

	// 三位货币种类代码
	currency         CurrencyType

	// 商户订单号
	bill_no          string

	// 订单标题
	title            string

	// 信用卡信息 BCPayPalCreditCard
	credit_card_info BCPayPalCreditCard

	// 信用卡id
	credit_card_id   string

	// 同步返回页面
	return_url       string
}

type BCPayPalCreditCard struct {
	// 卡号
	card_number  string

	// 过期时间中的月
	// 考虑使用time包???
	expire_month string

	// 过期时间中的年
	// 考虑使用time包???
	expire_year  string

	// 信用卡的三位cvv码
	cvv          string

	// 用户名字
	first_name   string

	// 用户的姓
	last_name    string

	// 卡类别
	// 什么类别???VISA MASTER?
	card_type    string
}

