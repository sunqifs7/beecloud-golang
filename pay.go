package bcGolang

import (
	"fmt"
	"reflect"
	"math"
	"time"
	"encoding/json"
)

type BCPay struct {
	bcApp 	BCApp
	//timeout	int
}

func (this *BCPay) RegisterApp(bcApp BCApp) {
	this.bcApp = bcApp
}

func (this *BCPay) Pay(payParam BCPayParams) BCPayResult {
	var bcPayResult BCPayResult
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

	response, ok := HttpPost(this.getBillPayUrl(), para)
	if  !ok {
		fmt.Println("Error returned. Should return a BCResult with error code")
		bcPayResult.BCResult = HandleInvalidResp(response)
		return bcPayResult
	}
	if err := json.Unmarshal(response, &bcPayResult); err != nil {
		fmt.Println("json.Unmarshal error")
		// should be an exception, since this might be BC-related
	}

	return bcPayResult

	/*
	bcPayResult.ResultCode = resMap["result_code"]
	bcPayResult.ResultCode = resMap["result_msg"]
	bcPayResult.ResultCode = resMap["err_detail"]
	bcPayResult.ResultCode = resMap["id"]
	if bcPayResult.ResultCode != 0 {
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
	*/

}

func (this *BCPay) Refund(refundParam BCRefundReqParams) BCRefundResult {
	var bcRefundResult BCRefundResult
	if this.bcApp.IsTestMode {
		bcRefundResult.BCResult = NotSupportedTestError("refund")
		return bcRefundResult
	}
	AttachAppSign(&refundParam.BCReqParams, REFUND, this.bcApp)
	fmt.Println(refundParam.BCReqParams)
	para := constructRefundParamMap(refundParam)

	content, ok := HttpPost(this.getBillRefundUrl(), para)
	if !ok {
		fmt.Println("Error returned. Should return a BCResult with error code")
		bcRefundResult.BCResult = HandleInvalidResp(content)
		return bcRefundResult
	}
	if err := json.Unmarshal(content, &bcRefundResult); err != nil {
		fmt.Println("json.Unmarshal error")
		// should be an exception
	}

	return bcRefundResult

	/*
	bcRefundResult.ResultCode = resMap["result_code"]
	bcRefundResult.ResultMsg = resMap["result_msg"]
	bcRefundResult.ErrDetail = resMap["err_detail"]
	id, ok := resMap["id"]
	if ok {
		bcRefundResult.Id = id
	}
	if refundParam.Channel == ALI && refundParam.NeedApproval != true && bcRefundResult.ResultCode == 0 {
		bcRefundResult.Url = resMap["url"]
	}
	*/

}

func (this *BCPay) AuditPreRefunds(preRefundParam BCPreRefundParams) BCPreRefundResult {
	var bcPreRefundResult BCPreRefundResult
	if this.bcApp.IsTestMode {
		bcPreRefundResult.BCResult = NotSupportedTestError("audit_pre_refunds")
		return bcPreRefundResult
	}
	// PAY here??? need to check
	AttachAppSign(&preRefundParams.BCReqParams, PAY, this.bcApp)
	fmt.Println(preRefundParams.BCReqParams)
	para := constructPreRefundParamMap(preRefundParam)

}

// private methods
func (this *BCPay) getBillPayUrl() string {
	if this.bcApp.IsTestMode {
		return GetRandomHost() + "rest/sandbox/bill"
	} else {
		return GetRandomHost() + "rest/bill"
	}
}

func (this *BCPay) getBillRefundUrl() string {
	return GetRandomHost() + "rest/refund"
}

func constructBCReqParamMap(bcReqParam BCReqParams) MapObject {
	para := make(MapObject)
	para["app_id"] = bcReqParam.AppId
	para["app_sign"] = bcReqParam.AppSign
	para["timestamp"] = bcReqParam.Timestamp //毫秒数
	return para
}

func constructPayReqParamMap(payParam BCPayReqParams) MapObject {
	para := constructBCReqParamMap(payParam.BCReqParams)
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

func constructRefundParamMap(refundParam BCRefundReqParams) MapObject {
	para := constructBCReqParamMap(refundParam.BCReqParams)
	if !StrEmpty(refundParam.Channel) {
		para["channel"] = refundParam.Channel
	}
	para["refund_no"] = refundParam.RefundNo
	para["bill_no"] = refundParam.BillNo
	para["refund_fee"] = refundParam.RefundFee
	para["optional"] = refundParam.Optional // what happends if optional is empty?
	if refundParam.NeedApproval != nil {
		para["need_approval"] = refundParam.NeedApproval
	}
}

func constructPreRefundParamMap(preRefundParam BCPreRefundParams) MapObject {
	para := constructBCReqParamMap(preRefundParam.BCReqParams)
	para["channel"] = preRefundParam.Channel // problem here
	para["ids"] = preRefundParam.Ids
	para["agree"] = preRefundParam.Agree
	para["deny_reason"] = preRefundParam.DenyReason

	return para
}
