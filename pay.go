package bcGolang

import (
	"fmt"
	"reflect"
	"math"
	"time"
)

type BCPay struct {
	bcApp 	BCApp
	//timeout	int
}

func (this *BCPay) RegisterApp(bcApp BCApp) {
	this.bcApp = bcApp
}

func (this *BCPay) Pay(payParam BCPayParams) BCPayResult {
	AttachAppSign(&payParam.BCReqParams, PAY, this.bcApp)
	fmt.Println(payParam.BCReqParams)
	para := constructPayReqParamMap(payParam.BCPayReqParams)
	switch payParam.Channel {
	case WX_JSAPI:
		para["openid"] = payParam.Openid
	case ALI_WEB:
		if !StrEmpty(payParam.ShowUrl) {
			para["show_url"] = payParam.ShowUrl
		}
	case ALI_WAP:
		para["use_app"] = payParam.UseApp
	case ALI_QRCODE:
		para["qr_pay_mode"] = payParam.QrPayMode
	case YEE_WAP:
		para["identity_id"] = payParam.IdentifyId
	case YEE_NOBANKCARD:
		para["cardno"] = payParam.CardNo
		para["cardpwd"] = payParam.CardPwd
		para["frqid"] = payParam.FrqId
	}

	resMap := HttpPost(this.getBillPayUrl(), para)
	var bcPayResult BCPayResult
	bcPayResult.ResultCode = resMap["result_code"]
	bcPayResult.ResultCode = resMap["result_msg"]
	bcPayResult.ResultCode = resMap["err_detail"]
	bcPayResult.ResultCode = resMap["id"]
	if resMap["result_code"] != 0 {
		return bcPayResult
	}
	switch payParam.Channel {
	case WX_APP:
		bcPayResult.AppId = resMap["app_id"]
		bcPayResult.PartnerId = resMap["partner_id"]
		bcPayResult.Package = resMap["package"]
		bcPayResult.NonceStr = resMap["nonce_str"]
		bcPayResult.Timestamp = resMap["timestamp"]
		bcPayResult.PaySign = resMap["pay_sign"]
		bcPayResult.PrepayId = resMap["prepay_id"]
	case WX_NATIVE:
		bcPayResult.CodeUrl = resMap["code_url"]
	case WX_JSAPI:
		bcPayResult.AppId = resMap["app_id"]
		bcPayResult.Package = resMap["package"]
		bcPayResult.NonceStr = resMap["nonce_str"]
		bcPayResult.Timestamp = resMap["timestamp"]
		bcPayResult.PaySign = resMap["pay_sign"]
		bcPayResult.SignType = resMap["sign_type"]
	case ALI_APP:
		bcPayResult.OrderString = resMap["order_string"]
	case ALI_WEB,ALI_WAP,ALI_QRCODE:
		bcPayResult.Html = resMap["html"]
		bcPayResult.Url = resMap["url"]
	case ALI_OFFLINE_QRCODE:
		bcPayResult.QrCode = resMap["qr_code"]
	case UN_APP:
		bcPayResult.Tn = resMap["tn"]
	case UN_WEB,UN_WAP,JD_WAP,JD_WEB,KUAIQIAN_WAP,KUAIQIAN_WEB:
		bcPayResult.Html = resMap["html"]
	case YEE_WAP,YEE_WEB,BD_WAP,BD_WEB:
		bcPayResult.Url = resMap["url"]
	case BD_APP:
		bcPayResult.OrderInfo = resMap["orderInfo"]
	default:
		fmt.Println("Wrong channel!")
	}

	return bcPayResult
}

// private methods
func (this *BCPay) getBillPayUrl() string {
	if this.bcApp.IsTestMode {
		return GetRandomHost() + "rest/sandbox/bill"
	} else {
		return GetRandomHost() + "rest/bill"
	}
}

func constructPayReqParamMap(payParam BCPayReqParams) MapObject {
	para := make(MapObject)
	para["app_id"] = payParam.AppId
	para["app_sign"] = payParam.AppSign
	para["timestamp"] = payParam.Timestamp
	para["channel"] = payParam.Channel
	para["total_fee"] = payParam.TotalFee
	para["bill_no"] = payParam.BillNo
	para["title"] = payParam.Title
	para["optional"] = payParam.Optional
	para["analysis"] = payParam.Analysis
	para["return_url"] = payParam.ReturnUrl
	para["bill_timeout"] = payParam.BillTimeout
	
	return para
}
